package matrix

/*
"Elements" interface is an iterator for elements of matrix.
Some implementations of "Elements" iterate all elements,
and others iterates elements satisfying conditions.
*/
type Elements interface {
	// Read the next element and return "true".
	// If the next element doesn't exist, return "false".
	HasNext() bool

	// Return the current read element.
	Get() (element float64, row, column int)
}

/*
"ElementMatcher" is a type of functions to be used check an element satisfies arbitary condition.
*/
type ElementMatcher func(element float64, row, column int) bool
