package iteraciones

import "fmt"

func Iterar() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i += 5 {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
