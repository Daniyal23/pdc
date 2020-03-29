package main

import "fmt"

func main() {
   var check1 int
	 var check2 int
	 var check3 int
	 var check4 int
	 check1=0
	 check2=0
	 check3=0
	 check4=0
  var input int
	for i:=0;i<4;{

	str1 := "1 Print Covid Cases in Pakistan"
	str3 := "2 Print Covid Cases in South Korea"
	str4 := "3 Print Covid Cases in France"
	str5 := "4 Print My Message to CoronaVirus"
	str6 := "0 To Exit"
	fmt.Println(str1)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Scanln( &input)
	if input== 1{
		fmt.Println("1000")
		check1=1
	}
	if input==2 {
		fmt.Println("200")
		check2=1
	}
	if input==3 {
		fmt.Println("300")
		check3=1
	}
	if input==4 {
		fmt.Println("400")
		check4=1
	}
	if input==0 {
		i=5
	}
	if check1==1{
		if check2==1{
			if check3==1{
				if check4==1{
					i=5
				}
			}
		}
	}

	}
}
