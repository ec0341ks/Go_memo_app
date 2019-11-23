package actions

import (
	"fmt"
	"coke/models"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	// "database/sql"
)

// BoardsNew default implementation.
func BoardsNew(c buffalo.Context) error {
	board := &models.Board{}
	c.Set("board", board)
	return c.Render(200, r.HTML("boards/new.html"))
}

// func BoardShow(){
// 	db := sql.Open("mysql")
// 	log.Println("Connected to mysql.")
// }

func BoardsCreate(c buffalo.Context) error {
	// Allocate an empty Widget
	board := &models.Board{}
	// Bind board to the html form elements
	if err := c.Bind(board); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(board)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		// return c.Render(200, r.HTML("/boards/index.html"))
		// return c.Render(422, r.Auto(c, board))
		return c.Redirect(307, "/")
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "新しい掲示板が作成されました")
	// and redirect to the widgets index page
	// return c.Render(201, r.Auto(c, board))
	return c.Redirect(303, "/")
}

// BoardsDelete default implementation.
func BoardsDelete(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)
	fmt.Println("ここだよ")
	fmt.Println(c.Params())
	id := c.Params().Get("id")
	fmt.Println(id)
	err := tx.RawQuery("DELETE FROM boards WHERE id = ?", id).Exec()
	fmt.Println(err)
	// fmt.Println(boards[0])
	return c.Redirect(303, "/")
	// return c.Render(200, r.HTML("boards/delete.html"))
}

