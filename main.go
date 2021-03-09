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
	commit = `git commit -m %q`
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
	defer fmt.Println("all done")

	daysBack := 365
	for i := daysBack; i >= 1; i-- {
		fmt.Printf("commiting %d days ago\n", i)
		for j := 0; j < 30; j++ {
			changeTxt := faker.Paragraph() + "\n" + fmt.Sprint(rand.Int()) + "\n"
			if err := os.WriteFile("tmp.txt", []byte(changeTxt), 0o777); err != nil {
				log.Fatalln("writing tmp file failed: ", err)
			}
			if _, err := exec.Command("bash", bashArgs(add)...).Output(); err != nil {
				log.Fatalln("git add failed: ", err)
			}

			day := time.Now().AddDate(0, 0, -i)
			timeStr := day.Format("Mon, 02 Jan 2006 15:04:05 -0700")
			os.Setenv("GIT_AUTHOR_DATE", timeStr)
			os.Setenv("GIT_COMMITTER_DATE", timeStr)

			commitMsg := faker.Sentence()
			commitCmd := fmt.Sprintf(commit, commitMsg)
			if _, err := exec.Command("bash", bashArgs(commitCmd)...).Output(); err != nil {
				log.Fatalln("git commit failed: ", err)
			}
		}
	}
}
