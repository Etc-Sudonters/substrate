package reiterate

import (
	"fmt"
	"testing"
)

func TestZipTwo(t *testing.T) {
	t.Run("SameSize", func(t *testing.T) {
		arr := makeArr(9)
		rev := MakeReversed(arr)
		zip := ZipTwo(arr, rev)

		n := len(arr) - 1
		for zip.Next() {
			pair := zip.Current()
			if pair.A+pair.B != n {
				t.Logf("%d + %d != %d", pair.A, pair.B, n)
				t.Fail()
			}
		}
	})

	t.Run("DiffSizes", func(t *testing.T) {
		n := 8
		zip := ZipTwo(makeArr(n+1), makeArr((n^2)+1))
		for zip.Next() {
			pair := zip.Current()
			if pair.A > n || pair.B > n {
				t.Fatalf("iterated out of bounds like a test speedrun or something")
			}
		}

	})
}

func TestInPlaceReverse(t *testing.T) {

	sizes := []int{4, 9, 67, 98, 12734}

	for i := range sizes {
		n := sizes[i]
		t.Run(fmt.Sprintf("[%d]int", n), func(T *testing.T) {

			arr := makeArr(n)

			InPlaceReverse(arr)

			if !isBackwards(arr) {
				t.Log(fmt.Sprintf("[%d]int did not reverse", n))
				t.Fail()
			}

		})

	}
}

func TestMakeReversed(t *testing.T) {
	sizes := []int{4, 9, 67, 98, 12735}

	for i := range sizes {
		n := sizes[i]
		t.Run(fmt.Sprintf("[%d]int", n), func(T *testing.T) {

			if !isBackwards(MakeReversed(makeArr(n))) {
				t.Log(fmt.Sprintf("[%d]int did not reverse", n))
				t.Fail()
			}

		})

	}
}

func isBackwards(rev []int) bool {
	n := len(rev)
	for i := 0; i < n; i++ {
		if rev[i] != n-i-1 {
			return false
		}
	}

	return true
}

func makeArr(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}
