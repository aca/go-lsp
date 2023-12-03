package main

import (
	"context"
	"log"
	"os"

	"github.com/aca/go-lsp/gopls_internal/lsp/cache"
	"github.com/aca/go-lsp/gopls_internal/lsp/lsprpc"
	"github.com/aca/go-lsp/gopls_internal/lsp/protocol"
	"github.com/aca/go-lsp/tools_internal/fakenet"
	"github.com/aca/go-lsp/tools_internal/jsonrpc2"
)

type Server struct {
	Conn   *jsonrpc2.Conn
	client protocol.Client
	// stateMu sync.Mutex

	lifetime context.Context
}

func (s *Server) Progress(context.Context, *protocol.ProgressParams) error { log.Println("1"); return nil }
func (s *Server) SetTrace(context.Context, *protocol.SetTraceParams) error { log.Println("1"); return nil }
func (s *Server) IncomingCalls(context.Context, *protocol.CallHierarchyIncomingCallsParams) ([]protocol.CallHierarchyIncomingCall, error) { log.Println("1"); return nil, nil }
func (s *Server) OutgoingCalls(context.Context, *protocol.CallHierarchyOutgoingCallsParams) ([]protocol.CallHierarchyOutgoingCall, error) { log.Println("1"); return nil, nil }
func (s *Server) ResolveCodeAction(context.Context, *protocol.CodeAction) (*protocol.CodeAction, error) { log.Println("1"); return nil, nil }
func (s *Server) ResolveCodeLens(context.Context, *protocol.CodeLens) (*protocol.CodeLens, error) { log.Println("1"); return nil, nil }
func (s *Server) ResolveCompletionItem(context.Context, *protocol.CompletionItem) (*protocol.CompletionItem, error){ return nil, nil }
func (s *Server) ResolveDocumentLink(context.Context, *protocol.DocumentLink) (*protocol.DocumentLink, error) { log.Println("1"); return nil, nil }
func (s *Server) Exit(context.Context) error { log.Println("1"); return nil }
func (s *Server) Initialize(context.Context, *protocol.ParamInitialize) (*protocol.InitializeResult, error) {
	return &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			CallHierarchyProvider: &protocol.Or_ServerCapabilities_callHierarchyProvider{Value: true},
			// CodeActionProvider:    codeActionProvider,
			CodeLensProvider:      &protocol.CodeLensOptions{}, // must be non-nil to enable the code lens capability
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: []string{"."},
			},
			DefinitionProvider:         &protocol.Or_ServerCapabilities_definitionProvider{Value: true},
			TypeDefinitionProvider:     &protocol.Or_ServerCapabilities_typeDefinitionProvider{Value: true},
			ImplementationProvider:     &protocol.Or_ServerCapabilities_implementationProvider{Value: true},
			DocumentFormattingProvider: &protocol.Or_ServerCapabilities_documentFormattingProvider{Value: false},
			DocumentSymbolProvider:     &protocol.Or_ServerCapabilities_documentSymbolProvider{Value: false},
			WorkspaceSymbolProvider:    &protocol.Or_ServerCapabilities_workspaceSymbolProvider{Value: false},
			// ExecuteCommandProvider: &protocol.ExecuteCommandOptions{
			// 	Commands: protocol.NonNilSlice(options.SupportedCommands),
			// },
			FoldingRangeProvider:      &protocol.Or_ServerCapabilities_foldingRangeProvider{Value: false},
			HoverProvider:             &protocol.Or_ServerCapabilities_hoverProvider{Value: false},
			DocumentHighlightProvider: &protocol.Or_ServerCapabilities_documentHighlightProvider{Value: false},
			DocumentLinkProvider:      &protocol.DocumentLinkOptions{},
			InlayHintProvider:         protocol.InlayHintOptions{},
			ReferencesProvider:        &protocol.Or_ServerCapabilities_referencesProvider{Value: false},
			// RenameProvider:            renameOpts,
			SelectionRangeProvider:    &protocol.Or_ServerCapabilities_selectionRangeProvider{Value: false},
			// SemanticTokensProvider: protocol.SemanticTokensOptions{
			// 	Range: &protocol.Or_SemanticTokensOptions_range{Value: true},
			// 	Full:  &protocol.Or_SemanticTokensOptions_full{Value: true},
			// 	Legend: protocol.SemanticTokensLegend{
			// 		TokenTypes:     protocol.NonNilSlice(options.SemanticTypes),
			// 		TokenModifiers: protocol.NonNilSlice(options.SemanticMods),
			// 	},
			// },
			// SignatureHelpProvider: &protocol.SignatureHelpOptions{
			// 	TriggerCharacters: []string{"s", "i", "l"},
			// },
			// TextDocumentSync: &protocol.TextDocumentSyncOptions{
			// 	Change:    protocol.Incremental,
			// 	OpenClose: true,
			// 	Save: &protocol.SaveOptions{
			// 		IncludeText: false,
			// 	},
			// },
			// Workspace: &protocol.WorkspaceOptions{
			// 	WorkspaceFolders: &protocol.WorkspaceFolders5Gn{
			// 		Supported:           true,
			// 		ChangeNotifications: "workspace/didChangeWorkspaceFolders",
			// 	},
			// },
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "my",
			// Version: "",
		},
	}, nil

}
func (s *Server) Initialized(context.Context, *protocol.InitializedParams) error { log.Println("1"); return nil }
func (s *Server) Resolve(context.Context, *protocol.InlayHint) (*protocol.InlayHint, error) { log.Println("1"); return nil, nil }
func (s *Server) DidChangeNotebookDocument(context.Context, *protocol.DidChangeNotebookDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidCloseNotebookDocument(context.Context, *protocol.DidCloseNotebookDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidOpenNotebookDocument(context.Context, *protocol.DidOpenNotebookDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidSaveNotebookDocument(context.Context, *protocol.DidSaveNotebookDocumentParams) error { log.Println("1"); return nil }
func (s *Server) Shutdown(context.Context) error { log.Println("1"); return nil }
func (s *Server) CodeAction(context.Context, *protocol.CodeActionParams) ([]protocol.CodeAction, error) { log.Println("1"); return nil, nil }
func (s *Server) CodeLens(context.Context, *protocol.CodeLensParams) ([]protocol.CodeLens, error) { log.Println("1"); return nil, nil }
func (s *Server) ColorPresentation(context.Context, *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) { log.Println("1"); return nil, nil }
func (s *Server) Completion(context.Context, *protocol.CompletionParams) (*protocol.CompletionList, error) {
	completionList := &protocol.CompletionList{
		IsIncomplete: false,
		ItemDefaults: &protocol.CompletionItemDefaults{},
		Items:        []protocol.CompletionItem{
			{
				Label:               "randomtext",
				LabelDetails:        &protocol.CompletionItemLabelDetails{
					Detail:      "randomtext from lsp",
					Description: "randomtext from lsp",
				},
				// Kind:                0,
				// Tags:                []protocol.CompletionItemTag{},
				// Detail:              "",
				// Documentation:       &protocol.Or_CompletionItem_documentation{},
				// Deprecated:          false,
				// Preselect:           false,
				// SortText:            "",
				// FilterText:          "",
				InsertText:          "randomtext",
				// InsertTextFormat:    &0,
				// InsertTextMode:      &0,
				// TextEdit:            &protocol.TextEdit{},
				// // TextEditText:        "silence",
				// AdditionalTextEdits: []protocol.TextEdit{},
				// CommitCharacters:    []string{},
				// Command:             &protocol.Command{},
				// Data:                nil,
			},
		},
	}
	return completionList, nil
}
func (s *Server) Declaration(context.Context, *protocol.DeclarationParams) (*protocol.Or_textDocument_declaration, error) { log.Println("1"); return nil, nil }
func (s *Server) Definition(context.Context, *protocol.DefinitionParams) ([]protocol.Location, error) { log.Println("1"); return nil, nil }
func (s *Server) Diagnostic(context.Context, *string) (*string, error) { log.Println("1"); return nil, nil }
func (s *Server) DidChange(context.Context, *protocol.DidChangeTextDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidClose(context.Context, *protocol.DidCloseTextDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidOpen(context.Context, *protocol.DidOpenTextDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DidSave(context.Context, *protocol.DidSaveTextDocumentParams) error { log.Println("1"); return nil }
func (s *Server) DocumentColor(context.Context, *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) { log.Println("1"); return nil, nil }
func (s *Server) DocumentHighlight(context.Context, *protocol.DocumentHighlightParams) ([]protocol.DocumentHighlight, error) { log.Println("1"); return nil, nil }
func (s *Server) DocumentLink(context.Context, *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) { log.Println("1"); return nil, nil }
func (s *Server) DocumentSymbol(context.Context, *protocol.DocumentSymbolParams) ([]interface{}, error) { log.Println("1"); return nil, nil }
func (s *Server) FoldingRange(context.Context, *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) { log.Println("1"); return nil, nil }
func (s *Server) Formatting(context.Context, *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) Hover(context.Context, *protocol.HoverParams) (*protocol.Hover, error) { log.Println("1"); return nil, nil }
func (s *Server) Implementation(context.Context, *protocol.ImplementationParams) ([]protocol.Location, error) { log.Println("1"); return nil, nil }
func (s *Server) InlayHint(context.Context, *protocol.InlayHintParams) ([]protocol.InlayHint, error) { log.Println("1"); return nil, nil }
func (s *Server) InlineCompletion(context.Context, *protocol.InlineCompletionParams) (*protocol.Or_Result_textDocument_inlineCompletion, error) { log.Println("1"); return nil, nil }
func (s *Server) InlineValue(context.Context, *protocol.InlineValueParams) ([]protocol.InlineValue, error) { log.Println("1"); return nil, nil }
func (s *Server) LinkedEditingRange(context.Context, *protocol.LinkedEditingRangeParams) (*protocol.LinkedEditingRanges, error) { log.Println("1"); return nil, nil }
func (s *Server) Moniker(context.Context, *protocol.MonikerParams) ([]protocol.Moniker, error) { log.Println("1"); return nil, nil }
func (s *Server) OnTypeFormatting(context.Context, *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) PrepareCallHierarchy(context.Context, *protocol.CallHierarchyPrepareParams) ([]protocol.CallHierarchyItem, error) { log.Println("1"); return nil, nil }
func (s *Server) PrepareRename(context.Context, *protocol.PrepareRenameParams) (*protocol.PrepareRenameResult, error) { log.Println("1"); return nil, nil }
func (s *Server) PrepareTypeHierarchy(context.Context, *protocol.TypeHierarchyPrepareParams) ([]protocol.TypeHierarchyItem, error) { log.Println("1"); return nil, nil }
func (s *Server) RangeFormatting(context.Context, *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) RangesFormatting(context.Context, *protocol.DocumentRangesFormattingParams) ([]protocol.TextEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) References(context.Context, *protocol.ReferenceParams) ([]protocol.Location, error) { log.Println("1"); return nil, nil }
func (s *Server) Rename(context.Context, *protocol.RenameParams) (*protocol.WorkspaceEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) SelectionRange(context.Context, *protocol.SelectionRangeParams) ([]protocol.SelectionRange, error) { log.Println("1"); return nil, nil }
func (s *Server) SemanticTokensFull(context.Context, *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) { log.Println("1"); return nil, nil }
func (s *Server) SemanticTokensFullDelta(context.Context, *protocol.SemanticTokensDeltaParams) (interface{}, error) { log.Println("1"); return nil, nil }
func (s *Server) SemanticTokensRange(context.Context, *protocol.SemanticTokensRangeParams) (*protocol.SemanticTokens, error) { log.Println("1"); return nil, nil }
func (s *Server) SignatureHelp(context.Context, *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) { log.Println("1"); return nil, nil }
func (s *Server) TypeDefinition(context.Context, *protocol.TypeDefinitionParams) ([]protocol.Location, error) { log.Println("1"); return nil, nil }
func (s *Server) WillSave(context.Context, *protocol.WillSaveTextDocumentParams) error { log.Println("1"); return nil }
func (s *Server) WillSaveWaitUntil(context.Context, *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) Subtypes(context.Context, *protocol.TypeHierarchySubtypesParams) ([]protocol.TypeHierarchyItem, error) { log.Println("1"); return nil, nil }
func (s *Server) Supertypes(context.Context, *protocol.TypeHierarchySupertypesParams) ([]protocol.TypeHierarchyItem, error) { log.Println("1"); return nil, nil }
func (s *Server) WorkDoneProgressCancel(context.Context, *protocol.WorkDoneProgressCancelParams) error { log.Println("1"); return nil }
func (s *Server) DiagnosticWorkspace(context.Context, *protocol.WorkspaceDiagnosticParams) (*protocol.WorkspaceDiagnosticReport, error) { log.Println("1"); return nil, nil }
func (s *Server) DidChangeConfiguration(context.Context, *protocol.DidChangeConfigurationParams) error { log.Println("1"); return nil }
func (s *Server) DidChangeWatchedFiles(context.Context, *protocol.DidChangeWatchedFilesParams) error { log.Println("1"); return nil }
func (s *Server) DidChangeWorkspaceFolders(context.Context, *protocol.DidChangeWorkspaceFoldersParams) error { log.Println("1"); return nil }
func (s *Server) DidCreateFiles(context.Context, *protocol.CreateFilesParams) error { log.Println("1"); return nil }
func (s *Server) DidDeleteFiles(context.Context, *protocol.DeleteFilesParams) error { log.Println("1"); return nil }
func (s *Server) DidRenameFiles(context.Context, *protocol.RenameFilesParams) error { log.Println("1"); return nil }
func (s *Server) ExecuteCommand(context.Context, *protocol.ExecuteCommandParams) (interface{}, error) { log.Println("1"); return nil, nil }
func (s *Server) Symbol(context.Context, *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) { log.Println("1"); return nil, nil }
func (s *Server) WillCreateFiles(context.Context, *protocol.CreateFilesParams) (*protocol.WorkspaceEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) WillDeleteFiles(context.Context, *protocol.DeleteFilesParams) (*protocol.WorkspaceEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) WillRenameFiles(context.Context, *protocol.RenameFilesParams) (*protocol.WorkspaceEdit, error) { log.Println("1"); return nil, nil }
func (s *Server) ResolveWorkspaceSymbol(context.Context, *protocol.WorkspaceSymbol) (*protocol.WorkspaceSymbol, error) { log.Println("1"); return nil, nil }
func (s *Server) NonstandardRequest(ctx context.Context, method string, params interface{}) (interface{}, error) { log.Println("1"); return nil, nil }

func main() {
	ctx := context.Background()
	s := new(Server)

	stream := jsonrpc2.NewHeaderStream(fakenet.NewConn("stdio", os.Stdin, os.Stdout))
	conn := jsonrpc2.NewConn(stream)

	conn.Go(ctx,
		protocol.Handlers(
			protocol.ServerHandler(s, jsonrpc2.MethodNotFound)))
	ss := lsprpc.NewStreamServer(cache.New(nil), false, nil)
	err := ss.ServeStream(ctx, conn)
	if err != nil {
		panic(err)
	}
}
