// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"github.com/shuLhan/ciigo"
)

func main() {
	srv := ciigo.NewServer("./content", ":5000", "./templates/html.tmpl")

	srv.Start()
}
