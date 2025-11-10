\## commands used and practices by me



| Command                              | Used For                                             | Example Output / Effect                                       |

| `mkdir git-hour`                     | Create a new folder for your repo                    | Creates folder `git-hour/`                                    |

| `cd git-hour`                        | Move into that folder                                | You’re now “inside” the repo directory                        |

| `git init`                           | Initialize a new local Git repo                      | Shows: `Initialized empty Git repository in ...`              |

| `git status`                         | See which files are new/modified/untracked           | Lists files in red (untracked) or green (staged)              |

| `git add README.md`                  | Stage a single file to be committed                  | Moves file from “untracked” → “ready to commit”               |

| `git add .`                          | Stage \*all\* changes in current folder                | Adds all modified/new files                                   |

| `git commit -m "message"`            | Save a snapshot of staged files with a message       | Creates a new commit in local repo                            |

| `git log`                            | Show full history of commits                         | Lists commits with author, date, and message                  |

| `git log --oneline`                  | Compact history view                                 | Shows short commit IDs like `a5b2c3d initial commit`          |

| `git branch`                         | List all branches and show current one               | `\* main`, `feature2`, etc.                                    |

| `git checkout -b feature/newbranch1` | Create \*\*and switch\*\* to a new branch                | `Switched to a new branch 'feature/newbranch1'`               |

| `git switch -c feature2`             | Same as above (modern syntax)                        | Creates and switches to branch `feature2`                     |

| `git push origin main`               | Push commits to remote GitHub repo                   | Uploads your local commits to GitHub                          |

| `git push origin feature2`           | Push commits from a feature branch                   | Adds new branch to GitHub with changes                        |

| `git remote add origin <URL>`        | Link your local repo to GitHub                       | Connects your local Git repo to GitHub remote                 |

| `git branch -a`                      | List all local + remote branches                     | Shows `main`, `origin/main`, etc.                             |

| `git merge feature2`                 | Merge another branch into the current one            | Combines changes into your current branch                     |

| `git mv old new`                     | Rename or move a file/folder (tracks rename cleanly) | e.g. `git mv fucntions functions`                             |

| `git diff`                           | Show differences not yet staged                      | Displays code differences between working and staged versions |

| `git diff --staged`                  | Show what’s staged but not committed                 | Shows exact lines that will be committed                      |

| `git restore <file>`                 | Undo unstaged local changes in a file                | Reverts file to last committed state                          |

| `git reset <file>`                   | Unstage a file (but keep its edits)                  | Moves file from “staged” → “modified”                         |

| `git reset --soft HEAD~1`            | Undo the \*last commit\* but keep changes staged       | Useful if commit message was wrong                            |

| `git status`                         | Check everything before pushing                      | Tells what’s staged, untracked, and which branch you’re on    |

| `git push origin your-branch-name`   | Push that branch to GitHub                           | Makes your feature branch visible on GitHub                   |





\## to do commands



| Command                                          | Used For                                              | Example / Effect                             |

| ------------------------------------------------ | ----------------------------------------------------- | -------------------------------------------- |

| `git clone <URL>`                                | Copy a repo from GitHub to your computer              | Downloads entire repo into a local folder    |

| `git pull origin main`                           | Pull latest changes from GitHub                       | Updates your local repo with any new commits |

| `git rm <file>`                                  | Delete a tracked file and stage the deletion          | Removes file from both local and next commit |

| `git stash`                                      | Temporarily save uncommitted changes                  | Stores your changes safely for later         |

| `git stash pop`                                  | Restore stashed changes                               | Brings back your last stashed edits          |

| `git tag -a v1.0 -m "First release"`             | Create a version tag                                  | Useful for marking versions (e.g. releases)  |

| `git show <commitID>`                            | Show detailed info about a specific commit            | Displays the commit diff and metadata        |

| `git reflog`                                     | Show all branch and commit movements                  | Lets you recover lost commits/branches       |

| `git blame <file>`                               | Show who last modified each line                      | Great for tracking code history              |

| `git revert <commitID>`                          | Undo a commit \*safely\* by adding a new inverse commit | Keeps history intact (unlike reset)          |

| `git fetch origin`                               | Download branch info without merging                  | Syncs data about remote branches             |

| `git remote -v`                                  | Show linked remote repo URLs                          | Confirms your repo is connected to GitHub    |

| `git config --global user.name "Your Name"`      | Set Git username (once)                               | Associates commits with your name            |

| `git config --global user.email "you@email.com"` | Set Git email (once)                                  | Associates commits with your GitHub account  |



\## shell commands - used 

| Command                   | Meaning                              

| ------------------------- | ------------------------------------------------------

| `dir`                     | List files and folders (Windows)     

| `mkdir`                   | make new folders                     

| `cd foldername`           | Move into a folder - Enters that directory              

| `cd ..`                   | Move up one directory -  Goes back to parent folder         

| `code .`                  | Open current folder in VS Code       

| `echo "text" >> file.txt` | Append text to a file-  Adds “text” at end of file         

| `type filename`           | Print file content in Command Prompt 







