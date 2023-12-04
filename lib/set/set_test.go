package set

import (
	"reflect"
	"testing"
)
func TestIntersection(t *testing.T) {
	s1 := CreateWithValues[int]([]int{10,20,30,40,50})
	s2 := CreateWithValues[int]([]int{20,40,60,80,100})
	s3 := CreateWithValues[int]([]int{20,40})

	if !reflect.DeepEqual(s1.Intersection(s2),s3) {
		t.Errorf("Incorrect intersection.")
	}
}
