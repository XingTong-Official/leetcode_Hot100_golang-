package main

type BigHeapStruct struct {
	//Heap[0]不使用,方便计算
	Heap   []int
	length int
	m      map[int][]int
}

func (BigHeapStruct) left(i int) int {
	return 2 * i
}
func (BigHeapStruct) right(i int) int {
	return 2*i + 1
}
func (BigHeapStruct) parent(i int) int {
	return i / 2
}
func build(nums []int) BigHeapStruct {

	var heapStruct BigHeapStruct
	heapStruct.m = make(map[int][]int)
	heapStruct.length = len(nums)
	heapStruct.Heap = make([]int, heapStruct.length+1)
	for i, num := range nums {
		heapStruct.Heap[i+1] = num
		heapStruct.m[num] = append(heapStruct.m[num], i+1)
	}
	for i := heapStruct.length / 2; i >= 1; i-- {
		heapStruct.MAX_HEAPIFY(i)
	}

	return heapStruct
}
func (heapStruct *BigHeapStruct) add(key int) {
	heapStruct.Heap = append(heapStruct.Heap, key)
	heapStruct.length++
	heapStruct.addM(heapStruct.length)
	heapStruct.increaseKey(heapStruct.length, key)
}
func (heapStruct *BigHeapStruct) delete(key int) {
	index := heapStruct.m[key][0]
	heapStruct.changeM(index, heapStruct.length)
	heapStruct.Heap[heapStruct.length], heapStruct.Heap[index] = heapStruct.Heap[index], heapStruct.Heap[heapStruct.length]
	heapStruct.deleteM(heapStruct.length)
	heapStruct.length--
	if index > 1 && heapStruct.Heap[index] > heapStruct.Heap[heapStruct.parent(index)] {
		heapStruct.increaseKey(index, heapStruct.Heap[index])
	} else if index >= 1 {
		heapStruct.MAX_HEAPIFY(index)
	}
	heapStruct.Heap = heapStruct.Heap[:heapStruct.length+1]

}
func (heapStruct *BigHeapStruct) getFront() int {
	return heapStruct.Heap[1]
}
func (heapStruct *BigHeapStruct) increaseKey(i, key int) {
	heapStruct.Heap[i] = key
	for i > 1 && heapStruct.Heap[heapStruct.parent(i)] < heapStruct.Heap[i] {
		heapStruct.changeM(i, heapStruct.parent(i))
		heapStruct.Heap[heapStruct.parent(i)], heapStruct.Heap[i] = heapStruct.Heap[i], heapStruct.Heap[heapStruct.parent(i)]
		i = heapStruct.parent(i)
	}
}
func (heapStruct *BigHeapStruct) MAX_HEAPIFY(i int) {
	max := i
	if heapStruct.left(i) <= heapStruct.length && heapStruct.Heap[heapStruct.left(i)] > heapStruct.Heap[max] {
		max = heapStruct.left(i)
	}
	if heapStruct.right(i) <= heapStruct.length && heapStruct.Heap[heapStruct.right(i)] > heapStruct.Heap[max] {
		max = heapStruct.right(i)
	}
	if max != i {
		heapStruct.changeM(i, max)
		heapStruct.Heap[i], heapStruct.Heap[max] = heapStruct.Heap[max], heapStruct.Heap[i]
		heapStruct.MAX_HEAPIFY(max)
	}
}
func (heapStruct *BigHeapStruct) changeM(i, j int) {
	vi := heapStruct.Heap[i]
	vj := heapStruct.Heap[j]
	for idx, num := range heapStruct.m[vi] {
		if num == i {
			heapStruct.m[vi][idx] = j
			break
		}
	}
	for idx, num := range heapStruct.m[vj] {
		if num == j {
			heapStruct.m[vj][idx] = i
			break
		}
	}
}

func (heapStruct *BigHeapStruct) deleteM(i int) {
	for index, num := range heapStruct.m[heapStruct.Heap[i]] {
		if num == i {
			heapStruct.m[heapStruct.Heap[i]] = append(heapStruct.m[heapStruct.Heap[i]][:index], heapStruct.m[heapStruct.Heap[i]][index+1:]...)
			if len(heapStruct.m[heapStruct.Heap[i]]) == 0 {
				delete(heapStruct.m, heapStruct.Heap[i])
			}
		}
	}
}
func (heapStruct *BigHeapStruct) addM(i int) {
	heapStruct.m[heapStruct.Heap[i]] = append(heapStruct.m[heapStruct.Heap[i]], i)
}
