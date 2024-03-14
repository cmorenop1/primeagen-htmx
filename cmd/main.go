package main

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// THE-PRIMEAGEN: This is a pattern that is all over the place on the internet. For Echo + Templ
type Templates struct {
	templates *template.Template
}

type Count struct {
	Count int
}

// === TYPES ===

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()

	count := Count{
		Count: 0,
	}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", count)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		return c.Render(200, "count", count)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
