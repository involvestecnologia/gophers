package trem

import "fmt"

type Trem struct {
	Vagoes     int
	maquinista string
}

func (t *Trem) SetMaquinistaName(name string) {
	t.maquinista = name
	fmt.Println("O nome do maquinista é " + t.maquinista)
}
