# Dsa
> 一个新手的数据结构与算法工具库
## 数据结构(Data Structure)
目前计划是数组、链表、栈、队列，期望下个版本可以是跳表、哈希和各种树
### 数组(datastructures.MapArray)
基于切片实现，除了NewArray和NewMapArrayFormMap创建一个数组之外，其他的都是实例.xx。如:
```go
	fmt.Println("hello word")
	arr1 := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 5, 6, 3,
	)
	arr2 := datastructures.NewMapArray[int](10)
	arr1.Push(1)
	arr2.Push(1)
	fmt.Println(arr1.ToString(), arr2.ToString())
// 输出：
// hello word
// [1,2,3,4,5,5,6,3,1] [1]
```
MapArray的数据结构是这样的
```go
// MapArray 是一个带有容量管理的动态数组。
type MapArray[T constraints.Ordered] struct {
	Data        []T // Data 保存数组元素。
	SectionSize int // SectionSize 是容量管理的阈值。
}

```
添加或者删除会自动调用容量管理，查看当前是否满足扩缩容条件。当数组长度%小于阈值本身的时候会自动缩容，当当前数组长度小于阈值的20%会自动扩容一个阈值。
> 如果传入了0就交给原生slice本身管理。
#### Push 

在数组末尾添加一个元素

func (list *MapArray[T]) Push(data T)
```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
fmt.Println(list.ToString())
// 此时输出
[1]
list.Push(2)
fmt.Println(list.ToString())
// 此时输出
// [1,2]
list.CapacityChanges()
// 此时输出
// 当前容量：10, 当前大小：2, 扩缩容阈值：10
```

#### Pop 

从数组中删除最后一个元素

func (list *MapArray[T]) Pop()

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.Pop()
fmt.Println(list.ToString())
// 此时输出
[1,2]
```

#### Unshift 

在数组开头添加一个元素。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)

list.Unshift(0)
fmt.Println(list.ToString())
// 此时输出
[0,1,2]
```

#### Splice 

在指定索引处插入一个元素。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.Splice(1, 100)
fmt.Println(list.ToString())
// 此时输出
[1,100,2,3]
```

#### FindAll 

查找数组中所有与指定元素相等的元素，并返回它们的索引。

```go
list := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 6, 7, 2, 6, 8, 7,
)
indexList := list.FindAll(2)
fmt.Println(indexList)
// 此时输出
[1 7]
```

#### Del 

删除数组中指定索引的元素。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.Del(1)
fmt.Println(list.ToString())
// 此时输出
[1,3]
```

#### Length 

返回数组中元素的个数。

```go
list := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 6, 7, 2, 6, 8, 7,
)

length := list.Length()
fmt.Println(length)
// 此时输出
11
```

#### Capacity 

返回当前数组的容量。

```go
list := datastructures.NewMapArrayFormMap[int](
		10,
		1, 2, 3, 4, 5, 6, 7, 2, 6, 8, 7,
	)
fmt.Println(list.Capacity())
// 此时输出
200
```

#### CapacityManager 

根据 SectionSize 阈值来管理数组的容量。

> 这里一般是在添加或者删除的时候自动触发，无需调用。无论在使用MapArray负责扩容还是slice在负责扩容时都是这样

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.CapacityManager()
```

#### CapacityChanges 

用于打印容量和大小的变化信息，以便调试。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.CapacityChanges()
// 此时输出
// 当前容量：10, 当前大小：3, 扩缩容阈值：10
```

#### ToString 

返回数组的字符串表示。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

str := list.ToString()
fmt.Println(str)
// 此时输出
"[1,2,3]"
```

#### Join 

将数组元素连接为一个字符串。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

str := list.Join(",")
fmt.Println(str)
// 此时输出
"1,2,3"
```

#### IndexOf 

返回数组中第一个匹配项的索引，若不存在则返回-1。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

index := list.IndexOf(2)
fmt.Println(index)
// 此时输出
1
```

#### Includes 

判断数组是否包含特定元素，返回布尔值。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

included := list.Includes(2)
fmt.Println(included)
// 此时输出
true
```

#### Reverse 

反转数组中元素的顺序。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.Reverse()
fmt.Println(list.ToString())
// 此时输出
[3,2,1]
```

#### Count

统计元素出现次数。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

count := list.Count(2)
fmt.Println(count)
// 此时输出
1
```

#### Concat

合并两个数组，返回一个新数组，不改变现有数组。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.Reverse()
fmt.Println(list.ToString())
// 此时输出
[3,2,1]
```

#### ForEach

对数组中的每个元素应用给定的函数。

```go
list := datastructures.NewMapArray[int](10)
list.Push(1)
list.Push(2)
list.Push(3)

list.ForEach(func(data int) {
	fmt.Println(data)
})
// 此时输出
// 1
// 2
// 3

```

