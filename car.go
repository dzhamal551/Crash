package main

import (
	"fmt"
	"strconv"
)

type car struct {
	model string
	color string
	crash bool
	speed int
}

func (м *car) acceleration() {
	fmt.Println("скорость на спидометре:", strconv.Itoa(м.speed))
	fmt.Println("разгон машины на 100км/ч")
	м.speed += 100
	fmt.Println("новая скорость:", strconv.Itoa(м.speed))
	fmt.Println("машину начало заносить......")
}
func (м *car) crashed() {
	fmt.Println("дитэкъой разбилась :(")
	м.crash = true
}
