// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"log"

	"git.sr.ht/~shulhan/ciigo"
	"github.com/shuLhan/share/lib/memfs"
)

var memFS *memfs.MemFS

func main() {
	var port string

	flag.StringVar(&port, "port", "5000", "HTTP port server")
	flag.Parse()

	serveOpts := &ciigo.ServeOptions{
		ConvertOptions: ciigo.ConvertOptions{
			Root:         "_content",
			HtmlTemplate: "_content/html.tmpl",
		},
		Address: "127.0.0.1:" + port,
		Mfs:     memFS,
	}
	err := ciigo.Serve(serveOpts)
	if err != nil {
		log.Fatal(err)
	}
}
