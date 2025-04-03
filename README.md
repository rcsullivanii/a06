# Assignment 7: Sleeping Barber with Go Processes

## Team
Robert Sullivan and Cole Vita

## Assignment Details 
Write a Go program to solve the Sleeping Barber problem. No need in the Go version to worry about hot-swapping code, but follow the other specs for the previous Elixir assignment (Assn. 5)
- Make each new customer a goroutine
- Make the waiting room a goroutine
- Make the receptionist a goroutine
  - The receptionist greets each arriving customer
  - If there is room to wait, the receptionist sends the customer to the waiting room; if no room, the customer is turned away.
- Make the barber a goroutine

This basic structure gives a lot of places that behavior can be parameterized or varied. Some initial parameters/behaviors:
- The waiting room has 6 waiting chairs
- The barber has one cutting chair
- Customers arrive at the shop at random times
- The barber takes a random time to complete each haircut
- The wait room is a FIFO queue
- Customers who wait stay until they are served
- The simulation operates "forever" generating new customers

## Design Rationale
What decisions did you make to get it running? Explain your overall design.

An atomic counter is used for to prevent race conditions when multiple goroutines try to increment the global customer ID at the same time. This is the only instance of "shared data" in the file.

The select statement used in the receptionist goroutine allows us to attempt to send a value and if send is blocked, then default statement gets executed. This works because we know if the send is blocked, there must already be 6 Customers in the buffered channel and therefore a full waiting room. Otherwise, the customer is added to the waiting room.

In our waiting room FIFO, we attempt to send the first customer in the waiting channel into the unbuffered "chair" channel. This send operation will block whenthe barber is busy (i.e. sleeping while simulating a haircut). It will, however, continue to try to send until eventually the barber completes the haircut and is ready for the next customer.

New customers are spawned randomly between 0-3 seconds and haircuts take some random time between 0-5 seconds. This is to show proper receptionist behavior when there is acculumation in the waiting room.

Adding for{} to the end of the main goroutine ensures that the main program and the goroutines that it has spawned continue until eventually the user exits the program using control + c.

## Sources
https://go.dev/tour/flowcontrol/4

https://gobyexample.com/atomic-counters 

https://go.dev/tour/basics/11

https://gobyexample.com/channels

https://gobyexample.com/channel-buffering 

https://go.dev/tour/concurrency/6 
