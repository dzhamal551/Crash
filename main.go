package main

import "time"

func main() {
	car := &car{
		model: "Granta",
		color: "Синий",
		crash: false,
		speed: 100,
	}
	car.acceleration()
	time.Sleep(5 * time.Second)
	car.crashed()
}
