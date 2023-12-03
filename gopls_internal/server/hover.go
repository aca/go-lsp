// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/gopls_internal/lsp/source"
	"github.com/aca/go-lsp/gopls_internal/mod"
	"github.com/aca/go-lsp/gopls_internal/telemetry"
	"github.com/aca/go-lsp/gopls_internal/template"
	"github.com/aca/go-lsp/gopls_internal/work"
	"github.com/aca/go-lsp/tools_internal/event"
	"github.com/aca/go-lsp/tools_internal/event/tag"
)

func (s *server) Hover(ctx context.Context, params *protocol.HoverParams) (_ *protocol.Hover, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("hover")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.hover", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, file.UnknownKind)
	defer release()
	if !ok {
		return nil, err
	}
	switch snapshot.FileKind(fh) {
	case file.Mod:
		return mod.Hover(ctx, snapshot, fh, params.Position)
	case file.Go:
		return source.Hover(ctx, snapshot, fh, params.Position)
	case file.Tmpl:
		return template.Hover(ctx, snapshot, fh, params.Position)
	case file.Work:
		return work.Hover(ctx, snapshot, fh, params.Position)
	}
	return nil, nil
}
