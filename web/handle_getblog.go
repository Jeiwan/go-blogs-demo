package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
)

type getBlogResponse struct {
	Posts []getBlogPost `json:"posts"`
}

type getBlogPost struct {
	Date time.Time `json:"date"`
	Post string    `json:"post"`
}

func (s Server) getBlog(c echo.Context) error {
	tenantDB, ok := c.Get("tenantdb").(DB)
	if !ok {
		return s.internalServerError(c, errors.New("tenant db is not set"))
	}

	posts, err := tenantDB.GetPosts(db.GetPostsRequest{})
	if err != nil {
		return s.internalServerError(c, err)
	}

	response := getBlogResponse{Posts: []getBlogPost{}}
	for _, post := range posts {
		response.Posts = append(response.Posts, getBlogPost{
			Date: post.Date,
			Post: post.Post,
		})
	}

	return c.JSON(http.StatusOK, response)
}
