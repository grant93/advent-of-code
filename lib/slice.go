package slice

func RemoveIndex[S ~[]E, E any](s S, index int) S {
	ret := make(S, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
