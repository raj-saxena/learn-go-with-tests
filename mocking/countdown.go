package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownFrom = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownFrom; i > 0; i-- {
		fmt.Fprintln(w, i)
		// time.Sleep(1 * time.Second)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
