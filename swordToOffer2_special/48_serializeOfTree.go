package main

import (
	"LeetCode/utils"
	"strconv"
	"strings"
)

type Codec struct {
}

func ConstructorCodec() Codec {
	return Codec{}
}

// 前序遍历 用'*'隔开每个节点 叶子结点下的两个nil节点也会插入到字符串中
// Serializes a tree to a single string.
func (this *Codec) serialize(root *utils.TreeNode) string {
	sb := &strings.Builder{}
	var dfs func(*utils.TreeNode)
	dfs = func(root *utils.TreeNode) {
		if root == nil {
			sb.WriteString("nil*")
			return
		}
		sb.WriteString(strconv.Itoa(root.Val))
		sb.WriteString("*")
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return sb.String()
}

// 用'*'将字符串分割回原来的节点 递归构建树结构 因为原本树的叶子结点下的nil结点也在字符串中 所以根据nil来作为递归条件的判断即可
// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *utils.TreeNode {
	sl := strings.Split(data, "*")
	var build func() *utils.TreeNode
	build = func() *utils.TreeNode {
		if sl[0] == "nil" {
			sl = sl[1:]
			return nil
		}
		val, _ := strconv.Atoi(sl[0])
		sl = sl[1:]
		return &utils.TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}
