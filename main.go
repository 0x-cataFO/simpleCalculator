package main

import (
	"fmt"
	// "unicode/utf8"
)

func addition(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if x == 0 || y == 0 {
		fmt.Println("Undefined, cannot divide by zero")
		return 0
	}
	return x / y
}

func main() {

	/* REVERSE STRING INPUT
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	inputRune := []rune(input)
	var reverseRune []rune
	for i := len(inputRune) - 1; i >= 0; i-- {
		reverseRune = append(reverseRune, inputRune[i])
	}
	fmt.Printf("%v\n", string(reverseRune)) */

	// UNICODE STRING LENGTH, byte string length
	/* var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	bytelength := len(input)
	runeLength := utf8.RuneCountInString(input)
	fmt.Println("Rune length:", runeLength)
	fmt.Println("Byte length:", bytelength) */

	/* celcius to farenheit converter
	var input int
	fmt.Print("Enter Temperature in ∘C : ")
	fmt.Scanln(&input)
	farenheit := (input*9/5 + 32)
	fmt.Printf("%v∘C = %v∘F\n", input, farenheit) */
	logo := `
                       __                      __              __                         
                      /  |                    /  |            /  |                        
    _______   ______  $$ |  _______  __    __ $$ |  ______   _$$ |_     ______    ______  
   /       | /      \ $$ | /       |/  |  /  |$$ | /      \ / $$   |   /      \  /      \ 
  /$$$$$$$/  $$$$$$  |$$ |/$$$$$$$/ $$ |  $$ |$$ | $$$$$$  |$$$$$$/   /$$$$$$  |/$$$$$$  |
  $$ |       /    $$ |$$ |$$ |      $$ |  $$ |$$ | /    $$ |  $$ | __ $$ |  $$ |$$ |  $$/ 
  $$ \_____ /$$$$$$$ |$$ |$$ \_____ $$ \__$$ |$$ |/$$$$$$$ |  $$ |/  |$$ \__$$ |$$ |      
  $$       |$$    $$ |$$ |$$       |$$    $$/ $$ |$$    $$ |  $$  $$/ $$    $$/ $$ |      
   $$$$$$$/  $$$$$$$/ $$/  $$$$$$$/  $$$$$$/  $$/  $$$$$$$/    $$$$/   $$$$$$/  $$/       
`

	for {
		fmt.Println(logo)
		fmt.Println("1 for Addition")
		fmt.Println("2 for Subtraction")
		fmt.Println("3 for Division")
		fmt.Println("4 for Multiplication")
		fmt.Println("5 to Exit calculator")
		var choice int
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		if choice == 5 {
			break
		}

		var num1 float64
		fmt.Print("Enter first number: ")
		fmt.Scanln(&num1)

		var num2 float64
		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		switch choice {
		case 1:
			fmt.Printf("%v added to %v = %v\n", num1, num2, addition(num1, num2))
		case 2:
			fmt.Printf("%v subtracted from %v = %v\n", num2, num1, subtract(num1, num2))
		case 3:
			fmt.Printf("%v divided by %v = %v\n", num1, num2, divide(num1, num2))
		case 4:
			fmt.Printf("%v multiplied by %v = %v\n", num2, num1, multiply(num1, num2))
		default:
			fmt.Printf("Undefined\n")
		}

	}

}
