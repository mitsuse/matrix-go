package dense

import (
	"fmt"
)

func rowsShouldBePositiveNumber(rows int) {
	shouldBePositiveNumber(rows, "rows")
}

func columnShouldBePositiveNumber(columns int) {
	shouldBePositiveNumber(columns, "columns")
}

func shouldBePositiveNumber(x int, name string) {
	if x > 0 {
		return
	}

	message := fmt.Sprintf("%q should be a positive number.", name)
	panic(message)
}
