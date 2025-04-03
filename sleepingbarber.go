// Team: Robert Sullivan and Cole Vita
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// use atomic counter to prevent race conditions where multiple goroutines
// try to increment this global variable at the same time
var customerId atomic.Uint64

type Customer struct {
	id uint64
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("Opening barbershop...")

	walkins := make(chan Customer)
	// create a buffered channel for waiting room with capacity 6
	waiting := make(chan Customer, 6)
	chair := make(chan Customer)

	go receptionist(walkins, waiting)
	go waitingRoom(waiting, chair)
	go barber(chair)
	

	go func() {
		for {
			// customer arrival is a random event [0, 3) seconds
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

			// launches a goroutine for each customer
			go func() {
				customer := Customer{id: customerId.Add(1)}
				// fmt.Printf("Customer %d enters shop\n", customer.id)
				walkins <- customer
			}()
		}
	}()

	// keeps main goroutine alive "forever"
	// control + c to exit
	for {}
 }

 func receptionist(walkins chan Customer, waiting chan Customer) {
	// continuously checks for stream of customers entering the barbershop
	for customer := range walkins {
		select {
		// if send is not blocked, the customer is sent to the waiting room
		case waiting <- customer:
			fmt.Printf("Customer %d in waiting room\n", customer.id)
		// otherwise, the buffered channel waiting is full and customer is turned away
		default:
			fmt.Printf("Customer %d is turned away\n", customer.id)
		}
	}
}

func waitingRoom(waiting chan Customer, chair chan Customer) {
	// continuously try to send customer to the barber
	// if the barber is idle, the send succeeds
	for customer := range waiting {
		chair <- customer
	}
}

func barber(chair chan Customer) {
	for customer := range chair {
		fmt.Printf("Barber begins haircut for Customer %d\n", customer.id)
		// haircut is a random event [0, 5) seconds
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Printf("Barber finished haircut for Customer %d\n", customer.id)
	}
}