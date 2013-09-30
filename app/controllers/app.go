package controllers

import (
  "github.com/robfig/revel"
)

type Application struct {
  *revel.Controller
}

func (c Application) Index() revel.Result {
  return c.Render()
  //return c.Redirect(routes.Application.Log)
}

func (c Application) Log(page int) revel.Result {
  return c.Render(page)
}
