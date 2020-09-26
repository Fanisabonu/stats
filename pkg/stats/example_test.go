package stats

import (
	"github.com/Fanisabonu/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   100,
			Category: "card",
		},
	}

	amount := Avg(payments)
	fmt.Println(amount)
	//Output:
	//100
}