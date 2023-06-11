package premitive

import (
	"fmt"
	"strings"
	"sync"

	"github.com/cockroachdb/errors"
	"github.com/go-playground/validator/v10"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
)

type (
	UID      string
	Slug     string
	Title    string
	Name     string
	LongText string
	Email    string
	URL      string
)

func NewUID(s string) (UID, error) {
	return withValidate(
		s,
		func() UID { return UID(s) },
	)
}

func (u UID) String() string {
	return string(u)
}

func NewSlug(s string) (Slug, error) {
	return withValidate(
		s,
		func() Slug { return Slug(s) },
	)
}

func (s Slug) String() string {
	return string(s)
}

func NewTitle(s string) (Title, error) {
	return withValidate(
		s,
		func() Title { return Title(s) },
	)
}

func (t Title) String() string {
	return string(t)
}

func NewName(s string) (Name, error) {
	return withValidate(
		s,
		func() Name { return Name(s) },
	)
}

func (n Name) String() string {
	return string(n)
}

func NewLongText(s string) (LongText, error) {
	return withValidate(
		s,
		func() LongText { return LongText(s) },
	)
}

func (l LongText) String() string {
	return string(l)
}

func NewEmail(s string) (Email, error) {
	return withValidate(
		s,
		func() Email { return Email(s) },
		email(),
		max(254),
	)
}

func (e Email) String() string {
	return string(e)
}

func NewURL(s string) (URL, error) {
	return withValidate(
		s,
		func() URL { return URL(s) },
		url(),
	)
}

func (u URL) String() string {
	return string(u)
}

func email() string {
	return "email"
}

func url() string {
	return "url"
}

func max(m int) string {
	return fmt.Sprintf("max=%d", m)
}

func withValidate[Arg, Ret any](value Arg, genFunc func() Ret, rules ...string) (r Ret, _ error) {
	if err := getValidate().Var(value, strings.Join(rules, ",")); err != nil {
		return r, newValidationError(1, err)
	}
	return genFunc(), nil
}

func newValidationError(depth int, ve error) error {
	err := errors.WrapWithDepth(
		depth+1, derrors.Errors.Validation.Err, derrors.Errors.Validation.Message,
	)

	verrs, ok := ve.(validator.ValidationErrors)
	if !ok {
		return errors.WithDetail(err, ve.Error())
	}

	for _, ve := range verrs {
		err = errors.WithDetail(err, ve.Error())
	}
	return err
}

var (
	newValidateOnce sync.Once
	validate        *validator.Validate
)

func getValidate() *validator.Validate {
	newValidateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}
