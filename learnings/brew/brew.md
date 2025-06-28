# HomeBrew Explained

* Why now? - During brew install postgres, I was not able to use psql command post installation, I was getting command not found.
* Found a github issue which suggested that brew install postgresql@15/16/17 is provided a keg-only [Ref: https://github.com/Homebrew/homebrew-core/issues/121043]
* Wanted to understand keg-only term, resulting in learning basics of homebrew

### Learning
HomeBrew is a package manager on MAC that doesn't come installed by default.
* Formulae - They are text based or technical softwares like Java, Postgres etc..,
* Formulas are installed here: /usr/local/Cellar/<software>/<version>
* Casks - Casks are software with user interfaces eg.., cockroach db console, chrome, firefox
* brew install <software> (brew install java)
* brew install --cask <software> (brew install –cask chrome)
* There will be a symlink created: /usr/local/bin/<software. a file that points to another file or directory is a symlink
* In homebrew terms
  * /usr/local/Cellar is a Cellar
  * /usr/local/Cellar/<software> is a Rack
  * /usr/local/Cellar/<software>/<version> is a Keg
  * Cellar is the place where you keep everything, and each software is a rack in the cellar, and each version of the software is a keg.

### Glossary
* Cellar: The room below ground level in a house. Used for storing wine
* Cask: a large container like a barrel, made of wood, metal or plastic and used for storing liquids, typically alcoholic drinks (curved)
* Keg: a small barrel, especially one of less than 10 gallons or (in the US) 30 gallons. ( vertical straight)