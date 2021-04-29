package main
import "fmt"
import "strconv"
import "os"
import "math"

func main() {
	var args = os.Args[1:]

	a := [10] string {"Zero","One", "Two", "Three", "Four", "Five","Six","Seven","Eight","Nine"}

	var err error
    val := make([]int, len(args))
    for i := 0; i < len(args); i++ {
		t := float64(len(args[i]))
    	if val[i], err = strconv.Atoi(args[i]); err != nil {
        panic(err)
     }
	   no := make([]float64,len(args[i]))
	   for j:=0; j<len(args[i]) ; j++ {
		   no[j] = (float64(val[i]))/math.Pow(10,t-1)
           no[j] = math.Mod(no[j] , 10)
		   fmt.Print(a[int(no[j])])
           t = t-1
	 }
     if i == len(args)-1 {
         continue
     } else {
        fmt.Print(",")
     }     
   }	
}