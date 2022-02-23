package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type squareInt = uint32

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
const (
	SidesCircle   = iota
	SidesTriangle = iota + 2
	SidesSquare
)

func CalcSquare(sideLen float64, sidesNum squareInt) float64 {
	if sidesNum == SidesCircle {
		return math.Pi * (sideLen * sideLen)
	} else if sidesNum == SidesTriangle {
		return (math.Sqrt(3.00) / 4) * (math.Pow(sideLen, 2))
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	}

	return 0

}
