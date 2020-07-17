package dfutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/types"
	"testing"
)

func TestKeepColumns(t *testing.T) {
	df := qframe.New(map[string]types.DataSlice{
		"A": make([]float64, 10), "B": make([]float64, 10), "C": make([]float64, 10), "D": make([]float64, 10),
	})

	// invoke function
	resDF := KeepColumns(&df, "A")

	// check results
	columns := resDF.ColumnNames()
	assert.Len(t, columns, 1, "wrong length")
	assert.Contains(t, columns, "A", "not found the column")
}
