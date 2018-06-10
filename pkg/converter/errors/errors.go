package errors

import "fmt"

type MErr map[string]error
type AErr []error

func (e MErr) Error() string {
	return fmt.Sprintf("%+v", e)
}

func (a AErr) Error() string {
	return fmt.Sprintf("%+v", a)
}
