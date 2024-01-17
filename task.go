// package main

// import "fmt"

// func main() {
// 	// определите переменные ver, id, pi
// 	// ...
// 	var (
// 		ver = "v0.0.1"
// 		id  = 0
// 		pi  = 3.1415
// 	)

// 	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
// }

// package main

// import "fmt"

// type T int

// const (
// 	one T = iota*2 + 1 // укажите здесь формулу с iota
// 	three
// 	five
// 	seven
// 	nine
// 	eleven
// )

// func main() {
// 	fmt.Println(one, three, five, seven, nine, eleven)
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)

// 	switch a := date.Year(); {
// 	case a > 1946 && a < 1964:
// 		fmt.Println("Привет, бумер!")
// 	case a > 1965 && a < 1980:
// 		fmt.Println("Привет, представитель X!")
// 	case a > 1981 && a < 1996:
// 		fmt.Println("Привет, миллениал!")
// 	case a > 1997 && a < 2012:
// 		fmt.Println("Привет, зумер!")
// 	case a > 2013:
// 		fmt.Println("Привет, альфа!")
// 	}
// }

package main

import "fmt"

func main() {
	num := 1
	for i := 1; i < 101; i++ {
		switch {
		case num%3 == 0:
			fmt.Println("Fizz")
			num++
		case num%5 == 0:
			fmt.Println("Buzz")
			num++
		default:
			fmt.Println(num)
			num++
		}
	}

}
