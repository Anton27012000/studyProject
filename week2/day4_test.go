package week2

import "testing"

func TestDiscountedPrice(t *testing.T) {

	product := Product{
		0,
		"Milk",
		1000,
		0.2,
		true,
	}

	result := product.DiscountedPrice()
	expected := 800.0

	if result != expected {
		t.Errorf("expected %.1f, got %.1f", expected, result)
	}
}

func TestAveragePriceEmpty(t *testing.T) {

	products := []Product{}

	result := AveragePrice(products)

	expected := 0.0

	if result != expected {
		t.Errorf(
			"expected %.1f, got %.1f",
			expected,
			result,
		)
	}
}

func TestTaskIsDone(t *testing.T) {

	task := Task{
		ID:       1,
		Title:    "Learn Go",
		Done:     true,
		Priority: 1,
	}

	if !task.IsDone() {
		t.Error("expected task to be done")
	}
}
