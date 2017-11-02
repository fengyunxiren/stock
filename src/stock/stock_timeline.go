package stock

import (
  "encoding/json"
  "time"
  "strconv"
  "errors"
)

type AskOrBid struct{
  Price float64 `json: "price"`
  Volume int `json: "volume"`
}

type TimeLineData struct{
    Amount int `json: "amount"`
    AvgPrice int `json: "avgPrice"`
    Ccl int `json: "ccl"`
    Date int `json: "date"`
    NetChangeRatio float64 `json: "netChangeRatio"`
    PreClose float64 `json: "preClose"`
    Price float64 `json: "price"`
    Time int `json: "time"`
    Volume int `json: "volume"`
}

type TimeLineJson struct{
  Ask []AskOrBid `json: "ask"`
  Bid []AskOrBid `json: "bid"`
  ErrorNo int `json: "errorNo"`
  ErrorMsg string `json: "errorMsg"`
  ExchangeStatus int `json: "exchangeStatus"`
  LatestTimelineStamp string `json: "latestTimelineStamp"`
  PreClose float64 `json: "preClose"`
  TimeLine []TimeLineData `json: "timeline"`
  Version int `json: "version"`
}

func GetTimeLine(stockCode string) ([]TimeLineData, error) {
  queryUrl := TIME_LINE
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": stockCode,
    "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
  }
  ret, err := Http.Get(queryUrl, param)
  if err != nil {
    return nil, err
  }
  var timeline TimeLineJson
  err = json.Unmarshal(ret, &timeline)
  if err != nil {
    return nil, err
  }
  if timeline.ErrorNo != 0 {
    return nil, errors.New(timeline.ErrorMsg)
  }
  return timeline.TimeLine, nil
}

type TimeLineFiveJson struct{
  ErrorNo int `json: "errorNo"`
  ErrorMsg string `json: "errorMsg"`
  ExchangeStatus int `json: "exchangeStatus"`
  LatestTimelineStamp string `json: "latestTimelineStamp"`
  PreClose float64 `json: "preClose"`
  TimeLine []TimeLineData `json: "timeline"`
  Version int `json: "version"`
}

func GetTimeLineFive(stockCode string) ([]TimeLineData, error) {
  queryUrl := TIME_LINE_FIVE
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": stockCode,
    "step": "3",
    "timestamp": strconv.FormatInt(time.Now().Unix(), 10),
  }
  ret, err := Http.Get(queryUrl, param)
  if err != nil {
    return nil, err
  }
  var timeline TimeLineFiveJson
  err = json.Unmarshal(ret, &timeline)
  if err != nil {
    return nil, err
  }
  if timeline.ErrorNo != 0 {
    return nil, errors.New(timeline.ErrorMsg)
  }
  return timeline.TimeLine, nil
}
