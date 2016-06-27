// Golang Sorter
package main

import (
    "sort"
    "fmt"
    "os"
    "bufio"
    "time"
)

var file_unsorted = "unsorted.txt"
var file_length_sort = "len_sort.txt"
var file_alpha_sort = "alpha_sort.txt"

type ByLength []string
type ByAlphabet []string

// Methods to implement the ByLength sort
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
   s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// Methods to implement the ByAlphabet sort
func (s ByAlphabet) Len() int {
    return len(s)
}
func (s ByAlphabet) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByAlphabet) Less(i, j int) bool {
    return s[i] < s[j]
}

// File methods
func createFile(filename string) (f *os.File) {
    f, err := os.Create(filename)
    check(err)
    return
}
func openFile(filename string) (f *os.File) {
    f, err := os.Open(filename)
    check(err)
    return
}
func writeFile(f *os.File, arr []string) {
    for i := 0; i < len(arr); i++ {
        f.WriteString(arr[i] + " ")
    }
}
func readFile(f *os.File) (arr []string) {
    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        arr = append(arr, scanner.Text())
    }
    return
}
func closeFile(f *os.File) {
    defer f.Close()
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
        fmt.Println("start: ", time.Now())
        f := openFile(file_unsorted)
        fmt.Println("after open: ", time.Now())
        arr := readFile(f)
        fmt.Println("after read: ", time.Now())

	sort.Sort(ByLength(arr))
        fmt.Println("after LEN sort: ", time.Now())
//        fmt.Println("ByLenght: ",arr)
        f1 := createFile(file_length_sort)
        writeFile(f1, arr)
        closeFile(f1)
        fmt.Println("after write LEN file: ", time.Now())

        sort.Sort(ByAlphabet(arr))
        fmt.Println("after ALPH sort: ", time.Now())
//        fmt.Println("ByAlphabet: ",arr)

        f2 := createFile(file_alpha_sort)
//        fmt.Println("after create file: ",time.Now())
        writeFile(f2, arr)
        fmt.Println("after write ALPH file: ", time.Now())
        closeFile(f2)
}
