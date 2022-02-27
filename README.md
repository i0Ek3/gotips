# gofucks
A buntch of Go tips.

## Tips

- Go 中仅有值传递。
- 善用 `switch` 而不是多个 `if`。
- 使用 `chan struct{}` 来传递信号。
- 防止结构体字段用纯值方式初始化，可在结构体中添加 `_ struct{}` 字段。
- 允许对值为 `nil` 的 `slice` 添加元素，但对值为 `nil` 的 `map` 添加元素时，会造成运行时 `panic`。
- 检查 `map` 中的 `key` 是否存在，可以直接使用返回的第二个参数 `ok` 来判断。
- Go 中的可变参数作为函数参数时，必须放在最后一位。
- 想要修改 `string` 的值，需要借助 `[]byte`，修改后再转换为 `string`。
- 两个 `nil` 是不相等的，故无法通过 `!=` 进行比较。
- 对变量加锁后再进行复制，会将锁的状态一同复制。
- 在单独的 `for` 循环中，`break` 可以跳出循环。但在 `for select` 中，`break` 可以跳出 `select` 块，但不会跳出 `for` 循环。如需跳出 `for` 循环，可以配合 `goto` 使用 `label` 解决。
- `map` 是线程不安全的，可以通过 `sync.RWMutex` 加锁 或者使用线程安全的 `sync.Map` 来解决。
- `return` 会先于 `defer` 返回，且 `return` 不是原子操作，会分为赋值和返回两个部分。
- 在同一个作用域中，多次声明同一个变量名，后声明的变量仅在当前作用域产生效果。
- 使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响原来的值。如果想要在调用函数中修改原来的值，可以使用指针接收者定义的方法。
- `for range` 中获取到的值是原本元素的副本。如 `for-range` 切片时使用的是切片的副本，但不会复制底层数组。[参考](https://go.dev/play/p/-YYOfIFYF2v)
- 在函数调用里面修改返回的切片，将会影响到原切片。通常我们新建一个切片，然后将修改后的结果复制到该新切片，而不是改变旧有切片。
- Go 语言中不存在引用变量，每个变量都占用一个唯一的内存位置。
- Go 中的预定义标识符（如 `string`、`len` 等）是可以作为变量使用的，但关键字不行，如 `default`。
- 当使用 `type` 声明一个新类型时，它不会继承原有类型的方法集。
- `init()` 函数不能被其他函数调用，包括 `main()` 函数，它总是第一个执行。
- 非命名类型（unamed type，如 `struct{}`、`[]string`、`interface{}`、`map[string]bool` 等）不能作为方法的接收者。[参考](https://go.dev/play/p/Xbdnni_JasU)
- 不同类型的值是不能相互赋值的，即使底层类型一样。对于底层类型相同的变量可以相互赋值的一个重要条件是，至少有一个变量不是有名类型（named type，如内置类型和用 `type` 声明的类型）。
- 在接口中，只有静态值和动态类型都为 `nil` 的情况下，接口值才为 `nil`。
- 在拷贝切片时，`copy(dst, src)` 函数返回 `len(dst)`、`len(src)` 之间的最小值。如果想要将 src 完全拷贝至 dst，必须给 dst 分配足够的内存空间。[参考](https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA)
- `panic` 引发异常以后，程序停止执行，然后调用延迟 defer 函数。[参考](https://go.dev/play/p/pTe1wUxn73P)
- `nil` 切片`var s []T`和空切片`s := make([]T, 0)`是不同的切片。[参考](https://stackoverflow.com/questions/59349879/whats-the-difference-between-int-and-int-in-go)
- `byte` 是 `uint8` 的别名，`rune` 是 `int32` 的别名。其中，`byte` 表示一个字节，`rune` 表示四个字节。
- 双引号包含着的内容用来表示字符串，其实质是 `byte` 类型的数组；单引号包含着的内容则用来表示字符，其实质是 `rune` 类型的数组。
- 在一个常量声明代码块中，如果 `iota` 没有出现在第一行，则常量的初始值就是非 0 值。[参考](https://go.dev/play/p/C1jHFpACuT7)
- 只要有一个指针指向一个引用的变量，那么这个变量就不会被释放，因此在 Go 语言中返回函数参数或临时变量是安全的。
- 如果类型定义了 `String()` 方法，那么在使用 `Printf()`、`Print()`、`Println()`、`Sprintf()` 等格式化输出时，会自动使用 `String()` 方法。[参考](https://go.dev/play/p/8jOYDn0m2WY)
- 可变函数是指针传递。[参考](https://go.dev/play/p/apu9JTmorrp)
- `defer` 语句通常应该放到 `if err != nil` 后面。[参考](https://go.dev/play/p/H2nLDO9Q3za)
- `for {}` 循环会独占 CPU 资源导致其他 Goroutine 饿死，可以通过阻塞的方式避免 CPU 占用，如使用 `select {}`。[参考](https://go.dev/play/p/VgVSO6Edb_6)
- 如果想实现函数或者方法的链式调用，则返回该函数或者方法的指针值即可。
- `defer` 函数的参数（包括接收者）是在 `defer` 语句出现的位置做计算的，而不是在函数执行的时候计算的。[参考](https://go.dev/play/p/r_wvQjHDO8Q)
- 类型 `T` 的方法集是包含值类型接收者 `T` 的一组方法；类型 `*T` 的方法集是包含值类型接收者 `T` 和指针类型接收者 `*T` 的一组方法。
- `goto` 无法跳转到其他函数或者内层代码。[参考](https://go.dev/play/p/miz2pGthALx)
- `recover()` 必须在 `defer()` 函数中直接调用才有效。
- `defer()` 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 `defer()` 时取出。[参考](https://go.dev/play/p/9n-JMGhicmQ)
- 多重赋值分为两个步骤：计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；赋值。[参考](https://go.dev/play/p/bhPE4fqIr-D)



## 链接中的代码输出什么？是否需要修改？

- [https://go.dev/play/p/R4OP-836tDo](https://go.dev/play/p/R4OP-836tDo)
- [https://go.dev/play/p/1glmROGjWK5](https://go.dev/play/p/1glmROGjWK5)
- [https://go.dev/play/p/wJ4nY_UIAha](https://go.dev/play/p/wJ4nY_UIAha)
- [https://go.dev/play/p/zEIuT1d18k0](https://go.dev/play/p/zEIuT1d18k0)
- [https://go.dev/play/p/VIicWUM7ae1](https://go.dev/play/p/VIicWUM7ae1)
- [https://go.dev/play/p/BrTR1ztnNQc](https://go.dev/play/p/BrTR1ztnNQc)
- [https://go.dev/play/p/8jOYDn0m2WY](https://go.dev/play/p/8jOYDn0m2WY)
- [https://go.dev/play/p/XZzFXoO6gOW](https://go.dev/play/p/XZzFXoO6gOW) [参考](https://gfw.go101.org/article/line-break-rules.html)
- [https://go.dev/play/p/apu9JTmorrp](https://go.dev/play/p/apu9JTmorrp)
- [https://go.dev/play/p/RvfC7SdNYrQ](https://go.dev/play/p/RvfC7SdNYrQ)
- [https://go.dev/play/p/pTe1wUxn73P](https://go.dev/play/p/pTe1wUxn73P)
- [https://go.dev/play/p/VgVSO6Edb_6](https://go.dev/play/p/VgVSO6Edb_6)
- [https://go.dev/play/p/EG2BaGmhOxb](https://go.dev/play/p/EG2BaGmhOxb)
- [https://go.dev/play/p/N2XeOCcF9-k](https://go.dev/play/p/N2XeOCcF9-k)
- [https://go.dev/play/p/eCIBYSVnAFa](https://go.dev/play/p/eCIBYSVnAFa)
- [https://go.dev/play/p/ZfbOyRF9V5C](https://go.dev/play/p/ZfbOyRF9V5C)
- [https://go.dev/play/p/r_wvQjHDO8Q](https://go.dev/play/p/r_wvQjHDO8Q)
- [https://go.dev/play/p/SI7ABAfCeQp](https://go.dev/play/p/SI7ABAfCeQp)
- [https://go.dev/play/p/nQYNKuLRYo9](https://go.dev/play/p/nQYNKuLRYo9)
- [https://go.dev/play/p/Y9EZtxL_FZq](https://go.dev/play/p/Y9EZtxL_FZq)
- [https://go.dev/play/p/T7VWzLlajuS](https://go.dev/play/p/T7VWzLlajuS)
- [https://go.dev/play/p/9n-JMGhicmQ](https://go.dev/play/p/9n-JMGhicmQ)
- [https://go.dev/play/p/bhPE4fqIr-D](https://go.dev/play/p/bhPE4fqIr-D)
- [https://go.dev/play/p/H1Su5YAtdNu](https://go.dev/play/p/H1Su5YAtdNu)
- []()
- []()
- []()





## Reference

- [https://golang.design/go-questions/](https://golang.design/go-questions/)
- [https://gfw.go101.org/article/101.html](https://gfw.go101.org/article/101.html)
- [https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw](https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw)
- [https://www.practical-go-lessons.com](https://www.practical-go-lessons.com)
- []()
- []()
- []()
- []()
- []()
- []()



## Credit

You All Guys!
