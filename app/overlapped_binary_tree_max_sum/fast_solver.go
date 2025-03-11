package overlappedbinarytreemaxsum

func FindFastBTreeMaxSum (stacks [][]int) int {
    if len(stacks) == 0 {
        return 0
    }
    for i := len(stacks) - 2; i >= 0; i-- {
        for j := 0; j < len(stacks[i]); j++ {
            if stacks[i+1][j] > stacks[i+1][j+1] {
                stacks[i][j] += stacks[i+1][j]
            } else {
                stacks[i][j] += stacks[i+1][j+1]
            }
        }
    }
    return stacks[0][0]
}