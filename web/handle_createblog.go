package web

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
)

type createBlogRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s Server) createBlog(c echo.Context) error {
	var req createBlogRequest
	if err := c.Bind(&req); err != nil {
		return s.badRequest(c, err)
	}

	_, err := s.mainDB.GetTenant(db.GetTenantRequest{ByName: req.Name})
	if err == nil {
		return s.badRequest(c, errors.New("blog already exists"))
	}

	if err != sql.ErrNoRows {
		return s.internalServerError(c, err)
	}

	createRequest := db.CreateTenantRequest{Name: req.Name, Password: req.Password}
	if err := s.mainDB.CreateTenant(createRequest); err != nil {
		return s.internalServerError(c, err)
	}

	if err := s.tenantDB.CreateTenantDB(createRequest); err != nil {
		if err := s.mainDB.DeleteTenant(db.DeleteTenantRequest{ByName: req.Name}); err != nil {
			return s.internalServerError(c, err)
		}

		return s.internalServerError(c, err)
	}

	return c.JSON(http.StatusCreated, nil)
}
