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
	"github.com/go-enjin/be/features/restrict/basic-auth"
)

func main() {
	common := `
{{ if .BasicUsername }}
<p>Hello {{ .BasicUsername }}</p>
{{ end }}
<ul>
	<li><a href="/">Home</a></li>
	<li><a href="/users">Any User</a></li>
	<li><a href="/only-ones">Only Ones</a></li>
	<li><a href="/ones-twos">Ones Twos</a></li>
	<li>
		<p>Any authenticated user can visit any sub-path of: "/this-path*"</p>
		<ul>
			<li><a href="/this-path/one">this-path: one</a></li>
			<li><a href="/this-path/two">this-path: two</a></li>
			<li><a href="/this-path-too">this-path-too</a></li>
		</ul>
		<p>However, "/this-path/is-ignored" is configured to be ignored by Basic
		Auth completely.</p>
		<ul>
			<li><a href="/this-path/is-ignored/one">this-path/is-ignored: one</a></li>
			<li><a href="/this-path/is-ignored/two">this-path/is-ignored: two</a></li>
			<li><a href="/this-path/is-ignored-too">this-path/is-ignored-too</a></li>
		</ul>
	</li>
</ul>
`
	homepage := `+++
Title = "Home"
Format = "tmpl"
+++
<h2>Example Homepage</h2>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
` + common
	anyUser := `+++
Title = "Any User"
Format = "tmpl"
BasicAuthGroups = ["users"]
+++
<h2>Any User</h2>
<p>Any authenticated user can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
` + common

	onlyOnesTmpl := `+++
Title = "Only Ones"
Format = "tmpl"
BasicAuthGroups = ["only-ones"]
+++
<h2>Only Ones</h2>
<p>Only users in the group of "only-ones" can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
` + common
	onesTwosOrg := `+++
Title = "Ones Twos"
Format = "tmpl"
BasicAuthGroups = ["ones-twos"]
+++
<h2>Ones Twos</h2>
<p>Only users in the group of "ones-twos" can see this content.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
` + common
	testingTmpl := `+++
Format = "tmpl"
+++
<h2>testing</h2>
<p>This page may or may not require authentication.</p>
<p>This content is an html/template and this string should contain the value of
a context variable set at compile-time: "{{ .CustomVariable }}".</p>
` + common
	enjin := be.New().
		Set("CustomVariable", "not-empty").
		AddPageFromString("/", homepage).
		AddPageFromString("/users", anyUser).
		AddPageFromString("/only-ones", onlyOnesTmpl).
		AddPageFromString("/ones-twos", onesTwosOrg).
		AddPageFromString("/this-path/one", testingTmpl).
		AddPageFromString("/this-path/two", testingTmpl).
		AddPageFromString("/this-path-too", testingTmpl).
		AddPageFromString("/this-path/is-ignored/one", testingTmpl).
		AddPageFromString("/this-path/is-ignored/two", testingTmpl).
		AddPageFromString("/this-path/is-ignored-too", testingTmpl).
		AddFeature(public.New().MountPath("/", "public").Make()).
		AddFeature(
			auth.New().
				Restrict("/this-path", "users").
				IgnoreLeadingPaths("/this-path/is-ignored").
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