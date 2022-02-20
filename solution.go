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
		/* In This Case, SideLen Is Considered As A Radius */
		return math.Pi * math.Pow(sideLen, 2)
	} else if sidesNum == SidesTriangle {
		return (math.Sqrt(float64(sidesNum)) / 4) * sideLen
	} else if sidesNum == SidesSquare {
		return math.Pow(float64(sidesNum), 2)
	}

	return 0

}
