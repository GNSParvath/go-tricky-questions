package main

import "fmt"

func main() {
    fmt.Println("\ncase:1 Basic aliasing — both point to same array")
    sameUnderlyingArray()

    fmt.Println("\ncase:2 Append creates a new backing array (may break alias)")
    appendBreaksAliasing()

    fmt.Println("\ncase:3 Passing slice to function (aliasing continues)")
    functionModifiesOriginal()

    fmt.Println("\ncase:4 Aliasing inside a loop (common trap)")
    sliceInsideLoop()

    fmt.Println("\ncase:5 Slice created with make() using length and capacity")
	sliceMakeWithCap()

	fmt.Println("\n case:6 Iterating over an array using range")
	rangeOverArray()

	fmt.Println("\ncase:7 Appending beyond slice capacity")
	appendBeyondCapacity()
}

// Case 1: Basic aliasing — both point to same array
func sameUnderlyingArray() {
    fmt.Println("sameUnderlyingArray")
    s1 := []int{1, 2, 3}
    s2 := s1
    s2[1] = 99
    fmt.Println("s1:", s1) // [1 99 3]
    fmt.Println("s2:", s2) // [1 99 3]
    fmt.Println()
}

// Case 2: Append creates a new backing array (may break alias)
func appendBreaksAliasing() {
    fmt.Println("appendBreaksAliasing")
    s1 := []int{1, 2, 3}
    s2 := s1
    s2 = append(s2, 4) // If capacity is exceeded, s2 points to new array
    s2[0] = 100
    fmt.Println("s1:", s1) // [1 2 3]
    fmt.Println("s2:", s2) // [100 2 3 4]
    fmt.Println()
}

// Case 3: Passing slice to function (aliasing continues)
func functionModifiesOriginal() {
    fmt.Println("functionModifiesOriginal")
    s := []int{1, 2, 3}
    modifySlice(s)
    fmt.Println("Outside:", s) // [99 2 3]
    fmt.Println()
}

func modifySlice(s []int) {
    s[0] = 99
    fmt.Println("Inside :", s) // [99 2 3]
}

// Case 4: Aliasing inside a loop (common trap) 
func sliceInsideLoop() {
    fmt.Println("liceInsideLoop")
    s := make([][]int, 3)
    temp := []int{1, 2, 3}
    for i := 0; i < 3; i++ {
      s[i] = temp
    }
    s[0][1] = 999
    fmt.Println("All rows alias same slice:", s)
    fmt.Println()
}

// case 5 : sliceMakeWithCap demonstrates how slices behave when created using make with explicit length and capacity.
func sliceMakeWithCap() {
    // Creates a slice of strings with length 3 and capacity 5.
	s := make([]string, 3, 5)

	// The slice has 3 empty string elements.
	fmt.Println("Slice:", s)                     // Output: ["" "" ""]
	fmt.Println("Length:", len(s))               // Output: 3
	fmt.Println("Capacity:", cap(s))             // Output: 5

	// Gotcha:
	// Capacity is 5, but only 3 elements are initialized (default: "").
	// The remaining 2 spaces are unused unless appended.
}

// case:6 rangeOverArray shows how the range keyword works when iterating over an array
func rangeOverArray() {
	// Correct array declaration with fixed size.
	arr := [3]int{1, 2, 3}

	// Range returns the index by default when used as: for i := range arr
	for i := range arr {
		fmt.Println("Index:", i) // Output: 0, 1, 2
	}

	// Tip:
	// Use `for i, v := range arr` if you want both index and value.
}

// case:7 appendBeyondCapacity demonstrates what happens when appending elements beyond a slice’s capacity.
func appendBeyondCapacity() {
	// Create a slice with length 3 and capacity 5.
	s := make([]string, 3, 5)
	fmt.Println("Before append:", s)                            // Output: ["" "" ""]
	fmt.Println("Length:", len(s), "Capacity:", cap(s))         // Output: len: 3, cap: 5

	// Append 3 new elements (which exceeds the capacity).
	s = append(s, "e", "f", "g")

	fmt.Println("After append:", s)                             // Output: ["" "" "" "e" "f" "g"]
	fmt.Println("Length:", len(s), "Capacity:", cap(s))         // Output: len: 6, cap: 10 (reallocated)

	// Tricky Point:
	// When you append beyond capacity, Go allocates a new underlying array (usually with double capacity),
	// and copies all elements into it. The old backing array may be garbage collected if not referenced.
}
