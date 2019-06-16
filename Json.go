package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
	id     int
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B)) //这个数据结构将采用json默认的编码方式，成员属性的name当作key

	res2D := Response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		id:     2013} //id部分的内容不会被打印出来，因为在定义数据结构的时候没有写其编码方式
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	//上面的部分都是go数据编码成json包数据，下面展示json数据解码成go数据
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) //手动创建一个json结构的数据

	var dat map[string]interface{} //dat是用来存放一个json包解码后的数据的变量，[string]是对应的json的key，interface是解码的结构，因为go没有泛型所以用interface表示任意一种数据结构

	if err := json.Unmarshal(byt, &dat); err != nil { //Unmarshal就是json解码函数
		panic(err) //解码失败的话panic
	}
	fmt.Println(dat) //打印解码的结果

	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string) //到这里dat-value的类型还是interface，需要类型转换
	fmt.Println(str1)

	//上面是将json解码得到的结果用interface泛型保存起来，除此之外还可以用自定义类型Response保存，这样就消除了类型转换断言
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := &Response2{} //这个地方换成Response1好像也可以，虽然Response1里没有json编码声明，而且json包第一个key是page，而Response里是Page，但是还是可以对上（难道是因为不分大小写？）
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
