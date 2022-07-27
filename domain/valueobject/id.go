package valueobject

import (
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type ID string

func NewID() ID { return ID(uuid.NewString()) }

func (o ID) String() string { return string(o) }

func (o ID) Validate() error { return is.UUIDv4.Validate(o) }
