package util

func Push(a []int, v int) []int {
	return append(a, v)
}

func Pop(a []int) []int {
	return a[:len(a)-1] // slice from 0 to len-1
}

func Unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

func Shift(a []int) []int {
	return a[1:]
}

func Insert(a []int, v int, p int) []int {
	// sliceは参照取得なので、前後で変数に格納するなどしても参照ごと値が変わるため、うまくいかない
	a = append(a[:p+1], a[p:len(a)]...)
	a[p] = v
	return a
}

func Remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}

