package dfutils

import (
	"github.com/pkg/errors"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/config/newqf"
	"github.com/tobgu/qframe/types"
)

// LeftJoin is to do left join of the two data frames using values from the col column as a keys. The provided
// names map allows to rename columns of the right tables during join.
func LeftJoin(col string, left, right *qframe.QFrame, rNamesMap map[string]string) (*qframe.QFrame, error) {
	// sanity check
	if left.Len() < right.Len() {
		return nil, errors.Errorf("left data frame has less rows than the right, left [%d], right [%d]",
			left.Len(), right.Len())
	}
	if !right.Contains(col) {
		return nil, errors.Errorf("column [%s] not found in the right data frame", col)
	}
	if !left.Contains(col) {
		return nil, errors.Errorf("column [%s] not found in the left data frame", col)
	}
	// build data holders for the extended right data frame values
	rData := make(map[string]interface{})
	rightTypes := right.ColumnTypeMap()
	rows := left.Len()
	for _, name := range right.ColumnNames() {
		if name != col {
			switch rightTypes[name] {
			case types.String:
				rData[name] = make([]*string, rows, rows)
			case types.Int:
				rData[name] = make([]int, rows, rows)
			case types.Float:
				rData[name] = make([]float64, rows, rows)
			default:
				return nil, errors.Errorf("unsupported type [%s] of column [%s]", rightTypes[name], name)
			}
		}
	}
	// get data views from the right data frame
	rViews := make(map[string]interface{})
	var err error
	for _, name := range right.ColumnNames() {
		if name != col {
			switch rightTypes[name] {
			case types.String:
				if rViews[name], err = right.StringView(name); err != nil {
					return nil, errors.Wrapf(err, "failed to get string view from column [%s]", name)
				}
			case types.Int:
				if rViews[name], err = right.IntView(name); err != nil {
					return nil, errors.Wrapf(err, "failed to get int view from column [%s]", name)
				}
			case types.Float:
				if rViews[name], err = right.FloatView(name); err != nil {
					return nil, errors.Wrapf(err, "failed to get float view from column [%s]", name)
				}
			default:
				return nil, errors.Errorf("unsupported type [%s] of column [%s]", rightTypes[name], name)
			}
		}
	}
	// iterate and build the data series for the expanded right data frame
	leftJoinColumn, _ := left.StringView(col)
	rightJoinColumn, _ := right.StringView(col)
	var rKey, lKey string
	for i := 0; i < rightJoinColumn.Len(); i++ {
		rKey = *rightJoinColumn.ItemAt(i)
		for j := 0; j < leftJoinColumn.Len(); j++ {
			lKey = *leftJoinColumn.ItemAt(j)
			if lKey == rKey {
				for _, name := range right.ColumnNames() {
					if name != col {
						switch rightTypes[name] {
						case types.String:
							rData[name].([]*string)[j] = rViews[name].(qframe.StringView).ItemAt(i)
						case types.Int:
							rData[name].([]int)[j] = rViews[name].(qframe.IntView).ItemAt(i)
						case types.Float:
							rData[name].([]float64)[j] = rViews[name].(qframe.FloatView).ItemAt(i)
						default:
							return nil, errors.Errorf("unsupported type [%s] of column [%s]", rightTypes[name], name)
						}
					}
				}
			}
		}
	}
	// create joined data frame
	data := make(map[string]interface{})
	leftTypes := left.ColumnTypeMap()
	columns := make([]string, 0)
	for _, name := range left.ColumnNames() {
		switch leftTypes[name] {
		case types.String:
			if val, err := left.StringView(name); err != nil {
				return nil, errors.Wrapf(err, "failed to get string view from column [%s]", name)
			} else {
				data[name] = val.Slice()
			}
		case types.Int:
			if val, err := left.IntView(name); err != nil {
				return nil, errors.Wrapf(err, "failed to get int view from column [%s]", name)
			} else {
				data[name] = val.Slice()
			}
		case types.Float:
			if val, err := left.FloatView(name); err != nil {
				return nil, errors.Wrapf(err, "failed to get float view from column [%s]", name)
			} else {
				data[name] = val.Slice()
			}
		default:
			return nil, errors.Errorf("unsupported type [%s] of column [%s]", rightTypes[name], name)
		}
		columns = append(columns, name)
	}
	for key, val := range rData {
		if newKey, ok := rNamesMap[key]; ok {
			data[newKey] = val
			columns = append(columns, newKey)
		} else {
			data[key] = val
			columns = append(columns, key)
		}
	}
	df := qframe.New(data, newqf.ColumnOrder(columns...))
	if df.Err != nil {
		return nil, df.Err
	} else {
		return &df, nil
	}
}
