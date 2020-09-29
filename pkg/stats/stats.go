package stats

import (
	"github.com/Fanisabonu/bank/v2/pkg/types"
)

//Avg -
func Avg(payments []types.Payment) (money types.Money) {
	c := 0
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
func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	moneys := map[types.Category]types.Money{}
	c := map[types.Category]types.Money{}
	for _, payment := range payments {
		if payment.Status != types.StatusFail {
			moneys[payment.Category] += payment.Amount
			c[payment.Category]++
		}
	}
	for key := range moneys {
		moneys[key] = moneys[key] / c[key]
	}
	return moneys
}

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	moneys := map[types.Category]types.Money{}
	for key, amount:= range first{
		moneys [key]-= amount
	}
	for key, amount:= range second{
		moneys [key]+= amount
	}
	return moneys
}
