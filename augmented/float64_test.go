package augmented

import (
	"testing"
	"github.com/Kerah/goaugmented"
)

func TestAugmentedFloat64(t *testing.T) {
	tree := goaugmented.New(1)
	iv := goaugmented.SingleDimensionInterval(
		NewFloat64(10.0),
		NewFloat64(100.0),
		1,
	)

	iv2 := goaugmented.SingleDimensionInterval(
		NewFloat64(50.0),
		NewFloat64(100.0),
		2,
	)
	tree.Add(iv, iv2)
	q := NewFloat64VI(64)
	ivs := tree.Query(q)


	data := map[uint64]struct{}{}

	for _, i := range ivs {
		data[i.ID()] = struct{}{}
	}

	if _, ok := data[1]; !ok {
		t.Errorf("Expected ID=1 in interval")
	}

	if _, ok := data[2]; !ok {
		t.Errorf("Expected ID=2 in interval")
	}

}
