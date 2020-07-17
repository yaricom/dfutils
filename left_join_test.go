package dfutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tobgu/qframe"
	"strings"
	"testing"
)

func TestLeftJoin(t *testing.T) {
	leftDF := qframe.ReadCSV(strings.NewReader(leftData))
	rightDF := qframe.ReadCSV(strings.NewReader(rightDate))
	joinedGTDF := qframe.ReadCSV(strings.NewReader(groundTruth))

	// join tables with column renamed
	colNames := map[string]string{"C": "E"}
	joinedDF, err := LeftJoin("Category", &leftDF, &rightDF, colNames)
	require.Nil(t, err, "failed to join data frames")
	require.NotNil(t, joinedDF, "join result is nil")

	// test that result is equal to the ground truth
	equal, reason := joinedDF.Equals(joinedGTDF)
	assert.True(t, equal, "join result is not equal to the ground truth, reason: %s", reason)

	t.Log("\n", leftDF)
	t.Log("\n", rightDF)
	t.Log("\n", joinedDF)
}

const (
	leftData = `Category,Subcategory,A,B,C
cat1,sub1,1,101,201
cat1,sub2,2,102,202
cat1,sub3,3,103,203
cat1,sub4,4,104,204
cat2,sub5,5,105,205
cat2,sub6,6,106,206
cat2,sub7,7,107,207
cat2,sub8,8,108,208
cat2,sub9,9,109,209
cat4,sub9,9,109,209
`
	rightDate = `Category,C
cat1,1000
cat2,2000
cat3,3000
`
	groundTruth = `Category,Subcategory,A,B,C,E
cat1,sub1,1,101,201,1000
cat1,sub2,2,102,202,1000
cat1,sub3,3,103,203,1000
cat1,sub4,4,104,204,1000
cat2,sub5,5,105,205,2000
cat2,sub6,6,106,206,2000
cat2,sub7,7,107,207,2000
cat2,sub8,8,108,208,2000
cat2,sub9,9,109,209,2000
cat4,sub9,9,109,209,0
`
)
