package main

import "fmt"

// func Points(games []string) (score int) {
// return
// }

func main() {
	// score := 0

	games := []string{
		"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3",
	}
	for _, i := range games {
		d := []string{i}
		// x := 0
		for v, _ := range d {
			x := d[v]
			fmt.Println(x)
		}
		// fmt.Println(d, x)
		// if x > y {
		// 	score += 3
		// }
		// if x < y {
		// 	score += 1
		// }
	}

	fmt.Println(games)
}
