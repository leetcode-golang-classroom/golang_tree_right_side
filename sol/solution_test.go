package sol

import (
	"reflect"
	"strconv"
	"testing"
)

func CreateBinaryTree(input *[]string) *TreeNode {
	var tree *TreeNode
	arr := *input
	result := make([]*TreeNode, len(arr))
	for idx, val := range arr {
		if val != "null" {
			num, _ := strconv.Atoi(val)
			result[idx] = &TreeNode{Val: num}
		}
		if idx == 0 {
			tree = result[idx]
		}
	}
	for idx := range result {
		if 2*idx+1 < len(result) {
			result[idx].Left = result[2*idx+1]
		}
		if 2*idx+2 < len(result) {
			result[idx].Right = result[2*idx+2]
		}
	}
	return tree
}
func Test_rightSideView(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "root = [1,2,3,null,5,null,4]",
			args: args{root: CreateBinaryTree(&[]string{"1", "2", "3", "null", "5", "null", "4"})},
			want: []int{1, 3, 4},
		},
		{
			name: "root = [1,null,3]",
			args: args{root: CreateBinaryTree(&[]string{"1", "null", "3"})},
			want: []int{1, 3},
		},
		{
			name: "root = []",
			args: args{root: CreateBinaryTree(&[]string{})},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
