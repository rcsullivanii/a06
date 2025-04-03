# Assignment 7: Sleeping Barber with Go Processes

## Team
Robert Sullivan and Cole Vita

## Assignment Details 
Write a Go program to solve the Sleeping Barber problem. No need in the Go version to worry about hot-swapping code, but follow the other specs for the previous Elixir assignment (Assn. 5)
- Make each new customer a goroutine
- Make a goroutine for the waiting room
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
