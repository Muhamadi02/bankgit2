package stats

import (
	"fmt"
	"github.com/Muhamadi02/bankgit2/pkg/bankgit2/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   2000_00,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   3000_00,
			Category: "medicine",
		},
		{
			ID:       3,
			Amount:   1000_00,
			Category: "restraunts",
		},
	}

	fmt.Println(Avg(payments))
	// Output:
	// 200000
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   2000_00,
			Category: "auto",
		},
		{
			ID:       2,
			Amount:   3000_00,
			Category: "medicine",
		},
		{
			ID:       3,
			Amount:   1000_00,
			Category: "restraunts",
		},
		{
			ID:       4,
			Amount:   5000_00,
			Category: "auto",
		},
	}

	fmt.Println(TotalInCategory(payments, "auto"))
	// Output:
	// 700000
}
