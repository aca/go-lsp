// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/aca/go-lsp/gopls_internal/lsp/command/gen"
	"github.com/aca/go-lsp/tools_internal/testenv"
)

func TestGenerated(t *testing.T) {
	testenv.NeedsGoPackages(t)
	testenv.NeedsLocalXTools(t)

	onDisk, err := os.ReadFile("command_gen.go")
	if err != nil {
		t.Fatal(err)
	}

	generated, err := gen.Generate()
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(string(generated), string(onDisk)); diff != "" {
		t.Errorf("command_gen.go is stale -- regenerate (-generated +on disk)\n%s", diff)
	}
}
