// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The stringintconv command runs the stringintconv analyzer.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"github.com/aca/go-lsp/gopls_internal/analysis/unusedparams"
)

func main() { singlechecker.Main(unusedparams.Analyzer) }
