package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/m/internal/app"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		if data[0] == "close" || len(data) == 3 {
			switch data[0] {
			case "currency_rates":
				result, err := app.GetExchangeRate(data)
				if err != nil {
					fmt.Printf("failed to get exchange rate: %v\n", err)
					continue
				}
				fmt.Printf("%s (%s): %s \n", result[0], result[1], result[2])
			case "close":
				return
			default:
				fmt.Println("Wrong command!")
			}
		} else {
			fmt.Println("Wrong command!")
		}
	}
}
