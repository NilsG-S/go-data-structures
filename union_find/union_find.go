package unionfind

import "errors"

// TODO(NilsG-S): Add some custom errors

type Element struct {
	parent *Element
	size   int
}

func (v *Element) Size() int {
	return v.size
}

type UnionFind struct {
	elements map[interface{}]*Element
	count    int
}

func New() *UnionFind {
	u := new(UnionFind)

	u.elements = make(map[interface{}]*Element)
	u.count = 0

	return u
}

func (v *UnionFind) Insert(key interface{}) error {
	if _, ok := v.elements[key]; ok {
		return errors.New("Key already exists")
	}

	e := new(Element)

	e.parent = e
	e.size = 1

	v.elements[key] = e

	return nil
}

func (v *UnionFind) Find(key interface{}) (*Element, error) {
	e, ok := v.elements[key]

	if !ok {
		return nil, errors.New("Key doesn't exist")
	}

	root := e

	for root != root.parent {
		root = root.parent
	}

	for e != root {
		eNew := e.parent
		e.parent = root
		e = eNew
	}

	return root, nil
}

func (v *UnionFind) Connected(key1, key2 interface{}) (bool, error) {
	root1, err1 := v.Find(key1)

	if err1 != nil {
		return false, err1
	}

	root2, err2 := v.Find(key2)

	if err2 != nil {
		return false, err2
	}

	return root1 == root2, nil
}

func (v *UnionFind) Union(key1, key2 interface{}) error {
	root1, err1 := v.Find(key1)

	if err1 != nil {
		return err1
	}

	root2, err2 := v.Find(key2)

	if err2 != nil {
		return err2
	}

	if root1 == root2 {
		return errors.New("Elements already connected")
	}

	if root1.size < root2.size {
		root1.parent = root2
		root2.size += root1.size
	} else {
		root2.parent = root1
		root1.size += root2.size
	}

	v.count--

	return nil
}
