package web

import (
	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
)

func (s Server) authValidator(name, password string, c echo.Context) (bool, error) {
	blogName := c.Param("blog_name")
	if blogName == "" {
		return false, nil
	}

	if blogName != name {
		return false, nil
	}

	tenant, err := s.mainDB.GetTenant(db.GetTenantRequest{ByName: name})
	if err != nil {
		return false, nil
	}

	if tenant.Password != password {
		return false, nil
	}

	return true, nil
}
