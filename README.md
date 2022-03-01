# gotips
A buntch of Go tips.

## Tips

- Go 中仅有值传递。
- 善用 `switch` 而不是多个 `if`。
- 使用 `chan struct{}` 来传递信号。
- 防止结构体字段使用纯值方式初始化，可在结构体中添加 `_ struct{}` 字段。
- 允许对值为 `nil` 的 `slice` 添加元素，但对值为 `nil` 的 `map` 添加元素时，会造成运行时 `panic`。
- 检查 `map` 中的 `key` 是否存在，可以使用返回的第二个参数 `ok` 来判断。
- Go 中的可变参数作为函数参数时，必须放在最后一位。
- 想要修改 `string` 的值，需要借助 `[]byte`，修改后再转为 `string`。
- 两个 `nil` 是不相等的，故无法通过 `!=` 进行比较。
- 对变量加锁后再进行复制，会将锁的状态一同复制。
- 在单独的 `for` 循环中，`break` 可以跳出循环。但在 `for select` 中，`break` 可以跳出 `select` 块，但不会跳出 `for` 循环。如需跳出 `for` 循环，可以配合 `goto` 使用 `label` 解决。
- `map` 是线程不安全的，可以通过 `sync.RWMutex` 加锁或者使用线程安全的 `sync.Map` 来解决。
- `return` 会先于 `defer` 返回，且 `return` 不是原子操作。因此，`return` 语句会分为赋值和返回两个部分。
- 在同一个作用域中，多次声明同一个变量名，后声明的变量仅在当前作用域生效。
- 使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响原来的值。如果想要在调用函数中修改原来的值，可以使用指针接收者定义的方法。
- `for range` 中获取到的值是元素的副本。如用 `for range` 遍历切片时使用的是切片的副本，但不会复制底层数组。[参考](https://go.dev/play/p/-YYOfIFYF2v)
- 在函数调用里修改返回的切片，将会影响到原切片。通常我们新建一个切片，然后将修改后的结果复制到该新切片，而不是改变旧有切片。
- Go 语言中不存在引用变量，每个变量都占用一个唯一的内存位置。
- Go 中的预定义标识符（如 `string`、`len` 等）是可以作为变量使用的，但关键字不行，如 `default`。
- 当使用 `type` 声明一个新类型时，它不会继承原有类型的方法集。
- `init()` 函数不能被其他函数调用，包括 `main()` 函数，它总是第一个执行。
- 非命名类型（`unamed type`，如 `struct{}`、`[]string`、`interface{}`、`map[string]bool` 等）不能作为方法的接收者。[参考](https://go.dev/play/p/Xbdnni_JasU)
- 不同类型的值是不能相互赋值的，即使底层类型一样。对于底层类型相同的变量可以相互赋值的一个重要条件是，至少有一个变量不是有名类型（`named type`，如内置类型和用 `type` 声明的类型）。
- 在接口中，只有静态值和动态类型都为 `nil` 的情况下，接口值才为 `nil`。
- 在拷贝切片时，`copy(dst, src)` 函数返回 `len(dst)`、`len(src)` 之间的最小值。如果想要将 src 完全拷贝至 dst，必须给 dst 分配足够的内存空间。[参考](https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA)
- `panic` 引发异常以后，程序停止执行，然后调用延迟 `defer` 函数。[参考](https://go.dev/play/p/pTe1wUxn73P)
- `nil` 切片`var s []T`和空切片`s := make([]T, 0)`是不同的切片。[参考](https://stackoverflow.com/questions/59349879/whats-the-difference-between-int-and-int-in-go)
- `byte` 是 `uint8` 的别名，`rune` 是 `int32` 的别名。其中，`byte` 表示一个字节，`rune` 表示四个字节。
- 双引号包含着的内容用来表示字符串，其实质是 `byte` 类型的数组；单引号包含着的内容则用来表示字符，其实质是 `rune` 类型的数组。
- 在一个常量声明代码块中，如果 `iota` 没有出现在第一行，则常量的初始值就是非零值。[参考](https://go.dev/play/p/C1jHFpACuT7)
- 只要有一个指针指向一个引用的变量，那么这个变量就不会被释放。因此，在 Go 语言中返回函数参数或临时变量是安全的。
- 如果一个类型实现了 `String()` 方法，那么在使用 `Printf()`、`Print()`、`Println()`、`Sprintf()` 等格式化输出时，会自动使用 `String()` 方法（[参考](https://go.dev/play/p/9mYm3JTJEOJ) （[参考](https://stackoverflow.com/questions/60765066/golang-what-does-the-following-do?rq=1)））。因此，再次调用 `String()` 方法将导致递归调用。[参考 1](https://go.dev/play/p/8jOYDn0m2WY) [参考 2](https://go.dev/play/p/8jOYDn0m2WY)
- 可变函数是指针传递。[参考](https://go.dev/play/p/apu9JTmorrp)
- `defer` 语句通常应该放到 `if err != nil` 后面。[参考](https://go.dev/play/p/H2nLDO9Q3za)
- `for {}` 循环会独占 CPU 资源导致其他 Goroutine 饿死，可以通过阻塞的方式避免 CPU 占用，如使用 `select {}`。[参考](https://go.dev/play/p/VgVSO6Edb_6)
- 如果想实现函数或者方法的链式调用，则返回该函数或者方法的指针值即可。
- `defer` 函数的参数（包括接收者）是在 `defer` 语句出现的位置做计算的，而不是在函数执行的时候计算的。[参考](https://go.dev/play/p/r_wvQjHDO8Q)
- 类型 `T` 的方法集是包含值类型接收者 `T` 的一组方法；类型 `*T` 的方法集是包含值类型接收者 `T` 和指针类型接收者 `*T` 的一组方法。
- `goto` 无法跳转到其他函数或者内层代码。[参考](https://go.dev/play/p/miz2pGthALx)
- `recover()` 必须在 `defer()` 函数中直接调用才会生效。
- `defer()` 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，等到真正执行 `defer()` 时取出。[参考](https://go.dev/play/p/9n-JMGhicmQ)
- 多重赋值分为两个步骤：先分别计算等号左边的表达式和等号右边的表达式，然后再将右边的值赋值给左边。[参考](https://go.dev/play/p/bhPE4fqIr-D)
- 对空指针解引用会造成 `panic`。[参考](https://go.dev/play/p/MHb9socpfdk)
- 将 `Mutex` 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。也可以通过嵌入 `*Mutex` 来避免复制的问题，但需要初始化。[参考](https://go.dev/play/p/iL0qUgiiggH)
- 当目标方法的接收者是指针类型时，那么被复制的就是指针。[参考](https://go.dev/play/p/ttoONtuoHan)
- 当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象并与其绑定，以便在稍后执行时能隐式第传入接收者参数。[参考](https://go.dev/play/p/-n9RfLyVWVX)
- 递增运算符 `++` 和递减运算符 `--` 的优先级低于解引用运算符 `*` 和取址运算符 `&`，解引用运算符和取址运算符的优先级低于选择器 `.` 中的属性选择操作符。[参考](https://go.dev/play/p/-bGBSCUGbyG)
- 截取符号 `sl[i:j]`，如果 j 省略，默认是截取到原切片或者数组的长度，若 `j > len(sl)`，则 `panic`。[参考](https://go.dev/play/p/_CCiGbYAkZA)
- 从一个基础切片派生出的子切片的长度可能大于基础切片的长度。[参考](https://go.dev/play/p/X5l5_6rBTQx)
- 不能使用多级指针调用方法。[参考](https://go.dev/play/p/L6jl8KpI2D7)
- 不可寻址的结构体不能调用带结构体指针接收者的方法。[参考](https://go.dev/play/p/2ZWdotuYAy_2)
- 在方法中，指针类型的接收者必须是合法指针（包括 `nil`），或者能够获取到实例地址的表达式。[参考](https://go.dev/play/p/kW1kZeZMWBF)
- 函数只能与 `nil` 比较。[参考](https://go.dev/play/p/_vtECkR00ZZ)
- `map[key]struct` 中 `struct` 是不可寻址的，因此无法直接赋值。若想知道 `struct` 中的地址，可以考虑使用临时变量（[参考](https://go.dev/play/p/l5dIxdOHpn5)）或者修改数据结构（[参考](https://go.dev/play/p/0-0Xr_ytXUt)）。[参考](https://go.dev/play/p/tNum3CIxW7l) 
- `:=` 操作符不能用于结构体字段赋值。[参考](https://go.dev/play/p/6v8mUJzsP0M)
- `&^` 为按位置零。表达式 `z = x &^ y` 表示为如果 y 中的 bit 位为 1，则 z 对应的 bit 位为 0，否则 z 对应 bit 位等于 x 中相应的 bit 位的值。
- `|` 为或操作符。表达式 `z = x | y` 表示为如果 y 中的 bit 位为 1，则 z 对应 bit 位为 1，否则 z 对应 bit 位等于 x 中相应的 bit 位的值，与 `&^` 操作相反。
- Go 语言中不支持 `++i` 和 `--i` 操作。另外，`i++` 和 `i--` 在 Go 语言中是语句，不是表达式，因此不能赋值给其他变量。
- 同变量不同，Go 语言中常量未使用是能编译通过的，但不可以对常量取地址。
- 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同。[参考](https://go.dev/play/p/mkBzT7ttBfv)
- Go 的字符串类型是不能赋值为 `nil` 的，也不能跟 `nil` 比较。[参考](https://go.dev/play/p/7B7wbztcwlr)
- 不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的。[参考](https://go.dev/play/p/I8hZZ72QA8C)
- 用字面量初始化数组、`slice` 和 `map` 时，最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错。[参考](https://go.dev/play/p/D9v3aFCTRL0)
- `cap()` 函数适用于数组、数组指针、`slice` 和 `channel`，不适用于 `map`，可以使用 `len()` 返回 `map` 的元素个数。当使用 `make` 创建 `map` 变量时指定第二个参数会被忽略。
- `nil` 用于表示 `interface`、`func`、`map`、`slice` 和 `channel` 的零值。如果不指定变量的类型，编译器将无法得出变量的具体类型，导致编译错误。[参考 1](https://go.dev/play/p/3M-DfjOP-Vr) [参考 2](https://go.dev/play/p/Gp0wBvdndWs)
- 不能使用短变量声明设置结构体字段值。[参考](https://go.dev/play/p/Ui-VB7OVGOl)
- 不能对 `nil` 的 `map` 直接赋值，需要使用 `make()` 初始化。但可以使用 `append()` 函数对为 `nil` 的 `slice` 增加元素。[参考 1](https://go.dev/play/p/L_cYqV84DrN) [参考 2](https://go.dev/play/p/kfjcbO74u4Z)
- 如果有未使用的变量，代码将编译失败，但可以有未使用的全局变量。另外，函数的参数未使用也是可以的。当然，如无必要，可以注释掉或者移除未使用的变量。[参考](https://go.dev/play/p/xYxO9jOJNYg)
- Go 语言中，大括号不能放在单独的一行，否则会编译错误。
- 有方向的 `channel` 不可以被关闭。[参考](https://go.dev/play/p/sh0pUfwN5fy)
- `map` 必须初始化才能使用。
- 指针不支持索引。[参考](https://go.dev/play/p/5tsektd4No-)
- 当且仅当动态值和动态类型都为 `nil` 时，接口类型值才为 `nil`。[参考](https://go.dev/play/p/K-iX86rToeG)
- 当使用 `for range` 遍历切片时，若只有一个参数，如 `for v := range x`，则 v 表示索引；若有两个参数，第二个参数才表示具体的值。[参考](https://go.dev/play/p/Zdu4XlgJLUI)
- 读、写一个 `nil channel` 会造成永久阻塞；向已经关闭的 `channel` 发送数据，会造成 `panic`；从一个已经关闭的 `channel` 接收数据，如果缓冲区为空，则返回一个零值，否则读取出对应的值；关闭一个已经关闭的 `channel` 会 `panic`。[参考](https://go.dev/play/p/h7NnRmXbtEA)
- 使用 `make` 初始化切片时，需要补充 `len` 参数，`cap` 参数可选，否则无法编译。[参考](https://go.dev/play/p/xrKuOwRCQiU)
- `select` 会随机选择一个可用通道做收发操作，而不是按顺序进行。[参考](https://go.dev/play/p/5tsektd4No-)
- `for range` 使用短变量声明 `:=` 的形式迭代变量时，变量 i、v 在每次循环中都会被重用，而不是重新声明。[参考 1](https://go.dev/play/p/pPk92ad178b) [参考 2](https://go.dev/play/p/UghIn1HZ-l1)
- 当 `range` 表达式发生复制时，副本的指针依旧指向原底层数组，所以对切片的修改都会反应到底层数组上。[参考](https://go.dev/play/p/mDjwzkSxTt1)
- 循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。[参考](https://go.dev/play/p/I67nM8JFsPZ)
- Go 中不同类型时不能比较的，切片也是不能进行比较的。
- 基于类型创建的方法必须定义在同一个包内，或者定一个该已有类型的新类型（[参考](https://go.dev/play/p/u5fYzh-7b72)）。[参考](https://go.dev/play/p/gWdbC0S_Z-d)
- `return` 之后的 `defer` 是不能注册的。[参考](https://go.dev/play/p/IPAp4769FZc)
- `var a []T`声明的是 `nil` 切片，不会分配内存。`a := []T{}`声明的是 `len` 和 `cap` 都为 0 的空切片，会分配内存。
- 函数参数为 `interface{}` 时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也是可以的，而不是使用 `*interface{}`。[参考](https://go.dev/play/p/9Ekk_CUlYdI)
- 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。



## 链接中的代码输出什么？有什么问题，是否需要修改？

- [https://go.dev/play/p/mEsneHeNTxR](https://go.dev/play/p/mEsneHeNTxR)

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

- [https://go.dev/play/p/ttoONtuoHan](https://go.dev/play/p/ttoONtuoHan)

- [https://go.dev/play/p/aHKokiGjLWe](https://go.dev/play/p/aHKokiGjLWe)

- [https://go.dev/play/p/9dRDxU5Rkrl](https://go.dev/play/p/9dRDxU5Rkrl)

- [https://go.dev/play/p/b7t_wUoSrok](https://go.dev/play/p/b7t_wUoSrok)

- [https://go.dev/play/p/qaLDSDS2Udn](https://go.dev/play/p/qaLDSDS2Udn)

- [https://go.dev/play/p/L6jl8KpI2D7](https://go.dev/play/p/L6jl8KpI2D7)

- [https://go.dev/play/p/Zi0Y6NijGHZ](https://go.dev/play/p/Zi0Y6NijGHZ)

- [https://go.dev/play/p/tBmTMxnNiLY](https://go.dev/play/p/tBmTMxnNiLY)

- [https://go.dev/play/p/SNMLaylFSR4](https://go.dev/play/p/SNMLaylFSR4)

- [https://go.dev/play/p/NWhEcVpPCun](https://go.dev/play/p/NWhEcVpPCun)

- [https://go.dev/play/p/EPi1Y0YTNIn](https://go.dev/play/p/EPi1Y0YTNIn)

- [https://go.dev/play/p/5w635PFllaX](https://go.dev/play/p/5w635PFllaX)

- [https://go.dev/play/p/GO_HnTjv8zh](https://go.dev/play/p/GO_HnTjv8zh)

- [https://go.dev/play/p/Ui-VB7OVGOl](https://go.dev/play/p/Ui-VB7OVGOl)

- [https://go.dev/play/p/kBEEWF7IQ3U](https://go.dev/play/p/kBEEWF7IQ3U)

- [https://go.dev/play/p/djK-lxKDrmJ](https://go.dev/play/p/djK-lxKDrmJ)

- [https://go.dev/play/p/ZZqUsVaZsK5](https://go.dev/play/p/ZZqUsVaZsK5)

- [https://go.dev/play/p/pAlMBeSuZJA](https://go.dev/play/p/pAlMBeSuZJA)

- [https://go.dev/play/p/lunSEku2gYj](https://go.dev/play/p/lunSEku2gYj)

- [https://go.dev/play/p/QJsQbZrtY40](https://go.dev/play/p/QJsQbZrtY40)

- [https://go.dev/play/p/NnGkzIPWDeD](https://go.dev/play/p/NnGkzIPWDeD)

- [https://go.dev/play/p/mDjwzkSxTt1](https://go.dev/play/p/mDjwzkSxTt1), [https://go.dev/play/p/bk8DSR66qhY](https://go.dev/play/p/bk8DSR66qhY)

- [https://go.dev/play/p/Hd63R-qvHu3](https://go.dev/play/p/Hd63R-qvHu3)

- [https://go.dev/play/p/QV7VndMJMnD](https://go.dev/play/p/QV7VndMJMnD)

- [https://go.dev/play/p/9mYm3JTJEOJ](https://go.dev/play/p/9mYm3JTJEOJ)

- [https://go.dev/play/p/nUQa15qss8h](https://go.dev/play/p/nUQa15qss8h) [参考](https://studygolang.com/articles/2192)

- [https://go.dev/play/p/sSYx1mfw39Q](https://go.dev/play/p/sSYx1mfw39Q)

- [https://go.dev/play/p/lITM1Im27Ku](https://go.dev/play/p/lITM1Im27Ku)

- [https://go.dev/play/p/BotqJKnBsAO](https://go.dev/play/p/BotqJKnBsAO)

- [https://go.dev/play/p/89_i3Ok5SNq](https://go.dev/play/p/89_i3Ok5SNq)

- [https://go.dev/play/p/cyDqT8bHWLM](https://go.dev/play/p/cyDqT8bHWLM)

- [https://go.dev/play/p/IPAp4769FZc](https://go.dev/play/p/IPAp4769FZc)

- [https://go.dev/play/p/9Ekk_CUlYdI](https://go.dev/play/p/9Ekk_CUlYdI)

- []()

- []()

- []()

- []()

- []()

- []()

- []()

- []()

- []()

- []()

  





## Reference

- [https://golang.design/go-questions/](https://golang.design/go-questions/)
- [https://gfw.go101.org/article/101.html](https://gfw.go101.org/article/101.html)
- [https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw](https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw)
- [https://www.practical-go-lessons.com](https://www.practical-go-lessons.com)
- [https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/](https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/)



## Credit

You All Guys!
