/* 
- Using mutex with example
- let 3 object
- Revenue, Account, and finReport
// revenue get revue, Accounting will add tax and deduct
// and finReport will report
// we want Accounting and findReport access revenue
// simultanously
*/

package main

import (
       "fmt"
       "sync"
)

type Revenue struct {
   mu sync.Mutex // enable async update with goroutines	
   total int
   revenues map[string]int
}

type Accounting struct {
   tax float32
}

type Report struct {
  revenue int
}


func (rev *Revenue) getRevenue(year string) int {
    
     // lock mutex optimistic
     rev.mu.Lock()
     
     // defer release
     defer rev.mu.Unlock()
     return rev.revenues[year]

}

func (ac Accounting) computeTax(gross int) (float32, float32) {
	interest := float32(gross) * ac.tax
	net := float32(gross) - interest
	return net, interest
}

func (rep Report) showReport(ac Accounting, gross int, year string) {
    net, paidInt := ac.computeTax(gross)
    fmt.Printf("Year:\t\t%s\n", year)
    fmt.Printf("Gross:\t\t%d\n", gross)
    fmt.Printf("Tax:\t\t%d%% or $%d\n", int(ac.tax * float32(100)), int(paidInt))
    fmt.Printf("net income:\t%d\n\n", int(net))
}

func main() {

  entr := Revenue{
      total: 0,	  
      revenues: map[string]int{"2001": 5000, "2002": 8000, "2003": 10000 },
   }
  
  acc := Accounting{tax: 0.25}
  report := Report{revenue: 0}

  // create a sync work group
  var wg sync.WaitGroup

  getFinance := func() {
     for year, _ := range entr.revenues {
	     gross := entr.getRevenue(year)
	     acc.computeTax(gross)

	     //report
	     report.showReport(acc, gross, year)

     }
   // wait until jobs done  
   wg.Done()
  }

  wg.Add(1)

  go getFinance()
  //go getFinance()
  
  wg.Wait()
  fmt.Println("Task completed")

}
