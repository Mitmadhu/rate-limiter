package main

import (
	"errors"
	"time"
)

type RateLimiter struct {
	interval time.Duration
	maxToken int
	token    chan bool
}

func NewLimiter(rps int, maxToken int) *RateLimiter {
	rl := &RateLimiter{
		interval: time.Second / time.Duration(rps),
		maxToken: maxToken,
		token:    make(chan bool, maxToken),
	}

	// create max token
	for i := 0; i < rl.maxToken; i++ {
		rl.token <- true
	}

	go rl.refillToken()
	return rl
}

func (rl *RateLimiter) refillToken() {
	timer := time.NewTicker(rl.interval)
	defer timer.Stop()

	// refill
	for {
		<-timer.C
		select {
		case rl.token <- true:
		default:
		}

	}

}

func (rl *RateLimiter) IsAllowed() bool {
	select {
	case <-rl.token:
		return true
	default:
		return false
	}
}

func (rl *RateLimiter) MakeCall() error {
	if !rl.IsAllowed() {
		return errors.New("rate limit exceeded")
	}

	println("200 ok")
	return nil
}

func main() {
	rate := NewLimiter(2, 10)

	for i := 0; i < 1000; i += 1 {
		err := rate.MakeCall()
		if err != nil {
			println(err.Error())
		}
		time.Sleep(time.Millisecond * 200)
	}
}
