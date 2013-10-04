package controllers

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo"
  //"otiai10/logServer/app/routes"
  "fmt"
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

type Person struct {
  Name  string
  Phone string
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

  fmt.Printf("submit      => %v\n", submit)
  fmt.Printf("imgURI      => %v\n", imgURI)
  fmt.Printf("createdTime => %v\n", createdTime)
  fmt.Printf("userAgent   => %v\n", userAgent)
  fmt.Printf("rawText     => %v\n", rawText)
  fmt.Printf("assuredText => %v\n", assuredText)
  fmt.Printf("result      => %v\n", result)

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  // Optional. Switch the session to a monotonic behavior.

  report := &OcrReport{
    ImgURI:      imgURI,
    CreatedTime: createdTime,
    UserAgent:   userAgent,
    RawText:     rawText,
    AssuredText: assuredText,
    Result:      result,
  }

  //if result {
  //  return c.RenderJson(report)
  //}
  session.SetMode(mgo.Monotonic, true)
  con := session.DB("kcwidget").C("logOcr")
  err = con.Insert(report)
  if err != nil {
    panic(err)
  }

  return c.RenderJson(report)
}
