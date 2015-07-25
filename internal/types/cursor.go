package types

/*
"Cursor" interface is an iterator for elements of matrix.
Some implementations of "Cursor" iterate all elements,
and others iterates elements satisfying conditions.
*/
type Cursor interface {
	// Read the next element and return "true".
	// If the next element doesn't exist, return "false".
	HasNext() bool

	// Return the current read element.
	Get() (element float64, row, column int)
}
