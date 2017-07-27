<p align="center">
  <img alt="GitHub Logo" src="docs/github_logo.png" height="140" />
  <img alt="Golang Logo" src="docs/golang_logo.png" height="140" style="margin-left: 2em;" />
  <img alt="Bitbucket Logo" src="docs/bitbucket_logo.png" height="140" style="margin-left: 2em;" />
  <h3 align="center">Git Backup</h3>
  <p align="center">Backup your git repositories programmatically to BitBucket.</p>
  <p align="center">
    <a href="https://github.com/oshalygin/git-backup/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/oshalygin/git-backup.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/oshalygin/git-backup"><img alt="Travis" src="https://travis-ci.org/oshalygin/git-backup.svg?branch=master"></a>
    <a href="/LICENSE.md"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
    <a href="https://coveralls.io/github/oshalygin/git-backup?branch=master"><img alt="Coveralls" src="https://coveralls.io/repos/github/oshalygin/git-backup/badge.svg?branch=master"></a>
    <a href="https://codeclimate.com/repos/596c01297de38412b7000136/feed"><img alt="Code Climate Issue Count" src="https://codeclimate.com/repos/596c01297de38412b7000136/badges/d8e88772201d137ea8b7/issue_count.svg"></a>
    <a href="https://goreportcard.com/report/github.com/oshalygin/git-backup"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/oshalygin/git-backup"></a>
    <a href="https://godoc.org/github.com/oshalygin/git-backup"><img src="https://godoc.org/github.com/oshalygin/git-backup?status.svg" alt="GoDoc"></a>
  </p>
</p>

# Introduction

This is a simple and straightforward CLI utility that allows you to backup your git projects to BitBucket.

# Motivation

It is common to be working on public projects on GitHub that you also want another copy of to your BitBucket account.  There are a number of use cases that come to mind that might prompt you to start doing this, such as:

* Ensuring you always have a backup in case other collaborates have force push access.
* If you're working with microservices and you want to pull down the latest code from all of the different repositories and have a backup on BitBucket

# Requirements

Without bloating the utility with a ton of _required_ options, the `git-backup` depends on certain aspects being present in each of your git repositories.

1. Make sure you have an origin remote configured.  Branches will be fetched/pulled from this source.
2. Make sure that you have a bitbucket remote configured.  All branches will be pushed to this source.

```bash
# If you dont have your sources configured, you can do something along these lines:
git remote add origin git@github.com:oshalygin/git-backup.git
git remote add bitbucket https://oshalygin@bitbucket.org/oshalygin/git-backup.git

```

# Installation

```bash
go get -u github.com/oshalygin/git-backup
```

# Usage

```
Usage of git-backup:
  --path string
        The path locations where every directory will be iterated through and backed up.
```

### Example

```bash
# Call the utility at the specified path with the 
# appropriate command line arguments.

git-backup --path=../projects

# Or use the shorthand
git-backup .  # This will run the utility and iterate over all of the folders in the current path

```

### Command Line Arguments

**path**: The path that contains all of the git projects that are to be backed up. 

# Limitations
The utility depends on the requirements, in the future this will be expanded into CLI options.

# License

[MIT](LICENSE)
