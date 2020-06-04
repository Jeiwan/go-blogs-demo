package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/loads"
	smiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
	"github.com/toqueteos/webbrowser"
)

func newUICmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ui",
		RunE: func(cmd *cobra.Command, args []string) error {
			specDoc, err := loads.Spec("./blogs.yaml")
			if err != nil {
				return err
			}

			b, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
			if err != nil {
				return err
			}

			e := echo.New()
			e.Use(middleware.Logger(), middleware.CORS())

			// handler := smiddleware.Redoc(smiddleware.RedocOpts{
			// 	BasePath: "/",
			// 	SpecURL:  "/swagger.json",
			// 	Path:     "docs",
			// }, http.NotFoundHandler())
			handler := smiddleware.Spec("/", b, http.NotFoundHandler())

			e.GET("/swagger.json", echo.WrapHandler(handler))

			errCh := make(chan error)
			go func() {
				errCh <- e.Start(":8081")
			}()

			_ = webbrowser.Open(fmt.Sprintf("http://petstore.swagger.io/?url=http://localhost:8081/swagger.json"))

			return <-errCh
		},
	}

	return cmd
}
