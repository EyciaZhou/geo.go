package geo

import (
	"fmt"
	"math"
)

var (
	eps = 1e-8
)

type Mat3x3 [9]float64

func (m *Mat3x3) Mul(m2 *Mat3x3) *Mat3x3 {
	var r Mat3x3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				r[i*3+j] += m[i*3+k] * m2[k*3+j]
			}
		}
	}
	return &r
}

func Mul(m1 *Mat3x3, m2 *Mat3x3) *Mat3x3 {
	return m1.Mul(m2)
}

func (m *Mat3x3) WithPo(x float64, y float64) *Mat3x3 {
	return Move(-x, -y).Mul(m).Mul(Move(x, y))
}

func (m *Mat3x3) Fix() {
	for i := 0; i < 9; i++ {
		if math.Abs(m[i]) < eps {
			m[i] = 0
		}
	}
}

func (m *Mat3x3) String() string {
	var st [9]string
	for i := 0; i < 9; i++ {
		st[i] = fmt.Sprintf("%.3f", m[i])
	}
	le := 0
	for i := 0; i < 9; i++ {
		if len(st[i]) > le {
			le = len(st[i])
		}
	}
	format := "%" + fmt.Sprintf("%d", le+2) + "s"
	for i := 0; i < 9; i++ {
		st[i] = fmt.Sprintf(format, st[i])
	}
	resultFormat := "/ %s %s %s \\\n| %s %s %s |\n\\ %s %s %s /"
	return fmt.Sprintf(resultFormat, st[0], st[1], st[2], st[3],
		st[4], st[5], st[6], st[7], st[8])
}

func (m *Mat3x3) Apply(x float64, y float64) (float64, float64) {
	t1 := x*m[0] + y*m[3] + m[6]
	t2 := x*m[1] + y*m[4] + m[7]
	return t1, t2
}

func (m *Mat3x3) Deter() float64 {
	return m[0]*(m[4]*m[8]-m[5]*m[7]) + m[1]*(m[5]*m[6]-m[3]*m[8]) + m[2]*(m[3]*m[7]-m[4]*m[6])
}

func Cross(x1, y1, x2, y2 float64) float64 {
	return x1*y2 - y1*x2
}

func (m *Mat3x3) Inv() *Mat3x3 {
	return NewMat3x3(
		Cross(m[4], m[5], m[7], m[8]), -Cross(m[1], m[2], m[7], m[8]), Cross(m[1], m[2], m[4], m[5]),
		-Cross(m[3], m[5], m[6], m[8]), Cross(m[0], m[2], m[6], m[8]), -Cross(m[0], m[2], m[3], m[5]),
		Cross(m[3], m[4], m[6], m[7]), -Cross(m[0], m[1], m[6], m[7]), Cross(m[0], m[1], m[3], m[4])).Div(m.Deter())
}

func (m *Mat3x3) Div(a float64) *Mat3x3 {
	invA := 1 / a
	return NewMat3x3(m[0]*invA, m[1]*invA, m[2]*invA, m[3]*invA, m[4]*invA, m[5]*invA, m[6]*invA, m[7]*invA, m[8]*invA)
}

func NewMat3x3(a1, a2, a3, b1, b2, b3, c1, c2, c3 float64) *Mat3x3 {
	return &Mat3x3{
		a1, a2, a3, b1, b2, b3, c1, c2, c3,
	}
}

func One() *Mat3x3 {
	return &Mat3x3{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func Move(tx, ty float64) *Mat3x3 {
	return NewMat3x3(1, 0, 0, 0, 1, 0, tx, ty, 1)
}

func Scale(sx, sy float64) *Mat3x3 {
	return NewMat3x3(sx, 0, 0, 0, sy, 0, 0, 0, 1)
}
func Rotate(theta float64) *Mat3x3 {

	return NewMat3x3(math.Cos(theta), math.Sin(theta), 0, -math.Sin(theta), math.Cos(theta), 0, 0, 0, 1)
}
func RotateWithPo(theta float64, x, y float64) *Mat3x3 {
	return Rotate(theta).WithPo(x, y)
}

func RotateClockwise(theta float64) *Mat3x3 {

	return NewMat3x3(-math.Cos(theta), math.Sin(theta), 0, -math.Sin(theta), -math.Cos(theta), 0, 0, 0, 1)
}

func (m *Mat3x3) Move(tx, ty float64) *Mat3x3 {
	return m.Mul(Move(tx, ty))
}

func (m *Mat3x3) Scale(sx, sy float64) *Mat3x3 {
	return m.Mul(Scale(sx, sy))
}
func (m *Mat3x3) Rotate(theta float64) *Mat3x3 {
	return m.Mul(Rotate(theta))
}
func (m *Mat3x3) RotateWithPo(theta float64, x, y float64) *Mat3x3 {
	return m.Mul(RotateWithPo(theta, x, y))
}

func (m *Mat3x3) RotateClockwise(theta float64) *Mat3x3 {
	return m.Mul(RotateClockwise(theta))
}

var (
	SymAboutX   = NewMat3x3(1, 0, 0, 0, -1, 0, 0, 0, 1)
	SymAboutY   = NewMat3x3(-1, 0, 0, 0, 1, 0, 0, 0, 1)
	SymAboutO   = NewMat3x3(-1, 0, 0, 0, -1, 0, 0, 0, 1)
	SymAboutXY  = NewMat3x3(0, 1, 0, 1, 0, 0, 0, 0, 1)
	SymAboutXfY = NewMat3x3(0, -1, 0, -1, 0, 0, 0, 0, 1)
)
