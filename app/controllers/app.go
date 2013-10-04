package controllers

import (
  "github.com/robfig/revel"
  "labix.org/v2/mgo"
)

type Application struct {
  *revel.Controller
}

type Ocr struct {
  *revel.Controller
}

type Person struct {
  Name  string
  Phone string
}

func (c Application) Index() revel.Result {

  //----------------->>>
  session, err := mgo.Dial("localhost")
  if err != nil {
          panic(err)
  }
  defer session.Close()
  // Optional. Switch the session to a monotonic behavior.
  session.SetMode(mgo.Monotonic, true)
  con := session.DB("test").C("hoge")
  //err = con.Insert(&Person{"Ale", "+55 53 8116 9639"},
  //           &Person{"Cla", "+55 53 8402 8510"})
  //if err != nil {
  //        panic(err)
  //}
  people := []Person{}
  err = con.Find(nil).All(&people)
  if err != nil {
          panic(err)
  }
  //----------------<<<

  return c.Render(people)
}

func (c Ocr) Index() revel.Result {
  return c.Render()
}

func (c Ocr) Show(page int) revel.Result {
  return c.Render(page)
}
