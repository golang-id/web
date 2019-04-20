// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go

package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	ciigo.Generate("./content", "cmd/golangid/static.go")
}
