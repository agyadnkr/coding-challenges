# Challenge 6

# Find Maximum Sum of 5 Consecutive Numbers in A Matrix

Write a function that take two-dimensional array (matrix) of integers as input. The dimension of the matrix will be n x n, where n >= 5. Find the largest sum of any 5 consecutive numbers in the matrix. These 5 consecutive numbers can be aligned vertically or horizontally but not diagonally. Return the largest sum.

Example 1:
input:
[1, 1, 1, 1, 1],
[1, 1, 2, 1, 1],
[1, 1, 3, 1, 1],
[1, 1, 4, 1, 1],
[1, 1, 5, 1, 1]
output: 15
explanation: the sum of numbers from the 3rd column: 1 + 2 + 3 + 4 + 5 = 15

Example 2: 
input:
[1, 1, 1, 1, 1],
[1, 1, 2, 1, 1],
[3, 4, 3, 6, 12],
[1, 1, 4, 1, 1],
[1, 1, 5, 1, 1]
output: 28
explanation: the sum of numbers from the 3rd row: 3, 4, 3, 6, 12 = 28

note: 
1. The matrix will be at least 5 x 5,
2. the matrix will be always symmetrical.