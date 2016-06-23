package main

import (
    "time"
    "fmt"
    "os"
    "bytes"
    "strconv"
)

var filename = "LOG.txt"
var buffer_size = 10
var num_requests = 10

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func createFile() (f *os.File) {
    f, err := os.Create(filename)
    check(err)
    return
}

func writeFile(f *os.File, s string) {
    // A `WriteString` is also available.
    n3, _ := f.WriteString(s+"\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // Issue a `Sync` to flush writes to stable storage.
    f.Sync()
}

func makeRequests(f *os.File, arg string, requests_total int) {
    requests := make(chan int, num_requests)
    for i := 1; i <= num_requests; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(time.Millisecond * 1000)

    for req := range requests {
        <-limiter
        fmt.Println("request", req+requests_total, time.Now())
    }

    burstyLimiter := make(chan time.Time, buffer_size)

    for i := 0; i < buffer_size; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Millisecond * 1000) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, num_requests)
    for i := 1; i <= num_requests; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req+requests_total, time.Now())
        var buffer bytes.Buffer
        buffer.WriteString("process ")
        buffer.WriteString(os.arg)
        buffer.WriteString(" - request ")
        buffer.WriteString(strconv.Itoa(req+requests_total))
        buffer.WriteString(": ")
        buffer.WriteString(time.Now().String())
        writeFile(f, buffer.String())
    }
}

func main() {
    arg := ""
    if len(os.Args) > 1 {
        arg = os.Args[1]
    }
    f := createFile()
    requests_total := 0
    for {
        makeRequests(f, arg, requests_total)
        requests_total += num_requests
    }
}
