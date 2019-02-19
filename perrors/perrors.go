package perrors

type Error struct {
	msg    string
	status int
}

func NewError(msg string, status int) error {
	return &Error{
		msg:    msg,
		status: status,
	}
}

func (e Error) Error() string { return e.msg }

func (e Error) Status() int { return e.status }
