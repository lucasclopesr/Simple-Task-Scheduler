package meta

// JobRequest representa um request de criação de job
type JobRequest struct {
	User string `yaml:"user"  json:"user"`
	Job  Job    `yaml:"job"  json:"job"`
}
