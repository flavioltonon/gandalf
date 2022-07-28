package valueobject

import (
	"regexp"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type Username string

func NewUsername(value string) Username { return Username(value) }

func (o Username) String() string { return string(o) }

func (o Username) Validate() error {
	rules := []ozzo.Rule{
		ozzo.Length(8, 32),
		ozzo.Match(regexp.MustCompile("(?i)[a-z0-9 ]")),
	}

	for _, rule := range rules {
		if err := rule.Validate(o); err != nil {
			return err
		}
	}

	return nil
}
