package infra

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(test *testing.T) {
	tt := []struct {
		argOne int
		argTwo int
		result int
	}{
		{
			argOne: 1,
			argTwo: 2,
			result: 3,
		},
		{
			argOne: -1,
			argTwo: 1,
			result: 0,
		},
		{
			argOne: math.MaxInt8,
			argTwo: 1,
			result: 1 << 7,
		}, {
			argOne: math.MaxInt16,
			argTwo: 1,
			result: 1 << 15,
		},
	}
	for _, t := range tt {
		var result int
		result = Add(t.argOne, t.argTwo)
		if result != t.result {
			test.Errorf("wrong: result=%d actual=%d", result, t.result)
		}
	}
}

func TestAdd2(t *testing.T) {
	var result int
	result = Add(1, 2)
	if result != 3 {
		t.Errorf("wrong: result=%d actual=%d", result, 3)
	}
}

func TestAdd3(t *testing.T) {
	Convey("Testing Add", t, func() {
		tt := []struct {
			a int
			b int
			c int
		}{
			{
				a: 1,
				b: 2,
				c: 3,
			},
			{
				a: 4,
				b: 5,
				c: 9,
			},
		}
		So(Add(tt[0].a, tt[0].b), ShouldEqual, tt[0].c)
		So(Add(tt[1].a, tt[1].b), ShouldEqual, tt[1].c)
	})
}
