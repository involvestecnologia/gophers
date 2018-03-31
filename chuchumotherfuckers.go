package main

import (
	"fmt"
	"time"
)

var sensor1 = false

func main() {
	if sensor1 {
		fmt.Println("A estrada está fechada ** TIN DIN DIN DIN DIN DIN ** ** CHOOO CHOOOO MOTHERFUCKERS **")
	}

	fmt.Println("A estrada está liberada!")
	time.Sleep(1 * time.Second)
	fmt.Println("E la vem o trem!")
	for counter := 0; counter < 15; counter++ {
		if counter != 14 {
			sensor1 = true
			fmt.Printf("A estrada está fechada para o vagão %d passar\n", counter+1)
			time.Sleep(1 * time.Second)
		} else {
			sensor1 = false
			fmt.Println("A estrada está liberada!")
		}
	}

}
