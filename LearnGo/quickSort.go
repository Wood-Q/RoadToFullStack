package main

//func partition(a []int) int {
//	pivot := a[len(a)-1]
//	i := -1
//	for j := 0; j < len(a); j++ {
//		if a[j] < pivot {
//			i++
//			a[j], a[i] = a[i], a[j]
//		}
//	}
//	a[len(a)-1], a[i+1] = a[i+1], a[len(a)-1]
//	return i + 1
//}
//
//func quickSort(a []int) {
//	if len(a) <= 1 {
//		return
//	}
//	pivot := partition(a)
//	quickSort(a[:pivot])
//	quickSort(a[pivot+1:])
//}
//
//func main() {
//	arr := []int{5, 2, 3, 4, 6, 7}
//	quickSort(arr)
//	fmt.Println(arr)
//}

func partition(arr []int) int {
	pivot := arr[len(arr)-1]
	i := -1
	for j := 0; j < len(arr); j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	pivot := partition(arr)
	quickSort(arr[:pivot])
	quickSort(arr[pivot+1:])
}
