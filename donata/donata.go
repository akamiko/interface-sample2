package donata

import (
	"github.com/akamiko/interface-sample7/animal"
)

type Donata struct {
	animal animal.Animal
}

func NewDonata(d animal.Animal) *Donata {
	return &Donata{
		animal: d,
	}
}

func (d *Donata) Hoge() {
	d.animal.Cry()
}
