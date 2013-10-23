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
  ExtVer      string
}

func Page(page int, countPerPage int) []OcrReport {
  count := countPerPage
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
    result bool,
    extVer string) *OcrReport {

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
    ExtVer:      extVer,
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

func FindOlder(ts int64, result bool) []OcrReport {
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

  err = collection.Find(bson.M{"createdtime" : bson.M{ "$lt" : ts }, "result" : result}).All(&reports)
  if err != nil {
    panic(err)
  }

  return reports
}

func Truncate(ts int64) *mgo.ChangeInfo {

  // {{{ TODO : DRY
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  defer session.Close()
  session.SetMode(mgo.Monotonic, true)
  collection := session.DB("kcwidget").C("logOcr")
  // }}}

  changeInfo, err2 := collection.RemoveAll(bson.M{"createdtime" : bson.M{"$lt": ts}})
  if err2 != nil {
    panic(err2)
  }

  return changeInfo
}
