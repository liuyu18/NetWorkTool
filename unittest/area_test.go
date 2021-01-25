package unittest

// 必须导入testing模块，并且方法的接受者为(t *testing.T)
import (
	"fmt"
	"testing"

)

// 测试1: 测试名称是否符合预期
func TestSetSomething(t *testing.T) {
	box := Newbox()
	box.SetName("bgbiao")
	if box.GetName() == "bgbiao" {
		fmt.Println("the rectangular name's result is ok")
	}
}

// 测试2: 测试计算出来的体积是否符合预期
func TestGetSomething(t *testing.T) {
	box := Newbox()
	box.SetSize(3, 4, 5)
	if box.GetVolume() == 60 {
		fmt.Println("the rectangular volume's result is ok")
	}
}