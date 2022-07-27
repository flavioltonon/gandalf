package presenter

type Error struct {
	Message string `json:"message"`
}

func NewError(cause error) *Error {
	return &Error{
		Message: cause.Error(),
	}
}
