package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	source string
	Amount int
}

// how much many going to make in 52 weeks

func main() {
	// variable fo rbank balance

	var bankBalance int
	var balance sync.Mutex
	// starting values

	fmt.Println("Intial account balance is ", bankBalance)
	// define weekly revenue
	incomes := []Income{
		{source: "main job", Amount: 500},
		{source: "Gifts", Amount: 10},
		{source: "Parttime job", Amount: 50},
		{source: "Investment", Amount: 30},
	}

	wg.Add(len(incomes))

	//loop through 52 weeks
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				t := bankBalance
				t += income.Amount
				bankBalance = t
				balance.Unlock()

				fmt.Printf("on week %d, you earned %d.00 from %s\n", week, income.Amount, income.source)
			}
		}(i, income)
	}
	wg.Wait()
	// print out final balance
	fmt.Println("Total balance:", bankBalance)
}



// output:
//  go run syncMutex\bank\main.go
// Intial account balance is  0
// on week 1, you earned 30.00 from Investment
// on week 2, you earned 30.00 from Investment    
// on week 3, you earned 30.00 from Investment    
// on week 4, you earned 30.00 from Investment    
// on week 5, you earned 30.00 from Investment    
// on week 6, you earned 30.00 from Investment    
// on week 7, you earned 30.00 from Investment    
// on week 8, you earned 30.00 from Investment    
// on week 9, you earned 30.00 from Investment    
// on week 10, you earned 30.00 from Investment   
// on week 11, you earned 30.00 from Investment   
// on week 12, you earned 30.00 from Investment   
// on week 13, you earned 30.00 from Investment   
// on week 14, you earned 30.00 from Investment   
// on week 15, you earned 30.00 from Investment   
// on week 16, you earned 30.00 from Investment   
// on week 17, you earned 30.00 from Investment   
// on week 18, you earned 30.00 from Investment   
// on week 19, you earned 30.00 from Investment   
// on week 20, you earned 30.00 from Investment   
// on week 21, you earned 30.00 from Investment   
// on week 22, you earned 30.00 from Investment   
// on week 23, you earned 30.00 from Investment   
// on week 24, you earned 30.00 from Investment   
// on week 25, you earned 30.00 from Investment   
// on week 26, you earned 30.00 from Investment   
// on week 27, you earned 30.00 from Investment   
// on week 28, you earned 30.00 from Investment   
// on week 29, you earned 30.00 from Investment   
// on week 30, you earned 30.00 from Investment   
// on week 31, you earned 30.00 from Investment   
// on week 1, you earned 500.00 from main job     
// on week 2, you earned 500.00 from main job     
// on week 3, you earned 500.00 from main job     
// on week 4, you earned 500.00 from main job     
// on week 5, you earned 500.00 from main job     
// on week 6, you earned 500.00 from main job     
// on week 7, you earned 500.00 from main job     
// on week 8, you earned 500.00 from main job     
// on week 9, you earned 500.00 from main job     
// on week 10, you earned 500.00 from main job    
// on week 11, you earned 500.00 from main job    
// on week 12, you earned 500.00 from main job    
// on week 13, you earned 500.00 from main job    
// on week 14, you earned 500.00 from main job    
// on week 15, you earned 500.00 from main job    
// on week 16, you earned 500.00 from main job    
// on week 17, you earned 500.00 from main job    
// on week 18, you earned 500.00 from main job    
// on week 19, you earned 500.00 from main job    
// on week 20, you earned 500.00 from main job    
// on week 21, you earned 500.00 from main job    
// on week 22, you earned 500.00 from main job    
// on week 23, you earned 500.00 from main job    
// on week 24, you earned 500.00 from main job    
// on week 25, you earned 500.00 from main job    
// on week 26, you earned 500.00 from main job    
// on week 27, you earned 500.00 from main job    
// on week 28, you earned 500.00 from main job    
// on week 29, you earned 500.00 from main job    
// on week 30, you earned 500.00 from main job    
// on week 31, you earned 500.00 from main job    
// on week 32, you earned 500.00 from main job    
// on week 33, you earned 500.00 from main job    
// on week 34, you earned 500.00 from main job    
// on week 35, you earned 500.00 from main job    
// on week 36, you earned 500.00 from main job    
// on week 37, you earned 500.00 from main job    
// on week 38, you earned 500.00 from main job    
// on week 39, you earned 500.00 from main job    
// on week 40, you earned 500.00 from main job    
// on week 41, you earned 500.00 from main job    
// on week 42, you earned 500.00 from main job    
// on week 43, you earned 500.00 from main job    
// on week 44, you earned 500.00 from main job    
// on week 45, you earned 500.00 from main job    
// on week 46, you earned 500.00 from main job    
// on week 47, you earned 500.00 from main job    
// on week 48, you earned 500.00 from main job    
// on week 49, you earned 500.00 from main job    
// on week 50, you earned 500.00 from main job    
// on week 51, you earned 500.00 from main job    
// on week 52, you earned 500.00 from main job    
// on week 1, you earned 10.00 from Gifts
// on week 2, you earned 10.00 from Gifts
// on week 3, you earned 10.00 from Gifts
// on week 4, you earned 10.00 from Gifts
// on week 5, you earned 10.00 from Gifts
// on week 6, you earned 10.00 from Gifts
// on week 7, you earned 10.00 from Gifts
// on week 8, you earned 10.00 from Gifts
// on week 9, you earned 10.00 from Gifts
// on week 10, you earned 10.00 from Gifts        
// on week 11, you earned 10.00 from Gifts        
// on week 12, you earned 10.00 from Gifts        
// on week 13, you earned 10.00 from Gifts        
// on week 14, you earned 10.00 from Gifts        
// on week 15, you earned 10.00 from Gifts        
// on week 16, you earned 10.00 from Gifts        
// on week 17, you earned 10.00 from Gifts        
// on week 18, you earned 10.00 from Gifts        
// on week 32, you earned 30.00 from Investment   
// on week 33, you earned 30.00 from Investment   
// on week 34, you earned 30.00 from Investment   
// on week 35, you earned 30.00 from Investment   
// on week 36, you earned 30.00 from Investment   
// on week 37, you earned 30.00 from Investment   
// on week 38, you earned 30.00 from Investment   
// on week 39, you earned 30.00 from Investment   
// on week 40, you earned 30.00 from Investment   
// on week 41, you earned 30.00 from Investment   
// on week 42, you earned 30.00 from Investment   
// on week 43, you earned 30.00 from Investment   
// on week 44, you earned 30.00 from Investment   
// on week 1, you earned 50.00 from Parttime job
// on week 2, you earned 50.00 from Parttime job  
// on week 3, you earned 50.00 from Parttime job  
// on week 4, you earned 50.00 from Parttime job  
// on week 5, you earned 50.00 from Parttime job  
// on week 6, you earned 50.00 from Parttime job  
// on week 7, you earned 50.00 from Parttime job  
// on week 8, you earned 50.00 from Parttime job  
// on week 9, you earned 50.00 from Parttime job  
// on week 10, you earned 50.00 from Parttime job 
// on week 11, you earned 50.00 from Parttime job 
// on week 12, you earned 50.00 from Parttime job 
// on week 13, you earned 50.00 from Parttime job 
// on week 14, you earned 50.00 from Parttime job 
// on week 15, you earned 50.00 from Parttime job 
// on week 16, you earned 50.00 from Parttime job 
// on week 17, you earned 50.00 from Parttime job 
// on week 18, you earned 50.00 from Parttime job 
// on week 19, you earned 50.00 from Parttime job 
// on week 20, you earned 50.00 from Parttime job 
// on week 21, you earned 50.00 from Parttime job 
// on week 22, you earned 50.00 from Parttime job 
// on week 23, you earned 50.00 from Parttime job 
// on week 24, you earned 50.00 from Parttime job 
// on week 25, you earned 50.00 from Parttime job 
// on week 26, you earned 50.00 from Parttime job 
// on week 27, you earned 50.00 from Parttime job 
// on week 28, you earned 50.00 from Parttime job 
// on week 29, you earned 50.00 from Parttime job 
// on week 30, you earned 50.00 from Parttime job 
// on week 31, you earned 50.00 from Parttime job 
// on week 32, you earned 50.00 from Parttime job 
// on week 33, you earned 50.00 from Parttime job 
// on week 34, you earned 50.00 from Parttime job 
// on week 35, you earned 50.00 from Parttime job 
// on week 36, you earned 50.00 from Parttime job 
// on week 37, you earned 50.00 from Parttime job 
// on week 38, you earned 50.00 from Parttime job 
// on week 39, you earned 50.00 from Parttime job 
// on week 40, you earned 50.00 from Parttime job 
// on week 41, you earned 50.00 from Parttime job 
// on week 42, you earned 50.00 from Parttime job 
// on week 43, you earned 50.00 from Parttime job 
// on week 44, you earned 50.00 from Parttime job 
// on week 45, you earned 50.00 from Parttime job 
// on week 46, you earned 50.00 from Parttime job 
// on week 47, you earned 50.00 from Parttime job 
// on week 48, you earned 50.00 from Parttime job 
// on week 49, you earned 50.00 from Parttime job 
// on week 50, you earned 50.00 from Parttime job 
// on week 51, you earned 50.00 from Parttime job 
// on week 52, you earned 50.00 from Parttime job 
// on week 45, you earned 30.00 from Investment   
// on week 19, you earned 10.00 from Gifts        
// on week 20, you earned 10.00 from Gifts        
// on week 21, you earned 10.00 from Gifts        
// on week 22, you earned 10.00 from Gifts        
// on week 23, you earned 10.00 from Gifts        
// on week 24, you earned 10.00 from Gifts        
// on week 46, you earned 30.00 from Investment   
// on week 47, you earned 30.00 from Investment   
// on week 48, you earned 30.00 from Investment   
// on week 49, you earned 30.00 from Investment   
// on week 50, you earned 30.00 from Investment   
// on week 51, you earned 30.00 from Investment   
// on week 52, you earned 30.00 from Investment   
// on week 25, you earned 10.00 from Gifts        
// on week 26, you earned 10.00 from Gifts        
// on week 27, you earned 10.00 from Gifts        
// on week 28, you earned 10.00 from Gifts        
// on week 29, you earned 10.00 from Gifts        
// on week 30, you earned 10.00 from Gifts        
// on week 31, you earned 10.00 from Gifts        
// on week 32, you earned 10.00 from Gifts        
// on week 33, you earned 10.00 from Gifts        
// on week 34, you earned 10.00 from Gifts        
// on week 35, you earned 10.00 from Gifts        
// on week 36, you earned 10.00 from Gifts        
// on week 37, you earned 10.00 from Gifts        
// on week 38, you earned 10.00 from Gifts
// on week 39, you earned 10.00 from Gifts        
// on week 40, you earned 10.00 from Gifts        
// on week 41, you earned 10.00 from Gifts        
// on week 42, you earned 10.00 from Gifts        
// on week 43, you earned 10.00 from Gifts        
// on week 44, you earned 10.00 from Gifts        
// on week 45, you earned 10.00 from Gifts        
// on week 46, you earned 10.00 from Gifts        
// on week 47, you earned 10.00 from Gifts        
// on week 48, you earned 10.00 from Gifts        
// on week 49, you earned 10.00 from Gifts        
// on week 50, you earned 10.00 from Gifts        
// on week 51, you earned 10.00 from Gifts        
// on week 52, you earned 10.00 from Gifts        
// Total balance: 30680
