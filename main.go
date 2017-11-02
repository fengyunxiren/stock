package main

import (
  "fmt"
  "./src/stock"
)

func main() {
  ret, _ := stock.GetTimeLine("usBIDU")
  for _, value := range ret {
    fmt.Println(value.Price)
  }
}
