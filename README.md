# Overview
The collection of utilities to work with [QFrame][1] implementation of the data frames in GO language.
It is aimed to provide implementation of functions missed in original QFrame but which is necessary for easy data manipulation 
using data frames.

## Installation
```bash
go get github.com/yaricom/dfutils
```

# Joining Data Frames
The data frames join operations is necessary for complex data manipulation routines involving multistage data processing.

## Left Join
The left join allows joining of two data frames by keys in the specific column. It is similar to the SQL LEFT JOIN function.
All rows from the left data frame will remain and matched columns from the right data frame will be added.

### Example
The left join of two data frames can be done as following:

```go
    import (
        "github.com/tobgu/qframe"
        "github.com/yaricom/dfutils"
        "strings"
    )

    leftDF := qframe.ReadCSV(strings.NewReader(leftData))
    rightDF := qframe.ReadCSV(strings.NewReader(rightDate))
    
    colNames := map[string]string{"C": "E"}
    joinedDF, err := dfutils.LeftJoin("Category", &leftDF, &rightDF, colNames)
```
#### Left Table
Category(s) | Subcategory(s) |  A(i) |  B(i) |  C(i)
----------- | -------------- | ----- | ----- | -----
cat1 | sub1 | 1 | 101 | 201
cat1 | sub2 | 2 | 102 | 202
cat1 | sub3 | 3 | 103 | 203
cat1 | sub4 | 4 | 104 | 204
cat2 | sub5 | 5 | 105 | 205
cat2 | sub6 | 6 | 106 | 206
cat2 | sub7 | 7 | 107 | 207
cat2 | sub8 | 8 | 108 | 208
cat2 | sub9 | 9 | 109 | 209

#### Right Table
Category(s) | C(i)
----------- | -----
cat1 | 1000
cat2 | 2000
cat3 | 3000

#### Join Table
Category(s) | Subcategory(s) |  A(i) |  B(i) |  C(i) |  E(i)
----------- | -------------- | ----- | ----- | ----- | -----
cat1 | sub1 | 1 | 101 | 201 | 1000
cat1 | sub2 | 2 | 102 | 202 | 1000
cat1 | sub3 | 3 | 103 | 203 | 1000
cat1 | sub4 | 4 | 104 | 204 | 1000
cat2 | sub5 | 5 | 105 | 205 | 2000
cat2 | sub6 | 6 | 106 | 206 | 2000
cat2 | sub7 | 7 | 107 | 207 | 2000
cat2 | sub8 | 8 | 108 | 208 | 2000
cat2 | sub9 | 9 | 109 | 209 | 2000

# Utilities
Multiple utilities also provided.

## Keep Columns
Allows keeping only a few columns and remove other ones. It is handy when you have many columns and want
to keep only a few:
```go
	df := qframe.New(map[string]types.DataSlice{
		"A": make([]float64, 10), "B": make([]float64, 10), "C": make([]float64, 10), "D": make([]float64, 10),
	})

	// invoke function
	resDF := dfutils.KeepColumns(&df, "A")
```

## Credits
* The **QFrame** (*Data Frames for GO language*) created and maintained by Tobias Gustafsson, see: [QFrame][1]

This source code maintained and managed by [Iaroslav Omelianenko][2]

[1]:https://github.com/tobgu/qframe
[2]:https://io42.space