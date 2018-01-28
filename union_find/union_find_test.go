package unionfind

import "testing"

func TestNew(t *testing.T) {
	u := New()

	if u.count != 0 {
		t.Errorf("Expected count 0, got count: %v", u.count)
	}

	if len(u.elements) != 0 {
		t.Errorf("Expected 0 elements, got %v elements", len(u.elements))
	}
}

func TestInsert(t *testing.T) {
	u := New()

	u.Insert(1)

	if u.count != 1 {
		t.Errorf("Expected 1 sets, got %v sets", u.count)
	}

	if len(u.elements) != 1 {
		t.Errorf("Expected 1 elements, got %v elements", len(u.elements))
	}

	if u.elements[1].parent != u.elements[1] {
		t.Errorf("Expected %v to be its own parent", 1)
	}

	if u.elements[1].size != 1 {
		t.Errorf("Expected %v to have size 1", 1)
	}
}
