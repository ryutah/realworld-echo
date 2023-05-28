package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/api/rest/gen"
)

type Server struct {
	*Article
}

var _ gen.ServerInterface = (*Server)(nil)

func NewServer(article *Article) *Server {
	return &Server{
		Article: article,
	}
}

// Get recent articles globally
// (GET /articles)
func (s *Server) GetArticles(ctx echo.Context, params gen.GetArticlesParams) error {
	panic("not implemented") // TODO: Implement
}

// Create an article
// (POST /articles)
func (s *Server) CreateArticle(ctx echo.Context) error {
	req := gen.NewArticleRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(echo.ErrBadGateway.Code, gen.GenericError{
			Errors: struct {
				Body []string `json:"body"`
			}{
				Body: []string{err.Error()},
			},
		})
	}
	panic("not implemented") // TODO: Implement
}

// Get recent articles from users you follow
// (GET /articles/feed)
func (s *Server) GetArticlesFeed(ctx echo.Context, params gen.GetArticlesFeedParams) error {
	panic("not implemented") // TODO: Implement
}

// Delete an article
// (DELETE /articles/{slug})
func (s *Server) DeleteArticle(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Update an article
// (PUT /articles/{slug})
func (s *Server) UpdateArticle(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Get comments for an article
// (GET /articles/{slug}/comments)
func (s *Server) GetArticleComments(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Create a comment for an article
// (POST /articles/{slug}/comments)
func (s *Server) CreateArticleComment(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Delete a comment for an article
// (DELETE /articles/{slug}/comments/{id})
func (s *Server) DeleteArticleComment(ctx echo.Context, slug string, id int) error {
	panic("not implemented") // TODO: Implement
}

// Unfavorite an article
// (DELETE /articles/{slug}/favorite)
func (s *Server) DeleteArticleFavorite(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Favorite an article
// (POST /articles/{slug}/favorite)
func (s *Server) CreateArticleFavorite(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Get a profile
// (GET /profiles/{username})
func (s *Server) GetProfileByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Unfollow a user
// (DELETE /profiles/{username}/follow)
func (s *Server) UnfollowUserByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Follow a user
// (POST /profiles/{username}/follow)
func (s *Server) FollowUserByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Get tags
// (GET /tags)
func (s *Server) GetTags(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Get current user
// (GET /user)
func (s *Server) GetCurrentUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Update current user
// (PUT /user)
func (s *Server) UpdateCurrentUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// (POST /users)
func (s *Server) CreateUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Existing user login
// (POST /users/login)
func (s *Server) Login(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}
