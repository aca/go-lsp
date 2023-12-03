// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"context"

	"golang.org/x/mod/modfile"
	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/cache"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/tools_internal/event"
)

func Format(ctx context.Context, snapshot *cache.Snapshot, fh file.Handle) ([]protocol.TextEdit, error) {
	ctx, done := event.Start(ctx, "work.Format")
	defer done()

	pw, err := snapshot.ParseWork(ctx, fh)
	if err != nil {
		return nil, err
	}
	formatted := modfile.Format(pw.File.Syntax)
	// Calculate the edits to be made due to the change.
	diffs := snapshot.Options().ComputeEdits(string(pw.Mapper.Content), string(formatted))
	return protocol.EditsFromDiffEdits(pw.Mapper, diffs)
}
