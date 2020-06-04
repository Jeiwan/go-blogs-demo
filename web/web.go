package web

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Jeiwan/goblogs/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type DB interface {
	CreatePost(db.CreatePostRequest) (int, error)
	GetPosts(db.GetPostsRequest) ([]db.Post, error)

	CreateTenant(db.CreateTenantRequest) error
	CreateTenantDB(db.CreateTenantRequest) error
	GetTenant(db.GetTenantRequest) (*db.Tenant, error)

	DeleteTenant(db.DeleteTenantRequest) error
}

type Server struct {
	commonDatasource *db.Datasource
	dbs              map[string]DB
	mainDB           DB
	tenantDB         DB
}

func New(commonDatasource *db.Datasource, mainDB DB, tenantDB DB) *Server {
	return &Server{
		commonDatasource: commonDatasource,
		dbs:              make(map[string]DB),
		mainDB:           mainDB,
		tenantDB:         tenantDB,
	}
}

func (s Server) Run(address string) error {
	e := echo.New()

	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "${method} ${uri}	${status} ${error} ${latency_human}\n",
		}),
		middleware.Gzip(),
		middleware.Recover(),
		middleware.RemoveTrailingSlash(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			Skipper: func(c echo.Context) bool {
				// TODO: restrict
				return false
			},
		}),
	)
	basicAuth := middleware.BasicAuth(s.authValidator)

	e.Add("POST", "/blog", s.createBlog)
	e.Add("GET", "/blog/:blog_name", s.getBlog, s.middlewareTenantDB)
	e.Add("POST", "/blog/:blog_name/posts", s.createPost, s.middlewareTenantDB, basicAuth)

	printRoutes(e)

	return e.Start(address)
}

func printRoutes(e *echo.Echo) {
	if logrus.GetLevel() != logrus.DebugLevel {
		return
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	for _, route := range e.Routes() {
		fmt.Fprintf(w, "\n %s\t%s\t%s\t", route.Method, route.Path, route.Name)
	}
	fmt.Fprintf(w, "\n\n")
}
