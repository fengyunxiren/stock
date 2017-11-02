package stock

import (
  "../request"
)

var Http request.Request
const (
  BASE_URL = "https://gupiao.baidu.com/api"
  TIME_LINE = BASE_URL + "/stocks/stocktimeline"
  THEME_LIST = BASE_URL + "/concept/getthemelist"
  FOLLOW_NUM = BASE_URL + "/rails/stockfollownum"
  BASIC_BATCH = BASE_URL + "/rails/stockbasicbatch"
  TIME_LINE_FIVE = BASE_URL + "/stocks/stocktimelinefive"
  DAY_BAR = BASE_URL + "/stocks/stockdaybar"
  WEEK_BAR = BASE_URL + "/stocks/stockweekbar"
  MONTH_BAR = BASE_URL + "/stocks/stockmonthbar"
)
