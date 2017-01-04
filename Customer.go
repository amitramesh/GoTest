package main

import "fmt"
import "sync"
import "runtime"

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Hello, My name is amit")
	fmt.Println("This is my first Go program i am really existed now")
	//wg.Add(2)
	//go foo()
	//go bar()
	wg.Wait()
}
