package augmented

import (
	"testing"
	"github.com/Kerah/goaugmented"
)

func TestAugmentedInt64(t *testing.T) {
	tree := goaugmented.New(1)
	iv := goaugmented.SingleDimensionInterval(
		NewInt64(10),
		NewInt64(100),
		1,
	)

	iv2 := goaugmented.SingleDimensionInterval(
		NewInt64(50),
		NewInt64(100),
		2,
	)
	tree.Add(iv, iv2)
	q := NewInt64VI(64)
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

