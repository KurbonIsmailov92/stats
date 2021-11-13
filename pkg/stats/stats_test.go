package stats

import (
	"testing"

	"github.com/KurbonIsmailov92/bank/v2/pkg/types"
)


func TestFilterByCtegory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")
	
	if len(result) != 0 {
		t.Error("result len  !=0"  )
	}
}

func TestFilterByCtegory_empty(t *testing.T) {
	payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len  !=0"  )
	}
}