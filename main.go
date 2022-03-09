package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/bxcodec/faker/v3"
)

var (
	add    = `git add --all`
	commit = `MY_TIME=$(date --date=\"%d days ago\" -R) GIT_AUTHOR_DATE=$MY_TIME GIT_COMMITTER_DATE=$MY_TIME git commit -m %q`
)

// func main() {
// 	commitMsg := faker.Sentence()
// 	daysBack := 5 * 365
// 	for i := 1; i < daysBack; i++ {
// 		os.WriteFile("tmp.txt", []byte(), 0777)
// 	}
// }
func bashArgs(cmd string) []string {
	return []string{"-c", cmd}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	changeTxt := faker.Paragraph() + "\n" + fmt.Sprint(rand.Int()) + "\n"
	if err := os.WriteFile("tmp.txt", []byte(changeTxt), 0o777); err != nil {
		log.Fatalln("writing tmp file failed: ", err)
	}
	if _, err := exec.Command("bash", bashArgs(add)...).Output(); err != nil {
		log.Fatalln("git add failed: ", err)
	}
	commitCmd := fmt.Sprintf(commit, 2, "boom")
	if _, err := exec.Command("bash", bashArgs(commitCmd)...).Output(); err != nil {
		log.Fatalln("git commit failed: ", err)
	}
}
