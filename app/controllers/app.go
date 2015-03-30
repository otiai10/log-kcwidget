package controllers

import (
  "github.com/revel/revel"
  "github.com/otiai10/log-kcwidget/app/models"
  "time"
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

func calcVelocity(reports []ocrReport.OcrReport) int {
  rowCount := len(reports)
  head := reports[0:1][0]
  tail := reports[len(reports)-1:][0]
  minRange := (head.CreatedTime - tail.CreatedTime)/int(time.Minute/time.Millisecond)
  return rowCount/minRange
}

func (c Ocr) Index() revel.Result {
  countPerPage := 40
  reports := ocrReport.Page(0, countPerPage)
  count := ocrReport.Count()
  _s, _f := 0, 0
  for _,r := range reports {
    if r.Result {
      _s++
      continue
    }
    _f++
  }
  latestSuccessRate := _s*100 /(_s + _f)
  latestVelocity := calcVelocity(reports)
  return c.Render(reports, count, countPerPage, latestSuccessRate, latestVelocity)
}

func (c Ocr) Page(page int) revel.Result {
  countPerPage := 40
  if page < 0 {
    page = 0
  }
  before := page - 1
  after  := page + 1
  if before < 0 {
    before = 0
  }
  reports := ocrReport.Page(page, countPerPage)
  count   := len(reports)
  return c.Render(page, reports, count, before, after, countPerPage)
}

func (c Ocr) Show(id int) revel.Result {
  // id って書くけどtargetCreatedTimeなんだよね
  report := ocrReport.Get(id)
  //if len(report) == 0 {
  //  return c.Redirect(routes.Ocr.NotFound())
  //}
  return c.Render(report)
}

func (c Ocr) NotFound() revel.Result {
  return c.Render()
}

func (c Ocr) Upload(
    submit string,
    imgURI string,
    createdTime int,
    userAgent string,
    rawText string,
    assuredText string,
    result bool,
    extVer string,
    ocrVer string) revel.Result {

  added := ocrReport.Add(imgURI, createdTime, userAgent, rawText, assuredText, result, extVer, ocrVer)
  return c.RenderJson(added)
}

func (c Ocr) Delete(target int) revel.Result {
  deleted := ocrReport.Delete(target)
  return c.RenderJson(deleted)
}

func (c Ocr) Summary(datehour int) revel.Result {
  var summary []ocrReport.OcrSummary
  if datehour == 0 {
    summary = ocrReport.FindAllSummary()
  } else {
    summary = ocrReport.FindSummary(datehour)
  }
  return c.RenderJson(summary)
}
