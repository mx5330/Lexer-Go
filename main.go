package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Token struct {
	Type  string
	Value any
}

var IdentifierPattern = regexp.MustCompile("[A-Za-z_]")
var IdentifierPatternAfter = regexp.MustCompile("[A-Za-z_0-9]")
var NumberPattern = regexp.MustCompile("[0-9]")

var Keywords map[string]string = map[string]string{
	"var":   "VARIABLE",
	"print": "PRINT",
}

var SpecialCharacters map[string]string = map[string]string{
	"=":  "ASSIGN",
	"==": "EQUALS",
	"+":  "ADDITION",
	"-":  "SUBSTRACTION",
	"/":  "DIVISION",
	"*":  "MULTIPLICATION",
	"(":  "LEFT_PAREN",
	")":  "RIGHT_PAREN",
	",":  "COMMA",
	".":  "DOT",
}

func GetTokens(input string) []Token {
	TokenList := []Token{}
	Code := []rune(input)
	TotalLetters := len(Code)

	CurrentPosition := 0
	for {
		if CurrentPosition >= TotalLetters {
			break
		}
		CurrentCharacter := Code[CurrentPosition]

		if CurrentCharacter == '"' {
			StartPosition := CurrentPosition
			for {
				CurrentPosition++
				if CurrentPosition >= TotalLetters {
					break
				}
				CurrentCharacter = Code[CurrentPosition]

				if CurrentCharacter == '"' {
					CurrentPosition++ // Consume last "
					break
				}
			}

			FinalString := string(Code[StartPosition:CurrentPosition])
			TokenList = append(TokenList, Token{
				Type:  "STRING",
				Value: FinalString,
			})
		} else if IdentifierPattern.MatchString(string(CurrentCharacter)) {
			StartPosition := CurrentPosition
			for {
				CurrentPosition++
				if CurrentPosition >= TotalLetters {
					break
				}
				CurrentCharacter = Code[CurrentPosition]

				if !IdentifierPatternAfter.MatchString(string(CurrentCharacter)) {
					break
				}
			}

			FinalIdentifier := string(Code[StartPosition:CurrentPosition])
			keyword, isKeyword := Keywords[strings.ToLower(FinalIdentifier)]

			if isKeyword {
				TokenList = append(TokenList, Token{
					Type:  "KEYWORD",
					Value: keyword,
				})
			} else {
				TokenList = append(TokenList, Token{
					Type:  "IDENTIFIER",
					Value: FinalIdentifier,
				})
			}
		} else if NumberPattern.MatchString(string(CurrentCharacter)) {
			StartPosition := CurrentPosition
			for {
				CurrentPosition++
				if CurrentPosition >= TotalLetters {
					break
				}
				CurrentCharacter = Code[CurrentPosition]

				if !NumberPattern.MatchString(string(CurrentCharacter)) {
					break
				}
			}

			FinalNumber := Code[StartPosition:CurrentPosition]
			Number, err := strconv.Atoi(string(FinalNumber))
			if err != nil {
				fmt.Println("Lexer: Error converting string to number.")
			}

			TokenList = append(TokenList, Token{
				Type:  "NUMBER",
				Value: Number,
			})
		} else {
			if CurrentPosition+2 < TotalLetters {
				if char, ok := SpecialCharacters[string(Code[CurrentPosition:CurrentPosition+2])]; ok {
					TokenList = append(TokenList, Token{
						Type:  "CHARACTER",
						Value: char,
					})
					CurrentPosition += 2
					continue
				}
			}
			if char, ok := SpecialCharacters[string(Code[CurrentPosition:CurrentPosition+1])]; ok {
				TokenList = append(TokenList, Token{
					Type:  "CHARACTER",
					Value: char,
				})
				CurrentPosition++
				continue
			}

			CurrentPosition++
		}

	}

	return TokenList
}
