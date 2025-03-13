package minrl

import "testing"

type testCase struct {
    TestName string
    Input string
    Expected string
}

var testCases = []testCase{
    {
        TestName: "all R",
        Input: "RRRRRR",
        Expected: "0123456",
    },
    {
        TestName: "all L",
        Input: "LLLLLL",
        Expected: "6543210",
    },
    {
        TestName: "all =",
        Input: "======",
        Expected: "0000000",
    },
    {
        TestName: "1",
        Input: "LLRR=",
        Expected: "210122",
    },
    {
        TestName: "2",
        Input: "==RLL",
        Expected: "000210",
    },
    {
        TestName: "3",
        Input: "=LLRR",
        Expected: "221012",
    },
    {
        TestName: "4",
        Input: "RRL=R",
        Expected: "012001",
    },
    {
        TestName: "5",
        Input: "RRRL===LLL",
        Expected: "01243333210",
    },
}

func TestSolver(t *testing.T) {
    RunTestCases(t, testCases)
}

func RunTestCases(t *testing.T, testCases []testCase) {
    for _, testCase := range testCases {
        t.Run(testCase.TestName, func(t *testing.T) {
            actual := SolveMinRL(testCase.Input)
            if actual != testCase.Expected {
                t.Errorf("test case [%s] fail, expected %s, got %s", testCase.TestName, testCase.Expected, actual)
            }
        })
    }
}