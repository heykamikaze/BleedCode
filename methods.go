package main
import "fmt"

type celsius  float64

func (c celsius) fahrenheit() fahrenheit { //принимает Цельсии переводит в Фаренгейты
      return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// func (c celsius) kelvin() kelvin { //принимает Цельсии переовдит в Кельвины
//       return kelvin(c + 273.15)
// }

type fahrenheit float64

func (f fahrenheit) celsius() celsius {//принимает Фарингейты переводит в Цельсии
      return celsius((f - 32.0)*5.0 / 9.0)
}

// func (f fahrenheit) kelvin() kelvin { //принимает Фаренгейты, вызывает цельсии переводящий в кельвины
//     return f.celsius().kelvin()
// }

// type kelvin float64
//
// func (k kelvin) celsius() celsius {
//     return celsius(k - 273.15)
// }
//
// func (k kelvin) fahrenheit() fahrenheit{ //принимает кельвины, вызывает цельсий переводящий в фарингейты
//     return k.celsius().fahrenheit()
// }

const(
  line          = "======================"
  rowFormat     = "| %8v | %8v |\n"
  numberFormat  = "%.1f"
)

type  getRowFn  func(row int) (string, string) //

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn)  {
    fmt.Println(line)
    fmt.Printf(" %8v  %8v \n", hdr1, hdr2)
    fmt.Println(line)
    for row := 0; row < rows; row++{
      cell1, cell2 := getRow(row)
      fmt.Printf(rowFormat, cell1, cell2)
    }
    fmt.Println(line)
}

func ctof(row int) (string, string)  {
  c := celsius(row*5 - 40)
  f := c.fahrenheit()
  cell1 := fmt.Sprintf(numberFormat, c)
  cell2 := fmt.Sprintf(numberFormat, f)
  return cell1, cell2
}

func ftoc(row int) (string, string)  {
  f := fahrenheit(row*5 - 40)
  c := f.celsius()
  // cell1 :=fmt.Sprintf("| %8v | %8v |\n", f)
  cell1 := fmt.Sprintf(numberFormat, f)
  cell2 := fmt.Sprintf(numberFormat, c)
  return cell1, cell2
}

func main()  {
  // var k kelvin = 294.0
  // var f fahrenheit = 19.9
  // c := k.celsius()
  // fmt.Println(k, " K is", c, " C")
  // c := k.fahrenheit()
  // fmt.Println(k, " K is", c, " F")
  // c := f.celsius()
  // fmt.Println(f.celsius())
  drawTable("C", "F", 5, ctof)
  fmt.Println()
  drawTable("F", "C", 5, ftoc)
}
