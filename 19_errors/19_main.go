package main

import(
	"fmt"
	"errors"
	"math/rand/v2"
)

var stockError = fmt.Errorf("There are not more units of the item")
var purchaseError = fmt.Errorf("Error during purchase process")

var currentQty = 3


func main() {
	for iter := range(5) {
		err := processPurchase(uint8(iter))

		if err != nil {
			if errors.Is(err, stockError) {
				fmt.Println("Out of stock error", err)
				} else if errors.Is(err, purchaseError) {
				fmt.Println("Payment process error", err)
			}
			continue
		}

		fmt.Println("Payment success")

	}
}

func processPurchase(items uint8) error {

	if items > uint8(currentQty) {
		return stockError
	}

	rnd := rand.IntN(10)

	if rnd % 2 == 0 {
		return fmt.Errorf("Processing purchase... %w", purchaseError)
	}

	return nil
}