package main

import "fmt"

type personInfo struct {
	id      string
	name    string
	Address string
}

func main(){
	searchName := "Snow"
	personDB := make(map[string]personInfo)
	personDB["Jack"] = personInfo{id:"002",name:"Jack Son"}
	personDB["Snow"] = personInfo{id:"005",name:"Jon Snow"}
	person, isExist :=  personDB[searchName]
	if isExist {
		fmt.Printf("Found person,: %v\n", person)
	} else {
		fmt.Printf("Person %v is not existing", searchName)
	}

	delete(personDB,"Jack") // delete函数不会返回任何信息
	delete(personDB,"A Non-existent Key") // delete一个不存在的key并不会发生任何事情，也不会报错
	//delete(nil,"Snow") // 但是delete函数的map参数不能是nil，会直接panic报错崩掉

	personDB2 := make(map[*personInfo]string)
	delete(personDB2,nil) // 这种key是nil并不会报错
}
