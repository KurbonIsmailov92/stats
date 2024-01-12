package stats

import (
	"log"

	"github.com/KurbonIsmailov92/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)
	count := int(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		sum += payment.Amount
		count++
	}

	return sum / types.Money(count)

}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {

		if payment.Status == types.StatusFail {
			continue
		}

		if payment.Category == category {
			sum += payment.Amount
		}
	}
	log.Print(sum)
	return sum
}

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment  {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
