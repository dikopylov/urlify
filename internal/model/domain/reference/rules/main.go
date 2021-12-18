package rules

import (
	"errors"
)

type LinkIsNotEmpty struct {
}

func (l *LinkIsNotEmpty) Validate(link string) error {
	if link == "" {
		return errors.New("input link is empty")
	}

	return nil
}
