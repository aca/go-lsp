// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/gopls_internal/lsp/source"
	"github.com/aca/go-lsp/gopls_internal/mod"
	"github.com/aca/go-lsp/tools_internal/event"
	"github.com/aca/go-lsp/tools_internal/event/tag"
)

func (s *server) InlayHint(ctx context.Context, params *protocol.InlayHintParams) ([]protocol.InlayHint, error) {
	ctx, done := event.Start(ctx, "lsp.Server.inlayHint", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, file.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	switch snapshot.FileKind(fh) {
	case file.Mod:
		return mod.InlayHint(ctx, snapshot, fh, params.Range)
	case file.Go:
		return source.InlayHint(ctx, snapshot, fh, params.Range)
	}
	return nil, nil
}
