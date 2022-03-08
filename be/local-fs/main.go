// Copyright (c) 2022  The Go-Enjin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"github.com/go-enjin/be"
	"github.com/go-enjin/be/features/fs/locals/content"
	"github.com/go-enjin/be/features/fs/locals/public"
)

func main() {

	enjin := be.New().
		AddFeature(content.New().MountPath("/", "content").Make()).
		AddFeature(public.New().MountPath("/", "public").Make()).
		AddThemes("themes").
		SetTheme("custom-theme").
		Set("Copyright", "Copyright (c) 2022").
		Set("CustomVariable", "compile-time text").
		AddPageFromString("/example", `+++
Format = "tmpl"
+++
<h2>Example Content</h2>
<p>Demonstrating adding content as a plain string with no metadata.</p>
<p>The following should not be empty: "{{ .CustomVariable }}"</p>
<p>shasum for hello.txt: {{ fsHash "/hello.txt" }}</p>
<p>url for styles.css: {{ fsUrl "/styles.css" }}</p>
<p>mime for gopher.png: {{ fsMime "/gopher.png" }}</p>
`).
		Build()
	if err := enjin.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}