package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("hello world")

	// arraySign([]int{2, 1})                    // 1
	// arraySign([]int{-2, 1})                   // -1
	// arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1

	// isAnagram("anak", "kana")       // true
	// isAnagram("anak", "mana")       // false
	// isAnagram("anagram", "managra") // true

	// findTheDifference("abcd", "abcde") // 'e'
	// findTheDifference("abcd", "abced") // 'e'
	// findTheDifference("", "y")         // 'y'

	// canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	// canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	// canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2

	// tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code her
	sign := 1 // default value (positive value)

	for _, v := range nums {
		if v == 0 {
			return 0
		}

		if v < 1 {
			sign *= -1
		}

	}

	fmt.Println(sign)
	return sign
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	s = strings.ToLower(s)
	t = strings.ToLower(t)

	if len(s) != len(t) {
		fmt.Println(false)
		return false
	}

	// Convert strings to slices of runes
	str1 := []rune(s)
	str2 := []rune(t)

	// Sort the slices
	sort.Slice(str1, func(i, j int) bool {
		return str1[i] < str1[j]
	})
	sort.Slice(str2, func(i, j int) bool {
		return str2[i] < str2[j]
	})

	// Compare sorted slices
	for i := range str1 {
		if str1[i] != str2[i] {
			fmt.Println(false)
			return false
		}
	}

	fmt.Println(true)
	return true
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

	for i := 0; ; i++ {
		ch := t[i]
		cnt[ch-'a']--
		if cnt[ch-'a'] < 0 {
			fmt.Println(string(ch))
			return ch
		}
	}
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// Sort the array
	sort.Ints(arr)

	// Compute the common difference
	diff := arr[1] - arr[0]

	// Check the difference between each pair of consecutive elements
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return false
		}
	}

	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	return nil
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	return nil
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()

	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
