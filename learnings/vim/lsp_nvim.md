## Goal

- Setup up Go local development environment with auto-completions, linting and necessary integrations with Go Tools.

## Context
- Currently vim has many LSP client support like coc.vim (Conquer of Completion), ALE.
- Vim uses VimScript for its plugin system and I have utilising VimPlug as the plugin manager for Vim.
- Understand why neovim is smoother to use over vim with coc.vim and other lsp clients.

## Language Server Protocols

- Language server protocol is the product of standardising the communication between the development tools and the language server process.
- Initially development tools was written in their own languages and when it had to support other languages a domain model (scanner, parser, type checker, builder and more) had to be written.
- Eg.., Visual Studio Code is written in TypeScript, for it to support C/C++ a domain model of C/C++ has to be written in TypeScript. Similarly Visual Studio is written in C#, for it to support C/C++ a domain model has to be written in C#.
- Thought its technically possible to achieve this, it is difficult one and tedius effort between for the development tool providers. 
- The thought of language servers popped up, were language servers run as independent processes. Development tools can communicate to the language servers via JSON RPC following a protocol.
- Using Language Servers/daemons was not a novel idea, editors like VIM/Emacs were already using those. The core idea was to introduce a protocol to standardise the communication.
References
- LSP by Microsoft: https://learn.microsoft.com/en-us/visualstudio/extensibility/language-server-protocol?view=vs-2022
- https://www.vikasraj.dev/blog/lsp-neovim-retrospective

## Lua Guide

- https://learnxinyminutes.com/lua/
- https://www.youtube.com/watch?v=NneB6GX1Els


## Installing NVIM in Custom Path

- Had a requirement to install nvim in custom path which is $HOME dir.

- Download tar from list of nvim binaries

```
curl -LO https://github.com/neovim/neovim/releases/download/nightly/nvim-macos-arm64.tar.gz
tar xzf nvim-macos-arm64.tar.gz
./nvim-macos-arm64/bin/nvim
```
- Export to the PATH, 
```
echo 'export PATH="/Users/venktesh.k/nvim/nvim-macos-arm64/bin:$PATH"' << ~/.zshrc
where nvim
nvim --version
```
