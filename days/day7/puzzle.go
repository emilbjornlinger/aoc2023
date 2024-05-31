package day7

import (
    "fmt"
    "os"
    "path/filepath"
    "aoc2023/input"
    "sort"
    "strings"
    "strconv"
)

type HandType int

const (
    HighCard        HandType = 0
    OnePair         HandType = 1
    TwoPair         HandType = 2
    ThreeOfAKind    HandType = 3
    FullHouse       HandType = 4
    FourOfAKind     HandType = 5
    FiveOfAKind     HandType = 6
)

type Card int

const (
    Two     Card = 2
    Three   Card = 3
    Four    Card = 4
    Five    Card = 5
    Six     Card = 6
    Seven   Card = 7
    Eight   Card = 8
    Nine    Card = 9
    Ten     Card = 10
    Knight  Card = 11
    Queen   Card = 12
    King    Card = 13
    Ace     Card = 14
)

type Hand struct {
    cards []Card
    handType HandType
}

type CardCounter struct {
    card Card 
    numOfCards int
}

type Entry struct {
    hand Hand 
    bid int
}

// Define type that implements sort.Interface
type EntryList []Entry

const dayName string = "day7"

// 255365857 too high
func Puzzle1() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    entries := make([]Entry, 0)

    // Populate entries
    for _, line := range inputSlice {
        entries = append(entries, CreateEntryFromString(line))
    }

    // Sort the entries
    sort.Stable(EntryList(entries))

    // After sorted loop through and use loop index to determine rank and sum up the winnings
    totalWinnings := 0
    for i, entry := range(entries) {
        totalWinnings = totalWinnings + (i + 1) * entry.bid
    }

    fmt.Printf("Output: %v\n", totalWinnings)
}

func Puzzle2() {
    wd, err  := os.Getwd()
    if err != nil {
        panic (err)
    }
    filename := filepath.Join(wd, "days", dayName, "input.txt")
    inputSlice := input.GetInputSlice(filename)

    // Implementation
    for _, line := range inputSlice {
        fmt.Println(line)
    }
    
    // For puzzle 2, create a new function CreateEntryFromStringWithSwappingJs 
    // that will do the same thing as CreateEntryFromString but after GetHandFromCards has been called will try to substitute 
    // The J's to be one of the other existing cards, try this for all combinations available and then check what the 
    // highest HandType returned is, set the final handType to be this but keep the cards as they were

    output := "Hello from " + dayName
    fmt.Printf("Output: %v\n", output)
}

func CreateEntryFromString(line string) Entry {
    newEntry := Entry{}
    cards, bid := GetCardsAndBid(line)

    // Set the bid of the entry
    newEntry.bid = bid

    // Create and set the hand of the entry
    newEntry.hand = GetHandFromCards(cards)

    return newEntry
}

func GetCardsAndBid(line string) (string, int) {
    words := strings.Split(line, " ")

    cards := words[0]
    bid, err := strconv.Atoi(words[1])
    if err != nil {
        panic (err)
    }

    return cards, bid
}

func GetHandFromCards(cardString string) Hand {
    newHand := Hand{}

    newHand.cards = GetCardsFromString(cardString)

    // Create list of cardCounter and loop through each card
    cardCounter := make([]CardCounter, 0)
    for _, card := range newHand.cards {
        indexOfCard := -1
        for i := 0; i < len(cardCounter); i++ {
            if card == cardCounter[i].card {
                indexOfCard = i
                break
            }
        }

        if indexOfCard != -1 {
            cardCounter[indexOfCard].numOfCards = cardCounter[indexOfCard].numOfCards + 1
        } else {
            cardCounter = append(cardCounter, CardCounter{card: card, numOfCards: 1})
        }
    }

    // Perform logic to discern type of hand
    switch len(cardCounter) {
    case 1:
        newHand.handType = FiveOfAKind
    case 2:
        if IsFourOfAKind(cardCounter) {
            newHand.handType = FourOfAKind
        } else {
            newHand.handType = FullHouse
        }
    case 3:
        if IsThreeOfAKind(cardCounter) {
            newHand.handType = ThreeOfAKind
        } else {
            newHand.handType = TwoPair
        }
    case 4:
        newHand.handType = OnePair
    case 5:
        newHand.handType = HighCard
    }

    return newHand
}

func GetCardsFromString(cardString string) []Card {
    cards := make([]Card, 0)
    
    for i := 0; i < len(cardString); i++ {
        switch cardString[i] {
        case '2':
            cards = append(cards, Two)
        case '3':
            cards = append(cards, Three)
        case '4':
            cards = append(cards, Four)
        case '5':
            cards = append(cards, Five)
        case '6':
            cards = append(cards, Six)
        case '7':
            cards = append(cards, Seven)
        case '8':
            cards = append(cards, Eight)
        case '9':
            cards = append(cards, Nine)
        case 'T':
            cards = append(cards, Ten)
        case 'J':
            cards = append(cards, Knight)
        case 'Q':
            cards = append(cards, Queen)
        case 'K':
            cards = append(cards, King)
        case 'A':
            cards = append(cards, Ace)
        }
    }

    return cards
}

func IsFourOfAKind(c []CardCounter) bool {
    if len(c) != 2 {
        panic ("Called IsFourOfAKind on cardCounter that did not have 2 types of cards")
    }

    if c[0].numOfCards == 4 || c[1].numOfCards == 4 {
        return true
    } else {
        return false
    }
}

func IsThreeOfAKind(c []CardCounter) bool {
    if len(c) != 3 {
        panic ("Called IsThreeOfAKind on cardCounter that did not have 3 types of cards")
    }

    if c[0].numOfCards == 3 || c[1].numOfCards == 3  || c[2].numOfCards == 3 {
        return true
    } else {
        return false
    }
}

func (e EntryList) Len() int {
    return len(e)
}

func (e EntryList) Swap(i, j int) {
    e[i], e[j] = e[j], e[i]
}

func (e EntryList) Less(i, j int) bool {
    // First compare handtypes
    if e[i].hand.handType < e[j].hand.handType {
        return true
    } else if e[i].hand.handType > e[j].hand.handType {
        return false
    } else {
        // If equal handtype, loop through cards and return true
        // if i has lower card in first differing spot
        for idx := 0; idx < len(e[i].hand.cards); idx++ {
            if e[i].hand.cards[idx] < e[j].hand.cards[idx] {
                return true
            } else if e[i].hand.cards[idx] > e[j].hand.cards[idx] {
                return false
            }
        }
    }

    // Return false if same hand
    fmt.Println("Same card present")
    fmt.Println(e[i])
    fmt.Println(e[j])
    return false
}
