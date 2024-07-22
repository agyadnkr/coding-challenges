#Challenge 5

Given a list of commands ["Add", "Subtract", "Multiply", "Divide"]. Implement a function that can interpret list of commands as mathematical operator and returns the final result.

Example 1:
input: []
output: 0

Example 2: 
input: ["Add 5", "Add 10", "Multiply 3", "Subtract 10"]
output: 35

Note: 
1. Operations are executed sequentially without considering mathematical precedence,
2. The command should be case-sensitive, meaning "add" or "suBStraCt" will be counted as 0. 