package types

type Cursor interface {
	// Read the next element and return "true".
	// If the next element doesn't exist, return "false".
	HasNext() bool

	// Return the current read element.
	Get() (element float64, row, column int)
}
