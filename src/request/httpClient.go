package request

import (
  "fmt"
  "net/http"
  "net/http/cookiejar"
  "net/url"
  "io/ioutil"
  "strings"
  "errors"
)

type Request struct{
  Client http.Client
  IsInit bool
}

func (r *Request) Init() {
  jar, err := cookiejar.New(nil)
  if err != nil {
    fmt.Println("Cookie set error!")
    return
  }
  r.Client.Jar = jar
  r.IsInit = true
}

func (r *Request) Post(queryUrl string, param map[string]string) ([]byte, error) {
  if !r.IsInit {
    r.Init()
    if !r.IsInit {
      return nil, errors.New("Init fail!")
    }
  }
  form := url.Values{}
  for key, value := range param {
    form.Set(key, value)
  }
  b := strings.NewReader(form.Encode())
  req, err := http.NewRequest("POST", queryUrl, b)
  if err != nil {
    fmt.Println("set request error!")
    return nil, err
  }
  req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
  resp, err := r.Client.Do(req)
  if err != nil {
    fmt.Println("response error!")
    return nil, err
  }
  defer resp.Body.Close()
  result, err := ioutil.ReadAll(resp.Body)
  if err != nil {
        fmt.Println("data error!")
        return nil, err
  }
  return result, nil
}

func (r *Request) Get(queryUrl string, param map[string]string) ([]byte, error) {
  if !r.IsInit {
    r.Init()
    if !r.IsInit {
      return nil, errors.New("Init fail!")
    }
  }
  u, err := url.Parse(queryUrl)
  if err != nil {
    fmt.Println("url parse error!")
    return nil, err
  }
  q := u.Query()
  for key, value := range param {
    q.Set(key, value)
  }
  u.RawQuery = q.Encode()
  req, err := http.NewRequest("GET", u.String(), nil)
  if err != nil {
    fmt.Println("Request set error!")
    return nil, err
  }
  resp, err := r.Client.Do(req)
  if err != nil {
    fmt.Println("Request error!")
    return nil, err
  }
  defer resp.Body.Close()
  result, err := ioutil.ReadAll(resp.Body)
  if err != nil {
        fmt.Println("data error!")
        return nil, err
  }
  return result, nil
}
