package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Romiz01/bank1/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Book",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "dog",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "Cats",
			Status:   types.StatusFail,
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID:       2,
			Amount:   53_00,
			Category: "Merve",
			Status:   types.StatusOk,
		},
		{
			ID:       1,
			Amount:   51_00,
			Category: "BBQ",
			Status:   types.StatusOk,
		},
		{
			ID:       3,
			Amount:   52_00,
			Category: "smart",
			Status:   types.StatusFail,
		},
	}

	inCategory := types.Category("Merve")
	totalInCategory := TotalInCategory(payments, inCategory)
	fmt.Println(totalInCategory)
	//Output:  5300

}

func TestCategoriesAvgUser(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 1_000_00},
		{ID: 1, Category: "food", Amount: 2_000_00},
		{ID: 1, Category: "auto", Amount: 3_000_00},
		{ID: 1, Category: "auto", Amount: 4_000_00},
		{ID: 1, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto": 266666,
		"food": 2_000_00,
		"fun":  5_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual %v", expected, result)
	}
}

func TestCategoriesAvgUser_nil(t *testing.T) {
	var payments []types.Payment
	res := CategoriesAvg(payments)

	if len(res) != 0 {
		t.Errorf("\n got > %v want > nil", res)
	}
}
func TestCategoriesAvg_one(t *testing.T) {
	payments := []types.Payment{
		{
			ID:       1,
			Category: "cafe",
			Amount:   100_00,
		},
		{
			ID:       2,
			Category: "cafe",
			Amount:   100_00,
		},
	}
	expected := map[types.Category]types.Money{
		"cafe": 100_00,
	}

	res := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, res) {
		t.Errorf("\n got > %v want > nil", res)
	}
}

func TestPeriodsDynamicUser_OneMoreElem(t *testing.T) {
	first := map[types.Category]types.Money{
		"cafe": 20,
		"auto": 14,
	}
	second := map[types.Category]types.Money{
		"cafe":   35,
		"auto":   17,
		"mobile": 17,
	}
	amount := map[types.Category]types.Money{
		"cafe":   15,
		"auto":   3,
		"mobile": 17,
	}

	got := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(amount, got) {
		t.Errorf("\n got > %v \n amount > %v", got, amount)
	}
}
