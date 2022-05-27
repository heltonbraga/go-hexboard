package main

import (
	"fmt"
	"hex/board"
)

func main() {
	fmt.Println("--------------------")
	b := board.MakeBoard(3)
	m := board.GetMatrix(b)
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			if m[j][i] == 0 {
				fmt.Printf(" -- ")
			} else {
				fmt.Printf(" %2d ", m[j][i])
			}
		}
		fmt.Println("")
	}
	fmt.Println("--------------------")
}
