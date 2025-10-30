package main

import (
	"CleanMyWorkSpace/clean"
	"fmt"

	c "github.com/Mentor-Paris/CleanMyWorkspace"
)

func main() {
	souk := c.GenererateWorkSpace()

	clean.CleanWorkSpace(souk)

	afficherEspaceTravail(souk)

}

func afficherEspaceTravail(espaceTravail *[][]*string) {
	for _, ligne := range *espaceTravail {
		for _, objet := range ligne {
			if objet == nil {
				fmt.Print("|nil")
			} else {
				fmt.Print("|" + *objet)
			}
		}
		fmt.Println("|")
	}
}
