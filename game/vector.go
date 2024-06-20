package game

import "math"

/*
* 	     X
*     ------
*	 |\ ) = Angle
*   Y| \
*    |  \
*    |	 \
*	     __|
 */
type Vector struct {
	X float64
	Y float64
}

func (v *Vector) getAngle() float64 {
	return math.Acos(v.X / math.Hypot(v.X, v.Y))
}

func (v *Vector) getMagnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
