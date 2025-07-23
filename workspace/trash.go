package main

import (
	"fmt"
	"log"
)

type futureError float64

func (f futureError) Error() string {
	return fmt.Sprint("this value (%0.2f) is too hot", f)
}

func overHeatDefiner(now float64, base float64) error {
	if base < now {
		return futureError(now)
	}
	return nil
}

func main() {
	var err error = overHeatDefiner(55.0, 25.0)
	if err != nil {
		log.Fatal(err)
	}
}
