package main

import(
  "net/http"
  "log"
  "flag"
)

var port = flag.String("port", "1337", "http service address")

func main() {
  flag.Parse()
  http.HandleFunc("/", handleRoot)
  err := http.ListenAndServe(":" + *port, nil)
  log.Println(err.Error())
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
  if req.Method != "GET" {
    http.Error(w, "Method not allowed", 405)
    log.Println("Wrong Method Bro")
    return
  } else if req.URL.Path == "/" {
    http.ServeFile(w, req, "index.html")
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    return
  } else if req.URL.Path == "/hellno2.m4a" {
    http.ServeFile(w, req, "hellno2.m4a")
    w.Header().Set("Content-Type", "audio/mp4")
    return
  } else if req.URL.Path == "/maverick.png" {
    http.ServeFile(w, req, "maverick.png")
    w.Header().Set("Content-Type", "image/png")
    return
  } else if req.URL.Path == "/favicon.ico" {
    http.ServeFile(w, req, "favicon.ico")
    w.Header().Set("Content-Type", "image/x-icon")
    return
  } else {
    http.Error(w, "Not found", 404)
    return
  }
}