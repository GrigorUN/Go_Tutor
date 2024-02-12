package main

import "fmt"

func main() {
	dim := 100
	s := make([]int, 0, dim)
	// заполняем слайс числами
	for i := 0; i < dim; i++ {
		s = append(s, i+1)
	}
	// оставляем первые и последние 10 элементов
	s = append(s[:10], s[dim-10:]...)
	dim = len(s)
	// разворачиваем слайс
	for i := range s[:dim/2] {
		s[i], s[dim-i-1] = s[dim-i-1], s[i]
	}
	fmt.Println(s)
}

// func main() {
// 	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
// 	workDaysSlice := weekTempArr[:5]
// 	weekendSlice := weekTempArr[3:]
// 	fromTuesdayToThursDaySlice := weekTempArr[1:6]
// 	weekTempSlice := weekTempArr[:6]

// 	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
// 	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
// 	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
// 	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func f(cnt *int) {
// 	*cnt++
// }

// func main() {
// 	// Получаем читателя пользовательского ввода
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Interaction counter")

// 	cnt := 0
// 	for {
// 		fmt.Print("-> ")
// 		// Считываем введённую пользователем строку. Программа ждёт, пока пользователь введёт строку
// 		_, err := reader.ReadString('\n')
// 		if err != nil {
// 			panic(err)
// 		}

// 		f(&cnt)

// 		fmt.Printf("User input %d lines\n", cnt)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	num := 1
// 	for i := 1; i < 101; i++ {
// 		switch {
// 		case num%3 == 0:
// 			fmt.Println("Fizz")
// 			num++
// 		case num%5 == 0:
// 			fmt.Println("Buzz")
// 			num++
// 		default:
// 			fmt.Println(num)
// 			num++
// 		}
// 	}

// }

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
