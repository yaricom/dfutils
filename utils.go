package dfutils

import "github.com/tobgu/qframe"

// KeepColumns is keep only specified columns in the provided data frame. All other columns will be dropped. It is handy
// when many columns should be dropped at once by keeping only a few.
func KeepColumns(df *qframe.QFrame, cols ...string) qframe.QFrame {
	names := df.ColumnNames()
	dropCols := make([]string, 0)
	for _, dfCol := range names {
		if !ContainedIn(dfCol, cols) {
			dropCols = append(dropCols, dfCol)
		}
	}
	return df.Drop(dropCols...)
}

// ContainedIn checks if a value in a list of a strings
func ContainedIn(value string, list []string) bool {
	for _, x := range list {
		if x == value {
			return true
		}
	}
	return false
}
