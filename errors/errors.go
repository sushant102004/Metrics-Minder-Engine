// cErrors -> Custom Errors

package cErrors


type CustomError struct {
  Msg string
}

func (e *CustomError) Error() string {
  return e.Msg
}

func NewError(msg string) *CustomError {
  return &CustomError{
    Msg: msg,
  }
}
