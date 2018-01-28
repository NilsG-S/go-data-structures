package unionfind

type Element struct {
	parent interface{}
	size   int
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
