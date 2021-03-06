package testHarness

import (
	"fmt"
	r "github.com/thomaseb191/go-coloring/reductions"
	//d "../display" //TODO: IMPORT
	g "github.com/thomaseb191/go-coloring/graphs"
	"time"
)

// TestData is a struct to handle metadata and an output graph
//		Name: the name of the test, following the convention of graphName_algorithmName
//		DurationMillis: the Duration of the reduction, designed to be converted to millis in post-processing
//		Output: the graph produced by the output of the algorithm
//		NumColors: the number of colors in the output graph. Its correctness should be asserted in post-processing
//		IsSafe: the result of running g.IsSafe() on the output
type TestData struct {
	Name string
	DurationMillis time.Duration
	Output g.Graph
	NumColors int
	IsSafe bool
}

// RunTest runs any number of color-reducing algorithms on a given graph file.
// 		fileName: the string name of the file for the graph
// 		algos: an array of IDs mapping to algorithm
// 		poolSize: the number of worker goroutines allowed for parallel algorithms
// 		debug: 0 if just generate output, 1 if allow prints, 2 if just graph and output, 3 if allow graph and prints
func RunTest(fileName string, algos []int, poolSize int, debug int) []TestData {
	//Parse and build the graph. Initialize the colors manually after asserting not safe
	var testDatas []TestData
	initGraph := ParseFile(fileName, false)
	if debug % 2 == 1 {
		fmt.Printf("Initial IsSafe() for %s without color init: %t\n", initGraph.Name, g.IsSafe(&initGraph))
	}
	g.RunColorInit(&initGraph)


	if len(algos) == 0 {
		algos = r.AllAlgIds
	}

	//Start the time, run algorithms
	for _, algo := range algos {
		copiedGraph := g.DeepCopy(&initGraph)
		start := time.Now()
		outGraph, algoName := r.RunReduction(copiedGraph, algo, poolSize, debug)

		//Stop the time, check the algorithm
		elapsed := time.Since(start)
		//fmt.Println(start, time.Now(), elapsed.Milliseconds(), elapsed.Nanoseconds())
		numColors := g.CountColors(&outGraph)
		isSafe := g.IsSafe(&outGraph)

		if debug % 2 == 1 {
			fmt.Printf("Output IsSafe() for %s_%s in %d: %t\n", initGraph.Name, algoName, elapsed.Nanoseconds(), isSafe)
			fmt.Printf("\t\tNum Colors: %d\n", numColors)
		}
		testName := initGraph.Name + "_" + algoName


		newTest := TestData{
			Name: testName,
			DurationMillis: elapsed,
			Output: outGraph,
			NumColors: numColors,
			IsSafe: isSafe,
		}
		testDatas = append(testDatas, newTest)

		//Render if desired
		if debug >= 2 {
			graphs := make([]*g.Graph, 2)
			graphs[0] = &g.Graph{
				Name: initGraph.Name + "_Original",
				Description: initGraph.Description + " Original",
				Nodes: initGraph.Nodes,
				MaxDegree: initGraph.MaxDegree,
			}
			graphs[1] = &g.Graph{
				Name: outGraph.Name + "_After_Reduction",
				Description: outGraph.Description + " After Reduction",
				Nodes: outGraph.Nodes,
				MaxDegree: outGraph.MaxDegree,
			}
			g.GenerateHTMLForMany(graphs)
			//g.GenerateHTMLForOne(&outGraph, testName)
		}
	}
	return testDatas
}