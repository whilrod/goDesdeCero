package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10)
	intde64 := int64(100)

	fmt.Println("Intcomun: ", intcomun)
	fmt.Println("Int de 32: ", intde32)
	fmt.Println("Int de 64: ", intde64)
}
