/**
administrator: Li yuan
datetime     : 2018/7/4 19:15
*/
/*package main
import "fmt"
func main() {
	fmt.Println("开始演示defer执行顺序")
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")  改些之前内容，也适当的增补遗留的内容。现在拆分成六章了。基于市面上的最新教材要写。
	fmt.Println("演示defer执行顺序结束")
}

*/
/**
administrator: Li yuan
datetime     : 2018/7/5 19:07
*/
/*package main
import (
	"errors"
	"fmt"
)
var notZeroErr = errors.New("除数不能为零")
func div(a,b int) (int,error){
	if b == 0{
		return 0,notZeroErr
	}
	return a/b,nil
}
func main() {
	result, e := div(10, 0)
	fmt.Println(result,e)
}*/
/**
administrator: Li yuan
datetime     : 2018/7/6 19:08
*/
package main
import "fmt"
func main() {
	fmt.Println("这是一个输出语句")
	panic("程序被终止，结束")
}