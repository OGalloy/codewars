package middle_character

//You are going to be given a word.
//Your job is to return the middle character of the word.
//If the word's length is odd, return the middle character.
//If the word's length is even, return the middle 2 characters.

//#Examples:
//Kata.getMiddle("test") should return "es"
//Kata.getMiddle("testing") should return "t"

func GetMiddle(str string) string {
	if len(str)%2 == 0 {
		lastIndexLeftPart := (len(str) / 2) - 1
		firstIndexRightPart := lastIndexLeftPart + 1
		return string(str[lastIndexLeftPart]) + string(str[firstIndexRightPart])
	}
	if len(str)%2 == 1 {
		middleIndex := len(str) / 2
		return string(str[middleIndex])
	}
	return str
}
