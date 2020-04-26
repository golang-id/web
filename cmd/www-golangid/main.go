// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"

	"github.com/shuLhan/ciigo"
)

func main() {
	var port string

	flag.StringVar(&port, "port", "5000", "HTTP port server")
	flag.Parse()

	srv := ciigo.NewServer("./_content", ":"+port, "./_templates/html.tmpl")
	srv.Start()
}
