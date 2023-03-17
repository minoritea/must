package must

import "fmt"

func ExampleMust1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, error) { return 0, fmt.Errorf("Must1 error") }

	Must1(f())
	// Output:
	// Must1 error
}

func ExampleMust2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, error) { return 0, 0, fmt.Errorf("Must2 error") }

	Must2(f())
	// Output:
	// Must2 error
}

func ExampleMust3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, error) { return 0, 0, 0, fmt.Errorf("Must3 error") }

	Must3(f())
	// Output:
	// Must3 error
}

func ExampleMust4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, int, error) { return 0, 0, 0, 0, fmt.Errorf("Must4 error") }

	Must4(f())
	// Output:
	// Must4 error
}

func ExampleMust5() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, int, int, error) { return 0, 0, 0, 0, 0, fmt.Errorf("Must5 error") }

	Must5(f())
	// Output:
	// Must5 error
}

func ExampleMust6() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, int, int, int, error) { return 0, 0, 0, 0, 0, 0, fmt.Errorf("Must6 error") }

	Must6(f())
	// Output:
	// Must6 error
}

func ExampleMust7() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, int, int, int, int, error) {
		return 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("Must7 error")
	}

	Must7(f())
	// Output:
	// Must7 error
}

func ExampleMust8() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	f := func() (int, int, int, int, int, int, int, int, error) {
		return 0, 0, 0, 0, 0, 0, 0, 0, fmt.Errorf("Must8 error")
	}

	Must8(f())
	// Output:
	// Must8 error
}
