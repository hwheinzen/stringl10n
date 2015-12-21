// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"fmt"
	"sort"
)

type Elem struct {
	Key   string
	Value string
}

// String formats Elem for printing.
func (p Elem) String() string {
	return fmt.Sprintf("%s: %s", p.Key, p.Value)
}

// ByKey implements sort.Interface for []Elem based on
// the Key field.
type ByKey []Elem

func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

// ByKeySort sorts a slice of Elems in increasing Key order.
// (see the standard library: sort.go)
func ByKeySort(x []Elem) { sort.Sort(ByKey(x)) }
