package valueobject

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type Password string

func (o Password) String() string { return string(o) }

func (o Password) Validate() error {
	return ozzo.Length(8, 32).Validate(o)
}
