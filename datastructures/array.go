// Package datastructures 提供了 MapArray，一个带有容量管理的动态数组。
package datastructures

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// MapArray 是一个带有容量管理的动态数组。
type MapArray[T constraints.Ordered] struct {
	Data        []T // Data 保存数组元素。
	SectionSize int // SectionSize 是容量管理的阈值。
}

// NewMapArray 创建一个新的 MapArray，并指定 SectionSize。
func NewMapArray[T constraints.Ordered](size int) *MapArray[T] {
	initSize := size
	sectionSize := size
	if size < 0 {
		sectionSize = 0
		initSize = 10
	}

	return &MapArray[T]{
		Data:        make([]T, 0, initSize),
		SectionSize: sectionSize,
	}
}

// NewMapArrayFormMap 创建一个新的 MapArray，并指定阈值和初始数据。
func NewMapArrayFormMap[T constraints.Ordered](size int, data ...T) *MapArray[T] {
	sectionSize := size
	if size < 0 {
		sectionSize = 0
	}
	list := MapArray[T]{Data: data, SectionSize: sectionSize}
	list.CapacityManager()
	return &list
}

// Push 在数组末尾添加一个元素。
func (list *MapArray[T]) Push(data T) {
	list.CapacityManager()
	list.Data = append(list.Data, data)
}

// Pop 从数组中删除最后一个元素。
func (list *MapArray[T]) Pop() {
	if len(list.Data) == 0 {
		return
	}
	list.Data = list.Data[:len(list.Data)-1]
	list.CapacityManager()
}

// Unshift 在数组开头添加一个元素。
func (list *MapArray[T]) Unshift(data T) int {
	list.Data = append([]T{data}, list.Data...)
	return list.Length()
}

// Splice 在指定索引处插入一个元素。
func (list *MapArray[T]) Splice(index int, data T) {
	if index < 0 {
		list.Unshift(data)
	} else if index > list.Length() {
		list.Push(data)
	} else {
		list.Data = append(
			list.Data[:index],
			append([]T{data}, list.Data[index:]...)...,
		)
	}
}

// FindAll 查找数组中所有与指定元素相等的元素，并返回它们的索引。
func (list *MapArray[T]) FindAll(data T) []int {
	indexList := make([]int, 0)
	for i := 0; i < list.Length(); i++ {
		if list.Data[i] == data {
			indexList = append(indexList, i)
		}
	}
	return indexList
}

// Del 删除数组中指定索引的元素。
func (list *MapArray[T]) Del(index int) {
	if len(list.Data) == 0 || index < 0 || index >= len(list.Data) {
		return
	}
	copy(list.Data[index:], list.Data[index+1:])
	list.Data = list.Data[:len(list.Data)-1]
	list.CapacityManager()
}

// Length 返回数组中元素的个数。
func (list *MapArray[T]) Length() int {
	return len(list.Data)
}

// Capacity 返回当前数组的容量。
func (list *MapArray[T]) Capacity() int {
	return cap(list.Data)
}

// CapacityManager 根据 SectionSize 阈值来管理数组的容量。
// 如果 SectionSize 是零，则让 slice 扩缩容接管。
func (list *MapArray[T]) CapacityManager() {
	if list.SectionSize == 0 {
		return // 如果 SectionSize 是零，则让 slice 扩缩容接管
	}

	// 计算当前容量和大小之间的差距
	diff := list.Capacity() - list.Length()

	// 计算扩容或缩容的阈值
	sc := list.Length() / list.SectionSize

	// 如果容量和大小的差超过了一个 SectionSize，则缩容数组
	if diff/list.SectionSize > 1 {
		newSlice := make([]T, len(list.Data), sc*list.SectionSize)
		copy(newSlice, list.Data)
		list.Data = newSlice
	}
	// 如果容量-大小小于一个 SectionSize，则扩容数组
	if diff < int(float64(list.SectionSize)*0.2) {
		newSlice := make([]T, len(list.Data), sc*list.SectionSize+list.SectionSize)
		copy(newSlice, list.Data)
		list.Data = newSlice
	}
}

// CapacityChanges 用于打印容量和大小的变化信息，以便调试。
func (list *MapArray[T]) CapacityChanges() {
	fmt.Println(fmt.Sprintf(
		`当前容量：%d, 当前大小：%d, 扩缩容阈值：%d`,
		list.Capacity(),
		list.Length(),
		list.SectionSize,
	))
}

// ToString 返回数组的字符串表示。
func (list *MapArray[T]) ToString() string {
	return fmt.Sprintf("[%s]", list.Join(","))
}

// Join 将数组元素连接为一个字符串。
func (list *MapArray[T]) Join(key string) string {
	str := ""
	for i := 0; i < list.Length(); i++ {
		str = fmt.Sprintf("%s%v%s", str, list.Data[i], key)
	}
	return str[:len(str)-1]
}

// IndexOf 返回数组中第一个匹配项的索引，若不存在则返回-1。
func (list *MapArray[T]) IndexValueOf(data T) int {
	result := -1
	for index, value := range list.Data {
		if value == data {
			result = index
			break
		}
	}
	return result
}

// Includes 判断数组是否包含特定元素，返回布尔值。
func (list *MapArray[T]) Includes(data T) bool {
	result := list.IndexValueOf(data)
	return result > -1
}

// Reverse 反转数组中元素的顺序。
func (list *MapArray[T]) Reverse() {
	for i := list.Length()/2 - 1; i >= 0; i-- {
		opp := list.Length() - 1 - i
		list.Data[i], list.Data[opp] = list.Data[opp], list.Data[i]
	}
}

// Count 统计元素出现次数。
func (list *MapArray[T]) Count(data T) int {
	result := 0
	for i := 0; i < list.Length(); i++ {
		if data == list.Data[i] {
			result += 1
		}
	}
	return result
}

// Concat 方法用于合并两个或多个数组。此方法不会更改现有数组，而是返回一个新数组。
func (list *MapArray[T]) Concat(data MapArray[T]) *MapArray[T] {
	sectionSize := 10
	if list.SectionSize > data.SectionSize {
		sectionSize = list.SectionSize
	} else {
		sectionSize = data.SectionSize
	}
	newArray := &MapArray[T]{
		SectionSize: sectionSize,
		Data:        list.Data,
	}
	newArray.Data = append(newArray.Data, data.Data...)
	return newArray
}

func (list *MapArray[T]) ForEach(f func(T)) {
	for _, v := range list.Data {
		f(v)
	}
}
