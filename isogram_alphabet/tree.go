package main

import "log"

type Node struct {
  cards []*Card
  letters []rune
  children []*Node
}

func (n *Node) getNodeScore() int {
  if len(n.letters) == 23 {
    log.Println("DEBUG:", string(n.letters))
  }
  return len(n.letters);
}

func buildNewTree(allCards []*Card) *Node {
  log.Println("Building the tree...")
  result := buildTree(nil, allCards)
  log.Println("Finished building tree.")

  return result
}

func makeNode(rootCards []*Card) *Node {
  return &Node {
    cards: rootCards,
    letters: getLettersInAllCards(rootCards),
    children: []*Node{},
  }
}

func buildTree(rootCards, allCards []*Card) *Node{
  root := makeNode(rootCards)

  for i, card := range allCards {
    if isCandidate(root.letters, card) {
      if rootCards == nil {
        // Log the cards being processed on the first level
        // to see some output
        log.Println("Now building tree for", card.name)
      }

      subtree := buildTree(append(rootCards, card), allCards[i+1:])
      root.children = append(root.children, subtree)
    }
  }

  return root
}

func isCandidate(existingLetters []rune, card *Card) bool {
  return noDuplicates(append(existingLetters, card.letters...))
}

func noDuplicates(runes []rune) bool {
  letters := make(map[rune]struct{})
  for _, rune := range runes {
    _, exists := letters[rune]
    if exists {
      return false
    } else {
      letters[rune] = struct{}{}
    }
  }

  return true
}

func findAllLeaves(root *Node) []*Node {
  result := []*Node{}

  if len(root.children) < 1 {
    //it's a leaf
    return []*Node{root}
  }

  for _, child := range root.children {
    leaves := findAllLeaves(child)
    result = append(result, leaves...)
  }

  return result
}
