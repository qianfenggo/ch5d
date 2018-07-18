/*package main
import (
	"encoding/json"
	"fmt"
)
type Books struct {
	Title   string `json:"title"`
	IsOnSale  bool `json:"is_on_sale,omitempty"`
	Price     float64 `json:"price,"`
	Book_id int `json:"book_id"`
}
func main() {
	book := &Books{}
	book.Title = "Go语言快乐入门"
	book.IsOnSale  = false
	book.Price = 9.9
	book.Book_id=1
	data, _ := json.Marshal(book)
	fmt.Println(string(data))
}*/
/*package main
import (
	"encoding/json"
	"fmt"
)
type Books struct {
	Title   string
	IsOnSale  bool
	Price     float64
	Book_id int
}
func main() {
	book := &Books{}
	book.Title = "Go语言快乐入门"
	book.IsOnSale  = true
	book.Price = 9.9
	book.Book_id=1
	data, _ := json.Marshal(book)
	fmt.Println(string(data))
}*/
package main
import (
	"encoding/json"
	"fmt"
)
type Books struct {
	Title   string
	IsOnSale  bool
	Price     float64
	Book_id int
}
func main() {
	var data = `{"Title":"Go语言快乐入门","IsOnSale":true,
		"Price":9.9,"Book_id":1}`
	p := &Books{}
	err := json.Unmarshal([]byte(data), p)
	fmt.Println(err)
	fmt.Println(*p)
}