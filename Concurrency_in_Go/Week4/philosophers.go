package main
import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct { sync.Mutex }

type Philo struct {
	id, count int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup, ch chan *Philo) {
	for i := 0; i < 3; i++ {
		ch <- &p
		
		if p.count < 3 {
			p.leftCS.Lock()
			p.rightCS.Lock()
	
			fmt.Println("starting to eat", p.id)
			p.count++
			fmt.Println("finishing eating", p.id)

			p.rightCS.Unlock()
			p.leftCS.Unlock()

			wg.Done()
		}
	}
}

func host(wg *sync.WaitGroup, ch chan *Philo) {
	for {
		if len(ch) == 2 {
			<- ch
			<- ch

			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan *Philo, 2)

	wg.Add(15)

	chopsticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i< 5; i++ {
		philos[i] = &Philo{i + 1, 0, chopsticks[i], chopsticks[(i+1)%5]}
	}

	go host(&wg, channel)

	for i := 0; i < 5; i++ {
		go philos[i].eat(&wg, channel)
	}
	
	wg.Wait()
}