package extended

// {P: len(arr) > 0}
// {Q: result = max(arr) ^ result E arr}

// I (инвариант цикла): maxVal = max(arr[0..i-1]) ∧ 1 ≤ i ≤ len(arr)

func findMax(arr []int) int {
	// Шаг 1: Инициализация
	// Доказываем: { len(arr) > 0 } => { maxVal = arr[0] ∧ i=1 } => I выполняется
	maxVal := arr[0] // max(arr[0..0]) тривиально

	// Шаг 2: Цикл
	// Инвариант I: maxVal = max(arr[0..i-1]) ∧ 1 ≤ i ≤ len(arr)
	for i := 1; i < len(arr); i++ {
		// Сохраняем инвариант перед итерацией (I ∧ i < len(arr))
		current := arr[i]

		// Шаг 2.1: Обновление максимума
		if current > maxVal {
			// Доказываем: maxVal = max(arr[0..i-1]) ∧ current > maxVal =>
			// => newMaxVal = max(arr[0..i])
			maxVal = current
		}
		// После обновления: maxVal = max(arr[0..i])
	}
	// Шаг 3: Завершение
	// I ∧ i >= len(arr) =>
	// => maxVal = max(arr[0..len(arr)-1]) = max(arr)
	return maxVal
}
