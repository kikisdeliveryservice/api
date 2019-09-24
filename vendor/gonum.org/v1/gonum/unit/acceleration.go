// Code generated by "go generate gonum.org/v1/gonum/unit”; DO NOT EDIT.

// Copyright ©2014 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unit

import (
	"errors"
	"fmt"
	"math"
	"unicode/utf8"
)

// Acceleration represents an acceleration in metres per second squared.
type Acceleration float64

// Unit converts the Acceleration to a *Unit
func (a Acceleration) Unit() *Unit {
	return New(float64(a), Dimensions{
		LengthDim: 1,
		TimeDim:   -2,
	})
}

// Acceleration allows Acceleration to implement a Accelerationer interface
func (a Acceleration) Acceleration() Acceleration {
	return a
}

// From converts the unit into the receiver. From returns an
// error if there is a mismatch in dimension
func (a *Acceleration) From(u Uniter) error {
	if !DimensionsMatch(u, Acceleration(0)) {
		*a = Acceleration(math.NaN())
		return errors.New("Dimension mismatch")
	}
	*a = Acceleration(u.Unit().Value())
	return nil
}

func (a Acceleration) Format(fs fmt.State, c rune) {
	switch c {
	case 'v':
		if fs.Flag('#') {
			fmt.Fprintf(fs, "%T(%v)", a, float64(a))
			return
		}
		fallthrough
	case 'e', 'E', 'f', 'F', 'g', 'G':
		p, pOk := fs.Precision()
		w, wOk := fs.Width()
		const unit = " m s^-2"
		switch {
		case pOk && wOk:
			fmt.Fprintf(fs, "%*.*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), p, float64(a))
		case pOk:
			fmt.Fprintf(fs, "%.*"+string(c), p, float64(a))
		case wOk:
			fmt.Fprintf(fs, "%*"+string(c), pos(w-utf8.RuneCount([]byte(unit))), float64(a))
		default:
			fmt.Fprintf(fs, "%"+string(c), float64(a))
		}
		fmt.Fprint(fs, unit)
	default:
		fmt.Fprintf(fs, "%%!%c(%T=%g m s^-2)", c, a, float64(a))
	}
}