package tree

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Config struct {
	Array []int `json:"array"`
}

// 从json文件中读取数组
func readArray(fileName string) Config {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Unable to load config file!")
		return Config{}
	}

	var config Config
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		fmt.Println("JSON decode error!")
		return Config{}
	}

	return config
}

// BinaryTreeGenerator 按json文件中的数组生成二叉树
func BinaryTreeGenerator(fileName string) *TreeNode {
	config := readArray(fileName)
	head := &TreeNode{config.Array[0], nil, nil}
	queue := []*TreeNode{head}
	var temp *TreeNode
	index := 1
	for index < len(config.Array) {
		n := len(queue)
		for i := 0; i < n; i++ {
			temp = queue[0]
			queue = queue[1:]
			if index >= len(config.Array) {
				break
			}
			if config.Array[index] != 1001 {
				temp.Left = &TreeNode{config.Array[index], nil, nil}
				queue = append(queue, temp.Left)
			}
			index++
			if index >= len(config.Array) {
				break
			}
			if config.Array[index] != 1001 {
				temp.Right = &TreeNode{config.Array[index], nil, nil}
				queue = append(queue, temp.Right)
			}
			index++
		}
	}
	return head
}
