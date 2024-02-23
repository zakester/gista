package main

import (
	"os/exec"
	"strconv"
	"strings"
	"unicode"
)

type Author struct {
	Name       string
	Commits    uint32
	Insertions uint64
	Deletions  uint64
}

func Authors() []Author {
	var authors []Author
	cmd := exec.Command("git", "shortlog", "HEAD", "-snu")

	output, err := cmd.Output()
	if err != nil {
		println("Error: ", err)
		return authors
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		commits, err := strconv.ParseUint(fields[0], 10, 32)
		if err != nil {
			println("Error: ", err)
			continue
		}

		authors = append(authors, Author{
			Name:    fields[1],
			Commits: uint32(commits),
		})
	}

	return authors
}

func ChangesInit(author *Author) {
	cmd := exec.Command("git", "log", "--numstat", "--pretty=\"%s\"", "--author="+author.Name)

	output, err := cmd.CombinedOutput()
	if err != nil {
		println(string(output))
		println(err.Error())
		return
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) != 0 && unicode.IsDigit(rune(fields[0][0])) {

			insertions, err := strconv.ParseUint(fields[0], 10, 64)
			if err != nil {
				println("Error: ", err)
				return
			}

			deletions, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				println("Error: ", err)
				return
			}

			author.Insertions += insertions
			author.Deletions += deletions
		}
	}
}
