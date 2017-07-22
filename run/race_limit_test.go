package run

import (
	"testing"

	floc "github.com/workanator/go-floc"
	"github.com/workanator/go-floc/flow"
)

const numOfRacers = 10

func TestRaceLimit(t *testing.T) {
	for no := 1; no <= numOfRacers; no++ {
		runRaceTest(t, no)
	}
}

func TestRaceLimitPanic(t *testing.T) {
	// Panic on zero limit
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s must panic because of invalid limit", t.Name())
			}
		}()

		runRaceTest(t, 0)
	}()

	// Panic on big limit
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("%s must panic because of invalid limit", t.Name())
			}
		}()

		runRaceTest(t, numOfRacers+1)
	}()
}

func runRaceTest(t *testing.T, limit int) {
	// Construct the flow control object.
	theFlow := flow.New()

	// Construct the state object which as data contains the counter.
	state := floc.NewStateContainer(new(int))
	defer state.Release()

	// Counstruct the result job.
	racers := make([]floc.Job, numOfRacers)
	for i := 0; i < numOfRacers; i++ {
		racers[i] = jobIncrement
	}

	theJob := RaceLimit(limit, racers...)

	// Run the job.
	floc.Run(theFlow, state, updateCounter, theJob)

	v := getCounter(state)
	if v != limit {
		t.Fatalf("%s expects counter value to be %d but get %d", t.Name(), limit, v)
	}
}
