package adminecho

import (
	"fmt"
	"github.com/go-advanced-admin/admin"
	"github.com/labstack/echo/v4"
	"mime"
	"net/http"
	"path/filepath"
)

// Integrator struct wraps an Echo group to handle routes and assets for the admin panel
type Integrator struct {
	group *echo.Group
}

// NewIntegrator creates a new Integrator instance for the given Echo group
func NewIntegrator(g *echo.Group) *Integrator {
	return &Integrator{group: g}
}

// HandleRoute adds a route to the Echo group and uses the admin.HandlerFunc for processing requests
func (i *Integrator) HandleRoute(method, path string, handler admin.HandlerFunc) {
	i.group.Add(method, path, func(c echo.Context) error {
		code, body := handler(c)
		if code == http.StatusFound || code == http.StatusMovedPermanently || code == http.StatusSeeOther {
			return c.Redirect(int(code), body)
		}
		return c.HTML(int(code), body)
	})
}

// ServeAssets serves static assets by using the provided admin.TemplateRenderer interface
func (i *Integrator) ServeAssets(prefix string, renderer admin.TemplateRenderer) {
	i.group.GET(fmt.Sprintf("/%s/*", prefix), func(c echo.Context) error {
		fileName := c.Param("*")
		fileData, err := renderer.GetAsset(fileName)
		if err != nil {
			return c.NoContent(http.StatusNotFound)
		}

		contentType := mime.TypeByExtension(filepath.Ext(fileName))
		if contentType == "" {
			contentType = "application/octet-stream"
		}
		return c.Blob(http.StatusOK, contentType, fileData)
	})
}

// GetQueryParam extracts a query parameter from the context
func (i *Integrator) GetQueryParam(ctx interface{}, name string) string {
	ec := ctx.(echo.Context)
	return ec.QueryParam(name)
}

// GetPathParam extracts a path parameter from the context
func (i *Integrator) GetPathParam(ctx interface{}, name string) string {
	ec := ctx.(echo.Context)
	return ec.Param(name)
}

// GetRequestMethod retrieves the HTTP request method from the context
func (i *Integrator) GetRequestMethod(ctx interface{}) string {
	ec := ctx.(echo.Context)
	return ec.Request().Method
}

// GetFormData retrieves form data from the request context
func (i *Integrator) GetFormData(ctx interface{}) map[string][]string {
	ec := ctx.(echo.Context)
	if err := ec.Request().ParseForm(); err != nil {
		return nil
	}
	return ec.Request().Form
}
