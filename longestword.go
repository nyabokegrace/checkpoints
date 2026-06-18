package main

import ("unicode"
"fmt")

func LongestWord(s string)string {
	if s == "" {
		return ""
	}
	runes:=[]rune(s)
	longest, word := "", ""
	
	for i:=0; i<len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			word+=string(runes[i])
			for i+1 < len(runes) && unicode.IsLetter(runes[i+1]) {
				i++
				word+=string(runes[i])
			}
             if i+1 < len (runes) && (runes[i+1]) == '.' || runes[i+1] == ',' || runes[i+1] == '!' || runes[i+1] == '?' || runes[i+1] == ';' || runes[i+1] == ':'{
				word+=string(runes[i+1])
				i++
			 }
			 if len([]rune(word))> len([]rune(longest)) {
				longest=word
			 }
			}
			word=""
		}
		return longest
	}

	func main(){
		fmt.Println(LongestWord("Hello World!"))
	}