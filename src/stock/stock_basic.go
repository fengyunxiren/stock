package stock

import (
  "encoding/json"
  "time"
  "strconv"
  "errors"
  "strings"
)

type BasicInfo struct{
  AmplitudeRatio float64 `json: "amplitudeRatio"`
  Asset int `json: "asset"`
  Capitalization float64 `json: "capitalization"`
  Close float64 `json: "close"`
  Exchange string `json: "exchange"`
  FollowNum int `json: "followNum"`
  High float64 `json: "high"`
  Low float64 `json: "low"`
  NetChange float64 `json: "NetChange"`
  Open float64 `json: "open"`
  PreClose float64 `json: "preClose"`
  Stockcode string `json: "stockCode"`
  StockName string `json: "stockName"`
  StockStatus string `json: "stockStatus"`
  TurnoverRatio int `json: "turnoverRatio"`
  Volume int `json: "volume"`
}

type BasicInfoJson struct{
  ErrorMsg string `json: "errorMsg"`
  ErrorNo int `json: "errorNo"`
  Data []BasicInfo `json: "data"`
}

func GetBasicInfo(stockCode []string) ([]BasicInfo, error) {
  queryUrl := BASIC_BATCH
  param := map[string]string{
    "from": "pc",
    "os_ver": "1",
    "cuid": "xxx",
    "vv": "100",
    "format": "json",
    "stock_code": strings.Join(stockCode, ","),
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
  var basic BasicInfoJson
  err = json.Unmarshal(ret, &basic)
  if err != nil {
    return nil, err
  }
  if basic.ErrorNo != 0 {
    return nil, errors.New(basic.ErrorMsg)
  }
  return basic.Data, nil
}
