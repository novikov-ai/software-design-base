package grade_calculator

// GradeCalculator структура для расчёта среднего балла
type GradeCalculator struct{}

// CalculateAverage вычисляет среднее значение оценок
func (gc *GradeCalculator) CalculateAverage(grades []float64) float64 {
	if len(grades) == 0 {
		return 0.0 // Обработка пустого списка
	}
	
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	return sum / float64(len(grades))
}