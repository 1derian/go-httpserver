package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
)
const DEFAULTPORT = 7788
func main() {
	port := flag.Int("p", DEFAULTPORT, "Set The Http Port")
	flag.Parse()
	pwd, _ := os.Getwd()
	http.Handle("/", http.FileServer(http.Dir(pwd)))
	err := http.ListenAndServe(":"+ strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println(err)
	}
}
