Git Branches and Tags
Git branches and tags are essentially the same.
Use git branches when we want to evolve a codebase.
Use git tags when you want to mark significant milestones/changes or stable versions.
Essentially both git branch and git tag point to a specific commit, branch does not mean strings of commits. (They are nicknames of a commit). With a branch name (that is the commit) we can get the lineage of all commits (as git traverses backward A<-B-<C — master, C commit is master. If C is known I can traverse back).
Guide: https://www.atlassian.com/git/tutorials/inspecting-a-repository/git-tag
Common myths: Merging a branch does not merge the tags as they are completely unrelated. Tags and branches denote a commit as we mentioned before. We basically checkout a named commit called tag
