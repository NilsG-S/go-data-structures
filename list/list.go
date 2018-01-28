package list

import (
	"bytes"
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func (v *Node) nilify() {
	v.Next = nil
	v.Prev = nil
}

type List struct {
	Front *Node
	Back  *Node
	Len   int
}

func New() *List {
	l := new(List)

	l.Front = nil
	l.Back = nil
	l.Len = 0

	return l
}

func (v *List) Prepend(in int) *Node {
	temp := new(Node)

	temp.Value = in
	temp.Next = nil
	temp.Prev = v.Front

	if v.Len != 0 {
		v.Front.Next = temp
	} else {
		v.Back = temp
	}

	v.Front = temp
	v.Len++

	return temp
}

func (v *List) Append(in int) {
	temp := new(Node)

	temp.Value = in
	temp.Next = v.Back
	temp.Prev = nil

	if v.Len != 0 {
		v.Back.Prev = temp
	} else {
		v.Front = temp
	}

	v.Back = temp
	v.Len++
}

func (v *List) Insert(in, pos int) error {
	if pos < 0 || (pos+1) > v.Len {
		return errors.New("Invalid index")
	}

	if pos == 0 {
		v.Prepend(in)
		return nil
	}

	temp := new(Node)
	cur := v.Front

	for i := 0; i < pos; i++ {
		cur = cur.Prev
	}

	temp.Value = in
	temp.Next = cur.Next
	temp.Prev = cur

	// These are automatically dereferenced
	cur.Next.Prev = temp
	cur.Next = temp

	v.Len++

	return nil
}

func (v *List) Pop() (interface{}, error) {
	if v.Len == 0 {
		return 0, errors.New("List has no elements")
	}

	temp := v.Front

	if v.Len == 1 {
		v.Front = nil
		v.Back = nil
	} else {
		v.Front.Prev.Next = nil
		v.Front = v.Front.Prev
	}

	defer temp.nilify()
	v.Len--

	return temp.Value, nil
}

func (v *List) PopBack() (interface{}, error) {
	if v.Len == 0 {
		return 0, errors.New("List has no elements")
	}

	temp := v.Back

	if v.Len == 1 {
		v.Front = nil
		v.Back = nil
	} else {
		v.Back.Next.Prev = nil
		v.Back = v.Back.Next
	}

	defer temp.nilify()
	v.Len--

	return temp.Value, nil
}

func (v *List) Remove(pos int) (interface{}, error) {
	if pos < 0 || (pos+1) > v.Len {
		return 0, errors.New("Invalid index")
	}

	if v.Len == 0 {
		return 0, errors.New("List has no elements")
	}

	if pos == 0 {
		return v.Pop()
	}

	if pos == v.Len-1 {
		return v.PopBack()
	}

	cur := v.Front

	for i := 0; i < pos; i++ {
		cur = cur.Prev
	}

	cur.Next.Prev = cur.Prev
	cur.Prev.Next = cur.Next

	defer cur.nilify()
	v.Len--

	return cur.Value, nil
}

func (v List) toSlice() []interface{} {
	if v.Front == nil {
		return make([]interface{}, 0)
	}

	var (
		out  []interface{} = make([]interface{}, 0, v.Len)
		temp *Node         = v.Front
	)

	for temp != nil {
		out = append(out, temp.Value)

		temp = temp.Prev
	}

	return out
}

func (v List) String() string {
	var buffer bytes.Buffer

	for _, v := range v.toSlice() {
		buffer.WriteString(fmt.Sprintf("%v ", v))
	}

	return buffer.String()
}
