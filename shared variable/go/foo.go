// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    //"time"
)

type Request struct {
    op string 
    done chan bool
}


var i = 0

func server(requests chan Request) {
    for {
        select{
        case req:= <-requests:
            switch req.op {
            case "inc":
                i++
            case "dec":
                i--
            case "read":
                req.done <- true
            }
        }
    }
}

func incrementing(requests chan Request, done chan bool) { 
    //TODO: increment i 1000000 times
    for j := 0; j < 1000000; j++ {
        requests <- Request{op: "inc"}
    }
    requests <- Request{op: "read", done: done}
}

func decrementing(requests chan Request, done chan bool) {
    //TODO: decrement i 1000000 times
    for j := 0; j < 1000000; j++ {
        requests <- Request{op: "dec"}
    }
    requests <- Request{op: "read", done: done}
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    runtime.GOMAXPROCS(2)   
    // allows the program to run two goroutines at the same time (parallelism)
	

    // TODO: Spawn both functions as goroutines
    requests := make(chan Request)
    done := make(chan bool)

    go incrementing(requests, done)
    go decrementing(requests, done)
    go server(requests)
    <-done
    <-done
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    Println("The magic number is:", i)
}
