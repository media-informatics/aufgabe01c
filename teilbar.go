package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	var number string
	flag.StringVar(&number, "n", "343", "Auf Teilbarkeit zu prüfende natürliche Zahl")
	flag.Parse()

	num := []byte(strings.TrimSpace(number))
	if len(num) == 0 {
		log.Fatalf("Keine gültige Zahl")
	}
	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	divisible := [10]bool{}
	divisible[2] = (n&1 == 0)

	last := num[len(num)-1]
	divisible[5] = last == '5' || last == '0'

	divisible[9] = isBy9(num)
	divisible[3] = divisible[9] || isBy3(num)
	divisible[6] = divisible[2] && divisible[3]

	fmt.Printf("Die Zahl %d wird ", n)
	any := false
	for _, divides := range divisible {
		if divides {
			any = true
		}
	}
	if any {
		fmt.Print("von ")
		for k, divides := range divisible {
			if divides {
				fmt.Printf("%d, ", k)
			}
		}
	} else {
		fmt.Print("nicht von 2, 3, 5, 6, oder 9 ")
	}
	fmt.Println("geteilt")
}

func isBy9(num []byte) bool {
	s := quersumme(num)
	if s < 10 {
		return s == 9
	}
	return isBy9([]byte(fmt.Sprint(s)))
}

func isBy3(num []byte) bool {
	s := quersumme(num)
	if s < 10 {
		return s == 3 || s == 6 || s == 9
	}
	return isBy3([]byte(fmt.Sprint(s)))
}

func quersumme(num []byte) uint {
	const zero = byte('0')
	var sum uint = 0
	for _, d := range num {
		sum += uint(d - zero)
	}
	return sum
}
