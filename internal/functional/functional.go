package functional

import "sync"

type Order struct {
	ID     int
	Amount float64
	Active bool
}

// ProcessOrdersImperative - традиционный подход с мутацией состояния
func ProcessOrdersImperative(orders []*Order, discountRate float64) float64 {
	var total float64
	var wg sync.WaitGroup
	mu := sync.Mutex{}

	for _, order := range orders {
		if !order.Active {
			continue
		}

		wg.Add(1)
		go func(o *Order) {
			defer wg.Done()

			// Прямая мутация объекта
			o.Amount *= (1 - discountRate)

			mu.Lock()
			total += o.Amount
			mu.Unlock()
		}(order)
	}

	wg.Wait()
	return total
}

// ---------------------------------------------------------------------------

// ProcessOrdersFunctional - функциональный подход
func ProcessOrdersFunctional(orders []Order, discountRate float64) float64 {
	// Чистый конвейер обработки
	activeOrders := filterActive(orders)
	discountedOrders := applyDiscountConcurrent(activeOrders, discountRate)
	return calculateTotal(discountedOrders)
}

// FunctionalOrder - иммутабельная DTO для обработки
type FunctionalOrder struct {
	ID     int
	Amount float64
}

// Чистая функция фильтрации
func filterActive(orders []Order) []FunctionalOrder {
	result := make([]FunctionalOrder, 0, len(orders))
	for _, o := range orders {
		if o.Active {
			result = append(result, FunctionalOrder{
				ID:     o.ID,
				Amount: o.Amount,
			})
		}
	}
	return result
}

// Конкурентное применение скидки без мутаций
func applyDiscountConcurrent(orders []FunctionalOrder, discountRate float64) []FunctionalOrder {
	type result struct {
		index int
		order FunctionalOrder
	}

	ch := make(chan result, len(orders))
	var wg sync.WaitGroup

	for i, o := range orders {
		wg.Add(1)
		go func(idx int, order FunctionalOrder) {
			defer wg.Done()
			// Иммутабельное преобразование
			ch <- result{
				index: idx,
				order: FunctionalOrder{
					ID:     order.ID,
					Amount: order.Amount * (1 - discountRate),
				},
			}
		}(i, o)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Сохранение порядка обработки
	results := make([]FunctionalOrder, len(orders))
	for r := range ch {
		results[r.index] = r.order
	}

	return results
}

// Чистая агрегация
func calculateTotal(orders []FunctionalOrder) float64 {
	total := 0.0
	for _, o := range orders {
		total += o.Amount
	}
	return total
}