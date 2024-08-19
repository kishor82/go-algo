# Bubble Sort Algorithm

1. Introduction to Sorting Algorithms
   - Bubble Sort is introduced as a simple, easy-to-visualize sorting algorithm
   - It's described as so simple it could be written under pressure (e.g., on a crashing plane)

2. Definition of a Sorted Array
   - Mathematically: For any position i, the element at i is less than or equal to the element at i+1

3. Bubble Sort Algorithm
   - Start at the first element
   - Compare adjacent elements, swap if they're in the wrong order
   - Repeat until reaching the end of the array
   - After one pass, the largest element is at the end
   - Repeat the process, excluding the last (sorted) element each time

4. Example Walkthrough
   - Detailed step-by-step explanation of sorting the array [1, 3, 4, 2, 7]

5. Key Characteristics
   - After each pass, the largest unsorted element is in its correct position
   - The sorting area decreases by one element after each pass
   - Algorithm stops when only one element is left to sort

6. Time Complexity Analysis
   - First pass: n checks
   - Second pass: n-1 checks
   - Continues until reaching 1 check
   - Sum of checks follows the pattern of summing integers from 1 to n

7. Gauss Anecdote
   - Story about Gauss quickly summing integers from 1 to 100
   - Formula: n(n+1)/2

8. Deriving Time Complexity
   - Using the sum formula: n(n+1)/2
   - Simplifying in Big O notation
   - Final time complexity: O(n²)

9. Big O Notation Rules
   - Drop constants
   - Drop insignificant terms (lower order terms)
