package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/andrewlader/pancake/repo"
)

const asciiArt = `

:::::::::     :::     ::::    :::  ::::::::      :::     :::    ::: :::::::::: 
:+:    :+:  :+: :+:   :+:+:   :+: :+:    :+:   :+: :+:   :+:   :+:  :+:        
+:+    +:+ +:+   +:+  :+:+:+  +:+ +:+         +:+   +:+  +:+  +:+   +:+        
+#++:++#+ +#++:++#++: +#+ +:+ +#+ +#+        +#++:++#++: +#++:++    +#++:++#   
+#+       +#+     +#+ +#+  +#+#+# +#+        +#+     +#+ +#+  +#+   +#+        
#+#       #+#     #+# #+#   #+#+# #+#    #+# #+#     #+# #+#   #+#  #+#        
###       ###     ### ###    ####  ########  ###     ### ###    ### ########## 

`

var orgFlag = flag.String("org", "", "Specifies the organization")
var repoFlag = flag.String("repo", "", "Specifies the repository")

func init() {
	flag.StringVar(orgFlag, "o", "", "Specifies the organization")
	flag.StringVar(repoFlag, "r", "", "Specifies the repository")
}

func main() {
	const githubTokenEnvName = "GITHUB_TOKEN"

	flag.Parse()

	fmt.Println(asciiArt)

	token := os.Getenv(githubTokenEnvName)

	log.Print("Mixing batter...")
	releaseNotes, err := repo.GetChangelog(token, *orgFlag, *repoFlag)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}

	log.Print("On the griddle...")
	file, err := os.Create("CHANGELOG.md")
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	defer file.Close()

	log.Print("Ready to serve...")
	fmt.Fprintf(file, releaseNotes)

	log.Print("All done!")
}
