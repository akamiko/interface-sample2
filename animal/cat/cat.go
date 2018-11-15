package cat

import (
	"fmt"

	"github.com/akamiko/interface-sample7/animal"
)

// 猫の構造体定義
type Cat struct{}

// 猫のインスタンスを返す
func NewCat() animal.Animal {
	return new(Cat)
}

// 猫鳴く
func (c *Cat) Cry() {
	fmt.Println("nya-")
}
