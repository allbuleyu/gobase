package main

import (
	"github.com/allbuleyu/gobase/gopl.io/ch7"
	"fmt"
	"net/http"
	"log"
	"io"
	"github.com/allbuleyu/gobase/gopl.io/ch3"
	"runtime"
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

	http.HandleFunc("/", handler1)
	http.ListenAndServe("localhost:8080", nil)


	return
	for i := 0; i < 10; i++ {
		mmm = append(mmm, map[int]int{i:i})
	}

	for _, mm := range mmm {
		m := mm
		fmt.Printf("mm ----- %p === %v \n", mm, mm)

		fmt.Printf("m %p === %v \n", m, m)
		fff = append(fff, func() {
			fmt.Println(m)
			fmt.Printf("%p\n", m)
		})
	}

	for _, ff := range fff {
		ff()
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}