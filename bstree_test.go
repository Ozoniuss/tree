package tree

import (
	"fmt"
	"testing"
)

func TestFormatLinuxTree(t *testing.T) {

	type testCase struct {
		rootNode *BstNode[int]
		repr     string
	}

	testcases := []testCase{
		{
			rootNode: nil,
			repr:     "*",
		},
		{
			rootNode: Instantiate1Bst(),
			repr: `5
├── 2
└── 10
    ├── 8
    │   ├── 6
    │   └── 9
    └── 12`,
		},
		{
			rootNode: Instantiate2Bst(),
			repr: `4
├── 2
│   ├── 1
│   └── *
└── 10
    ├── 8
    │   ├── 6
    │   │   ├── 5
    │   │   └── *
    │   └── 9
    │       ├── *
    │       └── 11
    └── 12`,
		},
	}

	for _, tc := range testcases {
		out := tc.rootNode.Format(FormatLinuxTree)
		if out != tc.repr {
			t.Errorf("expected:\n%s (len: %d)\ngot:\n%s (len: %d)", tc.repr, len(tc.repr), out, len(out))
		}
	}
}

func getPrintInorder[T comparable]() (*string, func(n *BstNode[T])) {
	out := ""
	return &out, func(n *BstNode[T]) {
		if out == "" {
			out = fmt.Sprint(n.Key)
		} else {
			out = fmt.Sprintf("%s %v", out, n.Key)
		}
	}
}

func TestTraverseInorder(t *testing.T) {
	type testCase struct {
		rootNode   *BstNode[int]
		inorderStr string
	}

	testcases := []testCase{
		{
			rootNode:   nil,
			inorderStr: "",
		},
		{
			rootNode:   Instantiate1Bst(),
			inorderStr: "2 5 6 8 9 10 12",
		},
	}

	out, printInorder := getPrintInorder[int]()

	for _, tc := range testcases {
		tc.rootNode.TraverseInorder(printInorder)
		if *out != tc.inorderStr {
			t.Errorf("expected:(%s) got:(%s)", tc.inorderStr, *out)
		}
	}
}
