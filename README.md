# 🧹 Git Branch Cleaner

🚀 A simple and interactive CLI tool to help you delete local Git branches easily.

🔄 Instead of manually typing `git branch -D <branch-name>`, this tool provides an **interactive prompt** where you can **select multiple branches** and delete them in one go.

---

## 📥 Installation

### 🔹 Option 1: Install via Go _(Recommended)_

If you have Go installed, you can install `git-branch-cleaner` globally with:

```sh
go install github.com/jobe95/git-branch-cleaner@latest
```

### 🔹 Option 2: Install Manually

1. Clone this repository:

```sh
git clone https://github.com/jobe95/git-branch-cleaner.git
cd git-branch-cleaner
```

2. Install the binary:

```sh
make install
```

## ⚡ Usage

Simply run:

```sh
git-branch-cleaner
```

You'll see a list of all your local branches (except the current one). You can select multiple branches and confirm before deleting them.

```sh
Select branches to delete: [Use arrows to move, space to select, <right> to all, <left> to none, type to filter]
  [x] feature/refactor    5e215e0 (7 minutes ago)
  [x] feature/ui-fix      1a1deaf (31 minutes ago)
  [ ] feature/cleanup     daf691d (47 minutes ago)

🔥 Are you sure you want to delete 2 branch(es)? [y/N] Yes
🔥 Deleting selected branches...
Deleted branch feature/refactor (was 5e215e0).
Deleted branch feature/ui-fix (was 1a1deaf).
```

## 🔄 Uninstallation

To remove git-branch-cleaner from your system:

### 📌 If installed via Go:

```sh
rm "$(go env GOPATH)/bin/git-branch-cleaner"
```

### 📌 If installed manually:

```sh
make uninstall
```
