package main

import (
"fmt"
"net/http"
)

func main(){

http.handleFunc("/", func (w, http.ResponseWriter, r *http.Request){
  
  fmt.Fprintf(w, "Hello World")

})
}
