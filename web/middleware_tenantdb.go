package web

import (
	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
)

func (s Server) middlewareTenantDB(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		blogName := c.Param("blog_name")
		if blogName == "" {
			return s.notFound(c)
		}

		tenant, err := s.mainDB.GetTenant(db.GetTenantRequest{ByName: blogName})
		if err != nil {
			return s.notFound(c)
		}

		tenantDB, ok := s.dbs[tenant.Name]
		if !ok {
			tenantDatasource := *s.commonDatasource
			tenantDatasource.DBName = tenant.Name

			newDB, err := db.New(&tenantDatasource)
			if err != nil {
				return s.internalServerError(c, err)
			}
			s.dbs[tenant.Name] = newDB
			tenantDB = newDB
		}

		c.Set("tenantdb", tenantDB)

		return next(c)
	}
}
