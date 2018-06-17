package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

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

	fmt.Print("Mixing batter...\n")
	startTime := time.Now()
	releaseNotes, numberOfReleaseNotes, err := repo.GetChangelog(token, *orgFlag, *repoFlag)
	if err != nil {
		log.Fatalf("Error! %v", err)
	}
	totalTime := time.Now().Sub(startTime)

	fmt.Printf("Making %d pancakes...\n", numberOfReleaseNotes)

	fmt.Print("On the griddle...\n")
	file, err := os.Create("CHANGELOG.md")
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	defer file.Close()

	fmt.Print("Ready to serve...\n")
	fmt.Fprintf(file, releaseNotes)

	fmt.Printf("Completed processing in %f seconds.\nAll done!\n", totalTime.Seconds())
}
