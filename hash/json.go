package hash

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

const (
	version    = 0
	minVersion = 0
	maxVersion = 0
)

const (
	AlreadyInitializedError  = "AlreadyInitializedError"
	IncompatibleVersionError = "IncompatibleVersion"
)

type matrixJson struct {
	Version  int          `json:"version"`
	Base     *types.Shape `json:"base"`
	View     *types.Shape `json:"view"`
	Offset   *types.Index `json:"offset"`
	Elements []Element    `json:"elements"`
	Rewriter byte         `json:"rewriter"`
}
