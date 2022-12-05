package main

import "fmt"

func main() {

	i := 0
loop:
	for {
		if i == 4 {
			break loop
		}
		i++
		fmt.Println(i)
	}
}
