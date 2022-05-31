package main

import (
    "fmt"
	"strconv"
	"math/rand"
	"time"
	"log"
	"os"

    "github.com/xuri/excelize/v2"
)

func main() {
	players := map[string]int{}
	var namesInHat []string



    f, err := excelize.OpenFile("Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer func() {
        // Close the spreadsheet.
        if err := f.Close(); err != nil {
            fmt.Println(err)
        }
    }()
    // Get all the rows in the Sheet1.
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
        return
    }
	indicator := 0
	var playerName string
	var playerChances int

    for _, row := range rows {
		indicator = 0
        for _, colCell := range row {
			if indicator == 0 {
				playerName = colCell
			} else {
				playerChances, _ = strconv.Atoi(colCell)
			}

			indicator++
			players[playerName] = playerChances
        }
    }

	mapsize := len(players)

	fmt.Printf("How many players do you want to draw? ")
	var amountOfPlayers int
	amountOfPlayers = -1
	for amountOfPlayers < 0 || amountOfPlayers > mapsize {
		fmt.Printf("%d is the Highest amount: ", mapsize)
		fmt.Scanln(&amountOfPlayers)
		fmt.Println(amountOfPlayers)
		if amountOfPlayers < 0 || amountOfPlayers > mapsize {
			fmt.Println("Invalid input! Please enter an amount between 0 and ", mapsize)
		}
	}
	//fmt.Println(players)
	hatOfNames := putNamesInHat(namesInHat, players)
	shuffledNames := shuffleNames(hatOfNames)
	//fmt.Println(shuffledNames)
	finalNames := pickNames(shuffledNames, amountOfPlayers)
	//fmt.Println(finalNames)
	//fmt.Println(shuffledNames)
	//fmt.Println(finalNames)
	writeToFile(finalNames)

}

func remove(s []string, r string) []string {
	var correctedList []string
	//fmt.Println("Removing ", r)
    for i := 0; i < len(s); i++ {
        if s[i] != r {
            correctedList = append(correctedList, s[i])
			//fmt.Println(correctedList)
        }
    }
    return correctedList
}

func putNamesInHat(hat []string, chances map[string]int) []string {
	for key, element := range chances{
		for i := 0; i < element; i++ {
			hat = append(hat,key)
		}
	}
	return hat
}

func shuffleNames(hat []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(hat), func(i, j int) { hat[i], hat[j] = hat[j], hat[i] })
	return hat
}

func pickNames(hat []string, amount int) []string{
	var finalNames []string
	var currName string
	for i := 0; i < amount; i++ {
		rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
		currName = hat[rand.Intn(len(hat))]
		finalNames = append(finalNames, currName)
		hat = remove(hat, currName)
		//fmt.Println(hat)
	}

	return finalNames

}

func writeToFile(winners []string){

	f, err := os.Create("Winners.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    
	for i := range winners {
		_, err2 := f.WriteString(winners[i])

		if err2 != nil {
			log.Fatal(err2)
		}
		_, err3 := f.WriteString("\n")

		if err3 != nil {
			log.Fatal(err3)
		}
	}
	fmt.Println("done")
}
