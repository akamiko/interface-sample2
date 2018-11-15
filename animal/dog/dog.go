package dog

import (
	"fmt"

	"github.com/akamiko/interface-sample7/animal"
)

// 犬の構造体定義
type Dog struct{}

// 犬のインスタンスを返す
func NewDog() animal.Animal {
	return new(Dog)
}

// 犬鳴く
func (d *Dog) Cry() {
	fmt.Println("wanwan")
}
