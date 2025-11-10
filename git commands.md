\# 🧠 Git Commands Cheat Sheet



A simple list of all Git commands I learned and practiced while building my repository.



---



\## 🏗️ 1. Setup \& Initialization



```bash

git init

````



> Initialize a new Git repository in the current folder.



```bash

git remote add origin <repo\_url>

```



> Link your local project to a GitHub repository.



```bash

git remote -v

```



> Verify the remote repository connection.



---



\## 📂 2. Checking and Viewing Files



```bash

dir

```



> (Windows) List files in the current directory.



```bash

ls

```



> (Mac/Linux) List files in the current directory.



```bash

git status

```



> See which files are new, modified, or staged.



---



\## ✍️ 3. Adding \& Committing Changes



```bash

git add .

```



> Stage all files for commit.



```bash

git add <filename>

```



> Stage a specific file only.



```bash

git commit -m "your message"

```



> Save the staged changes with a message.



---



\## 🚀 4. Pushing to GitHub



```bash

git push origin main

```



> Push your commits to the `main` branch on GitHub.



```bash

git push origin <branch\_name>

```



> Push your branch to GitHub (e.g., `feature1`, `feature2`).



---



\## 🌿 5. Branching \& Merging



```bash

git branch

```



> List all branches.



```bash

git branch <branch\_name>

```



> Create a new branch.



```bash

git switch <branch\_name>

```



> Switch to another branch.



```bash

git switch -c <new\_branch\_name>

```



> Create and switch to a new branch.



```bash

git merge <branch\_name>

```



> Merge another branch into the current one.



---



\## 🔍 6. Checking History



```bash

git log

```



> Show detailed commit history.



```bash

git log --oneline

```



> Show compact commit history.



---



\## 🔄 7. Undo \& Restore



```bash

git restore <filename>

```



> Discard changes in a specific file (revert to last commit).



```bash

git reset --soft HEAD~1

```



> Undo the last commit but keep the changes staged.



```bash

git reset --hard HEAD~1

```



> Undo the last commit \*\*and\*\* remove all related changes.



---



\## ⚙️ 8. Common Combos \& Shortcuts



```bash

git add . \&\& git commit -m "update" \&\& git push

```



> Add, commit, and push all in one line.



```bash

git pull origin main

```



> Fetch and merge the latest changes from GitHub.



---





