// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simplifyslice_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/aca/go-lsp/gopls_internal/analysis/simplifyslice"
	"github.com/aca/go-lsp/tools_internal/typeparams"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	tests := []string{"a"}
	if typeparams.Enabled {
		tests = append(tests, "typeparams")
	}
	analysistest.RunWithSuggestedFixes(t, testdata, simplifyslice.Analyzer, tests...)
}
