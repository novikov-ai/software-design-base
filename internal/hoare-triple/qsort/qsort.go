package qsort

// {P: arr != nil ^ len(arr) > 0}
// {Q: arr отсортирован}
func quickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

// {P: low < high}
func sort(arr []int, low, high int) {
	if low < high {
		// Закон partition() {Q: p - правильная позиция}
		p := partition(arr, low, high)

		// Рекурсия для левой части
		// {P: подмассив [low, p-1] существует}
		sort(arr, low, p-1)
		// {Q: [low, p-1] отсортирован}

		// Рекурсия для правой части
		// {P: подмассив [p+1, high] существует}
		sort(arr, p+1, high)
		// {Q: [p+1, high] отсортирован}

		// Собираем результат:
		// 1. Слева отсортировано и ≤ arr[p]
		// 2. Справа отсортировано и ≥ arr[p]
		// 3. arr[p] на месте → весь [low, high] отсортирован!
	}
	// {Q: [low, high] отсортирован}
}

// {P: low < high}
func partition(arr []int, low, high int) int {
	pivot := arr[high] // опорный элемент
	i := low - 1

	// Инвариант цикла:
	// ∀k∈[low, i]: arr[k] ≤ pivot
	// ∀k∈[i+1, j): arr[k] > pivot
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Ставим pivot на правильное место
	arr[i+1], arr[high] = arr[high], arr[i+1]

	// {Q:
	//   ∀k∈[low, i+1]: arr[k] ≤ arr[i+1] ∧
	//   ∀k∈[i+1, high]: arr[k] ≥ arr[i+1]
	// }
	return i + 1
}
