package main

import "fmt"

// I is an interface that defines a method J().
// Any type that implements the J() method satisfies the I interface.
type I interface {
	J()
}

// T is a struct with a single string field S.
type T struct {
	S string
}

// J is a method associated with the *T (pointer to T) type.
// This method implements the J() method of the I interface,
// making *T a type that satisfies the I interface.
func (t *T) J() {
	// If the receiver 't' (the pointer to T) is nil,
	// it means the underlying struct is not allocated or points to nothing.
	if t == nil {
		fmt.Println("<nil>")
		return // Exit the method if t is nil
	}
	// If 't' is not nil, print the value of its S field.
	fmt.Println(t.S)
}

// main is the entry point of the program.
func main() {
	// Declare a variable 'i' of interface type I.
	// Initially, 'i' holds a nil concrete value and a nil concrete type.
	var i I

	// Declare a variable 't' of type *T (pointer to T).
	// Initially, 't' is a nil pointer.
	var t *T

	// Assign the nil pointer 't' to the interface variable 'i'.
	// At this point, 'i' holds a nil concrete value, but its concrete type
	// is *main.T because 't' is of type *main.T.
	i = t
	// Call the describe function to show the concrete value and type stored in 'i'.
	// Output: (<nil>, *main.T)
	describe(i)
	// Call the J() method on the interface variable 'i'.
	// Due to dynamic dispatch, the J() method of *T will be called.
	// Since the concrete value within 'i' is nil, the J() method's
	// 'if t == nil' condition will be true, printing "<nil>".
	i.J()

	// Assign a new, non-nil pointer to a T struct to the interface variable 'i'.
	// The struct is initialized with the string "hello".
	// Now, 'i' holds a non-nil concrete value (&T{"hello"}) and its
	// concrete type is still *main.T.
	i = &T{"hello"}
	// Call the describe function again to show the updated concrete value and type in 'i'.
	// Output: (&{hello}, *main.T)
	describe(i)
	// Call the J() method on the interface variable 'i' again.
	// The J() method of *T will be called.
	// Since the concrete value within 'i' is now non-nil, the J() method's
	// 'if t == nil' condition will be false, and it will print "hello".
	i.J()
}

// describe is a helper function that prints the concrete value and concrete type
// stored within an interface variable.
func describe(i I) {
	// %v prints the value in a default format.
	// %T prints the type of the value.
	fmt.Printf("(%v, %T)\n", i, i)
}