/*package main
import (
"fmt"
"io/ioutil"
)
func main() {
b, err := ioutil.ReadFile("qianfeng.text")
if err != nil {
fmt.Print(err)
}
fmt.Println(b)
str := string(b)
fmt.Println(str)
}*/
package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"errors"
)

type Mode int

const (
	aaa Mode =  iota
	bbb
	ccc
	ddd
)

func abc (mode Mode)  {
	fmt.Println(mode)

}

type File  struct {

}

func main() {
	d1 := []byte("hello\ngo\n")
	ioutil.WriteFile("qianfeng.text", d1, 0644)
	fmt.Println(os.ModeAppend)
	fmt.Println(os.ModeDir)
	fmt.Println(os.ModeDevice)
	fmt.Println(os.ModeNamedPipe)
	fmt.Println(os.ModeSetgid)
	fmt.Println(os.ModeSetuid)
	abc(ccc)
	errors.New()
}
