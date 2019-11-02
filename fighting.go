package main
import ("fmt"
		"math")

		var prime =make(chan int,100)
		var wait =make(chan bool,100)

		func work(i int){
			for j:=2.0;j<=math.Sqrt(float64(i));j++{
				if i%int(j) ==0{
					wait <-false
					return
				}
			}
			wait <-true
			prime<-i
		}
		func main(){
			for i:=1;i<10000;i++{
				go work(i)
			}
			count:=0
			for i:=1;i<10000;i++{
				ok:=<-wait
				if ok{
					count++
				}
			}
	for i:=0;i<count;i++{
		fmt.Println(<-prime)
	}
}