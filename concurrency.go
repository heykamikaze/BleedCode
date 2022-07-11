package main
import "fmt"

func download(n int, ch chan int) {
     ch<-n*(n+1)/2
}

func main() {
    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)
    defer close (ch1)
    defer close (ch2)
    defer close (ch3)
    var s1, s2, s3 int
    fmt.Scanln(&s1)
    fmt.Scanln(&s2)
    fmt.Scanln(&s3)

    go download(s1, ch1)
    go download(s2, ch2)
    go download(s3, ch3)
/*    var c, sum int
    for {
        select {
            case <-ch1:
               c++
               sum += <-ch1
            case <-ch2:
               c++
               sum += <-ch2
            case <-ch3:
               c++
	       sum += <-ch3
	 }
  	 if c == 3 {break}
 }*/
    fmt.Println(<-ch1 + <-ch2 + <-ch3)
}
