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

import "runtime"

type trees []*tree

func (t trees) split() []trees {
	numParts := runtime.NumCPU()
	parts := make([]trees, numParts)
	for i := 0; i < numParts; i++ {
		parts[i] = t[i*len(t)/numParts : (i+1)*len(t)/numParts]
	}
	return parts
}