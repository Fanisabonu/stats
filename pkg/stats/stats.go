package stats
import (
	"github.com/Fanisabonu/bank/pkg/types"
)

//Avg -
func Avg(payments []types.Payment)(money types.Money) {
	for _, payment := range payments {
		money += payment.Amount
	}
	return money / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) (money types.Money) {
	for _, payment := range payments {
		if payment.Category == category {
			money += payment.Amount
		}
	}
	return
}
