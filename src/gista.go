package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	var authors = Authors()

	var totalInsertions uint64
	var totalDeletions uint64
	var totalCommits uint64

	for i := range authors {
		ChangesInit(&authors[i])

		totalCommits += uint64(authors[i].Commits)
		totalInsertions += authors[i].Insertions
		totalDeletions += authors[i].Deletions
	}

	println("Authors:", len(authors), "| Total Commits:", totalCommits, "| Total Insertions:", totalInsertions, "| Total Deletions:", totalDeletions, "| (Insertions + Deletions):", totalDeletions+totalInsertions)
  println("--------------------------")

  w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', tabwriter.Debug)
  fmt.Fprintln(w, "| Author\t Insertions\t Deletions\t Commits\t (Insertions+Deletions) |")
	for _, author := range authors {
    fmt.Fprintln(w, "|", author.Name, "\t", author.Insertions, "\t", author.Deletions, "\t", author.Commits, "\t", author.Insertions + author.Deletions)
	}

  w.Flush()
}
