package controllers

import (
  "github.com/robfig/revel"
  "otiai10/logServer/app/models"
  //"otiai10/logServer/app/routes"
  //"fmt"
)

type Application struct {
  *revel.Controller
}
type Ocr struct {
  *revel.Controller
}

type OcrReport struct {
  ImgURI      string
  CreatedTime int
  UserAgent   string
  RawText     string
  AssuredText string
  Result      bool
}

func (c Application) Index() revel.Result {
  return c.Render()
}

func (c Ocr) Index() revel.Result {
  reports := ocrReport.Page(0)
  count := ocrReport.Count()
  return c.Render(reports, count)
}

func (c Ocr) Show(page int) revel.Result {
  if page < 0 {
    page = 0
  }
  before := page - 1
  after  := page + 1
  if before < 0 {
    before = 0
  }
  reports := ocrReport.Page(page)
  count   := len(reports)
  return c.Render(page, reports, count, before, after)
}

func (c Ocr) Upload(
    submit string,
    imgURI string,
    createdTime int,
    userAgent string,
    rawText string,
    assuredText string,
    result bool) revel.Result {

  added := ocrReport.Add(imgURI, createdTime, userAgent, rawText, assuredText, result)

  return c.RenderJson(added)
}

func (c Ocr) Delete(target int) revel.Result {
  deleted := ocrReport.Delete(target)
  return c.RenderJson(deleted)
}

