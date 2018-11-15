package main

import (
	// 動物 >> 猫の構造体をimport
	"github.com/akamiko/interface-sample7/animal/cat"
	"github.com/akamiko/interface-sample7/donata"
)

func main() {
	c := cat.NewCat()
	donata := donata.NewDonata(c)
	donata.Hoge()
	/*
		l := fuga.NewLogger()
		foo := NewFoo(l)
		foo.Hoge()
		// 猫のインスタンス生成
		c := cat.NewCat()
		// 猫鳴く
		c.Cry()

		// 犬のインスタンス生成
		d := dog.NewDog()
		// 犬鳴く
		d.Cry()
	*/
}
