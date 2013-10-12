/*
 * cliから直で呼ばれる
 * Usage : `go run app/proc/summaize.go {arg}`
 */
package main

import (
  "time"
  "otiai10/logServer/app/models"
  "fmt"
)

func main() {
  //period := "24h"
  period := "24h"
  d,err := time.ParseDuration(period)
  if err != nil {
    panic(err)
  }
  ts_1day_before := time.Now().Add(-d).Unix() * 1000
  fmt.Printf("Period\t%v\n", period)

  successReports := ocrReport.FindOlder(ts_1day_before, true)

  failureReports := ocrReport.FindOlder(ts_1day_before, false)

  // 現在時刻を取得
  now := time.Now()
  // 欲しいフォーマットは 2013101209なので
  var datehour int
  ye := now.Year()
  mo := int(now.Month())
  da := now.Day()
  ho := now.Hour()
  datehour = ye * 1000000 + mo * 10000 + da * 100 + ho

  summary := ocrReport.OcrSummary{
    DateHour: datehour,
    Year:     ye,
    Month:    mo,
    Date:     da,
    Hour:     ho,
    Success:  len(successReports),
    Failure:  len(failureReports),
  }

  fmt.Printf("SUMMARY\t: %v\n", summary)

  // 集計を記録する
  succeeded := ocrReport.AddSummary(summary)
  fmt.Printf("SUCCESS\t: %v\n", succeeded)

  // 集計に使ったレコードは削除する
  changeInfo := ocrReport.Truncate(ts_1day_before)
  fmt.Printf("TRUNCATED\t: %v\n", changeInfo)
}
