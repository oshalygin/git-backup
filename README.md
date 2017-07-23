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
    <a href="https://coveralls.io/github/oshalygin/k8s-config?branch=master"><img alt="Coveralls" src="https://coveralls.io/repos/github/oshalygin/k8s-config/badge.svg?branch=master"></a>
    <a href="https://codeclimate.com/repos/596c01297de38412b7000136/feed"><img alt="Code Climate Issue Count" src="https://codeclimate.com/repos/596c01297de38412b7000136/badges/d8e88772201d137ea8b7/issue_count.svg"></a>
    <a href="https://goreportcard.com/report/github.com/oshalygin/git-backup"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/oshalygin/git-backup"></a>
    <a href="https://godoc.org/github.com/oshalygin/git-backup"><img src="https://godoc.org/github.com/oshalygin/git-backup?status.svg" alt="GoDoc"></a>
  </p>
</p>

# Introduction

This is a simple and straightforward CLI that allows you to backup your git projects to bitbucket.  

# Requirements

In the early stages of this utility, it is assumed that you have your git remote setup, specifically that the remote for bitbucket is named _"bitbucket"_.  If one is not setup, it will throw an error intentionally.

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

```

### Command Line Arguments

**path**: The path that contains all of the git projects that are to be backed up. 

# Limitations
* TBD

# License

[MIT](LICENSE)
