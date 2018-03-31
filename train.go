package main

import (
	"fmt"
	"time"

	. "./trem"
)

type Cancela struct {
	Som    string
	Status string
}

type Sensor struct {
	Ativado bool
}

var (
	cancela Cancela = Cancela{"ding ding ding", "fechada"}
	trem    Trem    = Trem{Vagoes: 15}
	sensor          = Sensor{false}
)

func (c *Cancela) FecharCancela() {
	c.Status = "fechada"
	fmt.Println("A cancela foi fechada!")
}

func (c *Cancela) AbrirCancela() {
	c.Status = "Aberta"
	fmt.Println("Cancela foi aberta!")
}

func (c *Cancela) TocarSom() {
	fmt.Println(c.Som)
}

func (s *Sensor) AtivarSensor() {
	s.Ativado = true
	fmt.Println("O sensor foi ativado!")
}

func (s *Sensor) DesativarSensor() {
	s.Ativado = false
	fmt.Println("O sensor foi desativado!")
}

func main() {

	trem.SetMaquinistaName("joao")
	for contador := 0; contador < trem.Vagoes; contador++ {
		if contador == 0 {
			sensor.AtivarSensor()
			cancela.FecharCancela()
			cancela.TocarSom()
		}

		if contador != 14 {
			fmt.Printf("A estrada está fechada para o vagão %d passar\n", contador+1)
			time.Sleep(1 * time.Second)
		} else {
			sensor.DesativarSensor()
			cancela.AbrirCancela()
			fmt.Println("A estrada está liberada!")
		}
	}
}
