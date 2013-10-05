package controllers

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo"
  //"otiai10/logServer/app/routes"
)

type Application struct {
  *revel.Controller
}
type Ocr struct {
  *revel.Controller
}
type Test struct {
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

  session, err := mgo.Dial("localhost")
  if err != nil {
          panic(err)
  }
  defer session.Close()
  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)
  con := session.DB("kcwidget").C("logOcr")
  reports := []OcrReport{}
  err = con.Find(nil).All(&reports)
  if err != nil {
          panic(err)
  }

  return c.Render(reports)
}

func (c Ocr) Index() revel.Result {
  return c.Render()
}

func (c Ocr) Show(page int) revel.Result {
  return c.Render(page)
}

func (c Test) Index() revel.Result {
  return c.Render()
}

func (c Test) Upload(
    submit string,
    imgURI string,
    createdTime int,
    userAgent string,
    rawText string,
    assuredText string,
    result bool) revel.Result {

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  report := &OcrReport{
    ImgURI:      imgURI,
    CreatedTime: createdTime,
    UserAgent:   userAgent,
    RawText:     rawText,
    AssuredText: assuredText,
    Result:      result,
  }

  // TODO: more validation
  if result {
    return c.RenderJson(report)
  }
  session.SetMode(mgo.Monotonic, true)
  con := session.DB("kcwidget").C("logOcr")
  err = con.Insert(report)
  if err != nil {
    panic(err)
  }

  return c.RenderJson(report)
}
