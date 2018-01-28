package list

import "testing"

func TestPrepend(t *testing.T) {
	l := New()
	first := l.Prepend()

	if first != l.Front {
		t.Error("Prepend doesn't set list front")
	}

	if first != l.Back {
		t.Error("Prepend doesn't set list back")
	}
}

func main() {
	// List{} just returns a value, not a pointer
	test := &List{
		front:  nil,
		back:   nil,
		length: 0,
	}

	test.Prepend(2)
	test.Prepend(1)
	fmt.Println(test)

	fmt.Println(test.Pop())
	fmt.Println(test)

	test.Pop()

	fmt.Println(test.Pop())
	fmt.Println(test)

	test.Prepend(2)
	test.Prepend(1)
	fmt.Println(test)

	fmt.Println(test.PopBack())
	fmt.Println(test)

	test.PopBack()

	fmt.Println(test.PopBack())
	fmt.Println(test)

	test.Prepend(2)
	test.Prepend(1)
	fmt.Println(test)

	fmt.Println(test.Remove(1))
	fmt.Println(test)

	test.Remove(0)

	fmt.Println(test.Remove(0))
	fmt.Println(test)

	// Even functions with pointer receivers can work on values, in some cases
	test.Prepend(2)
	fmt.Println(test)

	test.Prepend(1)
	fmt.Println(test)

	test.Append(3)
	fmt.Println(test)

	test.Insert(4, 1)
	fmt.Println(test)

	test.Insert(5, 0)
	fmt.Println(test)

	test.Insert(6, 4)
	fmt.Println(test)

	return
}
