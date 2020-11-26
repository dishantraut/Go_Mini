package main 

import ( 
	"fmt"
	"time"
) 

func main() { 

	t := time.Now() 
	fmt.Println(t) 

	// We can get individual values 
	// of hour, minute, seconds, 
	// nanoseconds from time.Time 
	// datatype 

	// Returns wall clock time 
	// from time.Time datatype 
	fmt.Println(t.Clock()) 
	fmt.Println(t.Hour()) 
	fmt.Println(t.Minute()) 
	fmt.Println(t.Second()) 
	fmt.Println(t.Nanosecond()) 

	fmt.Println("---------------------------------") 

	// We can get individual values 
	// of day, month, year, yearday, 
	// and weekday from time.Time 
	// datatype 

	// Returns date from 
	// time.Time datatype 
	fmt.Println(t.Date()) 
	fmt.Println(t.Day()) 
	fmt.Println(t.Month()) 
	fmt.Println(t.Year()) 
	fmt.Println(t.YearDay()) 
	fmt.Println(t.Weekday()) 

	// week number 
	fmt.Println(t.ISOWeek()) 

	fmt.Println("---------------------------------") 

	// current time in string formats 
	fmt.Println(t.String()) 

	// nanoseconds passed 
	// from 1 january 1970 
	fmt.Println(t.Unix()) 

	// prints abbreviated timezone and 
	// its offset in seconds from 
	// east of UTC 
	fmt.Println(t.Zone()) 

	// prints nanoseconds 
	// elapsed from 1 january 1970 
	fmt.Println(t.UnixNano()) 
} 
