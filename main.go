package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_,err := http.Get(link) // blocking call so got frozen on this line
	if err != nil{
		fmt.Println(link,"might be down !")
		return 
	}
	fmt.Println(link ,"is up !")
	return
}
//we have to add go routines.

/* go automatically creates one go routine by deafult. 
To create a new go routine we write 'go' before that function
so go create a new thread....go routine.
So go scheduler works on one CPU (may be many CPU) and handle the different go routines.
If one what to at a time . -> one CPU
assign different routine to different cpu.

Concurrency is not parallism
*/