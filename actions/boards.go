package actions

import "github.com/gobuffalo/buffalo"

// BoardsNew default implementation.
func BoardsNew(c buffalo.Context) error {
	return c.Render(200, r.HTML("boards/new.html"))
}

