package orders

import (
	"errors"
	"testing"
)

// MockOrderStore — mock реалізація OrderStore для тестування.
type MockOrderStore struct {
	ExecFunc  func(query string, args ...any) error
	LastQuery string
	LastArgs  []any
	CallCount int
}

// Exec реалізує OrderStore інтерфейс для mock.
func (m *MockOrderStore) Exec(query string, args ...any) error {
	m.CallCount++
	m.LastQuery = query
	m.LastArgs = args
	if m.ExecFunc != nil {
		return m.ExecFunc(query, args...)
	}
	return nil
}

// TestPlaceOrderSuccess — перевіряє успішне виконання PlaceOrder.
func TestPlaceOrderSuccess(t *testing.T) {
	mock := &MockOrderStore{
		ExecFunc: func(query string, args ...any) error {
			return nil
		},
	}
	service := NewOrderService(mock)

	err := service.PlaceOrder("ORD123", 99.99)

	if err != nil {
		t.Fatalf("PlaceOrder() returned error: %v", err)
	}

	if mock.CallCount != 1 {
		t.Errorf("Exec called %d times, want 1", mock.CallCount)
	}

	if mock.LastQuery != "insert into orders values (?, ?)" {
		t.Errorf("Exec called with query %q, want %q", mock.LastQuery, "insert into orders values (?, ?)")
	}

	if len(mock.LastArgs) != 2 {
		t.Fatalf("Exec called with %d args, want 2", len(mock.LastArgs))
	}

	if mock.LastArgs[0] != "ORD123" {
		t.Errorf("First arg is %v, want %q", mock.LastArgs[0], "ORD123")
	}

	if mock.LastArgs[1] != 99.99 {
		t.Errorf("Second arg is %v, want %v", mock.LastArgs[1], 99.99)
	}
}

// TestPlaceOrderError — перевіряє обробку помилок з store.
func TestPlaceOrderError(t *testing.T) {
	expectedErr := errors.New("database error")
	mock := &MockOrderStore{
		ExecFunc: func(query string, args ...any) error {
			return expectedErr
		},
	}
	service := NewOrderService(mock)

	err := service.PlaceOrder("ORD456", 50.00)

	if err != expectedErr {
		t.Errorf("PlaceOrder() returned %v, want %v", err, expectedErr)
	}

	if mock.CallCount != 1 {
		t.Errorf("Exec called %d times, want 1", mock.CallCount)
	}
}

// TestPlaceOrderMultipleCalls — перевіряє кілька послідовних викликів.
func TestPlaceOrderMultipleCalls(t *testing.T) {
	mock := &MockOrderStore{}
	service := NewOrderService(mock)

	orders := []struct {
		id     string
		amount float64
	}{
		{"ORD001", 10.50},
		{"ORD002", 25.00},
		{"ORD003", 99.99},
	}

	for _, order := range orders {
		err := service.PlaceOrder(order.id, order.amount)
		if err != nil {
			t.Fatalf("PlaceOrder() returned error: %v", err)
		}
	}

	if mock.CallCount != 3 {
		t.Errorf("Exec called %d times, want 3", mock.CallCount)
	}

	// Перевіриння останнього виклику
	if mock.LastArgs[0] != "ORD003" {
		t.Errorf("Last call first arg is %v, want %q", mock.LastArgs[0], "ORD003")
	}
}
