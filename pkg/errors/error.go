package errs

type Err struct {
	info string
	err  error
}

func (e *Err) Error() string {
	if e == nil {
		return "nil err"
	}

	res := ""

	if e.info != "" {
		res += e.info + ":"
	}

	if e.err != nil {
		res += e.err.Error()
	}
	return res
}

func (e *Err) UnWrap() error {
	if e == nil {
		return nil
	}
	return e.err
}

func Wrap(info string, err error) error {
	return &Err{
		info: info,
		err:  err,
	}
}

func UnWrap(err error) error {
	ue, ok := err.(interface{ Unwrap() error })
	if ok {
		return UnWrap(ue.Unwrap())
	}

	return err
}
