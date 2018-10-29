/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package point

import "math"

type Point struct {
	x int
	y int
	Z int
}

/*
OLIWI
*/
func New(x, y int) Point {
	return Point{x, y, 0}
}

func (v Point) GetX() int {
	return v.x
}

func (v Point) GetY() int {
	return v.y
}

func (v Point) Abs() float64 {
	return math.Sqrt(float64(v.x*v.x + v.y*v.y))
}
