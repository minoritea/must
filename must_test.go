package must

import "fmt"

func ExampleMust() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() error { return fmt.Errorf("Must error") }

	Must(f())
	// Output:
	// Must error
}
