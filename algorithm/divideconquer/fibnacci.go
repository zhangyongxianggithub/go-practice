package divideconquer

/* 斐波那契数列： 递归*/
func fib(n int) int {
	// 终止条件
	if n == 1 || n == 2 {
		return n - 1
	}
	// 递归调用f(n)=f(n-1)+f(n-2)
	res := fib(n-1) + fib(n-2)
	// 返回结果
	return res
}
