package main

import (
	clean "CleanMyWorkspace/Clean"
	"fmt"

	b "github.com/Mentor-Paris/CleanMyWorkspace"
)

func printWorkspace(a *[][]*string) {
	if a == nil {
		fmt.Println("(workspace is nil)")
		return
	}
	for _, row := range *a {
		for _, cell := range row {
			if cell == nil {
				fmt.Print("|nil")
			} else {
				fmt.Print("|" + *cell)
			}
		}
		fmt.Print("|\n")
	}
}

func main() {

	workspace := b.GenererateWorkSpace()

	fmt.Println("\n\nLe souk après 8h de puissance 4")
	printWorkspace(workspace)

	cleaned := clean.CleanWorkSpace(workspace)

	fmt.Println("\n\nLe souk après la punition des mentors")
	printWorkspace(cleaned)

	fmt.Println("\n\nBien joué, vous possédez officiellement du respect pour les autres.")
}
