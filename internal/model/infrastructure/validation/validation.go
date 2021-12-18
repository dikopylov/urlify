package validation

import (
	"net/url"
	"sync"
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

func (v *Validator) generateRule() <-chan Rule {
	out := make(chan Rule, len(v.rules))

	go func() {
		for _, rule := range v.rules {
			out <- rule
		}
		close(out)
	}()

	return out
}

func (v *Validator) runRules(done <-chan struct{}, link string, rules []<-chan Rule) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	output := func(ruleCh <-chan Rule) {
		for rule := range ruleCh {
			select {
			case out <- rule.Validate(link):
			case <-done:
			}
		}
		wg.Done()
	}

	wg.Add(len(rules))
	for _, ruleCh := range rules {
		go output(ruleCh)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func (v *Validator) Validate(link string) error {
	done := make(chan struct{}, 1)
	rules := make([]<-chan Rule, len(v.rules))

	defer close(done)

	for i := 0; i < len(v.rules); i++ {
		rules = append(rules, v.generateRule())
	}

	for err := range v.runRules(done, link, rules) {
		if err != nil {
			for i := 0; i < len(v.rules); i++ {
				done <- struct{}{}
			}

			return err
		}
	}

	return nil
}

type LinkIsCorrect struct {
}

func (l *LinkIsCorrect) Validate(link string) error {
	_, err := url.Parse(link)

	return err
}
