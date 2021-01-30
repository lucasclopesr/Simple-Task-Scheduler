package simperr

// SimpError é um erro do SIMP
type SimpError struct {
	Code    int    `yaml:"code"  json:"code"`
	Message string `yaml:"message"  json:"message"`
}

// ErrorBuilder cria a estrutura de um SimpError
type ErrorBuilder struct {
	err *SimpError
}

// Error é o método necessário para identificar um tipo erro em Go
func (err *SimpError) Error() string {
	return err.Message
}

// NewError inicializa um novo SimpError
func NewError() *ErrorBuilder {
	return &ErrorBuilder{
		err: &SimpError{},
	}
}

// Message define a mensagem a ser exibida pelo SimpError
func (b *ErrorBuilder) Message(msg string) *ErrorBuilder {
	b.err.Message = msg
	return b
}

// Code define o código do SimpError
func (b *ErrorBuilder) Code(code int) *ErrorBuilder {
	b.err.Code = code
	return b
}

// Build retorna o SimpError criado
func (b *ErrorBuilder) Build() *SimpError {
	return b.err
}

// NotFound define o tipo de erro a ser usado quando uma estrutura não foi encontrada
func (b *ErrorBuilder) NotFound() *ErrorBuilder {
	b.err.Code = ErrorNotFound
	return b
}

// BadRequest define o tipo de erro a ser usado quando é feita uma requisição inválida
func (b *ErrorBuilder) BadRequest() *ErrorBuilder {
	b.err.Code = ErrorBadRequest
	return b
}

// Códigos de erro comuns
const (
	ErrorNotFound = iota
	ErrorAlreadyExists
	ErrorMemoryLimit
	ErrorJobLimit
	ErrorBadRequest
)
