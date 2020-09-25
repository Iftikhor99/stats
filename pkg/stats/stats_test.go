package stats

import (
	"reflect"
	"testing"
	"github.com/Iftikhor99/bank/v2/pkg/types"	
)

func TestPeriodsDynamic(t *testing.T) {
	
	first := map[types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}

	second := map[types.Category]types.Money{
		"auto": 10,
		"food": 25,
		"mobile": 5,
	}

	expected := map[types.Category]types.Money{
		"auto": 0,
		"food": 5,
		"mobile": 5,
	}
	result := PeriodsDynamic(first,second)

	if !reflect.DeepEqual(expected,result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
		
}