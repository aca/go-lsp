#!/usr/bin/env elvish

cp -rp /home/rok/src/github.com/golang/tools/internal tools_internal
cp -rp /home/rok/src/github.com/golang/tools/gopls/internal gopls_internal

fd --extension go | each { |x| sed -i 's|golang.org/x/tools/internal|github.com/aca/go-lsp/tools_internal|g' $x }
fd --extension go | each { |x| sed -i 's|golang.org/x/tools/gopls/internal|github.com/aca/go-lsp/gopls_internal|g' $x }
