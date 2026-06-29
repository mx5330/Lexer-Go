# Lexer-Go
A simple Programming Language lexer made in Golang. Supports Integer numbers, Strings, Characters, Keywords and Identifiers.

# Use
## Code:  
Token Structure:  
type Token struct {  
	Type  string  
	Value any  
}  
  

All the words in the Keywords map will be taken as Language keywords instead of identifiers.  
The string that is assigned to the keyword is what will be in the Token.  
All keywords should be in lowercase. The code compares the input in lowercase.

## Examples:
The output for the next code:  
"  
var my_var_1 = "hello world"  
var coolMATH = 200 - 30 * 54 + 10  
  
print(my_var_1)  
"  
Looks like: (Type Value)  
<img width="240" height="342" alt="image" src="https://github.com/user-attachments/assets/4cdc0468-5cca-4164-9909-15f4b6804b16" />


# License
You may use/modify this for your personal projects. However for commercial use you need to contact me and we will get to something.
