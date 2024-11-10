package rbtree

/*
Instantiate1 returns the following sample tree:

	  5
	 / \
	2  10
	   / \
	  8   12
	 / \
	6   9
*/
func Instantiate1() *RbtNode[int] {
	five := &RbtNode[int]{Key: 5, P: nil}
	five.Root = five

	six := &RbtNode[int]{
		Root:  five,
		Left:  nil,
		Right: nil,
		Key:   6,
	}
	nine := &RbtNode[int]{
		Root:  five,
		Left:  nil,
		Right: nil,
		Key:   9,
	}

	eight := &RbtNode[int]{
		Root:  five,
		Left:  six,
		Right: nine,
		Key:   8,
	}
	six.P = eight
	nine.P = eight

	twelve := &RbtNode[int]{
		Root:  five,
		Left:  nil,
		Right: nil,
		Key:   12,
	}
	ten := &RbtNode[int]{
		Root:  five,
		Left:  eight,
		Right: twelve,
		Key:   10,
	}
	twelve.P = ten
	eight.P = ten

	two := &RbtNode[int]{
		Root:  five,
		Left:  nil,
		Right: nil,
		Key:   2,
	}
	ten.P = five
	two.P = five

	five.Left = two
	five.Right = ten

	return five
}

/*
Instantiate2 returns the following sample tree:

	    4
	   / \
	  2   10
	 /   / \
	1   8   12
	   / \
	  6   9
	 /     \
	5       11
*/
func Instantiate2() *RbtNode[int] {
	four := &RbtNode[int]{Key: 4, P: nil}
	four.Root = four

	five := &RbtNode[int]{
		Root:  four,
		Left:  nil,
		Right: nil,
		Key:   5,
	}
	eleven := &RbtNode[int]{
		Root:  four,
		Left:  nil,
		Right: nil,
		Key:   11,
	}
	one := &RbtNode[int]{
		Root:  four,
		Left:  nil,
		Right: nil,
		Key:   1,
	}
	twelve := &RbtNode[int]{
		Root:  four,
		Left:  nil,
		Right: nil,
		Key:   12,
	}
	six := &RbtNode[int]{
		Root:  four,
		Left:  five,
		Right: nil,
		Key:   6,
	}
	five.P = six
	nine := &RbtNode[int]{
		Root:  four,
		Left:  nil,
		Right: eleven,
		Key:   9,
	}
	eleven.P = nine

	eight := &RbtNode[int]{
		Root:  four,
		Left:  six,
		Right: nine,
		Key:   8,
	}
	six.P = eight
	nine.P = eight

	ten := &RbtNode[int]{
		Root:  four,
		Left:  eight,
		Right: twelve,
		Key:   10,
	}
	eight.P = ten
	twelve.P = ten

	two := &RbtNode[int]{
		Root:  four,
		Left:  one,
		Right: nil,
		Key:   2,
	}
	one.P = two

	ten.P = four
	two.P = four

	four.Left = two
	four.Right = ten

	return four
}
