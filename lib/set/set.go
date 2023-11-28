package set

// Generic implementation of sets
type Set[K comparable] struct {
	items map[K]bool
}

func (set *Set[K]) Add(value K) {
	set.items[value] = true
}

func (set *Set[K]) Remove(value K) {
	set.items[value] = false
}

func (s1 *Set[K]) Intersection(s2 *Set[K]) Set[K] {
	intersection := Set[K]
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for k, _ := range s1 {
		if s2[k] {
			intersection[k] = true
		}
	}
	return intersection
}
