package premitive

type (
	Slug        string
	Title       string
	Description string
	LongBody    string
	Author      string
)

func NewSlug(s string) (Slug, error) {
	return Slug(s), nil
}
