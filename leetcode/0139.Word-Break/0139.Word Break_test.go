package leetcode

import (
	"fmt"
	"testing"
)

type question139 struct {
	para139
	ans139
}

// para 是参数
// one 代表第一个参数
type para139 struct {
	one string
	two []string
}

// ans 是答案
// one 代表第一个答案
type ans139 struct {
	one bool
}

func Test_Problem139(t *testing.T) {

	qs := []question139{

		{
			para139{"leetcode", []string{"leet", "code"}},
			ans139{true},
		},

		{
			para139{"applepenapple", []string{"apple", "pen"}},
			ans139{true},
		},

		{
			para139{"catsandog", []string{"cat", "sand"}},
			ans139{false},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 139------------------------\n")

	for _, q := range qs {
		_, p := q.ans139, q.para139
		fmt.Printf("【input】:%v       【output】:%v\n", p, wordBreak(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
