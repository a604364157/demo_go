//错误和异常
package main

import "errors"

//1:panic()和recover()

/*panic：
1、内建函数
2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行，这里的defer 有点类似 try-catch-finally 中的 finally
4、直到goroutine整个退出，并报告错误

recover：
1、内建函数
2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
3、一般的调用建议
a). 在defer函数中，通过recever来终止一个gojroutine的panicking过程，从而恢复正常代码的执行
b). 可以获取通过panic传递的error

简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。*/

//如何正确的使用错误和异常(避免各种 if err != nil 的出现)
//1:失败的原因只有一个时，不使用error
//(这个我们应该返回bool)
func T07Check1(i int) error {
	if i == 1 {
		return errors.New("这是一个测试错误")
	}
	return nil
}

//2:没有失败时，不使用error
//这里应该不返回任何值
func T07Check2() error {
	//这里某些逻辑
	return nil
}

//3:error应放在返回值列表的最后
//4:错误值要统一定义(自定义错误常量,统一使用)
//5:错误逐层传递时，层层都加日志(方便错误定位)
//6:错误处理使用defer
//7:当尝试几次可以避免失败时，不要立即返回错误
//8:当上层函数不关心错误时，建议不返回error
//9:当发生错误时，不忽略有用的返回值

func func35() {

}

func main() {

}
