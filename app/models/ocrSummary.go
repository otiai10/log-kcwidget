//package ocrSummary
package ocrReport

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type OcrSummary struct {
  DateHour    int
  Year        int
  Month       int
  Date        int
  Hour        int
  Success     int
  Failure     int
}
func AddSummary(summary OcrSummary) bool {

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("summaryOcr")

  err = collection.Insert(summary)
  if err != nil {
    panic(err)
  }
  return true
}

func FindAllSummary() []OcrSummary {

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("summaryOcr")

  summaries := []OcrSummary{}
  err = collection.Find(nil).All(&summaries)
  if err != nil {
    panic(err)
  }
  return summaries
}

func FindSummary(datehour int) []OcrSummary {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("summaryOcr")

  summaries := []OcrSummary{}

  err = collection.Find(bson.M{"datehour" : bson.M{ "$gt" : datehour }}).All(&summaries)
  if err != nil {
    panic(err)
  }

  return summaries
}
