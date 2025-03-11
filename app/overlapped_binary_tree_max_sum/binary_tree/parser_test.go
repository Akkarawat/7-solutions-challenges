package binarytree

import "testing"

type testCase struct {
    TestName string
    Input    [][]int
    Expected *Node
}

func TestNodeParser(t *testing.T) {
    testCases := []testCase{
        {
            TestName: "0 stack",
            Input: [][]int{},
            Expected: nil,
        },
        {
            TestName: "1 stack",
            Input: [][]int{
                {1},
            },
            Expected: &Node{
                value: 1,
            },
        },
        {
            TestName: "2 stacks",
            Input: [][]int{
                {1},
                {2, 3},
            },
            Expected: &Node{
                value: 1,
                left: &Node{
                    value: 2,
                },
                right: &Node{
                    value: 3,
                },
            },
        },
        {
            TestName: "3 stacks",
            Input: [][]int{
                {1},
                {2, 3},
                {4, 5, 6},
            },
            Expected: &Node{
                value: 1,
                left: &Node{
                    value: 2,
                    left: &Node{
                        value: 4,
                    },
                    right: &Node{
                        value: 5,
                    },
                },
                right: &Node{
                    value: 3,
                    left: &Node{
                        value: 5,
                    },
                    right: &Node{
                        value: 6,
                    },
                },
            },
        },
    }

    for _, tc := range testCases {
        t.Run(tc.TestName, func(t *testing.T) {
            actual, err := FromStackedList(tc.Input)
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            if !NodesEqual(actual, tc.Expected) {
                t.Fatalf("expected %v, got %v", tc.Expected, actual)
            }
        })
    }
}

