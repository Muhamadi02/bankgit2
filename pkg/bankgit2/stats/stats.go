package stats

import "github.com/Muhamadi02/bankgit2/pkg/bankgit2/types"

// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	sum := types.Money(0)
	t := 0
	for _, payment := range payments {
		sum += payment.Amount
		t++
	}
	return types.Money(sum / types.Money(t))
}

// TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}

	return sum
}
