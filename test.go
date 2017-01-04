// You can edit this code!
// Click here and start typing.
package main

import "fmt"
import "sync"
import "runtime"

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Hello, My name is amit")
	fmt.Println("This is my first Go program i am really existed now")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for index := 0; index < 45; index++ {
		fmt.Println("Foo:", index)
	}
	wg.Done()
}

func bar() {
	for index := 0; index < 45; index++ {
		fmt.Println("Bar", index)
		fmt.Println("Bar", index)
	}
	wg.Done()
}
