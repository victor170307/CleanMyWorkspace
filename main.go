package main

import (
	"CleanMyWorkspace/Clean"
	"fmt"

	"github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {

	workspace := CleanMyWorkspace.GenererateWorkSpace()

	fmt.Println("c'est pas propre du tout ici")
	printWorkspace(workspace)

	fmt.Println("\nc'est maintenant propre !!!")
	cleanedWorkspace := Clean.CleanWorkSpace(workspace)
	printWorkspace(cleanedWorkspace)
}

func printWorkspace(workspace *[][]*string) {
	if workspace == nil {
		return
	}

	for _, row := range *workspace {
		for i, cell := range row {
			if i > 0 {
				fmt.Print("|")
			}
			if cell == nil || *cell == "" {
				fmt.Print("nil")
			} else {
				fmt.Print(*cell)
			}
		}
		fmt.Println("|")
	}
}
