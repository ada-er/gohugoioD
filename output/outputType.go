// Copyright 2017-present The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package output

import (
	"fmt"
	"strings"

	"github.com/spf13/hugo/media"
)

var (
	// An ordered list of built-in output formats
	// See https://www.ampproject.org/learn/overview/
	AMPType = Type{
		Name:      "AMP",
		MediaType: media.HTMLType,
		BaseName:  "index",
	}

	CSSType = Type{
		Name:      "CSS",
		MediaType: media.CSSType,
		BaseName:  "styles",
	}

	HTMLType = Type{
		Name:      "HTML",
		MediaType: media.HTMLType,
		BaseName:  "index",
	}

	JSONType = Type{
		Name:        "JSON",
		MediaType:   media.HTMLType,
		BaseName:    "index",
		IsPlainText: true,
	}

	RSSType = Type{
		Name:      "RSS",
		MediaType: media.RSSType,
		BaseName:  "index",
	}
)

var builtInTypes = map[string]Type{
	strings.ToLower(AMPType.Name):  AMPType,
	strings.ToLower(CSSType.Name):  CSSType,
	strings.ToLower(HTMLType.Name): HTMLType,
	strings.ToLower(JSONType.Name): JSONType,
	strings.ToLower(RSSType.Name):  RSSType,
}

type Types []Type

// Type represents an output represenation, usually to a file on disk.
type Type struct {
	// The Name is used as an identifier. Internal output types (i.e. HTML and RSS)
	// can be overridden by providing a new definition for those types.
	Name string

	MediaType media.Type

	// Must be set to a value when there are two or more conflicting mediatype for the same resource.
	Path string

	// The base output file name used when not using "ugly URLs", defaults to "index".
	BaseName string

	// The protocol to use, i.e. "webcal://". Defaults to the protocol of the baseURL.
	Protocol string

	// IsPlainText decides whether to use text/template or html/template
	// as template parser.
	IsPlainText bool

	// Enable to ignore the global uglyURLs setting.
	NoUgly bool
}

func GetType(key string) (Type, bool) {
	found, ok := builtInTypes[key]
	if !ok {
		found, ok = builtInTypes[strings.ToLower(key)]
	}
	return found, ok
}

// TODO(bep) outputs rewamp on global config?
func GetTypes(keys ...string) (Types, error) {
	var types []Type

	for _, key := range keys {
		tpe, ok := GetType(key)
		if !ok {
			return types, fmt.Errorf("OutputType with key %q not found", key)
		}
		types = append(types, tpe)
	}

	return types, nil
}
