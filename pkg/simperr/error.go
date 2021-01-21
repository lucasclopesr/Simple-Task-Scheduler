package simperr

type SimpError struct {
	Code    int    `yaml:"code"  json:"code"`
	Message string `yaml:"message"  json:"message"`
}

func (s *SimpError) Error() string {
	return s.Message
}

const (
	ErrorNotFound = iota
	ErrorAlreadyExists
	ErrorMemoryLimit
	ErrorJobLimit
)
