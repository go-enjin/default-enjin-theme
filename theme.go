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

package theme

import (
	"embed"

	"github.com/go-enjin/be/pkg/log"
	"github.com/go-enjin/be/pkg/theme"
)

const Name = "default-enjin"

//go:embed default-enjin/**
//go:embed default-enjin/layouts/_default/**
var themeFS embed.FS

func Theme() (t *theme.Theme) {
	var err error
	if t, err = theme.NewEmbed(Name, Name, themeFS); err != nil {
		log.FatalDF(1, "error loading %v theme: %v", Name, err)
	}
	return
}
