# pancake

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
