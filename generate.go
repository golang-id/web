// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go

package main

import (
	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	opts := &ciigo.GenerateOptions{
		Root:           "_content",
		HTMLTemplate:   "_content/html.tmpl",
		GenPackageName: "main",
		GenVarName:     "memFS",
		GenGoFileName:  "cmd/www-golangid/static.go",
	}
	ciigo.Generate(opts)
}
