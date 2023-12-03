// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hooks adds all the standard gopls implementations.
// This can be used in tests without needing to use the gopls main, and is
// also the place to edit for custom builds of gopls.
package hooks // import "github.com/aca/go-lsp/gopls_internal/hooks"

import (
	"github.com/aca/go-lsp/gopls_internal/settings"
	"github.com/aca/go-lsp/tools_internal/diff"
	"mvdan.cc/xurls/v2"
)

func Options(options *settings.Options) {
	options.LicensesText = licensesText
	if options.GoDiff {
		switch options.NewDiff {
		case "old":
			options.ComputeEdits = ComputeEdits
		case "new":
			options.ComputeEdits = diff.Strings
		default:
			options.ComputeEdits = BothDiffs
		}
	}
	options.URLRegexp = xurls.Relaxed()
	updateAnalyzers(options)
	updateGofumpt(options)
}
