package stats

import 	"github.com/Iftikhor99/bank/v2/pkg/types"

var status = types.StatusFail
//Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money {
	average := types.Money(0)
	length := len(payments)
	for _, payment := range payments {
		if payment.Status == status {
			length = length -1
			continue
		}
		average += payment.Amount	
	}
	
	average = average / types.Money(length)
	return average
} 

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sum := types.Money(0)
		
	for _, payment := range payments {
		if payment.Status == status {
			continue
		}
		if payment.Category == category {
			sum += payment.Amount	
		}
	}	
	return sum
} 

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	
	categoriesTotal := map[types.Category]types.Money{}
	categoriesCount := map[types.Category]types.Money{}
	categoriesAverage := map[types.Category]types.Money{}

	for _, payment := range payments {
		if payment.Status == status {
			continue
		}	
		categoriesTotal[payment.Category] += payment.Amount
		categoriesCount[payment.Category] += 1
	}

	
	for key, category := range categoriesTotal {
		categoriesAverage[key] = category/categoriesCount[key]
		
	}

	return categoriesAverage
}