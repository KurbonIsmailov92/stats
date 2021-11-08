package stats

import (
	"fmt"

	"github.com/KurbonIsmailov92/bank/v2/pkg/types"
)

func ExampleAvg() {

	payments := []types.Payment{
		{
			ID: 123,
			Amount: 10,
			Category: "food",
			Status: types.StatusOk,
		},
		{
			ID: 124,
			Amount: 20,
			Category: "tax",
			Status: types.StatusOk,
		},
		{
			ID: 125,
			Amount: 15,
			Category: "car",
			Status: types.StatusFail,
		},
		{
			ID: 126,
			Amount: 10,
			Category: "food",
			Status: types.StatusOk,
		},

	}

	result:= Avg(payments)

	fmt.Println(result)

	// Output: 13
}

func ExampleTotalInCategory() {

	payments := []types.Payment{
		{
			ID:       123,
			Amount:   10,
			Category: "food",
			Status: types.StatusOk,
		},
		{
			ID:       124,
			Amount:   20,
			Category: "tax",
			Status: types.StatusInProgress,
		},
		{
			ID:       125,
			Amount:   15,
			Category: "car",
			Status: types.StatusOk,
		},
		{
			ID:       126,
			Amount:   10,
			Category: "food",
			Status: types.StatusOk,
		},
	}

	result := TotalInCategory(payments, "food")

	fmt.Println(result)

	// Output: 20
}