package hoare_triple

// {P: true}
// {Q: (x < 0 -> result = -x) ^ (x >= 0 -> result = x) ^ result >= 0}
func abs(x int) int {
	if x < 0 {
		return -x // Корректно: -x > 0 при x < 0
	}

	return x // Корректно: x >= 0
}

// {P: true}
// {Q: (result = a v result = b) ^ result >= a ^ result >= b}
func max(a, b int) int {
	if a >= b {
		return a // Корректно: a >= b
	}
	return b // Корректно: b < a
}

// {P: true}
// {Q: result = max(|a|, |b|) ^ result >= |a| ^ result >= |b|}
// При |a| >= |b| возвращаем a - корректно
// При |a| < |b| возвращаем b- корректно
func maxAbs(a, b int) int {
	return max(abs(a), abs(b))
}
