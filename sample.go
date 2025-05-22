package tree

// /*
// Instantiate1Bst returns the following sample tree:

// 	  5
// 	 / \
// 	2  10
// 	   / \
// 	  8   12
// 	 / \
// 	6   9
// */
// func Instantiate1Bst() *BstNode[int] {
// 	five := &BstNode[int]{Value: 5, P: nil}

// 	six := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 6,
// 	}
// 	nine := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 9,
// 	}

// 	eight := &BstNode[int]{
// 		Left:  six,
// 		Right: nine,
// 		Value: 8,
// 	}
// 	six.P = eight
// 	nine.P = eight

// 	twelve := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 12,
// 	}
// 	ten := &BstNode[int]{
// 		Left:  eight,
// 		Right: twelve,
// 		Value: 10,
// 	}
// 	twelve.P = ten
// 	eight.P = ten

// 	two := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 2,
// 	}
// 	ten.P = five
// 	two.P = five

// 	five.Left = two
// 	five.Right = ten

// 	return five
// }

// /*
// Instantiate2Bst returns the following sample tree:

// 	    4
// 	   / \
// 	  2   10
// 	 /   / \
// 	1   8   12
// 	   / \
// 	  6   9
// 	 /     \
// 	5       11
// */
// func Instantiate2Bst() *BstNode[int] {
// 	four := &BstNode[int]{Value: 4, P: nil}

// 	five := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 5,
// 	}
// 	eleven := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 11,
// 	}
// 	one := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 1,
// 	}
// 	twelve := &BstNode[int]{
// 		Left:  nil,
// 		Right: nil,
// 		Value: 12,
// 	}
// 	six := &BstNode[int]{
// 		Left:  five,
// 		Right: nil,
// 		Value: 6,
// 	}
// 	five.P = six
// 	nine := &BstNode[int]{
// 		Left:  nil,
// 		Right: eleven,
// 		Value: 9,
// 	}
// 	eleven.P = nine

// 	eight := &BstNode[int]{
// 		Left:  six,
// 		Right: nine,
// 		Value: 8,
// 	}
// 	six.P = eight
// 	nine.P = eight

// 	ten := &BstNode[int]{
// 		Left:  eight,
// 		Right: twelve,
// 		Value: 10,
// 	}
// 	eight.P = ten
// 	twelve.P = ten

// 	two := &BstNode[int]{
// 		Left:  one,
// 		Right: nil,
// 		Value: 2,
// 	}
// 	one.P = two

// 	ten.P = four
// 	two.P = four

// 	four.Left = two
// 	four.Right = ten

// 	return four
// }
