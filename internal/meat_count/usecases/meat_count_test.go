package usecases

import (
	"7-solutions-challenges/internal/meat_count/services/mocks"
	"reflect"
	"testing"
)

type testcase struct {
	TestName string
	Input    input
	Expected map[string]int
}

type input struct {
	MeatText         string
	MeatAndFillerText string
}

var testCases = []testcase{
	{
		TestName: "Basic",
		Input: input{
			MeatText:         "Bacon T-bone",
			MeatAndFillerText: "ipsum Bacon. T-bone bacon T-bone.",
		},
		Expected: map[string]int{
			"bacon": 2,
			"t-bone": 2,
		},
	},
	{
		TestName: "Mixed Case",
		Input: input{
			MeatText:         "Chicken Pork",
			MeatAndFillerText: "CHICKEN pork. Chicken pork chicken",
		},
		Expected: map[string]int{
			"chicken": 3,
			"pork": 2,
		},
	},
	{
		TestName: "Extra Separators",
		Input: input{
			MeatText:         "Beef",
			MeatAndFillerText: "....Beef;;;; beef  ..BEEF!?",
		},
		Expected: map[string]int{
			"beef": 3,
		},
	},
	{
		TestName: "No Meat Found",
		Input: input{
			MeatText:         "Lamb",
			MeatAndFillerText: "this is just some random text.",
		},
		Expected: map[string]int{},
	},
	{
		TestName: "Repeated Words",
		Input: input{
			MeatText:         "Duck",
			MeatAndFillerText: "Duck duck duck duck goose duck",
		},
		Expected: map[string]int{
			"duck": 5,
		},
	},
	{
		TestName: "Multiple Meats",
		Input: input{
			MeatText:         "Salmon Shrimp",
			MeatAndFillerText: "Salmon salmon. shrimp, shrimp! shrimp salmon.",
		},
		Expected: map[string]int{
			"salmon": 3,
			"shrimp": 3,
		},
	},
	{
		TestName: "Empty Input",
		Input: input{
			MeatText:         "",
			MeatAndFillerText: "",
		},
		Expected: map[string]int{},
	},
}

func TestGetMeatCount(t *testing.T) {
	for _, testcase := range testCases {
		t.Run(testcase.TestName, func(t *testing.T) {
			mockService := mocks.NewMockMeatTextService(testcase.Input.MeatText, testcase.Input.MeatAndFillerText)
			uc := NewMeatCountUsecase(mockService)

			// Execute function
			result, err := uc.GetMeatCount()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Use reflect.DeepEqual to compare entire map
			if !reflect.DeepEqual(result, testcase.Expected) {
				t.Errorf("Test %s failed.\nExpected: %v\nGot: %v", testcase.TestName, testcase.Expected, result)
			}
		})
	}
}
