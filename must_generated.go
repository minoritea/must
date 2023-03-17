package must

// Must1 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must1[T1 any](t1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return t1
}

// Must2 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

// Must3 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3
}

// Must4 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4
}

// Must5 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5
}

// Must6 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5, t6
}

// Must7 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must7[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, err error) (T1, T2, T3, T4, T5, T6, T7) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5, t6, t7
}

// Must8 panics if the function result contains an error,
// otherwise it returns the remainder of the return values excluding the error type.
func Must8[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any, T7 any, T8 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, t7 T7, t8 T8, err error) (T1, T2, T3, T4, T5, T6, T7, T8) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5, t6, t7, t8
}
