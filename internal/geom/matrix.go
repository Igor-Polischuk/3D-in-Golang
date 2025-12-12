package geom

import (
	"errors"
	"math"
	"reflect"
)

type Matrix struct {
	M    [][]float64
	R, C int8
}

func NewMatrix(c, r int8) Matrix {
	M := make([][]float64, r)

	for i := range r {
		M[i] = make([]float64, c)
	}

	return Matrix{M: M, C: c, R: r}
}

func InitMatrix(mArr [][]float64) Matrix {
	r := len(mArr)
	c := len(mArr[0])

	m := NewMatrix(int8(c), int8(r))

	m.M = mArr

	return m
}

func UnitMatrix(n int8) Matrix {
	m := NewMatrix(n, n)
	for i := range n {
		m.M[i][i] = 1
	}
	return m
}

func mulMatRowAndCol(r, c []float64) float64 {
	res := 0.0
	for i := range len(r) {
		res += r[i] * c[i]
	}

	return res
}

func (m1 Matrix) Mul(m2 *Matrix) Matrix {
	if m1.C != m2.R {
		panic("can't multiply matrices with incompatible dimensions")
	}

	res := NewMatrix(m2.C, m1.R)

	for i := range res.R {
		for j := range res.C {
			row := m1.M[i]
			column := m2.GetCol(j)
			res.M[i][j] = mulMatRowAndCol(row, column)

		}
	}

	return res
}

func (m *Matrix) Equals(other Matrix) bool {
	return reflect.DeepEqual(m.M, other.M)
}

func (m *Matrix) GetCol(c int8) []float64 {
	if c > m.C-1 || c < 0 {
		return []float64{}
	}

	res := make([]float64, m.R)

	for i := range m.R {
		res[i] = m.M[i][c]
	}

	return res
}

func (m *Matrix) Transpose() Matrix {
	res := NewMatrix(m.R, m.C)

	for i := range m.R {
		for j := range m.C {
			res.M[j][i] = m.M[i][j]
		}
	}
	return res
}

func (m *Matrix) Det() (float64, error) {
	if m.R != m.C {
		return 0, errors.New("cannot calculate determinant of a non-square matrix")
	}
	n := m.R

	if n == 2 {
		return m.M[0][0]*m.M[1][1] - m.M[0][1]*m.M[1][0], nil
	}

	if n == 3 {
		a, b, c := m.M[0][0], m.M[0][1], m.M[0][2]
		d, e, f := m.M[1][0], m.M[1][1], m.M[1][2]
		g, h, i := m.M[2][0], m.M[2][1], m.M[2][2]

		det := a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
		return det, nil
	}

	// TODO
	return 0, errors.New("determinant calculation for N > 3 not implemented")
}

func (m *Matrix) minor(r, c int8) Matrix {
	if m.R <= 1 || m.C <= 1 {
		return Matrix{}
	}

	res := NewMatrix(m.C-1, m.R-1)

	resRow, resCol := int8(0), int8(0)

	for i := range m.R {
		if i == r {
			continue
		}
		resCol = 0
		for j := range m.C {
			if j == c {
				continue
			}
			res.M[resRow][resCol] = m.M[i][j]
			resCol++
		}
		resRow++
	}
	return res
}

func (m *Matrix) CofactorMatrix() (Matrix, error) {
	if m.R != m.C {
		return Matrix{}, errors.New("cannot calculate cofactor matrix for a non-square matrix")
	}

	n := m.R
	cofactor := NewMatrix(n, n)

	for i := range n {
		for j := range n {
			subMatrix := m.minor(i, j)

			detMinor, err := subMatrix.Det()
			if err != nil {
				return Matrix{}, err
			}

			sign := 1.0
			if (i+j)%2 != 0 {
				sign = -1.0
			}

			cofactor.M[i][j] = sign * detMinor
		}
	}
	return cofactor, nil
}

// M⁻¹ = (1/det(M)) * adj(M)
func (m *Matrix) Inverse() (Matrix, error) {
	if m.R != m.C {
		return Matrix{}, errors.New("only square matrices can be inverted")
	}
	if m.R > 3 {
		return Matrix{}, errors.New("inversion implemented only for 2x2 and 3x3 matrices")
	}

	det, err := m.Det()
	if err != nil {
		return Matrix{}, err
	}

	if math.Abs(det) < 1e-9 {
		return Matrix{}, errors.New("matrix is singular and cannot be inverted")
	}

	cofactor, err := m.CofactorMatrix()
	if err != nil {
		return Matrix{}, err
	}

	adjugate := cofactor.Transpose()

	//  M⁻¹ = (1/det(M)) * adj(M)
	invDet := 1.0 / det

	inverse := NewMatrix(m.C, m.R)
	for i := range m.R {
		for j := range m.C {
			inverse.M[i][j] = adjugate.M[i][j] * invDet
		}
	}

	return inverse, nil
}
