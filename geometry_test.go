package geometry

import (
	"math"
	"testing"
)

func TestRect(t *testing.T) {
	r := rect{
		width:  5,
		height: 8,
	}
	if r.area() != 40 {
		t.Fatalf("area() = %f, want 40", r.area())
	}
	if r.perim() != 26 {
		t.Fatalf("perim() = %f, want 26", r.perim())
	}
}

func TestCircle(t *testing.T) {
	c := circle{
		radius: 10,
	}
	wantArea := math.Round(314.159265)
	gotArea := math.Round(c.area())
	if gotArea != wantArea {
		t.Fatalf("area() = %f, want %f", gotArea, wantArea)
	}
	wantPerim := math.Round(62.831853)
	gotPerim := math.Round(c.perim())
	if gotPerim != wantPerim {
		t.Fatalf("area() = %f, want %f", gotPerim, wantPerim)
	}
}
