# goaugmented
Golang augmented tree realisation for interfaces types.

```go
import (
    "github.com/Kerah/goaugmented"
    "github.com/Kerah/goaugmented/augmented"
    "fmt"
)

func main(){
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
		fmt.Printf("Founded: %d\n", i.ID())
	}
}
```
