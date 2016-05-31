package augmented

import (
	"github.com/Kerah/goaugmented"
	"math"
)

type augmentedFloat64 float64

func NewFloat64(in float64) (goaugmented.Value) {
	var v augmentedFloat64 = augmentedFloat64(in)
	return &v
}

func NewFloat64Interval(start, end float64, id uint64) goaugmented.Interval {
	return goaugmented.SingleDimensionInterval(
		NewFloat64(start),
		NewFloat64(end),
		id,
	)
}

func NewFloat64VI(value float64) goaugmented.Interval {
	return goaugmented.ValueInterval(NewFloat64(value))
}


func (af *augmentedFloat64)  typeCast(in goaugmented.Value) float64 {
	switch in.(type) {
	case *augmentedFloat64:
		data := in.(*augmentedFloat64)
		return float64(*data)
	}
	return 0
}

func (af *augmentedFloat64) Greater(in goaugmented.Value) (r bool) {
	if float64(*af) > af.typeCast(in) {
		r = true
	}
	return
}

func (af *augmentedFloat64) GreaterOrEq(in goaugmented.Value) (r bool) {
	if float64(*af) >= af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) Lesser(in goaugmented.Value) (r bool) {
	if float64(*af) < af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) LesserOrEq(in goaugmented.Value) (r bool) {
	if float64(*af) <= af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) Substract(in goaugmented.Value) (r int64) {

	return int64(float64(*af) - af.typeCast(in))
}

func (af *augmentedFloat64) Add(in int64) (goaugmented.Value) {
	*af = augmentedFloat64(float64(*af)+ float64(in))
	return af
}

func (af *augmentedFloat64) MinimalValue() goaugmented.Value {
	df := augmentedFloat64(-1*math.MaxFloat64)
	return &df
}
func (af *augmentedFloat64) MaximalValue() goaugmented.Value {
	df := augmentedFloat64(math.MaxFloat64)
	return &df
}
