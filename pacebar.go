// Package pacebar implements a very simple progress bar
// to be used by go programs.
//
// To set up pacebar, create a Pacebar object (pb), which you need
// to provide with the amount of work that you will be doing,
// and a name.
// For every amount of work done, call pb.Done(amount).
package pacebar

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const maxWidth int = 40
const alpha float64 = 0.15

// Pacebar is the type to interact with when showing
// progress through the pacebar package
type Pacebar struct {
	// mu is used for mutual exclusion
	mu sync.Mutex
	// Work is the amount of work to be done in total
	Work int
	// Name is the name of the work that will be done
	Name string

	// done is the amount of completed work
	done int

	// speed is the estimated completions per second
	speed float64

	// lastUpdate contains the last time Done() was called
	lastUpdate time.Time
}

func (pb *Pacebar) runningAverage(amount int) {
	if pb.speed == 0 {
		pb.speed = 10 //TODO: better initialization
	} else {
		thisUpdate := time.Now()
		currentSpeed := float64(amount) / (thisUpdate.Sub(pb.lastUpdate).Seconds())
		pb.speed = alpha*currentSpeed + pb.speed*(1-alpha)
		pb.lastUpdate = thisUpdate
	}
}

func (pb *Pacebar) showProgress() {
	var showRun, showMax int
	if pb.Work > maxWidth {
		showRun = pb.done * maxWidth / pb.Work
		showMax = maxWidth
	} else {
		showRun = pb.done
		showMax = pb.Work
	}

	fmt.Printf("\r\033[1m%s: \033[32m%s\033[31m%s \033[0m(%d / %d) (%4.f / s)", pb.Name,
		strings.Repeat("―", showRun), strings.Repeat("―", showMax-showRun), pb.done, pb.Work, pb.speed)

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
	pb.runningAverage(amount)
	pb.showProgress()
}
