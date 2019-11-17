package actions

import "github.com/gobuffalo/buffalo"

// BoardsShow default implementation.
func BoardsShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("boards/show.html"))
}

