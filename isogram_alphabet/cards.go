package main

import (
  "strings"
  "unicode"

  "golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Card struct {
  name string
  letters []rune
}

func makeCardFromText(text string) *Card {
  return &Card{
    name: text,
    letters: normalizeLetters(text),
  }
}

func normalizeLetters(text string) []rune {
  runes := []rune(strings.ToLower(text))
  runes = filterNonAlphabet(runes)
  runes = removeAccents(runes)

  return runes
}

func removeAccents(letters []rune) []rune {
  t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, string(letters))

  return []rune(s)
}

func filterNonAlphabet(runes []rune) []rune {
  result := []rune{}
  for _, r := range runes {
    if unicode.In(r, unicode.Ll) {
      result = append(result, r)
    }
  }

  return result
}

func getLettersInAllCards(cards []*Card) []rune {
  result := []rune{}
  for _, card := range cards {
    result = append(result, card.letters...)
  }
  return result
}
