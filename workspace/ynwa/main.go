package concurrency

import (
	"log"
	"os"
	"time"
)

// startGoroutineFn представляет сигнатуру функции для запуска горутин
type startGoroutineFn func(done <-chan interface{}, pulseInterval time.Duration) (heartbeat <-chan interface{})

// newSteward создает и возвращает функцию, которая мониторит здоровье горутины (`startGoroutineFn`)
// и перезапускает её, если она становится нездоровой
func newSteward(timeout time.Duration, startGoroutine startGoroutineFn) startGoroutineFn {
	return func(done <-chan interface{}, pulseInterval time.Duration) <-chan interface{} {
		heartbeat := make(chan interface{})

		go func() {
			defer close(heartbeat)

			var wardDone chan interface{}
			var wardHeartbeat <-chan interface{}

			startWard := func() {
				wardDone = make(chan interface{})
				wardHeartbeat = startGoroutine(orDone(wardDone, done), timeout/2)
			}

			startWard()

			pulse := time.Tick(pulseInterval)

		monitorLoop:
			for {
				timeoutSignal := time.After(timeout)

				for {
					select {
					case <-pulse:
						select {
						case heartbeat <- struct{}{}:
						default:
						}
					case <-wardHeartbeat:
						continue monitorLoop
					case <-timeoutSignal:
						log.Println("steward: ward unhealthy; restarting")
						close(wardDone)
						startWard()
						continue monitorLoop
					case <-done:
						return
					}
				}
			}
		}()

		return heartbeat
	}
}

// NewStewardCaller демонстрирует использование паттерна steward с примером горутины doWork
func NewStewardCaller() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWorkWithSteward := newSteward(4*time.Second, doWork)

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() {
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) {
	}

	log.Println("Done")
}
