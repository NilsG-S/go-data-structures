package unionfind

import "errors"

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
