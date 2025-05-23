package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "Gopher hotel"
	fmt.Println("Hotel " + hotelName)

	var roomsAvailable uint8
	rand.Seed(time.Now().UnixNano())
	var rooms, roomsOccupied uint8 = 100, uint8(rand.Intn(100))
	roomsAvailable = rooms - roomsOccupied
	fmt.Println("Rooms available: ", roomsAvailable)

	fmt.Print("rooms > roomsOccupied: ")
	fmt.Println(rooms > roomsOccupied)

	if rooms > roomsOccupied {
		fmt.Println("There are rooms available")
	} else {
		fmt.Println("No rooms available")
	}
}
