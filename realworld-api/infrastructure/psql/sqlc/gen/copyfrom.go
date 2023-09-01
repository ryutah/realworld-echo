// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: copyfrom.go

package gen

import (
	"context"
)

// iteratorForCreateArticleTag implements pgx.CopyFromSource.
type iteratorForCreateArticleTag struct {
	rows                 []CreateArticleTagParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateArticleTag) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateArticleTag) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ArticleSlug,
		r.rows[0].TagName,
		r.rows[0].CreatedAt,
		r.rows[0].UpdatedAt,
	}, nil
}

func (r iteratorForCreateArticleTag) Err() error {
	return nil
}

func (q *Queries) CreateArticleTag(ctx context.Context, arg []CreateArticleTagParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"article_tag"}, []string{"article_slug", "tag_name", "created_at", "updated_at"}, &iteratorForCreateArticleTag{rows: arg})
}
