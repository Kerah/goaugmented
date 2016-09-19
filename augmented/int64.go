package augmented

import (
	"github.com/Kerah/goaugmented"
	"math"
)

type aInt64 int64

func NewInt64(in int64) (goaugmented.Value) {
	var v aInt64 = aInt64(in)
	return &v
}

func NewInt64Interval(start, end int64, id uint64) goaugmented.Interval {
	return goaugmented.SingleDimensionInterval(
		NewInt64(start),
		NewInt64(end),
		id,
	)
}

func NewInt64VI(value int64) goaugmented.Interval {
	return goaugmented.ValueInterval(NewInt64(value))
}

func (af *aInt64)  typeCast(in goaugmented.Value) int64 {
	switch in.(type) {
	case *aInt64:
		data := in.(*aInt64)
		return int64(*data)
	}
	return 0
}

func (af *aInt64) Greater(in goaugmented.Value) (r bool) {
	if int64(*af) > af.typeCast(in) {
		r = true
	}
	return
}

func (af *aInt64) GreaterOrEq(in goaugmented.Value) (r bool) {
	if int64(*af) >= af.typeCast(in) {
		r = true
	}
	return
}
func (af *aInt64) Lesser(in goaugmented.Value) (r bool) {
	if int64(*af) < af.typeCast(in) {
		r = true
	}
	return
}
func (af *aInt64) LesserOrEq(in goaugmented.Value) (r bool) {
	if int64(*af) <= af.typeCast(in) {
		r = true
	}
	return
}
func (af *aInt64) Substract(in goaugmented.Value) (r int64) {

	return int64(int64(*af) - af.typeCast(in))
}

func (af *aInt64) Add(in int64) (goaugmented.Value) {
	*af = aInt64(int64(*af)+ int64(in))
	return af
}

func (af *aInt64) MinimalValue() goaugmented.Value {
	df := aInt64(-1*math.MaxInt64)
	return &df
}
func (af *aInt64) MaximalValue() goaugmented.Value {
	df := aInt64(math.MaxInt64)
	return &df
}
