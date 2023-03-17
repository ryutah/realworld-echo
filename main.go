package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/adapter/rest/gen"
)

func main() {
	e := echo.New()
	gen.RegisterHandlers(e, &server{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatal(err)
	}
}

type server struct {
	gen.ServerInterface
}

// Get recent articles globally
// (GET /articles)
func (s *server) GetArticles(ctx echo.Context, params gen.GetArticlesParams) error {
	return ctx.JSON(http.StatusOK, gen.SingleArticleResponse{
		Article: gen.Article{
			Author: gen.Profile{
				Bio:       "",
				Following: false,
				Image:     "",
				Username:  "",
			},
			Body:           "",
			CreatedAt:      time.Now(),
			Description:    "",
			Favorited:      false,
			FavoritesCount: 0,
			Slug:           "",
			TagList:        []string{},
			Title:          "",
			UpdatedAt:      time.Now(),
		},
	})
}

// Create an article
// (POST /articles)
func (s *server) CreateArticle(ctx echo.Context) error {
	var req gen.NewArticleRequest
	if err := ctx.Bind(req); err != nil {
		return err
	}
	panic("not implemented") // TODO: Implement
}

// Get recent articles from users you follow
// (GET /articles/feed)
func (s *server) GetArticlesFeed(ctx echo.Context, params gen.GetArticlesFeedParams) error {
	panic("not implemented") // TODO: Implement
}

// Delete an article
// (DELETE /articles/{slug})
func (s *server) DeleteArticle(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Get an article
// (GET /articles/{slug})
func (s *server) GetArticle(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Update an article
// (PUT /articles/{slug})
func (s *server) UpdateArticle(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Get comments for an article
// (GET /articles/{slug}/comments)
func (s *server) GetArticleComments(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Create a comment for an article
// (POST /articles/{slug}/comments)
func (s *server) CreateArticleComment(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Delete a comment for an article
// (DELETE /articles/{slug}/comments/{id})
func (s *server) DeleteArticleComment(ctx echo.Context, slug string, id int) error {
	panic("not implemented") // TODO: Implement
}

// Unfavorite an article
// (DELETE /articles/{slug}/favorite)
func (s *server) DeleteArticleFavorite(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Favorite an article
// (POST /articles/{slug}/favorite)
func (s *server) CreateArticleFavorite(ctx echo.Context, slug string) error {
	panic("not implemented") // TODO: Implement
}

// Get a profile
// (GET /profiles/{username})
func (s *server) GetProfileByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Unfollow a user
// (DELETE /profiles/{username}/follow)
func (s *server) UnfollowUserByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Follow a user
// (POST /profiles/{username}/follow)
func (s *server) FollowUserByUsername(ctx echo.Context, username string) error {
	panic("not implemented") // TODO: Implement
}

// Get tags
// (GET /tags)
func (s *server) GetTags(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Get current user
// (GET /user)
func (s *server) GetCurrentUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Update current user
// (PUT /user)
func (s *server) UpdateCurrentUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// (POST /users)
func (s *server) CreateUser(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}

// Existing user login
// (POST /users/login)
func (s *server) Login(ctx echo.Context) error {
	panic("not implemented") // TODO: Implement
}
