package usecase_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ryutah/realworld-echo/realworld-api/domain/repository"
	. "github.com/ryutah/realworld-echo/realworld-api/usecase"
)

func TestArticle_Get(t *testing.T) {
	type fields struct {
		errorHandler ErrorHandler
		outputPort   struct{ get GetArticleOutputPort }
		repository   struct{ article repository.Article }
	}
	type args struct {
		ctx     context.Context
		slugStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skip("not implemented yet")
			_ = gomock.NewController(t)

			a := NewArticle(tt.fields.outputPort.get, tt.fields.errorHandler)

			if err := a.Get(tt.args.ctx, tt.args.slugStr); (err != nil) != tt.wantErr {
				t.Errorf("Article.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
