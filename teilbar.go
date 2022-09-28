package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	const suffix = ", "
	var number string
	flag.StringVar(&number, "n", "343", "Auf Teilbarkeit zu prüfende natürliche Zahl")
	flag.Parse()

	n, err := strconv.Atoi(number)
	if err != nil {
		log.Fatalf("ungültige Zahl %w", err)
	}
	if n < 0 {
		n = -n
	}

	divisible := [10]bool{} // mit make(map[uint]bool) haben keys zufällige Reihenfolge

	divisible[2] = isBy2(n)
	divisible[5] = isBy5(n)
	divisible[9] = isBy9(n)
	divisible[3] = divisible[9] || isBy3(n)
	divisible[6] = divisible[2] && divisible[3]

	fmt.Printf("Die Zahl %d wird von ", n)
	builder := strings.Builder{}
	any := false
	for k, divides := range divisible {
		if divides {
			any = true
			builder.WriteString(fmt.Sprintf("%d%s", k, suffix))
		}
	}
	if !any {
		builder.WriteString("2, 3, 5, 6, 9 nicht")
	}
	fmt.Printf("%s geteilt.\n", strings.TrimSuffix(builder.String(), suffix))
}

func isBy2(n int) bool {
	return n&1 == 0
}

func isBy5(n int) bool {
	num := []byte(fmt.Sprint(n))
	last := num[len(num)-1]
	return last == '5' || last == '0'
}

func isBy9(n int) bool {
	s := quersumme(n)
	return s == 9 || (s > 9 && isBy9(s))
}

func isBy3(n int) bool {
	s := quersumme(n)
	return s == 3 || s == 6 || s == 9 || (s > 9 && isBy3(s))
}

func quersumme(n int) int {
	num := []byte(fmt.Sprint(n))
	const zero = byte('0')
	var sum int = 0
	for _, d := range num {
		sum += int(d - zero)
	}
	return sum
}
