package tree

import (
	"cmp"
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
		out := tc.rootNode.Format("FormatLinuxTree")
		if out != tc.repr {
			t.Errorf("expected:\n%s (len: %d)\ngot:\n%s (len: %d)", tc.repr, len(tc.repr), out, len(out))
		}
	}
}

func printNodes[T cmp.Ordered]() (*string, func(n *BstNode[T])) {
	out := ""
	return &out, func(n *BstNode[T]) {
		if out == "" {
			out = fmt.Sprint(n.Value)
		} else {
			out = fmt.Sprintf("%s %v", out, n.Value)
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

func TestEquals(t *testing.T) {
	type testCase struct {
		name  string
		tree1 *BstNode[int]
		tree2 *BstNode[int]
		equal bool
	}

	testcases := []testCase{
		{
			name:  "equal sample trees 1",
			tree1: Instantiate1Bst(),
			tree2: Instantiate1Bst(),
			equal: true,
		},
		{
			name:  "fairly different trees",
			tree1: Instantiate1Bst(),
			tree2: Instantiate2Bst(),
			equal: false,
		},
		{
			name:  "equal sample trees 2",
			tree1: Instantiate2Bst(),
			tree2: Instantiate2Bst(),
			equal: true,
		},
		{
			name:  "equal nil trees",
			tree1: nil,
			tree2: nil,
			equal: true,
		},
		{
			name:  "empty tree with nil tree",
			tree1: nil,
			tree2: &BstNode[int]{},
			equal: false,
		},
		// TODO: comparisons with equal elements but different shape,
		// and equal shapes but different elements.
	}

	for _, tc := range testcases {
		if tc.equal != Equal(tc.tree1, tc.tree2) {
			t.Errorf("failed tree comparison for: %s", tc.name)
		}
	}
}

func TestInsertEmptyTree(t *testing.T) {
	var root = &BstNode[int]{}
	Insert(root, 10)
	if root.Size != 1 {
		t.Errorf("expected size to be 1, got %d", root.Size)
	}
	if root.Value != 10 {
		t.Errorf("expected root to have value %v, got %v", 10, root.Value)
	}
}
