package stats

import (
	"github.com/Iftikhor99/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		types.Payment {
			Amount: 10_000_00,
			Status: types.StatusFail,
		},
		types.Payment {
			Amount: 20_000_00,
			Status: types.StatusOk,
		},
		types.Payment {
			Amount: 30_000_00,
			Status: types.StatusOk,
		},
	}
	fmt.Println(Avg(payments))
	//Output:2500000
	
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		types.Payment {
			Amount: 10_000_00,
			Category: "auto",
			Status: types.StatusFail,
		},
		types.Payment {
			Amount: 20_000_00,
			Category: "auto",
			Status: types.StatusOk,
		},
		types.Payment {
			Amount: 30_000_00,
			Category: "restaurant",
			Status: types.StatusOk,
		},
	}
	fmt.Println(TotalInCategory(payments, "auto"))
	//Output:2000000
	
}

func ExampleCategoriesAvg() {

payments := []types.Payment{
	types.Payment {
		Amount: 10_000_00,
		Category: "auto",
		Status: types.StatusFail,
	},
	types.Payment {
		Amount: 20_000_00,
		Category: "auto",
		Status: types.StatusOk,
	},
	types.Payment {
		Amount: 30_000_00,
		Category: "restaurant",
		Status: types.StatusOk,
	},
}
fmt.Println(CategoriesAvg(payments))
//map[auto:2000000 restaurant:3000000]

}

func ExamplePeriodsDynamic() {

	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	second := map[types.Category]types.Money{
		"auto": 5,
		"food": 3,
	}
	
	fmt.Println(PeriodsDynamic(first,second))
	//map[auto:2000000 restaurant:3000000]
	
}