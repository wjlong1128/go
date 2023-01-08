package main

// 函数
func main(){
	// 空白符 空白符用来匹配一些不需要的值，然后丢弃掉
	a,_ := GetNames2()
	println(a)
}
// 按值传递（call by value） 按引用传递（call by reference
/*
Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量

如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，比如 &variable）传递给函数，这就是按引用传递

此时传递给函数的是一个指针。如果传递给函数的是一个指针，指针的值（一个地址）会被复制，但指针的值所指向的地址上的值不会被复制；我们可以通过这个指针的值来修改这个值所指向的地址上的值。（译者注：指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递。）

几乎在任何情况下，传递指针（一个 32 位或者 64 位的值）的消耗都比传递副本来得少

在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）
*/
func SetName(name *string){
	*name = "new Name"
}


// 命名的返回值（named return variables）
/*
	命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，当需要返回的时候，我们只需要一条简单的不带参数的 return 语句
*/

// 普通的return 函数
func GetNames() (string,string){
	return "a","b"
}

// 命名返回值
func GetNames2() (a string,b string){
	a = "a"
	b = "b"
	return // return a,b   [a,b]可省略直接写return
}

// 尽量使用命名返回值：会使代码更清晰、更简短，同时更加容易读懂


// 改变外部变量
/*
传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不再需要使用 return 返回
*/

func reply(num *int) {
	*num = 1
}
/*
这仅仅是个指导性的例子，当需要在函数内改变一个占用内存比较大的变量时，性能优势就更加明显了。然而，如果不小心使用的话，传递一个指针很容易引发一些不确定的事，所以，我们要十分小心那些可以改变外部变量的函数，在必要时，需要添加注释以便其他人能够更加清楚的知道函数里面到底发生了什么
*/