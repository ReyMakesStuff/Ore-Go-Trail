package main

/*
	Ore-Go-Trail
	by RBurke
	Based on the 1978 BASIC code, I present to you....
	A terrible version in Go!

	Apologies to Don Rawitsch, Bill Heinemann, Paul Dillenberger
	(the original programmers)
*/

import (
	f "fmt"
	"strings"
)

func main() {
	//nothing to see here yet
	f.Println("Ore-Go-Trail!")

	//check for help; I put this in a block to free up the memory for the answer
	//may be useless, so may change later.
	{
		f.Print("Do you need help? (y/n) ")
		var answer string
		f.Scanln(&answer)
		if strings.ToUpper(answer) == "Y" {
			showHelp()
		}
	}
}

func showHelp() {
	//show help information
	f.Printf("\nThis is a Go interpretation of the original Oregon Trail\n\n")
	f.Println("It is a text-based simulation of a family of five making a perilous journey.")
	f.Printf("Traveling from Independence, Missouri to Oregon City, Oregon in 1847,\nthey will have to travel 2,040 miles in around 6 months.\n\n")
	f.Println("But they'll probably die.")
}
