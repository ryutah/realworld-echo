package rest

import (
	"context"

	"github.com/ryutah/realworld-echo/realworld-api/api/rest/gen"
)

type Server struct {
	*Article
}

var _ gen.StrictServerInterface = (*Server)(nil)

func NewServer(article *Article) *Server {
	return &Server{
		Article: article,
	}
}

// Create an article
// (POST /articles)
func (s *Server) CreateArticle(ctx context.Context, request gen.CreateArticleRequestObject) (gen.CreateArticleResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Get recent articles from users you follow
// (GET /articles/feed)
func (s *Server) GetArticlesFeed(ctx context.Context, request gen.GetArticlesFeedRequestObject) (gen.GetArticlesFeedResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Delete an article
// (DELETE /articles/{slug})
func (s *Server) DeleteArticle(ctx context.Context, request gen.DeleteArticleRequestObject) (gen.DeleteArticleResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Update an article
// (PUT /articles/{slug})
func (s *Server) UpdateArticle(ctx context.Context, request gen.UpdateArticleRequestObject) (gen.UpdateArticleResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Get comments for an article
// (GET /articles/{slug}/comments)
func (s *Server) GetArticleComments(ctx context.Context, request gen.GetArticleCommentsRequestObject) (gen.GetArticleCommentsResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Create a comment for an article
// (POST /articles/{slug}/comments)
func (s *Server) CreateArticleComment(ctx context.Context, request gen.CreateArticleCommentRequestObject) (gen.CreateArticleCommentResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Delete a comment for an article
// (DELETE /articles/{slug}/comments/{id})
func (s *Server) DeleteArticleComment(ctx context.Context, request gen.DeleteArticleCommentRequestObject) (gen.DeleteArticleCommentResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Unfavorite an article
// (DELETE /articles/{slug}/favorite)
func (s *Server) DeleteArticleFavorite(ctx context.Context, request gen.DeleteArticleFavoriteRequestObject) (gen.DeleteArticleFavoriteResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Favorite an article
// (POST /articles/{slug}/favorite)
func (s *Server) CreateArticleFavorite(ctx context.Context, request gen.CreateArticleFavoriteRequestObject) (gen.CreateArticleFavoriteResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Get a profile
// (GET /profiles/{username})
func (s *Server) GetProfileByUsername(ctx context.Context, request gen.GetProfileByUsernameRequestObject) (gen.GetProfileByUsernameResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Unfollow a user
// (DELETE /profiles/{username}/follow)
func (s *Server) UnfollowUserByUsername(ctx context.Context, request gen.UnfollowUserByUsernameRequestObject) (gen.UnfollowUserByUsernameResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Follow a user
// (POST /profiles/{username}/follow)
func (s *Server) FollowUserByUsername(ctx context.Context, request gen.FollowUserByUsernameRequestObject) (gen.FollowUserByUsernameResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Get tags
// (GET /tags)
func (s *Server) GetTags(ctx context.Context, request gen.GetTagsRequestObject) (gen.GetTagsResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Get current user
// (GET /user)
func (s *Server) GetCurrentUser(ctx context.Context, request gen.GetCurrentUserRequestObject) (gen.GetCurrentUserResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Update current user
// (PUT /user)
func (s *Server) UpdateCurrentUser(ctx context.Context, request gen.UpdateCurrentUserRequestObject) (gen.UpdateCurrentUserResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// (POST /users)
func (s *Server) CreateUser(ctx context.Context, request gen.CreateUserRequestObject) (gen.CreateUserResponseObject, error) {
	panic("not implemented") // TODO: Implement
}

// Existing user login
// (POST /users/login)
func (s *Server) Login(ctx context.Context, request gen.LoginRequestObject) (gen.LoginResponseObject, error) {
	panic("not implemented") // TODO: Implement
}
