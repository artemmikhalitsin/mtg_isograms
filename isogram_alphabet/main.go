package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "sort"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func main() {
  minScore := 23
  file, err := os.Open("isograms.txt")

  if err != nil {
    log.Fatal("Error reading input isogram file")
  }

  log.Println("Loading cards from file...")
  cards := readCards(file)
  log.Println("Cards loaded from file.")

  log.Println("Building the tree...")
  tree := buildNewTree(cards)
  log.Println("Finished building tree.")

  log.Println("Finding all leaves...")
  leaves := findAllLeaves(tree)
  log.Println("All leaves found.")

  rankings := buildRankings(leaves)
  scores := []int{}
  for s := range rankings {
    if s >= minScore { // lets reduce the output a little
      scores = append(scores, s)
    }
  }
  sort.Ints(scores)

  log.Println("----- RESULTS OF THE SEARCH ------")
  for _, score := range scores {
    log.Println("SCORE: ", score)
    combinations := rankings[score]
    log.Println("POSSIBLE CARDS:")
    for _, node := range combinations {
      fmt.Println(getCardNamesQuoted(node.cards), "(Missing letters: ", string(getMissingLetters(node.letters)), ")")
    }
  }
}

func readCards(file *os.File) []*Card {
  scanner := bufio.NewScanner(file)
  result := []*Card{}

  for scanner.Scan() {
    text := scanner.Text()
    result = append(result, makeCardFromText(text))
  }

  return result
}

func getCardNamesQuoted(cards []*Card) []string {
  result := []string{}
  for _, card := range cards {
    result = append(result, fmt.Sprintf("\"%s\"", card.name))
  }

  return result
}

/*
  Score -> slice of nodes that have the cards that make up that score
*/
func buildRankings(leaves []*Node) map[int][]*Node {
  result := make(map[int][]*Node)

  for _, leaf := range leaves {
    score := leaf.getNodeScore()

    _, exists := result[score]
    if !exists {
      result[score] = []*Node{}
    }

    result[score] = append(result[score], leaf)
  }

  return result
}

func getMissingLetters(letters []rune) []rune {
  result := []rune{}

  presence := make(map[rune]struct{})
  for _, l := range letters {
    presence[l] = struct{}{}
  }

  for _, l := range alphabet {
    _, exists := presence[l]
    if !exists {
      result = append(result, rune(l))
    }
  }

  return result
}
