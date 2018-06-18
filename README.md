# pancake

[![Go Report Card](https://goreportcard.com/badge/github.com/andrewlader/pancake)](https://goreportcard.com/report/github.com/andrewlader/pancake)
[![Build Status](https://travis-ci.org/andrewlader/pancake.svg?branch=master)](https://travis-ci.org/andrewlader/pancake)
[![codecov](https://codecov.io/gh/andrewlader/pancake/branch/master/graph/badge.svg)](https://codecov.io/gh/andrewlader/pancake)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/AndrewLader/pancake/blob/master/LICENSE)

```
:::::::::     :::     ::::    :::  ::::::::      :::     :::    ::: :::::::::: 
:+:    :+:  :+: :+:   :+:+:   :+: :+:    :+:   :+: :+:   :+:   :+:  :+:        
+:+    +:+ +:+   +:+  :+:+:+  +:+ +:+         +:+   +:+  +:+  +:+   +:+        
+#++:++#+ +#++:++#++: +#+ +:+ +#+ +#+        +#++:++#++: +#++:++    +#++:++#   
+#+       +#+     +#+ +#+  +#+#+# +#+        +#+     +#+ +#+  +#+   +#+        
#+#       #+#     #+# #+#   #+#+# #+#    #+# #+#     #+# #+#   #+#  #+#        
###       ###     ### ###    ####  ########  ###     ### ###    ### ########## 
```

## Why _**pancake**_?
Well, I wrote this just after eating pankcakes. _And_ the purpose of this application is to squash all of the GitHub release notes down into a single `CHANGELOG.md` file. In other words, it _pancakes_ the release notes into a changelog file. :smile:

## Why changelogs?
First, what is a changelog?

### What is a changelog?
A changelog is a curated compilation of all of the releases for a particular project/repository. Each individual release in the log contains the tag name, the date and author of the release, and a few notes about the critical / notable items in that release.

### Who cares?
_Everybody!_ Any one developing on the project, and everyone who is using the product of the project, will care about the changelog. A consumer of the project may want to know what was in the last release before choosing to update to that version. Or, a developer may want to know why something is broken, and will want to check the history. It matters!

And that's why there's _**pancake**_! It makes it easy to create easy-to-read changelogs, that can quickly be recreated as new releases are made.

## Installation
After cloning the repo, just use the following command at the root of the project to install it:
```
$ go install
```

## Usage
The _pancake_ app assumes that there is a `GITHUB_TOKEN` environment variable that is set to an `oauth token` for your GitHub account. To run _pancake_, use the following commands:
```
$ pancake --org ${github_organization_name} --repo ${repo_name}
```
