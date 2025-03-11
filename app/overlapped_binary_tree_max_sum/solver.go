package overlappedbinarytreemaxsum

import binarytree "7-solutions-challenges/app/overlapped_binary_tree_max_sum/binary_tree"

func FindBTreeMaxSum (stacks [][]int) (int, error) {
    topNode, err := binarytree.FromStackedList(stacks)
    if err != nil {
        return 0, err
    }
    return topNode.GetMaxSum(), nil
}