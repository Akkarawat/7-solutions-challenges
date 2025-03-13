package usecases

import (
	"7-solutions-challenges/internal/meat_count/services/mocks"
	"testing"
)

type testcase struct {
    TestName string
    Input    input
    Expected map[string]int
}

type input struct {
    MeatText string
    MeatAndFillerText string
}

var testCases = []testcase{
    {
        TestName: "Basic",
        Input: input{
            MeatText: "Bacon T-bone",
            MeatAndFillerText: "ipsum dolor amet, Bacon. Ipsum bacon dolor amet. T-bone t-bone. Ipsum dolor amet.",
        },
        Expected: map[string]int{
            "bacon": 2,
            "t-bone": 2,
        },
    },
}

func TestGetMeatCount(t *testing.T) {
    for _, testcase := range testCases {
        RunTestCase(t, testcase)
    }
}

func RunTestCase(t *testing.T, testcase testcase) {
    t.Run(testcase.TestName, func(t *testing.T) {
        mockService := mocks.NewMockMeatTextService(testcase.Input.MeatText, testcase.Input.MeatAndFillerText)
        uc := NewMeatCountUsecase(mockService)

        result, err := uc.GetMeatCount()
        if err != nil {
            t.Errorf("Error: %v", err)
        }

        if len(result) != len(testcase.Expected) {
            t.Errorf("Expected: %v, Got: %v", testcase.Expected, result)
        }

        for key, value := range result {
            if testcase.Expected[key] != value {
                t.Errorf("Expected: %v, Got: %v", testcase.Expected, result)
            }
        }
    })
}