package controllers

import (
	"github.com/revel/revel"
)

// App struct for the app
type App struct {
	*revel.Controller
}

// Index returns the home page
func (c App) Index() revel.Result {
	greeting := "hi there"
	return c.Render(greeting)
}

// Hello returns the hello {your name} page
func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}
