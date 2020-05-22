package common

type ErrorCode uint8

const (
	Unknown ErrorCode = iota
	AlreadyExists
	NotFound
	Internal
)

type DaoErr struct {
	Code ErrorCode
	Err  error
}

func (d *DaoErr) Error() string {
	return d.Err.Error()
}

func NewDaoErr(code ErrorCode, err error) *DaoErr {
	return &DaoErr{code, err}
}
