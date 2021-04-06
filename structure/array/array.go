package array

import "errors"

// Array
type Array struct {
	data   []int
	cap    int
	length int
}

var (
	InvalidIndex = errors.New("invalid index")
)

// NewArray
func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{data: make([]int, 0, capacity), length: 0}
}

// Length
func (a *Array) Length() int {
	return a.length
}

func (a *Array) Find(index int) (int, error) {
	if index < 0 || index > a.length {
		return 0, InvalidIndex
	}
	return 0, nil
}

// Insert insert data to head、middle and end
func (a *Array) Insert(index, data int) {

}

// Delete
func (a *Array) Delete(indexList int) {

}
