package simperr

// SimpError é um erro do SIMP
type SimpError struct {
	Code    int    `yaml:"code"  json:"code"`
	Message string `yaml:"message"  json:"message"`
}

func (s *SimpError) Error() string {
	return s.Message
}

// Códigos de erro comuns
const (
	ErrorNotFound = iota
	ErrorAlreadyExists
	ErrorMemoryLimit
	ErrorJobLimit
)
