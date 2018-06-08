package main

import (
	"github.com/allbuleyu/gobase/gopl.io/ch7"
	"fmt"
	"net/http"
	"log"
	"io"
	"github.com/allbuleyu/gobase/gopl.io/ch3"

	"github.com/allbuleyu/gobase/gopl.io/ch8"
)

const debug bool = false

func handler(w http.ResponseWriter, r *http.Request)  {
	xxx := ch7.PrintTracks(nil)
	//fmt.Fprintf(w, "url.path = %q \n", r.URL.Path)
	xxx.Execute(w, ch7.Tracks)
}

// http 测试
func httpHander() {
	http.HandleFunc("/", handler)
	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}

var mmm []map[int]int
var fff []func()

func handler1(w http.ResponseWriter, r *http.Request) {
	ch3.Mm(w)
}

func main() {
	url := "http://www.catjc.com"
	urls, err := ch8.Extract(url)
	for i := range urls {
		fmt.Println(urls[i])
	}
	fmt.Println(err)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}