package stageleft

import (
	"errors"
)

// decorates an error with an exit code
func AttachExitCode(err error, ec ExitCode) ExitCodeError {
	return ExitCodeError{err, ec}
}

type ExitCodeError struct {
	Err  error
	Code ExitCode
}

func (ece ExitCodeError) Unwrap() error {
	return ece.Err
}

func (ece ExitCodeError) Error() string {
	return ece.Err.Error()
}

// extract an exit code or use this one
func AsExitCode(r interface{}, exit ExitCode) ExitCode {
	if r == nil {
		return exit
	}

	switch r := r.(type) {
	case ExitCode:
		exit = r
		break
	case *ExitCode:
		exit = *r
		break
	case ExitCodeError:
		exit = r.Code
		break
	case error:
		exit = ExitCodeFromErr(r, exit)
		break
	case uint8:
		exit = ExitCode(r)
		break
	}

	return exit
}

// attempt to pull an exit code from an error, otherwise use the provided
func ExitCodeFromErr(err error, code ExitCode) ExitCode {
	var ec ExitCodeError

	if errors.As(err, &ec) {
		code = ec.Code
	}

	return code
}

// unix exit codes
type ExitCode uint8

const (
	// build your own like this
	ExitSuccess ExitCode = 0
)
