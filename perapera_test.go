package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/mattrosencrantz/gotools/templateset"
/*	"io/ioutil"*/
)

func Sum(a *[]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

type T struct {
	name string
}

type MyString string

func (ms MyString) String() string {
/*	fmt.Println("h")*/
	return (string(ms) + "!!")
}

func main() {
/*	list, _ := ioutil.ReadDir("/Users/sunfmin/Developments/perapera/templates/videos")

	for _, f := range list {
		fmt.Println(f.Name)
	}
*/
	ts := templateset.MustBuildFromDirectory("./templates/videos", nil)

	fmt.Println(ts)

	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println(([]int)("Hi Felix"))
	fmt.Println(string([]int{70, 101, 108, 105, 120}))
	fmt.Println(string(1212))
	fmt.Println(strconv.Itoa64(121212312312))
	fmt.Println(strconv.Itoa64(121))

	fmt.Println(MyString("hello cool"))

	fmt.Println(12<<1)

	fmt.Println(os.Stdout)

	f, _ := os.Open("perapera.go")
	var buf [48]byte

	fmt.Println(buf)

	f.Read(buf[0:32])

/*	fmt.Println(string([]byte{'1', '2'}))*/
	fmt.Println(string([]byte{112, 97}))
	fmt.Println(string(buf[0:]))

	fmt.Println(cap(buf[0:12]))
	fmt.Println(len(buf[0:12]))

	fmt.Println(os.Stdout)

	fmt.Println(buf)

	a := [...]float64{7.0, 8.5, 9.1, 10, 12}

	fmt.Println(a)

	fmt.Println(&T{"hi"})

	fmt.Println(&[]int{1,2})

	x := Sum(&[]float64{7.0, 8.5, 9.1, 10, 202})
	fmt.Println(x)
}

