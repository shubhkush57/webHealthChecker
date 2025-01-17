package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link,c) // a new go routine
		
	}
	for i := 0;i<len(links);i++{
		fmt.Println(i,<-c)
	}
	// checking infite time with some delays..
	for {
		// it pause the main routine for 2 seconds only
		// not good to put sleep int he main .
		// time.Sleep(time.Second * 2)
		// go checkLink(<-c,c) // this is also a blocking function....
		// we also need to pause this for the seconds..

		// adding the function literals
		go func(){
			time.Sleep(2*time.Second)
			checkLink(<-c,c)
		}()

	}
	// blocking channel first one get's and executes and exits
	// still main routine is waiting for some other data to come.

}

func checkLink(link string,c chan string) {
	_,err := http.Get(link) // blocking call so got frozen on this line
	if err != nil{
		fmt.Println(link,"might be down !")
		c <- "Might be down I think"
		return 
	}
	fmt.Println(link ,"is up !")
	c<- "Yep it is up"
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

Main-> Parent Routine: default
others -> Child Routine: we have to add go

Main routine doesn't care about child routines to complete  the process.

Channles is to communicate btw different go routines. only way to do this.

Channle is a type 

recieve a message thorught the channel is a blocking thing.

Function literals aare the anonymous function..
*/