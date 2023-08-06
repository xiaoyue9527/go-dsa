package datastructures

import (
	"testing"
	"time"
)

func TestMapArray(t *testing.T) {
	t.Run("Push and Length", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)

		if l := arr.Length(); l != 2 {
			t.Errorf("Expected length to be 2, but got %d", l)
		}
	})

	t.Run("Pop and Length", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Pop()

		if l := arr.Length(); l != 1 {
			t.Errorf("Expected length to be 1 after Pop, but got %d", l)
		}
	})

	t.Run("Unshift and Length", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Unshift(0)

		if l := arr.Length(); l != 3 {
			t.Errorf("Expected length to be 3 after Unshift, but got %d", l)
		}
	})

	t.Run("Splice", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.Splice(1, 100)

		if l := arr.Length(); l != 4 {
			t.Errorf("Expected length to be 4 after Splice, but got %d", l)
		}
		if d := arr.Data[1]; d != 100 {
			t.Errorf("Expected the second element to be 100 after Splice, but got %v", d)
		}
	})

	t.Run("FindAll", func(t *testing.T) {
		arr := NewMapArrayFormMap(int(10), 1, 2, 3, 4, 5, 6, 7, 2, 6, 8, 7)
		indexList := arr.FindAll(2)
		if len(indexList) != 2 {
			t.Errorf("Expected 2 occurrences of 2, but got %d", len(indexList))
		}
		if indexList[0] != 1 || indexList[1] != 7 {
			t.Errorf("Expected occurrences at index 1 and 7, but got %v", indexList)
		}
	})

	t.Run("Del and Length", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.Del(1)

		if l := arr.Length(); l != 2 {
			t.Errorf("Expected length to be 2 after Del, but got %d", l)
		}
	})

	t.Run("Capacity and CapacityManager", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.CapacityManager()

		if c := arr.Capacity(); c != 10 {
			t.Errorf("Expected capacity to be 10 after CapacityManager, but got %d", c)
		}
	})

	t.Run("CapacityChanges", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.CapacityChanges() // Print the capacity changes
	})

	t.Run("ToString and Join", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		str := arr.ToString()
		if str != "[1,2,3]" {
			t.Errorf("Expected ToString() to return '[1,2,3]', but got '%s'", str)
		}
	})

	t.Run("IndexValueOf and Includes", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		index := arr.IndexValueOf(2)
		included := arr.Includes(2)

		if index != 1 {
			t.Errorf("Expected IndexValueOf to return 1 for value 2, but got %d", index)
		}
		if !included {
			t.Errorf("Expected Includes to return true for value 2, but got false")
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		arr := NewMapArray[int](10)
		arr.Push(1)
		arr.Push(2)
		arr.Push(3)
		arr.Reverse()
		if str := arr.ToString(); str != "[3,2,1]" {
			t.Errorf("Expected Reverse to 返回数组 [3,2,1], but got '%s'", str)
		}
	})

	t.Run("Count", func(t *testing.T) {
		arr := NewMapArrayFormMap(int(10), 1, 2, 2, 3, 4, 2, 5)
		count := arr.Count(2)
		if count != 3 {
			t.Errorf("Expected Count to return 3 for value 2, but got %d", count)
		}
	})

	t.Run("Concat and Length", func(t *testing.T) {
		arr1 := NewMapArrayFormMap(int(10), 1, 2, 3)
		arr2 := NewMapArrayFormMap(int(10), 4, 5, 6)
		concatArr := arr1.Concat(*arr2)
		if l := concatArr.Length(); l != 6 {
			t.Errorf("Expected length to be 6 after Concat, but got %d", l)
		}
	})

	t.Run("ForEach", func(t *testing.T) {
		arr := NewMapArrayFormMap(int(10), 1, 2, 3)
		sum := 0
		arr.ForEach(func(data int) {
			sum += data
		})
		if sum != 6 {
			t.Errorf("Expected sum to be 6 after ForEach, but got %d", sum)
		}
	})

	t.Run("Performance Test", func(t *testing.T) {
		largeArr := NewMapArray[int](100)
		startTime := time.Now()
		for i := 0; i < 100000; i++ {
			largeArr.Push(i)
		}
		elapsedTime := time.Since(startTime)
		t.Logf("Time taken for pushing 100000 elements: %v", elapsedTime)
	})

}
