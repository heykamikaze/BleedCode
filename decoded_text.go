package main
import ("fmt"
"strings"
 )

// func main()  {
//   cipherText := "CSOITEUIWUIZNSROCNKFD"
//   keyword := "GOLANG"
//   message := ""
//   keyIndex := 0
//
//   for i := 0; i < len(cipherText); i++ {
//     //A=0, B=1, ... 25
//     c := cipherText[i] - 'A'
//     k := keyword[keyIndex] - 'A'
//
//     //coded letter - key letter
//     c = (c-k+26)%26 + 'A'
//     message += string(c)
//     keyIndex++
//     keyIndex %= len(keyword)
//   }
//   fmt.Println(message)
// }
func main()  {
  message := "abc de fg"
  keyword := "hijklm"
  keyIndex := 0
  cipherText := ""

  message = strings.ToUpper(strings.Replace(message, " ", "", -1))
  keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

  for i := 0; i < len(message); i++ {
    var c = message[i]
    if c >= 'A' && c <= 'Z' {
      c -= 'A'
      k := keyword[keyIndex] - 'A'
      c = (c+k)%26 + 'A'
      keyIndex++
      keyIndex %= len(keyword)
    }
    cipherText += string(c)
  }
  fmt.Println(cipherText)
}
