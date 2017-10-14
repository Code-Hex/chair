package chair

type causer interface {
	Cause() error
}

type exiter interface {
	ExitCode() int
}

type skipErr struct{}

func (skipErr) Error() string { return "skip" }

func makeSkipError() error { return skipErr{} }

// UnwrapErrors get important message from wrapped error message
func UnwrapErrors(err error) (int, error) {
	for e := err; e != nil; {
		switch e.(type) {
		case exiter:
			return e.(exiter).ExitCode(), e
		case causer:
			e = e.(causer).Cause()
		default:
			return 1, e // default error
		}
	}
	return 0, nil
}
