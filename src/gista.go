package main

func main() {
  var authors = Authors();

  ChangesInit(&authors[1])

  println("Author:", authors[1].Name)
  println("Commits:", authors[1].Commits)

  println("Insertions:", authors[1].Insertions)
  println("Deletions:", authors[1].Deletions)

}
