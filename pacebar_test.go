package pacebar

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	const work int = 100
	pb := Pacebar{
		Work: work,
		Name: "TestBar",
	}

	for i := 0; i < work; i++ {
		pb.Done(1)
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
}
