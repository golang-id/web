// Copyright 2019, The golang-id.org Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"strings"

	"git.sr.ht/~shulhan/ciigo"
	"git.sr.ht/~shulhan/pakakeh.go/lib/memfs"
)

var memFS *memfs.MemFS

const (
	cmdEmbed = `embed`

	dirRoot       = `_content`
	htmlTemplate  = `_content/html.tmpl`
	defListenAddr = `127.0.0.1:5000`
)

func main() {
	var (
		convertOpts = ciigo.ConvertOptions{
			Root:         dirRoot,
			HTMLTemplate: htmlTemplate,
		}
		embedOpts = ciigo.EmbedOptions{
			ConvertOptions: convertOpts,
			EmbedOptions: memfs.EmbedOptions{
				PackageName: `main`,
				VarName:     `memFS`,
				GoFileName:  `cmd/www-golangid/static.go`,
			},
		}
		serveOpts = ciigo.ServeOptions{
			ConvertOptions: convertOpts,
			Mfs:            memFS,
		}
	)

	flag.BoolVar(&serveOpts.IsDevelopment, `dev`, false,
		`Jalankan mode pengembangan.`)
	flag.StringVar(&serveOpts.Address, `http`, defListenAddr,
		`Alamat peladen HTTP.`)
	flag.Parse()

	var cmd = strings.ToLower(flag.Arg(0))

	var err error
	switch cmd {
	case cmdEmbed:
		err = ciigo.GoEmbed(embedOpts)
	default:
		err = ciigo.Serve(serveOpts)
	}
	if err != nil {
		log.Fatal(err)
	}
}
