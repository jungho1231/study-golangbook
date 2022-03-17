package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	got := Calc()

	t.Run("클로저 테스트 - 파라미터 값이 1 인 경우 ", func(t *testing.T) {

		result := got(1)
		want := 8

		assertNumber(t, result, want)
	})

	t.Run("클로저 테스트 - 파라미터 값이 2 인 경우 ", func(t *testing.T) {

		result := got(2)
		want := 11

		assertNumber(t, result, want)
	})

}

func assertNumber(t *testing.T, result int, want int) {
	if result != want {
		t.Errorf("result %d want %d", result, want)
	}
}
