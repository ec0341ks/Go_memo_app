package actions

import (
	"fmt"

	"coke/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	// "github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	// if !ok {
	// 	return errors.WithStack(errors.New("no transaction found"))
	// }
	boards := []models.Board{}
	err := tx.All(&boards)
	fmt.Println(boards)
	fmt.Println("here!!!!!!!!!!!!!!!!")
	fmt.Println(err)
	c.Set("boards", boards)
	fmt.Println(boards)

	// for i := 1, i<=boards.length; i++ {
	// 	fmt.Println(i)
	// }

	// err := board.All(&boards)
	return c.Render(200, r.HTML("boards/index.html"))
}
