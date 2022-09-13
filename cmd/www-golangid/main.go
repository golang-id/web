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

const (
	cmdEmbed = `embed`

	dirRoot      = `_content`
	htmlTemplate = `_content/html.tmpl`
	listenAddr   = `127.0.0.1:`
)

func main() {
	var (
		convertOpts = ciigo.ConvertOptions{
			Root:         dirRoot,
			HtmlTemplate: htmlTemplate,
		}
		embedOpts = &ciigo.EmbedOptions{
			ConvertOptions: convertOpts,
			EmbedOptions: memfs.EmbedOptions{
				PackageName: `main`,
				VarName:     `memFS`,
				GoFileName:  `cmd/www-golangid/static.go`,
			},
		}
		serveOpts = &ciigo.ServeOptions{
			ConvertOptions: convertOpts,
			Mfs:            memFS,
		}

		cmd  string
		port string
		err  error
	)

	flag.StringVar(&port, "port", "5000", "HTTP port server")
	flag.Parse()

	cmd = flag.Arg(0)

	switch cmd {
	case cmdEmbed:
		err = ciigo.GoEmbed(embedOpts)
	default:
		serveOpts.Address = listenAddr + port
		err = ciigo.Serve(serveOpts)
	}
	if err != nil {
		log.Fatal(err)
	}
}
