package main

import "fmt"

func main() {

	days := [...]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	fmt.Printf("days=%v\n\n", days) // or days[:]

	weekdays := days[:5] // or days[0:5] including 0, excluding 5; called 'slicing'
	fmt.Printf("weekdays=%v\n", weekdays)
	fmt.Printf("len(weekadays)=%d\n", len(weekdays))
	fmt.Printf("cap(weekdays)=%d\n", cap(weekdays)) // *********************************
	fmt.Printf("weekdays[0]=%s\n", weekdays[0])
	fmt.Printf("weekdays[4]=%s\n\n", weekdays[4])
	// fmt.Printf(weekdays[5])			// panic : runtime error : index out of range

	weekend := days[5:] // or days[5:7]
	fmt.Printf("weekend=%v\n", weekend)
	fmt.Printf("len(weekend)=%d\n", len(weekend))
	fmt.Printf("cap(weekend)=%d\n", cap(weekend))
	fmt.Printf("weekend[0]=%s\n", weekend[0])
	fmt.Printf("weekend[1]=%s\n\n", weekend[1])
	// fmt.Printf("weekend[2]")			// panic : runtime error : index out of range

	retired := days[:0]
	fmt.Printf("retired=%v\n", retired)
	fmt.Printf("len(retired)=%d\n", len(retired))
	fmt.Printf("cap(retired)=%d\n\n", cap(retired))
	// fmt.Printf("retired[0]")			// panic : runtime error : index out of range

	fmt.Printf("my vacation=%v\n", days[3:5])

}
