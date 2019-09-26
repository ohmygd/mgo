package merror

type Merr struct {
	code int
	msg  string
}

func New(code int) *Merr {
	return &Merr{
		code: code,
	}
}

func NewWM(code int, msg string) *Merr {
	return &Merr{
		code: code,
		msg:msg,
	}
}

func (m *Merr) GetCode() int {
	return m.code
}

func (m *Merr) GetMsg() string {
	return m.msg
}

func (m *Merr) Error() string {
	return m.msg
}
