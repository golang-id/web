// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run generate.go

package main

import (
	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

func main() {
	opts := &ciigo.EmbedOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/html.tmpl",
		},
		EmbedOptions: memfs.EmbedOptions{
			PackageName: "main",
			VarName:     "memFS",
			GoFileName:  "cmd/www-golangid/static.go",
		},
	}
	ciigo.GoEmbed(opts)
}
