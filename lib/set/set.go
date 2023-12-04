package set

// Generic implementation of sets
type Set[K comparable] map[K]struct{}

func Create[K comparable]() Set[K] {
	return Set[K]{}
}

func CreateWithValue[K comparable](value K) Set[K] {
	a := Set[K]{}
	a.Add(value)
	return a
}

func CreateWithValues[K comparable](values []K) Set[K] {
	a := Set[K]{}
	a.AddValues(values)
	return a
}

func (set Set[K]) Add(value K) {
	set[value] = struct{}{} 
}

func (set Set[K]) AddValues(value []K) {
	for _, v := range value {
		set.Add(v)
	}
}

func (set Set[K]) Remove(value K) {
	delete(set, value)
}

func (s1 Set[K]) Intersection(s2 Set[K]) Set[K] {
	intersection := Set[K]{}
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for k, _ := range s1 {
		if _, found := s2[k]; found {
			intersection.Add(k) 
		}
	}
	return intersection
}
