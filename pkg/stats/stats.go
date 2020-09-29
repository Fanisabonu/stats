package stats

import (
	"reflect"
	"testing"
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
	for key:= range moneys  {
		moneys[key]=moneys[key]/c[key]
	}
	return moneys 
}

func TestCategoriesAvg_manyCategory(t *testing.T) {
	payments := []types.Payment{
	  {
		Category: "car",
		Amount:   0,
		Status:   types.StatusOk,
	  },
	  {
		Category: "car",
		Amount:   100,
		Status:   types.StatusInProgress,
	  },
	  {
		Category: "food",
		Amount:   10000,
		Status:   types.StatusOk,
	  },
	  {
		Category: "fun",
		Amount:   100,
		Status:   types.StatusOk,
	  },
	  {
		Category: "fun",
		Amount:   100,
		Status:   types.StatusFail,
	  },
	}
	want := map[types.Category]types.Money{
	  "car":  50,
	  "food": 10000,
	  "fun":  100,
	}
	got := CategoriesAvg(payments)
	if !reflect.DeepEqual(want, got) {
	  t.Errorf("want: %v, got: %v", want, got)
	}
  }