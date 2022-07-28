package valueobject

import "github.com/go-ozzo/ozzo-validation/v4/is"

type ID string

func NewID(value string) ID { return ID(value) }

func (o ID) String() string { return string(o) }

func (o ID) Validate() error { return is.UUIDv4.Validate(o) }
