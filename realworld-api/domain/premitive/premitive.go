package premitive

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	derrors "github.com/ryutah/realworld-echo/realworld-api/domain/errors"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtime"
)

type (
	UID       string
	Slug      string
	Title     string
	Name      string
	LongText  string
	ShortText string
	Email     string
	URL       string
	JSTTime   time.Time
)

func NewUID(s string) (UID, error) {
	return withValidate(
		s,
		func() UID { return UID(s) },
		max(255),
	)
}

func (u UID) String() string {
	return string(u)
}

func NewSlug(s string) (Slug, error) {
	return withValidate(
		s,
		func() Slug { return Slug(s) },
		max(50),
	)
}

func (s Slug) String() string {
	return string(s)
}

func NewTitle(s string) (Title, error) {
	return withValidate(
		s,
		func() Title { return Title(s) },
		max(255),
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
		max(5000),
	)
}

func (l LongText) String() string {
	return string(l)
}

func NewShortText(s string) (ShortText, error) {
	return withValidate(
		s,
		func() ShortText { return ShortText(s) },
		max(255),
	)
}

func (s ShortText) String() string {
	return string(s)
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

func NewJSTTime(t time.Time) JSTTime {
	return JSTTime(xtime.JST(t))
}

func (j JSTTime) Time() time.Time {
	return time.Time(j)
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
		return r, derrors.NewValidationError(1, err)
	}
	return genFunc(), nil
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
