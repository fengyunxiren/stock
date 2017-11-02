package stock

import (
  "encoding/json"
  "time"
  "strconv"
  "errors"
)

type BarData struct{
  Date int `json: "date"`
  Kdj KDJ `json: "kdj"`
  Ma5 Ma `json: "ma5"`
  Ma10 Ma `json: "ma10"`
  Ma20 Ma `json: "ma20"`
  Macd MACD `json: "macd"`
  Rsi RSI `json: "rsi"`
}

type KDJ struct{
  D float64 `json: "d"`
  J float64 `json: "j"`
  K float64 `json: "k"`
}

type Kline struct{
  Amount int `json: "amount"`
  Ccl interface{} `json: "ccl"`
  Close float64 `json: "close"`
  High float64 `json: "high"`
  Low float64 `json: "low"`
  NetChangeRatio float64 `json: "netChangeRatio"`
  Open float64 `json: "open"`
  PreClose float64 `json: "preClose"`
  Volume float64 `json: "volume"`
}

type Ma struct{
  AvgPrice float64 `json: "avgPrice"`
  Ccl interface{} `json: "ccl"`
  Volume float64 `json: "volume"`
}

type RSI struct{
  Rsi1 float64 `json: "rsi1"`
  Rsi2 float64 `json: "rsi2"`
  Rsi3 float64 `json: "rsi3"`
}

type MACD struct{
  Dea float64 `json: "dea"`
  Diff float64 `json: "diff"`
  Macd float64 `json: "macd"`
}

type BarJson struct{
  ErrorMsg string `json: "errorMsg"`
  ErrorNo int `json: "errorNo"`
  MashData []BarData `json: "mashData"`
}

func GetDayBar(stockCode string) ([]BarData, error) {
  queryUrl := DAY_BAR
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": stockCode,
    "step": "3",
    "start": "",
    "count": "160",
    "fq_type": "no",
    "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
  }
  ret, err := Http.Get(queryUrl, param)
  if err != nil {
    return nil, err
  }
  var dayBar BarJson
  err = json.Unmarshal(ret, &dayBar)
  if err != nil {
    return nil, err
  }
  if dayBar.ErrorNo != 0 {
    return nil, errors.New(dayBar.ErrorMsg)
  }
  return dayBar.MashData, nil
}

func GetWeekBar(stockCode string) ([]BarData, error) {
  queryUrl := WEEK_BAR
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": stockCode,
    "step": "3",
    "start": "",
    "count": "160",
    "fq_type": "no",
    "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
  }
  ret, err := Http.Get(queryUrl, param)
  if err != nil {
    return nil, err
  }
  var weekBar BarJson
  err = json.Unmarshal(ret, &weekBar)
  if err != nil {
    return nil, err
  }
  if weekBar.ErrorNo != 0 {
    return nil, errors.New(weekBar.ErrorMsg)
  }
  return weekBar.MashData, nil
}

func GetMonthBar(stockCode string) ([]BarData, error) {
  queryUrl := MONTH_BAR
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": stockCode,
    "step": "3",
    "start": "",
    "count": "160",
    "fq_type": "no",
    "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
  }
  ret, err := Http.Get(queryUrl, param)
  if err != nil {
    return nil, err
  }
  var monthBar BarJson
  err = json.Unmarshal(ret, &monthBar)
  if err != nil {
    return nil, err
  }
  if monthBar.ErrorNo != 0 {
    return nil, errors.New(monthBar.ErrorMsg)
  }
  return monthBar.MashData, nil
}
