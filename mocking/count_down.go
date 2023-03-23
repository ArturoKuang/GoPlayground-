package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type Sleeper interface {
	Sleep()
}

type SpyTime struct {
	durationSlept time.Duration
}

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

type DefaultSleeper struct{}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Test(t string) (n int, err error) {
	fmt.Println("dsa")
	return
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
