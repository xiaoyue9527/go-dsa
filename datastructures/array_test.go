package datastructures

import (
	"testing"
	"time"
)

func TestMapArray(t *testing.T) {
	// 创建一个 MapArray
	arr := NewMapArray(5)

	// 添加元素
	arr.Push(1)
	arr.Push(2)
	arr.Push(3)

	// 检查长度是否正确
	if l := arr.Length(); l != 3 {
		t.Errorf("Expected length to be 3, but got %d", l)
	}

	// 检查容量是否正确
	if c := arr.Capacity(); c != 5 {
		t.Errorf("Expected capacity to be 5, but got %d", c)
	}

	// 测试 Unshift
	arr.Unshift(0)
	if l := arr.Length(); l != 4 {
		t.Errorf("Expected length to be 4 after Unshift, but got %d", l)
	}
	if d := arr.Data[0]; d != 0 {
		t.Errorf("Expected the first element to be 0 after Unshift, but got %v", d)
	}

	// 测试 Splice
	arr.Splice(1, 100)
	if l := arr.Length(); l != 5 {
		t.Errorf("Expected length to be 5 after Splice, but got %d", l)
	}
	if d := arr.Data[1]; d != 100 {
		t.Errorf("Expected the second element to be 100 after Splice, but got %v", d)
	}

	// 测试 FindAll
	arr.Push(2)
	arr.Push(3)
	indexList := arr.FindAll(2)
	if len(indexList) != 2 {
		t.Errorf("Expected 2 occurrences of 2, but got %d", len(indexList))
	}
	if indexList[0] != 3 || indexList[1] != 5 {
		t.Errorf("Expected occurrences at index 1 and 5, but got %v", indexList)
	}

	// 测试 Del
	arr.Del(2)
	if l := arr.Length(); l != 6 {
		t.Errorf("Expected length to be 5 after Del, but got %d", l)
	}
	if d := arr.Data[1]; d != 100 {
		t.Errorf("Expected the second element to be 100 after Del, but got %v", d)
	}

	// 测试 Pop
	arr.Pop()
	if l := arr.Length(); l != 5 {
		t.Errorf("Expected length to be 4 after Pop, but got %d", l)
	}
	if d := arr.Data[3]; d != 3 {
		t.Errorf("Expected the last element to be 2 after Pop, but got %v", d)
	}

	// 测试空切片的情况
	emptyArr := NewMapArray(5)
	emptyArr.Pop() // 测试 Pop 空切片
	if l := emptyArr.Length(); l != 0 {
		t.Errorf("Expected length to be 0 after Pop from empty list, but got %d", l)
	}

	// 测试超出容量阈值的情况
	capacityArr := NewMapArray(2)
	for i := 0; i < 105; i++ {
		capacityArr.Push(i)
	}
	if c := capacityArr.Capacity(); c <= 100 {
		t.Errorf("Expected capacity to be increased after adding elements, but got %d", c)
	}

	// 测试负面情况
	negativeArr := NewMapArray(5)
	negativeArr.Splice(-1, 100) // 测试负索引插入元素
	if l := negativeArr.Length(); l != 1 {
		t.Errorf("Expected length to be 1 after negative index Splice, but got %d", l)
	}
	negativeArr.Del(10) // 测试删除不存在的元素
	if l := negativeArr.Length(); l != 1 {
		t.Errorf("Expected length to be 1 after deleting non-existent element, but got %d", l)
	}

	// 性能测试
	largeArr := NewMapArray(100)
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		largeArr.Push(i)
	}
	elapsedTime := time.Since(startTime)
	t.Logf("Time taken for pushing 100000 elements: %v", elapsedTime)
}
