package convert

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/yckao/drone-convert-advanced/core"
	"github.com/yckao/drone-convert-advanced/handler/convert/signature"
	"github.com/yckao/drone-convert-advanced/logger"
)

type Server struct {
	Commits core.CommitService
	Repos   core.RepositoryService
	Convert core.ConvertService
	secret  string
}

func New(
	commits core.CommitService,
	repos core.RepositoryService,
	convert core.ConvertService,
	secret string,
) Server {
	return Server{
		Commits: commits,
		Repos:   repos,
		Convert: convert,
		secret:  secret,
	}
}

func (s Server) Handler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Use(logger.Middleware)
	r.Use(signature.HandleSignature(s.secret))

	r.Handle("/", HandleConvert(s.Convert))

	return r
}

func HandleConvert(converter core.ConvertService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.FromRequest(r).Debugf("converter: cannot read http.Request body")
			w.WriteHeader(400)
			return
		}

		req := &core.ConvertArgs{}
		err = json.Unmarshal(body, req)
		if err != nil {
			logger.FromRequest(r).Debugf("converter: cannot unmarshal http.Request body")
			http.Error(w, "Invalid Input", 400)
			return
		}

		res, err := converter.Convert(r.Context(), req)
		if err != nil {
			logger.FromRequest(r).Debugf("converter: cannot convert configuration: %s: %s: %s",
				req.Repo.Slug,
				req.Build.Target,
				err,
			)
			http.Error(w, err.Error(), 404)
			return
		}

		if res == nil {
			w.WriteHeader(204)
			return
		}
		out, _ := json.Marshal(res)
		logger.FromRequest(r).Traceln("converter: converted configuration json: %s",
			string(out),
		)
		w.Header().Set("Content-Type", "application/json")
		w.Write(out)
	}
}
