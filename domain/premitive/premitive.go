package premitive

type (
	Slug        string
	Title       string
	Name        string
	Description string
	LongBody    string
	Author      string
	Email       string
	URL         string
)

func NewSlug(s string) (Slug, error) {
	return Slug(s), nil
}
