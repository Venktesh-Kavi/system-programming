

- Packer package manager is not longer under active development.
- LazyVim package manager seems to be stable and has a large community around it.


Setup

* Create a nvim dir under ~/.config.
* Perform vim . opens a file explorer with netrw. Use :Ex for explore
* File structure

``` lua
    venkz
        core
            init.lua
            keymaps.lua
            options.lua
        plugins
            colorscheme.lua
            nvim-tree.lua
        lazy.lua
    init.lua
```

* Default lazyvim plugin load priority is 50.
* 
