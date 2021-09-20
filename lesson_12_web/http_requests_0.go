package main

import "net/http"
import "fmt"
import "io"


func main(){
  resp, _ := http.Get("http://example.com/")
  body := resp.Body
  headers := resp.Header

  status := resp.Status
  fmt.Println(status)

  fmt.Println(headers)
  fmt.Println(headers["Content-Type"])

  fmt.Println(body)
  body_data, _ := io.ReadAll(body)
  fmt.Println(body_data)
  fmt.Println(string(body_data))
}
