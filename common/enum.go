package common

type ErrorCode uint8

const (
	Unknown ErrorCode = iota
	AlreadyExists
	NotFound
	Internal
	InvalidArgument
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

func FromError(err error) (*DaoErr, bool) {
	if de, ok := err.(*DaoErr); ok {
		return de, true
	}
	return NewDaoErr(Unknown, err), false
}
