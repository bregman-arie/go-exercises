# Arrays 101

## Objectives

1. Create an empty array of integers of size 10
2. Create an arrays of strings and fill it with 4 different strings
3. Can you mix different data types in one array?
4. You've created an empty array with the size of 10. Once declared can you add items to it? Can you add 50 items?
5. There is an array called 'games'. Assign the string "Donkey Kong" to the first place in the array
6. There is an array called "movies". Print the following:
   1. The array itself
   2. The second item in the array
   3. The length of the array

## Solution

1. `var numbers = [10]int{}`
2. `var strings = [4]string{"some", "random", "strings", ":)"}`
3. No, you can't mix data types in the same array
4. Yes, you can add items to an array after it's declared. No, you can't at 50 items to an array of size 10. You can add items up to the size of the array.
5. `games[0] = "Donkey Kong"`
6. 
   1. `fmt.Printf("Array: %v\n", movies)`
   2. `fmt.Printf("Array's 2nd item: %v\n", movies[1])`
   3. `fmt.Printf("Array's Length: %v\n", len(movies))`