package thread

import (
	"forums/internal/models"
	"github.com/labstack/echo/v4"
)

type ThreadHandler interface {
	AddPosts(c echo.Context) error
}

type ThreadUsecase interface {
	AddPosts(posts models.Posts, slugOrId string) (models.Posts, error)
}

type ThreadRepo interface {
	AddPostsInThreadByID(posts models.Posts, id int32) (models.Posts, error)
	AddPostsInThreadBySlug(posts models.Posts, slug string) (models.Posts, error)
	GetThreadByID(id int32) (models.Thread, error)
	GetThreadBySlug(slug string) (models.Thread, error)
	GetPostByID(id int64) (models.Post, error)
}
