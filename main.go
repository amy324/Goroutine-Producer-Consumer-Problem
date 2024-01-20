package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <- ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza#%d. It will take %d seconds...\n", pizzaNumber, delay)
		// Delay
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** Order cannot be fulfilled as ingredients are out of stock for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** Order cannot be fulfilled due to staffing issues for pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready!", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}


func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza is being made
	var i = 0
	//continuously run until we recieve quit notification

	//try to make pizzas
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
				//data sent to data channel which symbolises that an attempt to make a pizza was made
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <- pizzaMaker.quit:
				//close channels 
				close(pizzaMaker.data)
				close(quitChan)
				//return exits go return
				return
			}
		}
	}
}

func main() {
	//seed the random number generator
	rand.NewSource(time.Now().UnixNano())


	//print the starting message
	color.Cyan("The Pizzeria is open for business")
	color.Cyan("---------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background
	go pizzeria(pizzaJob)

	//create and run consumer

	//print the ending message
}
