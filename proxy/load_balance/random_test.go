package load_balance

import (
	"fmt"
	"testing"
)

func TestLoadBalance(t *testing.T)  {
	rb := &RandomBalance{}

	rb.Add("http://172.0.0.1:2001")
	rb.Add("http://172.0.0.1:2002")
	rb.Add("http://172.0.0.1:2003")
	rb.Add("http://172.0.0.1:2004")
	rb.Add("http://172.0.0.1:2005")

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}