#Challenge 5

Given a list of commands ["Add", "Subtract", "Multiply", "Divide"]. Implement a function that can interpret list of
commands as mathematical operator and returns the final result.

Example 1:
input: []
output: 0

Example 2:
input: ["Add 5", "Add 10", "Multiply 3", "Subtract 10"]
output: 35

Note:
1. Operations are executed sequentially without considering mathematical precedence,
2. The command should be case-sensitive, meaning "add" or "suBStraCt" will be ignored,
3. The format of the command are strict (command followed by a whitespace followed by a number without thousand-separator). If a command do not follow that rule, it will be ignored,
4. Division by zero is ignored,
5. the number in the commands are always positive integer or zero.