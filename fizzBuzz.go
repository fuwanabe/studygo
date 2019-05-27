package main

import "fmt"

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	i := 1
	for i <= 100 {
		switch {
		case i%15 == 0:
			fmt.Printf("%v : FizzBuzz\n", i)

		case i%3 == 0:
			fmt.Printf("%v : Fizz\n", i)
		case i%5 == 0:
			fmt.Printf("%v : Buzz\n", i)
		}

		i++
	}

}
