package stats

import (
	"reflect"
	"testing"
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

  func TestPeriodsDynamic_manyParam(t *testing.T) {
	first, second := map[types.Category]types.Money{
	  "auto": 10,
	  "food": 20,
	}, map[types.Category]types.Money{
	  "auto":   10,
	  "food":   25,
	  "mobile": 5,
	}
	want := map[types.Category]types.Money{
	  "auto":   0,
	  "food":   5,
	  "mobile": 5,
	}
	got := PeriodsDynamic(first, second)
	if !reflect.DeepEqual(want, got) {
	  t.Errorf("want: %v, got: %v", want, got)
	}
  }