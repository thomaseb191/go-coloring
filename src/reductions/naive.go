package reductions

import (
	g "../graphs"
	"fmt"
	"time"
)

func RunNaive(gr g.Graph, poolSize int, debug int) g.Graph {
	if debug % 2 == 1 {
		fmt.Printf("Starting reduction for %s algorithm...\n", "Naive")
	}
	time.Sleep(100 * time.Millisecond)
	//TODO: IMPLEMENT REDUCTION

	return gr
}
