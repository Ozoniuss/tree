package tree

const (
	PREFIX_LEFT  = "├──"
	PREFIX_RIGHT = "└──"

	EXTRA_LEFT  = "│   "
	EXTRA_RIGHT = "    "
)

const (
	/*
	   FormatLinuxTree formats the tree nicely as a string.

	   	    4
	   	   / \
	   	  2   10
	   	 /   / \
	   	1   8   12
	   	   / \
	   	  6   9
	   	 /     \
	   	5       11

	   would get converted to

	   	4
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
	   		└── 12
	*/
	FormatLinuxTree = "FormatLinuxTree"
)
