package signature

import (
	"net/http"

	"github.com/99designs/httpsignatures-go"
	"github.com/yckao/drone-convert-advanced/logger"
)

func HandleSignature(secret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)

			signature, err := httpsignatures.FromRequest(r)
			if err != nil {
				log.Debugf("converter: invalid or missing signature in http.Request")
				http.Error(w, "Invalid or Missing Signature", 400)
				return
			}
			if !signature.IsValid(secret, r) {
				log.Debugf("converter: invalid signature in http.Request")
				http.Error(w, "Invalid Signature", 400)
				return
			}
			ctx = logger.WithContext(ctx, log)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
