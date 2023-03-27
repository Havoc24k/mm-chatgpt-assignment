package main

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

// Reverser component.
type Mathserve interface {
	Add(context.Context, int, int) (int, error)
	Sub(context.Context, int, int) (int, error)
}

type mathserve struct {
	weaver.Implements[Mathserve]
}

func (m *mathserve) Add(_ context.Context, x, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, fmt.Errorf("Zero values")
	}

	result := 0

	for i := 0; i < x; i++ {
		result += y
	}

	return result, nil
}

func (m *mathserve) Sub(_ context.Context, x, y int) (int, error) {
	if x == 0 || y == 0 {
		return 0, fmt.Errorf("Zero values")
	}

	result := 0

	for i := 0; i < x; i++ {
		result -= y
	}

	return result, nil
}
