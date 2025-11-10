Q1. **Difference between git pull and git fetch**

* `git fetch` downloads the latest changes from the remote repository (like new commits or branches) but doesn’t apply them to your current working files. It simply updates your local copy of the remote branches so you can review what changed.
* `git pull` actually does two things — it first performs a `fetch`, then immediately merges those fetched changes into your current branch.
* In simple terms:

  * **fetch** = “check what’s new”
  * **pull** = “get what’s new and merge it in”

---

Q2. **What is git cherry-pick?**

* `git cherry-pick` is used when you want to take one specific commit from another branch and apply it to your current branch.
* Instead of merging the whole branch, you pick only one “cherry” (commit) that you need.
* Example:

  * You made a small bug fix in `feature-x`, and now you want the same fix in `main` without merging everything.
  * You can use:

    ```
    git cherry-pick <commit-hash>
    ```
* It’s very useful when you only need specific changes, not the full branch.

---

Q3. **Is Go a value or reference language?**

* Go is mainly a **value-type** language, meaning variables are passed and assigned by copying values.
* When you assign one variable to another, a separate copy is created. Changing one doesn’t affect the other.
* Example:

  ```go
  a := 10
  b := a
  b = 20
  fmt.Println(a) // still prints 10
  ```
* But Go also allows you to use **pointers** for reference-type behavior.

  ```go
  p := &a
  *p = 30
  fmt.Println(a) // now prints 30
  ```
* So, Go is value-based by default but supports references using pointers when needed.

---

Q4. **Important concept – HEAD in Git**

* `HEAD` represents your current position in Git — it points to the latest commit on the branch you’re currently working on.
* Examples:

  * `HEAD -> main` means you’re on the main branch.
  * `HEAD -> feature/login` means you’re on that feature branch.
  * `HEAD~1` means one commit before the current one.
* You can think of `HEAD` as a marker that tells Git, “this is where I currently am in the project’s history.”

---

**What’s the difference between `git merge` and `git rebase`?**

* **Merge** combines two branches by creating a new “merge commit.” It keeps the full branch history.
* **Rebase** moves or “replays” your commits on top of another branch to make the history linear and cleaner.
* Example:

  * `merge` keeps all branches visible → good for collaboration.
  * `rebase` rewrites history → good for a tidy project log.
* In short: **merge = preserve history, rebase = rewrite history for clarity.**

---

**What happens when you clone a repository?**

* `git clone` copies an existing remote repository (from GitHub, GitLab, etc.) onto your local machine.
* It automatically sets up:

  * A `.git` folder (all Git tracking data)
  * A working directory with project files
  * A link to the remote (`origin`) so you can pull/push easily later
* So after cloning, you can start working immediately — you already have history, branches, and connection to the original repo.

---

**What is the difference between `git reset`, `git revert`, and `git restore`?**

* **`git reset`**: Moves the `HEAD` to a previous commit. It can change your commit history (used locally).

  * `--soft` keeps changes staged,
  * `--hard` removes all changes.
* **`git revert`**: Creates a new commit that *undoes* the changes made by a previous commit — safe for shared repos.
* **`git restore`**: Used to undo changes in working files without touching commit history.
* In short:

  * `reset` → move HEAD (dangerous)
  * `revert` → make a new undo commit (safe)
  * `restore` → discard working changes

