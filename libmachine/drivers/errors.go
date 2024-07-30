package drivers

// ErrResourcesCreated indicates that some external resources, such as VM, have been created when the error is returned
type ErrResourcesCreated string

func (e ErrResourcesCreated) Error() string {
	return string(e)
}

// ErrResourcesCreated indicates that some external resources, such as VM, have been created when the error is returned
// type ErrResourcesCreated struct {
// 	Err error
// }
//
// func (e *ErrResourcesCreated) Error() string {
// 	return fmt.Sprintf("(ErrResourcesCreated) %s", e.Err)
// }
//
// func (e *ErrResourcesCreated) Unwrap() error {
// 	return e.Err
// }
//
// func ErrResourcesCreatedf(format string, a ...any) ErrResourcesCreated {
// 	return ErrResourcesCreated(fmt.Sprintf(format, a...))
// }

// func (e *ErrResourcesCreated) Is(err error) bool {
// 	if _, ok := err.(*ErrResourcesCreated); ok {
// 		return true
// 	}
// 	return false
// }

// func IsErrResourcesCreated(err error) bool {
// 	var errWaiting ErrResourcesCreated
// 	return errors.As(err, &errWaiting)
// }
