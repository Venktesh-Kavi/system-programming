## Vim Notes


### Buffers
- In VIM everything is a buffer
- Typical file editing softwares, load the file into a buffer and application edits the file in the buffer. When the save button is hit it is saved to the disk.
- In VIM editing happens directly on the buffer, saving rewrites the file or creates a new file.

## Modes
- Vim has two modes Insert and Normal Mode

## Cheat Sheet

- Use ^W / ^B to skip words till the next white-space is found. (eg., don't if we do w here, it cursor stays at ').
- % - go to opening or closing braces
- { or } - to skip paragraphs, usefully for iterating methods
- Changing words
    1/ * the word
    2/ g+n: highlights the word
    3/ c+g+n: change to new word
    4/ go to next occurrence (n and . (repeat last command)
- Copying a method efficiently -> typically we do ^V select the lines via J or K and then Yank it. Better approach v%. % finds the matching brace and puts it in visual mode, now yank. Method is copied!!!
- Copying to registers - Vim Yanks does not copy data to system buffer. Use "*y or setup vim config to allow Yanks to system buffer.
- Mutli line editing -> (Ctrl + V) + (Shit + I) + ESC
- Multi line editing at the end with words having varying length
    - gv (go to last highlighted) + $ (go to last char of line) + ^A (start append moode) + tpye + ESC.
- Go to a line: (g) + (line number)
- Macros: q (start macro) + a (choose a register) + perform any action + q (stop macro)
        - replay macro eg.., (5aj)
- Mark and move: m (start mark) + a (choose a register) [marks the current line)
-       - go back to the marked line: ' + a
- Enable word wrap in Vim -> :set wrap -> untoggle :set wrap!
- Go to file: place cursor on file path (g + f)
- Go to url: place cursor on url (g + x)
- Re-indent then file: gg + (=) + G
- Collapse all blocks in code: z + M
- Reopen collapsed blocks: z + R

## References
- https://www.danielmsullivan.com/pages/tutorial_vim.html#:~:text=Vim%20has%20two%20main%20%22modes,spend%20most%20of%20your%20time.

