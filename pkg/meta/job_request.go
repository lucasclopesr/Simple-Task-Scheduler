package meta

type JobRequest struct {
	User string `yaml:"user"  json:"user"`
	Job  Job    `yaml:"job"  json:"job"`
}
