package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"

	"github.com/kaniusii/guessinggame/keeper"
	"github.com/kaniusii/guessinggame/seeker"
)

func main() {

	var min, max, n int
	var err error

	fmt.Print("Enter minimum value: ")
	_, err = fmt.Scanf("%d\n", &min)
	if err != nil {
		fmt.Println("Error reading minimum value:", err)
		return
	}

	fmt.Print("Enter maximum value: ")
	_, err = fmt.Scanf("%d\n", &max)
	if err != nil {
		fmt.Println("Error reading maximum value:", err)
		return
	}

	fmt.Print("Enter number of iterations (n): ")
	_, err = fmt.Scanf("%d\n", &n)
	if err != nil {
		fmt.Println("Error reading number of iterations:", err)
		return
	}

	totalBob, totalRobert, totalSam, totalEmily := 0, 0, 0, 0
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		k := keeper.Keeper{
			Min: min,
			Max: max,
			Num: min + rand.Intn(max-min+1),
		}
		fmt.Println(k.Num)
		s := seeker.Seeker{
			Keeper: &k,
		}

		wg.Add(4)

		go func() {
			s.Turns = 0
			s.BisectionBob()
			totalBob += s.Turns
			wg.Done()
		}()

		go func() {
			s.Turns = 0
			s.RandomRobert()
			totalRobert += s.Turns
			wg.Done()
		}()

		go func() {
			s.Turns = 0
			s.StartSam()
			totalSam += s.Turns
			wg.Done()
		}()

		go func() {
			s.Turns = 0
			s.EndEmily()
			totalEmily += s.Turns
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Total turns for Bob:", totalBob)
	fmt.Println("Total turns for Robert:", totalRobert)
	fmt.Println("Total turns for Sam:", totalSam)
	fmt.Println("Total turns for Emily:", totalEmily)
	fmt.Println()

	fmt.Println("Average turns for Bob:", strconv.FormatFloat(float64(totalBob)/float64(n), 'f', 2, 64))
	fmt.Println("Average turns for Robert:", strconv.FormatFloat(float64(totalRobert)/float64(n), 'f', 2, 64))
	fmt.Println("Average turns for Sam:", strconv.FormatFloat(float64(totalSam)/float64(n), 'f', 2, 64))
	fmt.Println("Average turns for Emily:", strconv.FormatFloat(float64(totalEmily)/float64(n), 'f', 2, 64))
}
