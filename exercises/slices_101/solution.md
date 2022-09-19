# Slices

## Objectives

1. Why do we need slices if we have arrays?
2. How to define an empty slice of strings called `someSlice`?
3. What the following lines does? `append(someSlice, 'abc')`
4. How to retrieve the first element of the slice `someSlice`?

## Solution

1. Slices, similarly to arrays, are also have size and are index based but as opposed to arrays their size can be resized when needed so it makes them a flexible option in case where you can't define what should be exactly the size of an array or when an array is expected to be added with items over time.
2. `var someSlice = []string{}`
3. Adds the string 'abc' to the end of the slice `someSlice`
4. Similarly to how it's done with arrays: `someSlice[0]`