package core

import (
	"fmt"
	//"io/ioutil"
	"os"
	"net/http"
	"time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func DatabaseLog(start time.Duration, r *http.Request) {

	// To start, here's how to dump a string (or just
	// bytes) into a file.
	//d1 := []byte("hello\ngo\n")
	//err := ioutil.WriteFile("./log/log", d1, 0644)
	//check(err)

	// For more granular writes, open a file for writing.
	f, err := os.OpenFile("./log/127", os.O_APPEND|os.O_WRONLY, 0600)
	check(err)

	// It's idiomatic to defer a `Close` immediately
	// after opening a file.
	defer f.Close()

	// You can `Write` byte slices as you'd expect.
	//d2 := []byte{115, 111, 109, 101, 10}
	//n2, err := f.Write(d2)
	//check(err)
	//fmt.Printf("wrote %d bytes\n", n2)

	// A `WriteString` is also available.
	strng := fmt.Sprintf("%s\t%f\t%s\t%s\n",
			start,
			time.Duration.Seconds(start),
			r.RequestURI,
			r.Method,)
	n3, err := f.WriteString(strng)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a `Sync` to flush writes to stable storage.
	f.Sync()

}
