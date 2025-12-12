package geom

import (
	"testing"
)

func TestMatrix(t *testing.T) {
	t.Run("Multiplication", func(t *testing.T) {
		m1 := InitMatrix([][]float64{
			{4, 2},
			{9, 0},
		})

		m2 := InitMatrix([][]float64{
			{3, 1},
			{-3, 4},
		})

		expectedRes := InitMatrix([][]float64{
			{6, 12},
			{27, 9},
		})

		res, err := m1.Mul(&m2)
		if err != nil {
			t.Errorf("Error multiplying matrices: %v", err)
		}

		if !res.Equals(expectedRes) {
			t.Errorf("Multiplication result is not equal to expected result: %v != %v", res, expectedRes)
		}

		m1 = InitMatrix([][]float64{
			{2, 1},
			{-3, 0},
			{4, -1},
		})

		m2 = InitMatrix([][]float64{
			{5, -1, 6},
			{-3, 0, 7},
		})

		expectedRes = InitMatrix([][]float64{
			{7, -2, 19},
			{-15, 3, -18},
			{23, -4, 17},
		})

		res, err = m1.Mul(&m2)
		if err != nil {
			t.Errorf("Error multiplying matrices: %v", err)
		}

		if !res.Equals(expectedRes) {
			t.Errorf("Multiplication result is not equal to expected result: %v != %v", res, expectedRes)
		}
	})

	t.Run("Inversion", func(t *testing.T) {
		m1 := InitMatrix([][]float64{
			{1, 2, -3},
			{3, 2, -4},
			{2, -1, 0},
		})
		t.Log(m1)

		m2, err := m1.Inverse()

		if err != nil {
			t.Errorf("Error inversing matrices: %v", err)
		}

		expectedRes := UnitMatrix(3)

		res, err := m1.Mul(&m2)

		if err != nil {
			t.Errorf("Error multiplying matrices: %v", err)
		}

		if !res.Equals(expectedRes) {
			t.Errorf("Multiplication result is not equal to expected result: %v != %v", res, expectedRes)
		}

	})
}
