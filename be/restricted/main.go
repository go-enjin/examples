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
	"github.com/go-enjin/be/features/fs/locals/public"
	"github.com/go-enjin/be/features/restrict/basic"
)

func main() {
	homepage := `+++
Title = "Home"
Format = "tmpl"
+++
<h2>Example Homepage</h2>
{{ if .BasicUsername }}
<p>Hello {{ .BasicUsername }}</p>
{{ end }}
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
<ul>
	<li><a href="/">Home</a></li>
	<li><a href="/users">Any User</a></li>
	<li><a href="/only-ones">Only Ones</a></li>
	<li><a href="/ones-twos">Ones Twos</a></li>
</ul>
`
	anyUser := `+++
Title = "Any User"
Format = "tmpl"
RestrictBasic = ["users"]
+++
<h2>Any User</h2>
{{ if .BasicUsername }}
<p>Hello {{ .BasicUsername }}</p>
{{ end }}
<p>Any authenticated user can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
<ul>
	<li><a href="/">Home</a></li>
	<li><a href="/users">Any User</a></li>
	<li><a href="/only-ones">Only Ones</a></li>
	<li><a href="/ones-twos">Ones Twos</a></li>
</ul>
`
	onlyOnesTmpl := `+++
Title = "Only Ones"
Format = "tmpl"
RestrictBasic = ["only-ones"]
+++
<h2>Only Ones</h2>
{{ if .BasicUsername }}
<p>Hello {{ .BasicUsername }}</p>
{{ end }}
<p>Only users in the group of "only-ones" can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
<ul>
	<li><a href="/">Home</a></li>
	<li><a href="/users">Any User</a></li>
	<li><a href="/only-ones">Only Ones</a></li>
	<li><a href="/ones-twos">Ones Twos</a></li>
</ul>
`
	onesTwosOrg := `+++
Title = "Ones Twos"
Format = "tmpl"
RestrictBasic = ["ones-twos"]
+++
<h2>Ones Twos</h2>
<p>Only users in the group of "ones-twos" can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
{{ if .BasicUsername }}
<p>Hello {{ .BasicUsername }}</p>
{{ end }}
<ul>
	<li><a href="/">Home</a></li>
	<li><a href="/users">Any User</a></li>
	<li><a href="/only-ones">Only Ones</a></li>
	<li><a href="/ones-twos">Ones Twos</a></li>
</ul>
`
	enjin := be.New().
		Set("CustomVariable", "not-empty").
		AddPageFromString("/", homepage).
		AddPageFromString("/users", anyUser).
		AddPageFromString("/only-ones", onlyOnesTmpl).
		AddPageFromString("/ones-twos", onesTwosOrg).
		AddFeature(public.New().MountPath("/", "public").Make()).
		AddFeature(
			basic.New().
				RestrictAllData(true).
				Htpasswd("auth/one.htpasswd", "auth/two.htpasswd").
				Htgroups("auth/one.htgroup", "auth/two.htgroup", "auth/many.htgroup").
				Make(),
		).
		Build()
	if err := enjin.Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}