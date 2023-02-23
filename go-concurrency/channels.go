package main

import "fmt"

func ChannelExample() {
	numberChannel := make(chan int)

	go func() {
		numberChannel <- 12
	}()
	nums := <-numberChannel

	fmt.Println(nums)
}

func BufferedChannelExample() {
	names := make(chan string, 3)

	names <- "Paul"
	names <- "George"
	names <- "Ringo"

	fmt.Println(<-names)
	fmt.Println(<-names)
	fmt.Println(<-names)

}

func RangeBufferedChannelExample() {

	powerpuffGirls := make(chan string, 3)

	powerpuffGirls <- "Blossom"
	powerpuffGirls <- "Bubbles"
	powerpuffGirls <- "Buttercup"

	close(powerpuffGirls)

	for p := range powerpuffGirls {
		fmt.Println(p)
	}
}

func IdiomaticBufferedChannelExample() {

	const numOfPositions = 8
	players := []string{"Mike Piazza", "Fred McGriff", "Craig Biggio",
		"Matt Williams", "Barry Larkin", "Dante Bichette", "Barry Bonds", "Tony Gwynn", "Todd Hundley",
		"Jeff Bagwell", "Eric Young", "Chipper Jones", "Ozzie Smith", "Gary Sheffield"}

	startingPlayers := make(chan string, len(players))

	for _, i := range players {
		startingPlayers <- i
	}

	startingLineup := make(chan string, numOfPositions)

	for i := 0; i < numOfPositions; i++ {
		go func() {
			p := <-startingPlayers
			startingLineup <- p
		}()
	}
	close(startingPlayers)

	var roster []string
	for i := 0; i < numOfPositions; i++ {
		roster = append(roster, <-startingLineup)
	}
	close(startingLineup)

	fmt.Println(roster)

}
