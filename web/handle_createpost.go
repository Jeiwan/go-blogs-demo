package web

import (
	"errors"
	"net/http"
	"time"

	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
)

type createPostRequest struct {
	Post string `json:"post"`
}

type createPostResponse struct {
	PostID int `json:"post_id"`
}

func (s Server) createPost(c echo.Context) error {
	var req createPostRequest
	if err := c.Bind(&req); err != nil {
		return s.badRequest(c, err)
	}

	tenantDB, ok := c.Get("tenantdb").(DB)
	if !ok {
		return s.internalServerError(c, errors.New("tenant db is not set"))
	}

	postID, err := tenantDB.CreatePost(db.CreatePostRequest{Date: time.Now(), Post: req.Post})
	if err != nil {
		return s.internalServerError(c, err)
	}

	resp := createPostResponse{PostID: postID}

	return c.JSON(http.StatusOK, resp)
}
