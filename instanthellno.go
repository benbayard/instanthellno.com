package main

import(
  "net/http"
  "log"
)

func main() {
  http.HandleFunc("/", handleRoot)
  err := http.ListenAndServe(":1337", nil)
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
  } else {
    http.Error(w, "Not found", 404)
  }
}