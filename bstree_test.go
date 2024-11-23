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

func printNodes[T comparable]() (*string, func(n *BstNode[T])) {
	out := ""
	return &out, func(n *BstNode[T]) {
		if out == "" {
			out = fmt.Sprint(n.Key)
		} else {
			out = fmt.Sprintf("%s %v", out, n.Key)
		}
	}
}

func TestTreeTraversals(t *testing.T) {
	type testCase struct {
		rootNode      *BstNode[int]
		inorderStr    string
		preorderStr   string
		postorderStr  string
		levelorderStr string
	}

	testcases := []testCase{
		{
			rootNode:      nil,
			inorderStr:    "",
			preorderStr:   "",
			postorderStr:  "",
			levelorderStr: "",
		},
		{
			rootNode:      Instantiate1Bst(),
			inorderStr:    "2 5 6 8 9 10 12",
			preorderStr:   "5 2 10 8 6 9 12",
			postorderStr:  "2 6 9 8 12 10 5",
			levelorderStr: "5 2 10 8 12 6 9",
		},
	}

	for _, tc := range testcases {
		outInorder, printInorder := printNodes[int]()
		outPreorder, printPreorder := printNodes[int]()
		outPostorder, printPostorder := printNodes[int]()
		outLevelorder, printLevelorder := printNodes[int]()
		tc.rootNode.TraverseInorder(printInorder)
		if *outInorder != tc.inorderStr {
			t.Errorf("[inorder] expected:(%s) got:(%s)", tc.inorderStr, *outInorder)
		}
		tc.rootNode.TraversePostorder(printPostorder)
		if *outPostorder != tc.postorderStr {
			t.Errorf("[postorder] expected:(%s) got:(%s)", tc.postorderStr, *outPostorder)
		}
		tc.rootNode.TraversePreorder(printPreorder)
		if *outPreorder != tc.preorderStr {
			t.Errorf("[preorder] expected:(%s) got:(%s)", tc.preorderStr, *outPreorder)
		}
		tc.rootNode.TraverseLevelorder(printLevelorder)
		if *outLevelorder != tc.levelorderStr {
			t.Errorf("[levelorder] expected:(%s) got:(%s)", tc.levelorderStr, *outLevelorder)
		}
	}
}
