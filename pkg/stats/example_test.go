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
