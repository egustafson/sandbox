package main

import "fmt"

func main() {

	runner()

	fmt.Println("done.")
}

func doPanic() {
	e := error(nil)
	e.Error() // cause nil dereference (a panic)
	//panic("intentional panic")
}

func runner() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic occured: %v\n", err)
		}
	}()

	doPanic()

}
