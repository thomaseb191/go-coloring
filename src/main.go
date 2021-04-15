package main

import (
	"fmt"
	g "github.com/thomaseb191/go-coloring/graphs"
	t "github.com/thomaseb191/go-coloring/testHarness"
)

func main() {
	gr := t.ParseFile("res/Sample01.txt", true)
	//gr := t.ParseFile("res/Error01.txt") //Errors due to directed edges
	//gr := t.ParseFile("res/Error02.txt") //Errors due to too small max degree

	//gr2 := g.DeepCopy(&gr)
	//gr2.Name = "A copy"
	//gr2.Description = "A copied description"
	//gr2.MaxDegree = 10
	//gr2.Nodes[0].Name = "Acopy"
	//gr2.Nodes[0].Neighbors = append(gr2.Nodes[0].Neighbors, &gr2.Nodes[1])

	g.PrintGraph(&gr)
	//g.PrintGraph(&gr2)

	fmt.Printf("\n\n\n")

	tResults := t.RunTest("res/Sample02.txt", []int{}, -1, 3)
	fmt.Printf("Duration of test %s: %d with %d colors\n", tResults[0].Name, tResults[0].DurationMillis.Milliseconds(), tResults[0].NumColors)
	g.PrintGraph(&tResults[0].Output)
}
