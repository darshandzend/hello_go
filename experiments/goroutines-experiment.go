/* Program to show goroutines aren't truly parallel by default
: The other goroutine will only execute when the currently
executing goroutine has blocked e.g. on Sleep or a channne-
l.

The following program will ALWAYS have a predicted output unless
GOMAXPROCS is manually set.
*/

package main

import "fmt"
import "time"
import "runtime"

func routine(id int) {

	fmt.Println("Entered Goroutine", id)

	/* this will explicitely yield the running goroutine.
	switch the comment below to understand more */

	//runtime.Gosched()
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine", id, "Loop", i)
	}

	time.Sleep(time.Second)

	fmt.Println("Exited Goroutine", id)
}

func main() {

	for i := 0; i < 10; i++ {
		go routine(i)
	}

	var in string
	fmt.Scanln(&in)
}
