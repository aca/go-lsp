// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/gopls_internal/lsp/source"
	"github.com/aca/go-lsp/gopls_internal/mod"
	"github.com/aca/go-lsp/gopls_internal/work"
	"github.com/aca/go-lsp/tools_internal/event"
	"github.com/aca/go-lsp/tools_internal/event/tag"
)

func (s *server) Formatting(ctx context.Context, params *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	ctx, done := event.Start(ctx, "lsp.Server.formatting", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, file.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	switch snapshot.FileKind(fh) {
	case file.Mod:
		return mod.Format(ctx, snapshot, fh)
	case file.Go:
		return source.Format(ctx, snapshot, fh)
	case file.Work:
		return work.Format(ctx, snapshot, fh)
	}
	return nil, nil
}
