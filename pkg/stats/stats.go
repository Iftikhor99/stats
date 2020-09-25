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

func PeriodsDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money{
	result := map[types.Category]types.Money{}
	mianKey := types.Category("")
	
	for key1 := range first {
		for key2 := range second {
			if key2 == key1 {
				continue				
			}		
			second[key1] += types.Money(0)
		}			
	}

	for key2 := range second {
		for key1 := range first {
			if key1 == key2 {
				continue				
			}		
			first[key2] += types.Money(0)
		}			
	}

	for key1, category1 := range first {
		mianKey = key1
		for key2, category2 := range second {
			if mianKey == key2 {
				result[mianKey] = category2 - category1
			}		
			continue
		}			
	}
	return result
}
