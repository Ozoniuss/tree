package rbtree

import "testing"

func TestToStringPreorder(t *testing.T) {

	type testCase struct {
		rootNode *RbtNode[int]
		repr     string
	}

	testcases := []testCase{
		{
			rootNode: nil,
			repr:     "*",
		},
		{
			rootNode: Instantiate1(),
			repr: `5
├── 2
└── 10
    ├── 8
    │   ├── 6
    │   └── 9
    └── 12`,
		},
		{
			rootNode: Instantiate2(),
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
		out := tc.rootNode.ToStringPreorder()
		if out != tc.repr {
			t.Errorf("expected:\n%s (len: %d)\ngot:\n%s (len: %d)", tc.repr, len(tc.repr), out, len(out))
		}
	}
}
