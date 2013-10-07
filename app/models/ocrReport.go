package ocrReport

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type OcrReport struct {
  // idとりたい
  //Id          string
  ImgURI      string
  CreatedTime int
  UserAgent   string
  RawText     string
  AssuredText string
  Result      bool
}

func init() {
}

func Page(page int) []OcrReport {
  count := 10
  skip  := count * page

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  reports := []OcrReport{}

  err = collection.Find(nil).Sort("-createdtime").Skip(skip).Limit(count).All(&reports)
  if err != nil {
    panic(err)
  }
  return reports
}

func Get(targetCreatedTime int) OcrReport {

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  report := OcrReport{}
  err = collection.Find(bson.M{"createdtime" : targetCreatedTime }).One(&report)
  return report
}

// 使っちゃだめよ！！
func All() []OcrReport {

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  reports := []OcrReport{}

  err = collection.Find(nil).Sort("-createdtime").Skip(0).Limit(5).All(&reports)
  if err != nil {
    panic(err)
  }
  return reports
}

func Count() int {

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  count, err2 := collection.Count()
  if err2 != nil {
    panic(err2)
  }
  return count
}

func Add(
    imgURI string,
    createdTime int,
    userAgent string,
    rawText string,
    assuredText string,
    result bool) *OcrReport {

  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()

  // Factory?
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
    return report
  }

  session.SetMode(mgo.Monotonic, true)
  con := session.DB("kcwidget").C("logOcr")
  err = con.Insert(report)
  if err != nil {
    panic(err)
  }

  return report
}

func Delete(targetCreatedTime int) OcrReport {

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  report := OcrReport{}
  err = collection.Find(bson.M{"createdtime" : targetCreatedTime }).One(&report)

  err = collection.Remove(bson.M{"createdtime" : targetCreatedTime })

  //if err == nil {
  //  return report
  //}
  return report
}
