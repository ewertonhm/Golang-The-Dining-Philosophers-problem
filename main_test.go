package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()

		if len(orderFinished) != 5 {
			t.Errorf("incorrect lenght of slice; expected 5 but got %d", len(orderFinished))
		}
	}
}

func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"zero daley", time.Second * 0},
		{"quarter second daley", time.Microsecond * 250},
		{"half second daley", time.Microsecond * 500},
	}
	for _, e := range theTests {
		orderFinished = []string{}

		eatTime = e.delay
		sleepTime = e.delay
		thinkTime = e.delay

		for i := 0; i < 10; i++ {
			orderFinished = []string{}
			dine()

			if len(orderFinished) != 5 {
				t.Errorf("%s: incorrect lenght of slice; expected 5 but got %d", e.name, len(orderFinished))
			}
		}
	}
}
