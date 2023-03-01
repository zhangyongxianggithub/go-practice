package main

import (
	"fmt"
)

func main() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
		8: "August",
		9: "September", 10: "October", 11: "November", 12: "December",
	}
	fmt.Println(months[1:])
	summer := months[6:9]
	Q2 := months[4:7]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
}
