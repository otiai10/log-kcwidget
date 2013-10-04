package controllers

import (
  "github.com/robfig/revel"
)

type Application struct {
  *revel.Controller
}

type Ocr struct {
  *revel.Controller
}

func (c Application) Index() revel.Result {
  return c.Render()
}

func (c Ocr) Index() revel.Result {
  return c.Render()
}

func (c Ocr) Show(page int) revel.Result {
  return c.Render(page)
}
