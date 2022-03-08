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

	"github.com/go-enjin/be/pkg/net/ip/ranges/atlassian"
	"github.com/go-enjin/be/pkg/net/ip/ranges/cloudflare"
)

func main() {
	if atlIpRanges, err := atlassian.GetIpRanges(); err != nil {
		fmt.Fprintf(os.Stderr, "error getting atlassian ip ranges: %v\n", err)
		os.Exit(1)
	} else {
		for idx, cidr := range atlIpRanges {
			fmt.Printf("atlassian[%d]: %v\n", idx, cidr)
		}
	}
	if cfIpRanges, err := cloudflare.GetIpRanges(); err != nil {
		fmt.Fprintf(os.Stderr, "error getting cloudflare ip ranges: %v\n", err)
		os.Exit(1)
	} else {
		for idx, cidr := range cfIpRanges {
			fmt.Printf("cloudflare[%d]: %v\n", idx, cidr)
		}
	}
}