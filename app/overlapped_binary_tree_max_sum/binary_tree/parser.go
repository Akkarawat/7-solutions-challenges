package binarytree

import "fmt"

func FromStackedList(stacks [][]int) (topNode *Node, err error) {
    if len(stacks) == 0 {
        return nil, nil
    }
    nodeStacks := [][]*Node{}
    for i, stack := range stacks {
        stackLength := len(stack)
        if stackLength != i + 1 {
            return nil, fmt.Errorf("stack %d should be of length %d", i+1, i + 1)
        }
        nodeStacks = append(nodeStacks, []*Node{})
        for j, value := range stack {
            newNode := NewNode(value)
            nodeStacks[i] = append(nodeStacks[i], &newNode)
            if i > 0 {
                if j < len(stack) - 1 {
                    nodeStacks[i-1][j].AddLeft(&newNode)
                }
                if j > 0 {
                    nodeStacks[i-1][j-1].AddRight(&newNode)
                }
            }
        }
    }
    return nodeStacks[0][0], nil
}