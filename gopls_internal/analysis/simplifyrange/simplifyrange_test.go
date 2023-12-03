// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplifyrange_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/aca/go-lsp/gopls_internal/analysis/simplifyrange"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.RunWithSuggestedFixes(t, testdata, simplifyrange.Analyzer, "a")
}
