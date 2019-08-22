package error

type merr struct {
	code int
	msg string
}

func New(code int, msg string) *merr {
	return &merr{
		code:code,
		msg:msg,
	}
}

func (m *merr)Error() string{
	return m.msg
}