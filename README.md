# proxy

[![Build Status](https://github.com/lspserver/proxy/workflows/ci/badge.svg?branch=main&event=push)](https://github.com/lspserver/proxy/actions?query=workflow%3Aci)
[![codecov](https://codecov.io/gh/lspserver/proxy/branch/main/graph/badge.svg?token=CPBX6iyWnS)](https://codecov.io/gh/lspserver/proxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/lspserver/proxy)](https://goreportcard.com/report/github.com/lspserver/proxy)
[![License](https://img.shields.io/github/license/lspserver/proxy.svg)](https://github.com/lspserver/proxy/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/lspserver/proxy.svg)](https://github.com/lspserver/proxy/releases/latest)



## Introduction

*proxy* is the proxy of [lspserver](https://github.com/lspserver) written in Go.



## Prerequisites

- Go >= 1.17.0



## Run

```bash
version=latest make build
./bin/proxy --config-file="$PWD/config/config.yml"
```



## Docker

```bash
version=latest make docker
docker run -v "$PWD"/config:/tmp ghcr.io/lspserver/proxy:latest --config-file="/tmp/config.yml"
```



## Usage

```
```



## Settings

*proxy* parameters can be set in the directory [config](https://github.com/lspserver/proxy/blob/main/config).

An example of configuration in [config.yml](https://github.com/lspserver/proxy/blob/main/config/config.yml):

```yaml
apiVersion: v1
kind: proxy
metadata:
  name: proxy
spec:
  server:
    - name: c
      run:
        - ccls
        - --init="{"clang": {"resourceDir": "/usr/lib/llvm-10/lib/clang/10.0.0/include"}}"
    - name: c++
      run:
        - ccls
        - --init="{"clang": {"resourceDir": "/usr/lib/llvm-10/lib/clang/10.0.0/include"}}"
      - name: css
        run:
          - npx
          - css-languageserver
          - --stdio
      - name: dockerfile
        run:
          - npx
          - docker-langserver
          - --stdio
      - name: go
        run:
          - gopls
      - name: html
        run:
          - npx
          - html-languageserver
          - --stdio
      - name: java
      - name: javascript
        run:
          - node
          - language-server-stdio.js
      - name: python
      - name: rust
        run:
          - rls
      - name: shell
        run:
          - npx
          - bash-language-server
          - start
```



## License

Project License can be found [here](LICENSE).



## Reference

- [codemirror-lsp-example](https://github.com/wylieconlon/codemirror-lsp-example)
- [codemirror-mode](https://github.com/codemirror/CodeMirror/tree/master/mode)
- [jsonrpc-ws-proxy](https://github.com/wylieconlon/jsonrpc-ws-proxy)
- [language-server-protocol](https://microsoft.github.io/language-server-protocol/)
- [lsp-editor-adapter-example](https://github.com/wylieconlon/lsp-editor-adapter/tree/master/example)
- [monaco-editor](https://microsoft.github.io/monaco-editor/)
- [monaco-languageclient](https://github.com/TypeFox/monaco-languageclient)
- [vscode-json-languageservice](https://github.com/microsoft/vscode-json-languageservice)
