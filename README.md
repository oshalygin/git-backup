[![Build Status](https://travis-ci.org/oshalygin/k8s-config.svg?branch=master)](https://travis-ci.org/oshalygin/k8s-config)
[![Coverage Status](https://coveralls.io/repos/github/oshalygin/k8s-config/badge.svg?branch=master)](https://coveralls.io/github/oshalygin/k8s-config?branch=master)
[![Code Climate](https://codeclimate.com/repos/59598bef371afb02870005c1/badges/d8e88772201d137ea8b7/gpa.svg)](https://codeclimate.com/repos/59598bef371afb02870005c1/feed)
[![Issue Count](https://codeclimate.com/repos/59598bef371afb02870005c1/badges/d8e88772201d137ea8b7/issue_count.svg)](https://codeclimate.com/repos/593e287a150338028600480b/feed)
# Introduction

This is a simple and straightforward CLI that allows you to backup your git projects to bitbucket.  

# Requirements

In the early stages of this utility, it is assumed that you have your git remote setup, specifically that the remote for bitbucket is named _"bitbucket"_.  If one is not setup, it will throw an error intentionally.

# Installation

```bash
go get -u github.com/oshalygin/k8s-config
```

# Usage

```
Usage of k8s-config:
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
