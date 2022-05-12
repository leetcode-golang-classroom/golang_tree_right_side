# golang_tree_right_side

Given the `root` of a binary tree, imagine yourself standing on the **right side**
 of it, return *the values of the nodes you can see ordered from top to bottom*.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2021/02/14/tree.jpg](https://assets.leetcode.com/uploads/2021/02/14/tree.jpg)

```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]

```

**Example 2:**

```
Input: root = [1,null,3]
Output: [1,3]

```

**Example 3:**

```
Input: root = []
Output: []

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 100]`.
- `100 <= Node.val <= 100`

## 解析

題目給了一個二元樹根結點 root

想要找到二元樹，每個level 最右邊的結點如下圖

![](https://i.imgur.com/M0L1xrK.png)

直覺的作法是透過 Breadth First Traversal 廣度優先演算法實作

廣度優先演算法透過一個 queue ，每次把同一層的 node 照順序(由左而右)放到 queue 內

然後這時最右邊的結點就是我們要尋找的結點了

而要繼續往下一層就是把 queue 內的結點 node 依次(從最左邊到右邊) shift 出來 ，把 shift 出來的 node 子結點再依序放入 queue

這樣等到 queue 為空時，整個tree 都走訪結束

因此時間複雜度是 O(n) ，空間複雜度也是 O(n)

![](https://i.imgur.com/LBC8UmF.png)

## 程式碼

```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		qLen := len(queue)
		var rightSide *TreeNode
		// loop over current queue find rightest
		for idx := 0; idx < qLen; idx++ {
			// shift left ele
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				// 更新
				rightSide = node
				queue = append(queue, node.Left, node.Right)
			}
		}
		if rightSide != nil {
			result = append(result, rightSide.Val)
		}
	}
	return result
}
```

## 困難點

1. 理解如何以 level order 來對 tree 做尋訪
2. 知道每次如何推進到下一層 sibling

## Solve Point

- [x]  Understand what problem to solve
- [x]  Analysis Complexity