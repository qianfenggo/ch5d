package main
import "fmt"
func main() {
	class := new(string)
	*class = "扣丁学堂"
	number := new(int)
	*number = 666
	fmt.Println(*class,*number)
}
/*func main() {
	class:="扣丁学堂"
	number:=888
	fmt.Println("字符串变量class的内存地址是：",&class)
	fmt.Println("数字变量number的内存地址是：",&number)
}*/
/*func main() {
	os.Mkdir("qianfeng", 0777)
	os.MkdirAll("qianfeng1/more1/more2", 0777)
	err := os.Remove("qianfeng1")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("qianfeng1")
}*/
/*func main() {
	ch := make(chan int)
	//go read(ch)
	go write(ch)
	os.Mkdir()

}
func write(in chan int) {
	in<-10
}
func read(out chan int) {
	value := <-out
	fmt.Println(value)
}*/

/*func main() {
	for i := 0; i < 3; i++ {
		go add(i, i)
	}
	stime := time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Second * 5)
	fmt.Println(stime)
}
func add(a int, b int) {
	c := a + b
	fmt.Println(c)
}
*/
/*func main() {
	defer fmt.Println("这是defer1所要执行的内容")
	defer fmt.Println("这是defer2所要执行的内容")
	fmt.Println("panic函数前面的代码会执行")
	panic("触发panic函数")
	fmt.Println("panic函数后面的代码不会执行")
	defer fmt.Println("panic函数后面的defer语句不会执行")
}*/
/*func main() {
	var a [1000]int
	for i := 0; i < 1000; i++{
		go func(i int) {
			for {
				a[i]++
			}
		}(i)

	}
	fmt.Println(a)
}
*/
