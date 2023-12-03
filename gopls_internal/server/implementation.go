// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"context"

	"github.com/aca/go-lsp/gopls_internal/file"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/gopls_internal/lsp/source"
	"github.com/aca/go-lsp/gopls_internal/telemetry"
	"github.com/aca/go-lsp/tools_internal/event"
	"github.com/aca/go-lsp/tools_internal/event/tag"
)

func (s *server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (_ []protocol.Location, rerr error) {
	recordLatency := telemetry.StartLatencyTimer("implementation")
	defer func() {
		recordLatency(ctx, rerr)
	}()

	ctx, done := event.Start(ctx, "lsp.Server.implementation", tag.URI.Of(params.TextDocument.URI))
	defer done()

	snapshot, fh, ok, release, err := s.beginFileRequest(ctx, params.TextDocument.URI, file.Go)
	defer release()
	if !ok {
		return nil, err
	}
	return source.Implementation(ctx, snapshot, fh, params.Position)
}
