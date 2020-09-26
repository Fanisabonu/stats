package stats

import (
	"github.com/Fanisabonu/bank/v2/pkg/types"
)

//Avg -
func Avg(payments []types.Payment) (money types.Money){
	c :=0 
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			c++ 
			money += payment.Amount
		}
	}
	return money / types.Money(c)
}

func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
		if payment.Category == category {
			if payment.Status != types.StatusFail {

				money += payment.Amount
			}
		}
	}
	return
}
