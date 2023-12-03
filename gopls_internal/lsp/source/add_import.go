// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package source

import (
	"context"

	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/cache"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/tools_internal/imports"
)

// AddImport adds a single import statement to the given file
func AddImport(ctx context.Context, snapshot *cache.Snapshot, fh file.Handle, importPath string) ([]protocol.TextEdit, error) {
	pgf, err := snapshot.ParseGo(ctx, fh, ParseFull)
	if err != nil {
		return nil, err
	}
	return ComputeOneImportFixEdits(snapshot, pgf, &imports.ImportFix{
		StmtInfo: imports.ImportInfo{
			ImportPath: importPath,
		},
		FixType: imports.AddImport,
	})
}
