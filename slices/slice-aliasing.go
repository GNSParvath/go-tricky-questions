package main

import "fmt"

func main() {
    sameUnderlyingArray()
    appendBreaksAliasing()
    functionModifiesOriginal()
    sliceInsideLoop()
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
