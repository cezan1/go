package model

import (
	"fmt"
	"testing"
)

func TestNewConfigWithFile(t *testing.T) {
	c, _ := NewConfigWithFile("D:\\MyGo\\cmdedu\\etc\\" + "model.json")
	fmt.Println("base_path: ", c)

}
