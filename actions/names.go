package actions

import "github.com/gobuffalo/buffalo"

// NamesDelete default implementation.
func NamesDelete(c buffalo.Context) error {
	return c.Render(200, r.HTML("names/delete.html"))
}

