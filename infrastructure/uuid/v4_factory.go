package uuid

import "github.com/google/uuid"

type V4Factory struct{}

func NewV4Factory() *V4Factory { return new(V4Factory) }

func (f *V4Factory) NewUUID() string { return uuid.NewString() }
