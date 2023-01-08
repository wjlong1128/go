package main

import "fmt"

/*
当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，
换句话说：一种可阅读性和打印性的输出。如果类型定义了 String() 方法，
它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。

// 说人话就是Java的ToString方法
*/

type User struct {
	name string
}

func (u *User) String() string {
	return "我就是 ~~~" + u.name + "~~~!!!"
}


func main() {
	u := &User{"张三"}
	fmt.Printf("u: %v\n", u) // u: 我就是 ~~~张三~~~!!!
}

/*
当你广泛使用一个自定义类型时，最好为它定义 String() 方法。
从上面的例子也可以看到，格式化描述符 %T 会给出类型的完全规格，%#v 会给出实例的完整输出，
包括它的字段（在程序自动生成 Go 代码时也很有用）

不要在 String() 方法里面调用涉及 String() 方法的方法，它会导致意料之外的错误
*/