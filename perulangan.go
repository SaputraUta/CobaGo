package main

import "fmt"

func main() {
	i := 1

	for {
		fmt.Println("Perulangan while ke ", i)
		i++
		if i > 5 {
			break
		}
	}
	fmt.Println("Ini coba stash")
}
