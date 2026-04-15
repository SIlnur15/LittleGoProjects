package main

import "time"

func main() {
	type Context interface {
		Deadline() (deadline time.Time, ok bool)
		Done() <-chan struct{}
		Err() error
		Value(key any) any
	}
}
