package overlappedbinarytreemaxsum

import (
	"testing"
)

type testCase struct {
    TestName string
    Input    [][]int
    Expected int
}

var testCases = []testCase{
    {
        TestName: "0 level",
        Input:    [][]int{},
        Expected: 0,
    },
    {
        TestName: "1 level",
        Input:    [][]int{{5}},
        Expected: 5,
    },
    {
        TestName: "2 levels",
        Input:    [][]int{{1}, {2, 3}},
        Expected: 4, // 1 -> 3
    },
    {
        TestName: "3 levels",
        Input:    [][]int{{1}, {2, 3}, {4, 5, 6}},
        Expected: 10, // 1 -> 3 -> 6
    },
    {
        TestName: "Negative values",
        Input:    [][]int{{-10}, {-20, -30}, {-1, -2, -3}},
        Expected: -31, // -10 -> -20 -> -1
    },
    {
        TestName: "Mixed values",
        Input:    [][]int{{10}, {-2, 7}, {8, -4, 3}},
        Expected: 20, // 10 -> 7 -> 3
    },
    {
        TestName: "hard",
        Input:    HARD_INPUT,
        Expected: 7273,
    },
}

func TestFindBTreeMaxSum(t *testing.T) {
    RunTestCasesForNormalSolver(t, testCases)
}

func TestFindFastBTreeMaxSum(t *testing.T) {
    RunTestCasesForFastSolver(t, testCases)
}


func RunTestCasesForNormalSolver(t *testing.T, testCases []testCase) {
    for _, testCase := range testCases {
        t.Run(testCase.TestName, func(t *testing.T) {
            actual, _ := FindBTreeMaxSum(testCase.Input)
            if actual != testCase.Expected {
                t.Errorf("test case [%s] fail, expected %d, got %d", testCase.TestName, testCase.Expected, actual)
            }
        })
    }
}

func RunTestCasesForFastSolver(t *testing.T, testCases []testCase) {
    for _, testCase := range testCases {
        t.Run(testCase.TestName, func(t *testing.T) {
            actual := FindFastBTreeMaxSum(testCase.Input)
            if actual != testCase.Expected {
                t.Errorf("test case [%s] fail, expected %d, got %d", testCase.TestName, testCase.Expected, actual)
            }
        })
    }
}