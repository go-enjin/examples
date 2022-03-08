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
	"github.com/go-enjin/be/features/requests/cloudflare"
)

func main() {
	homepage := `
<h2>Example Content</h2>
<p>Demonstrating adding content as a plain string with no metadata.</p>
<p>Check out the <a href="/tmpl">tmpl page</a></p>
`
	tmpl := `+++
Format = "tmpl"
+++
This content is an html/template and this string should contain
the value of a context variable set at compile-time: "{{ .CustomVariable }}".
`
	orgTmpl := `+++
Format = "org"
+++
* TODO start writing documentation for Go-Enjin
* DONE compile-time context variables: "{{ .CustomVariable }}"
`
	enjin := be.New().
		Set("CustomVariable", "not-empty").
		AddPageFromString("/", homepage).
		AddPageFromString("/tmpl", tmpl).
		AddPageFromString("/org", orgTmpl).
		AddFeature(cloudflare.New().AllowDirect().Make()).
		Build()
	if err := enjin.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}