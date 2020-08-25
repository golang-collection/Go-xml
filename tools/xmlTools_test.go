package tools

import "testing"

/**
* @Author: super
* @Date: 2020-08-25 14:18
* @Description:
**/

func TestParse(t *testing.T){
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	Parse(input)
}