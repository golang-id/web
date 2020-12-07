// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"

	"git.sr.ht/~shulhan/ciigo"
)

func main() {
	var port string

	flag.StringVar(&port, "port", "5000", "HTTP port server")
	flag.Parse()

	ciigo.Serve("./_content", ":"+port, "./_templates/html.tmpl")
}
