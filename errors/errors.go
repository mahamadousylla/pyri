package errors

type DetailedError struct {
	msg    string
	status int
}

func NewDetailedError(msg string, status int) error {
	return &DetailedError{
		msg:    msg,
		status: status,
	}
}

func (e DetailedError) Error() string { return e.msg }

func (e DetailedError) Status() int { return e.status }
