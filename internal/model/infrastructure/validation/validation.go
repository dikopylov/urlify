package validation

import (
	"net/url"
)

type Validator struct {
	rules []Rule
}

type Rule interface {
	Validate(link string) error
}

func (v *Validator) isEmptyRules() bool {
	return len(v.rules) == 0
}

func (v *Validator) SetRules(rules []Rule) *Validator {
	v.rules = make([]Rule, len(rules), cap(rules))
	v.rules = rules

	return v
}

func (v *Validator) Validate(link string) error {
	for _, rule := range v.rules {
		err := rule.Validate(link)

		if err != nil {
			return err
		}
	}

	return nil
}

type LinkIsCorrect struct {
}

func (l *LinkIsCorrect) Validate(link string) error {
	_, err := url.ParseRequestURI(link)

	return err
}
