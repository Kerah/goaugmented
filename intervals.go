/*
Copyright 2014 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package goaugmented

import (
	"sync"
	"fmt"
)

var intervalsPool = sync.Pool{
	New: func() interface{} {
		return make(Intervals, 0, 10)
	},
}

// Intervals represents a list of Intervals.
type Intervals []Interval

// Dispose will free any consumed resources and allow this list to be
// re-allocated.
func (ivs *Intervals) Dispose() {
	for i := 0; i < len(*ivs); i++ {
		(*ivs)[i] = nil
	}

	*ivs = (*ivs)[:0]
	intervalsPool.Put(*ivs)
}

type dimension struct {
	low, high Value
}

type interval struct {
	dimensions []*dimension
	id         uint64
}

func (mi *interval) checkDimension(dimension uint64) {
	if dimension > uint64(len(mi.dimensions)) {
		panic(fmt.Sprintf(`Dimension: %d out of range.`, dimension))
	}
}

func (mi *interval) LowAtDimension(dimension uint64) Value {
	return mi.dimensions[dimension-1].low
}

func (mi *interval) HighAtDimension(dimension uint64) Value {
	return mi.dimensions[dimension-1].high
}

func (mi *interval) OverlapsAtDimension(iv Interval, dimension uint64) bool {
	return mi.HighAtDimension(dimension).Greater(iv.LowAtDimension(dimension)) &&
		mi.LowAtDimension(dimension).Lesser(iv.HighAtDimension(dimension))
}

func (mi interval) ID() uint64 {
	return mi.id
}

func SingleDimensionInterval(low, high Value, id uint64) *interval {
	return &interval{[]*dimension{&dimension{low: low, high: high}}, id}
}

func MultiDimensionInterval(id uint64, dimensions ...*dimension) *interval {
	return &interval{dimensions: dimensions, id: id}
}

func ValueInterval(val Value) *interval {
	return SingleDimensionInterval(val, val, 0)
}