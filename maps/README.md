## This repository showcases GoLang code I developed while actively practicing and honing my GoLang skills.


### Questions with maps.

1. Word Frequency Counter: Write a function that takes a string as input and returns a map where keys are words in the string, and values are the number of times each word appears.

2. Key Existence Check: Write a program that initializes a map with some key-value pairs. Check if a given key exists in the map, and print the corresponding value if it exists.

3. Reverse a Map: Given a map of string-to-int (map[string]int), write a program to create a new map that reverses the key-value relationship, i.e., an int-to-string map (map[int]string).

4. Count Character Frequency: Write a function that counts the frequency of each character in a string and stores the result in a map.

5. Sum of Values: Write a function that calculates the sum of all integer values in a map where the map is map[string]int.

6. Merge Two Maps: Write a function that merges two maps. If both maps have the same key, the value from the second map should overwrite the value from the first map.

7. Find Keys by Value: Write a function that takes a map of map[string]int and a target value as inputs and returns a slice of all keys that have the target value.

8. Count Words in a Sentence: Write a function that takes a sentence and returns a map where keys are unique words (case-insensitive), and values are their counts.

9. Find Duplicates in a Slice: Given a slice of integers, create a map to count occurrences of each number and identify which numbers are duplicates.

10. Map of Maps: Write a program that uses a map of maps (e.g., map[string]map[string]int) to store and retrieve hierarchical data, like categories and subcategories.

11. Group by Length: Write a function that takes a slice of strings and groups them by their lengths using a map, where the key is the length and the value is a slice of strings with that length.

12. Flatten a Map of Slices: Given a map of map[string][]int, write a function to flatten it into a single slice containing all the integers.

13. Custom Sorting by Value: Write a function that sorts a map of map[string]int by its values in descending order and returns the sorted keys in a slice.

14. Convert Map to JSON: Write a program that takes a map and converts it into a JSON string using Go's encoding/json package.

15. Employee Records: Create a map where keys are employee IDs (int), and values are another map containing employee details (name, age, department). Write functions to:

`    Add a new employee.
    Update employee details.
    Retrieve details by ID.
    Delete an employee by ID.`

    
Answers are in this file: [mapsIn.go](mapsInGo/mapsIn.go)