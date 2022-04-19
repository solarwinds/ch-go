// Code generated by ./cmd/ch-gen-int, DO NOT EDIT.

package proto

// ColInt8 represents Int8 column.
type ColInt8 []int8

// Compile-time assertions for ColInt8.
var (
	_ ColInput  = ColInt8{}
	_ ColResult = (*ColInt8)(nil)
	_ Column    = (*ColInt8)(nil)
)

// Type returns ColumnType of Int8.
func (ColInt8) Type() ColumnType {
	return ColumnTypeInt8
}

// Rows returns count of rows in column.
func (c ColInt8) Rows() int {
	return len(c)
}

// Row returns i-th row of column.
func (c ColInt8) Row(i int) int8 {
	return c[i]
}

// Append int8 to column.
func (c *ColInt8) Append(v int8) {
	*c = append(*c, v)
}

// AppendArr appends slice of int8 to column.
func (c *ColInt8) AppendArr(v []int8) {
	*c = append(*c, v...)
}

// Reset resets data in row, preserving capacity for efficiency.
func (c *ColInt8) Reset() {
	*c = (*c)[:0]
}

// LowCardinality returns LowCardinality for Int8 .
func (c *ColInt8) LowCardinality() *ColLowCardinalityOf[int8] {
	return &ColLowCardinalityOf[int8]{
		index: c,
	}
}

// Array is helper that creates Array of int8.
func (c *ColInt8) Array() *ColArrOf[int8] {
	return &ColArrOf[int8]{
		Data: c,
	}
}

// NewArrInt8 returns new Array(Int8).
func NewArrInt8() *ColArr {
	return &ColArr{
		Data: new(ColInt8),
	}
}

// AppendInt8 appends slice of int8 to Array(Int8).
func (c *ColArr) AppendInt8(data []int8) {
	d := c.Data.(*ColInt8)
	*d = append(*d, data...)
	c.Offsets = append(c.Offsets, uint64(len(*d)))
}
