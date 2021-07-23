package main

import "fmt"

func main()  {
	fmt.Println("hello, go")

	// 数据类型
	/**
	1.布尔型：true false
	2.数字类型：
	  2.1 整型
	  2.2 浮点型 float32 float64
	  2.3 其它数字类型
	3.字符串类型：
	4.派生类型：
	  4.1 指针类型（Pointer）
	  4.2 数组类型
	  4.3 结构化类型(struct)
	  4.4 Channel 类型
	  4.5 函数类型
	  4.6 切片类型
	  4.7 接口类型（interface）
	  4.8 Map 类型
	 */

	// 变量的定义
	var a string = "hello"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = "go"
	fmt.Println(d)

	// 常量的定义
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)

	// 条件语句
	/* 定义局部变量 */
	var conditionA int = 100
	var conditionB int = 200

	/* 判断条件 */
	if conditionA == 100 {
		/* if 条件语句为 true 执行 */
		if conditionB == 200 {
			/* if 条件语句为 true 执行 */
			fmt.Printf("conditionA 的值为 100 ， conditionB 的值为 200\n" )
		}
	}
	fmt.Printf("conditionA 值为 : %d\n", conditionA )
	fmt.Printf("conditionB 值为 : %d\n", conditionB )

	result := max(conditionA, conditionB)
	fmt.Println(result)

	// for 循环
	/*
	for true  {
		fmt.Printf("这是无限循环。\n");
	}
	*/

	// 数组申明
	var balance1 [10] float32
	fmt.Println(balance1)

	// 数组初始化
	var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance2)

	mapFun()

}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if (num1 > num2) {
		result = num1
	} else {
		result = num2
	}
	return result
}

func mapFun() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
}