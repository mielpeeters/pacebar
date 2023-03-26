package pacebar

import (
	"fmt"
	"strings"
	"sync"
)

const maxWidth int = 40

// Pacebar is the type to interact with when showing
// progress through the pacebar package
type Pacebar struct {
	// mu is used for mutual exclusion
	mu sync.Mutex
	// Work is the amount of work to be done in total
	Work int
	// done is the amount of completed work
	done int
}

func (pb *Pacebar) showProgress() {

	maxStripes := 40

	var showRun, showMax int
	if pb.Work > maxStripes {
		showRun = pb.done * 40 / pb.Work
		showMax = 40
	} else {
		showRun = pb.done
		showMax = pb.Work
	}

	fmt.Printf("\r\033[1mPacebar: \033[32m%s\033[31m%s \033[0m(%d / %d)",
		strings.Repeat("―", showRun), strings.Repeat("―", showMax-showRun), pb.done, pb.Work)

	// end by adding a line for subsequent outputs
	if pb.done >= pb.Work {
		fmt.Printf("\n")
	}
}

// Done is used to indicate that an amount of work has been done
func (pb *Pacebar) Done(amount int) {
	pb.mu.Lock()
	defer pb.mu.Unlock()

	pb.done += amount
	pb.showProgress()
}
