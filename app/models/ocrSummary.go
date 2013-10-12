//package ocrSummary
package ocrReport

import (
  "labix.org/v2/mgo"
  //"labix.org/v2/mgo/bson"
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
