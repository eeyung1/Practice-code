package main


func CaesarCipher(text string, shift int) string {

    shift = shift % 26

    if shift < 0 {
        shift += 26
    }


    answer := []rune{}
    
    for _, ch := range text {
        if ch >= 'a' && ch <= 'z' {
            // Shift lowercase
            value := ch - 'a'
            result := value + rune(shift)
            finalResult := result % 26
            char := finalResult + 'a'
            answer = append(answer, char)
            
        } else if ch >= 'A' && ch <= 'Z' {
            // Shift uppercase
            value := ch - 'A'
            result := value + rune(shift)
            finalResult := result % 26
            char := finalResult + 'A'
            answer = append(answer, char)
            
        } else {
            // Preserve everything else
            answer = append(answer, ch)
        }
    }
    
    return string(answer)
}