package core

type (
	Build struct {
		ID           int64             `json:"id"`
		RepoID       int64             `json:"repo_id"`
		Trigger      string            `json:"trigger"`
		Number       int64             `json:"number"`
		Parent       int64             `json:"parent,omitempty"`
		Status       string            `json:"status"`
		Error        string            `json:"error,omitempty"`
		Event        string            `json:"event"`
		Action       string            `json:"action"`
		Link         string            `json:"link"`
		Timestamp    int64             `json:"timestamp"`
		Title        string            `json:"title,omitempty"`
		Message      string            `json:"message"`
		Before       string            `json:"before"`
		After        string            `json:"after"`
		Ref          string            `json:"ref"`
		Fork         string            `json:"source_repo"`
		Source       string            `json:"source"`
		Target       string            `json:"target"`
		Author       string            `json:"author_login"`
		AuthorName   string            `json:"author_name"`
		AuthorEmail  string            `json:"author_email"`
		AuthorAvatar string            `json:"author_avatar"`
		Sender       string            `json:"sender"`
		Params       map[string]string `json:"params,omitempty"`
		Cron         string            `json:"cron,omitempty"`
		Deploy       string            `json:"deploy_to,omitempty"`
		DeployID     int64             `json:"deploy_id,omitempty"`
		Started      int64             `json:"started"`
		Finished     int64             `json:"finished"`
		Created      int64             `json:"created"`
		Updated      int64             `json:"updated"`
		Version      int64             `json:"version"`
		Stages       []*Stage          `json:"stages,omitempty"`
	}

	Stage struct {
		ID        int64             `json:"id"`
		BuildID   int64             `json:"build_id"`
		Number    int               `json:"number"`
		Name      string            `json:"name"`
		Kind      string            `json:"kind,omitempty"`
		Type      string            `json:"type,omitempty"`
		Status    string            `json:"status"`
		Error     string            `json:"error,omitempty"`
		ErrIgnore bool              `json:"errignore"`
		ExitCode  int               `json:"exit_code"`
		Machine   string            `json:"machine,omitempty"`
		OS        string            `json:"os"`
		Arch      string            `json:"arch"`
		Variant   string            `json:"variant,omitempty"`
		Kernel    string            `json:"kernel,omitempty"`
		Limit     int               `json:"limit,omitempty"`
		Started   int64             `json:"started"`
		Stopped   int64             `json:"stopped"`
		Created   int64             `json:"created"`
		Updated   int64             `json:"updated"`
		Version   int64             `json:"version"`
		OnSuccess bool              `json:"on_success"`
		OnFailure bool              `json:"on_failure"`
		DependsOn []string          `json:"depends_on,omitempty"`
		Labels    map[string]string `json:"labels,omitempty"`
		Steps     []*Step           `json:"steps,omitempty"`
	}

	Step struct {
		ID        int64  `json:"id"`
		StageID   int64  `json:"step_id"`
		Number    int    `json:"number"`
		Name      string `json:"name"`
		Status    string `json:"status"`
		Error     string `json:"error,omitempty"`
		ErrIgnore bool   `json:"errignore,omitempty"`
		ExitCode  int    `json:"exit_code"`
		Started   int64  `json:"started,omitempty"`
		Stopped   int64  `json:"stopped,omitempty"`
		Version   int64  `json:"version"`
	}
)
