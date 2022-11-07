# gotips

A bunch tips for Go.

## Content

- [1 nil](https://github.com/i0Ek3/gotips#1-nil)

- [2 defer](https://github.com/i0Ek3/gotips#2-defer)

- [3 loop](https://github.com/i0Ek3/gotips#3-loop)

- [4 string](https://github.com/i0Ek3/gotips#4-string)

- [5 byte & rune](https://github.com/i0Ek3/gotips#5-byte--rune)

- [6 array & slice](https://github.com/i0Ek3/gotips#6-array--slice)

- [7 map](https://github.com/i0Ek3/gotips#7-map)

- [8 channel](https://github.com/i0Ek3/gotips#8-channel)

- [9 func](https://github.com/i0Ek3/gotips#9-func)

- [10 method](https://github.com/i0Ek3/gotips#10-method)

- [11 interface](https://github.com/i0Ek3/gotips#11-interface)

- [12 struct](https://github.com/i0Ek3/gotips#12-struct)

- [13 goroutine](https://github.com/i0Ek3/gotips#13-goroutine)

- [14 context](https://github.com/i0Ek3/gotips#14-context)

- [15 lock](https://github.com/i0Ek3/gotips#15-lock)

- [16 init](https://github.com/i0Ek3/gotips#16-init)

- [17 itoa](https://github.com/i0Ek3/gotips#17-itoa)

- [18 assign](https://github.com/i0Ek3/gotips#18-assign)

- [19 pointer](https://github.com/i0Ek3/gotips#19-pointer)

- [20 operator](https://github.com/i0Ek3/gotips#20-operator)

- [21 constant](https://github.com/i0Ek3/gotips#21-constant)

- [22 error](https://github.com/i0Ek3/gotips#22-error)

- [23 misc](https://github.com/i0Ek3/gotips#23-misc)

- [24 other](https://github.com/i0Ek3/gotips#24-other)

## Tips

> Classification Rule: determine the corresponding topic according to the subject in the tips.

### 1. nil

- 两个 `nil` 是不相等的，故无法通过 `!=` 进行比较。

- `nil` 切片（`var s []T`）和空切片（`s := make([]T, 0)` 或者 `s := []T{}`）是不同的切片，前者不会分配内存，而后者会分配内存。[参考](https://stackoverflow.com/questions/59349879/whats-the-difference-between-int-and-int-in-go)

- `nil` 只能赋值给指针、`chan`、`func`、`interface`、`map`、`slice` 等类型的变量。如果不指定变量的类型，编译器将无法得出变量的具体类型，导致编译错误。[参考 1](https://go.dev/play/p/3M-DfjOP-Vr) [参考 2](https://go.dev/play/p/Gp0wBvdndWs)

- 允许对值为 `nil` 的 `slice` 添加元素，但对值为 `nil` 的 `map` 添加元素时，会造成运行时 `panic`。[参考 1](https://go.dev/play/p/L_cYqV84DrN) [参考 2](https://go.dev/play/p/kfjcbO74u4Z)

### 2. defer

- `defer` 函数的参数（包括接收者）是在 `defer` 语句出现的位置做计算的，而不是在函数执行的时候计算的。即 `defer` 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，等到真正执行 `defer` 时取出。[参考 1](https://go.dev/play/p/9n-JMGhicmQ) [参考 2](https://go.dev/play/p/r_wvQjHDO8Q) [参考 3](https://go.dev/play/p/ZTrNCA8IclB) [参考 4](https://go.dev/play/p/LBr2jCRHmRU)

- `return` 会先于 `defer` 返回，且 `return` 不是原子操作（`return` 语句分为赋值和返回两个部分，命名返回值会被 `return` 后面的值覆盖掉 [参考](https://go.dev/play/p/igGUWZqHPTN)）。因此，在有名返回函数中，一定要注意 `return` 语句。

- `defer` 语句在前，且包含 `recover()` 时，当遇到 `panic` 程序将停止执行，然后调用 `defer` 函数。[参考](https://go.dev/play/p/pTe1wUxn73P)

- 当代码中不包含 `recover()` 时，出现 `panic` 语句的时候，会先按照 `defer` 后进先出的顺序执行，最后才会执行 `panic`。[参考](https://go.dev/play/p/2W16mXLG3H2)

- `recover()` 必须在 `defer` 函数中直接调用才会生效。

- `return` 之后的 `defer` 是不能注册的。[参考](https://go.dev/play/p/IPAp4769FZc)

- `defer` 语句通常应该放到 `if err != nil` 后面。[参考](https://go.dev/play/p/H2nLDO9Q3za)

- 匿名返回时，`return` 语句返回的值不会影响 `defer` 语句中的结果。

- 不要在 `for` 循环中使用 `defer`，因为 `defer` 只有在函数退出时才会执行。

- 对性能要求较高时，应避免使用 `defer Unlock()`。

### 3. loop

- 善用 `switch` 而不是多个 `if`，且 `switch` 中必须要有 `default` 子句。
- 在没有显示 `break` 的情况下，`switch` 中的 `case` 执行完会自动 `break`，若想继续执行下面的 `case`，可添加 `fallthrough` 关键字。[参考](https://go.dev/play/p/vfwgYe94fK_Y)
- 在单独的 `for` 循环中，`break` 可以跳出循环。但在 `for select` 中，`break` 可以跳出 `select` 块，但不会跳出 `for` 循环。如需跳出 `for` 循环，可以配合 `goto` 使用 `label` 解决。
- `for {}` 循环会独占 CPU 资源导致其他 goroutine 饿死，可以通过阻塞的方式避免 CPU 占用，如使用 `select {}`。[参考](https://go.dev/play/p/VgVSO6Edb_6)
- `goto` 无法跳转到其他函数或者内层代码，禁止在业务代码中使用。[参考](https://go.dev/play/p/miz2pGthALx)
- `for range` 中如果仅需要第一项，则丢弃第二个，如 `for key := range nums`；如果仅需要第二项，则第一项用下划线表示，如 `for _, val := range nums`。
- `for range` 使用短变量声明 `:=` 的形式迭代变量时，变量 `i`、`v` 在每次循环中都会被重用，而不是重新声明。[参考 1](https://go.dev/play/p/pPk92ad178b) [参考 2](https://go.dev/play/p/UghIn1HZ-l1)
- 当 `range` 表达式发生复制时，副本的指针依旧指向原底层数组，所以对切片的修改都会反应到底层数组上。[参考](https://go.dev/play/p/mDjwzkSxTt1)
- `for range` 中获取到的值是元素的副本，不会复制底层数组。[参考 1](https://go.dev/play/p/-YYOfIFYF2v) [参考 2](https://go.dev/play/p/R0oFiqzbKrj)
- 循环次数在循环开始前就已经确定，循环内改变切片的长度，不影响循环次数。[参考](https://go.dev/play/p/I67nM8JFsPZ)

### 4. string

- Go 语言中的字符串是只读的，所以想要修改 `string` 的值，需要先将 `string` 转为 `[]byte`，然后再转为 `string`。[参考](https://go.dev/play/p/o5T0H0YJfuO)

- Go 的字符串类型是不能赋值为 `nil` 的，也不能跟 `nil` 比较。[参考](https://go.dev/play/p/7B7wbztcwlr)

- 如果一个类型实现了 `String()` 方法，那么在使用 `fmt.Printf()` 等格式化输出方法时，会自动使用 `String()` 方法。因此，再次调用 `String()` 方法将导致递归调用。[参考](https://go.dev/play/p/8jOYDn0m2WY)

- 空字符串的判断应使用 `if len(s) == 0` 而不是 `if s == ""`。

- 尽量使用 `string.Builder` 进行字符串的拼接。

### 5. byte & rune

- `byte` 是 `uint8` 的别名，大小为 1 字节，代表 ASCII 码的一个字符。

- `rune` 是 `int32` 的别名，大小为 4 字节，代表一个 UTF-8 字符。当需要处理中文、日文或者其他复合字符时，则需用到 `rune` 类型。

### 6. array & slice

- 无法对 `nil` 切片进行赋值，会 `panic`。[参考](https://go.dev/play/p/9We3a-pUfXr)
- 在函数调用里修改返回的切片，将会影响到原切片。通常我们新建一个切片，然后将修改后的结果复制到该新切片，而不是改变旧有切片。
- 在拷贝切片时，`copy(dst, src)` 返回 `len(dst)`、`len(src)` 之间的最小值作为新内容的长度。如果想要将 `src` 完全拷贝至 `dst`，必须给 `dst` 分配足够的内存空间。[参考 1](https://mp.weixin.qq.com/s/3qguB_V6mwPl-G2q-TjnfA) [参考 2](https://go.dev/play/p/zEIuT1d18k0)
- 从一个基础切片派生出的子切片的长度可能大于基础切片的长度。若基础切片是 `slice`，使用操作符 `[low:high]`，只要满足 `0 <= low <= high <= cap(slice)`，下标 `low` 和 `high` 都可以大于 `len(slice)`。若其中的 `high` 省略，且 `low < len(slice)`，则默认是截取到原切片的长度；若 `low > len(slice)`，则 `panic`。否则，截取之后的切片长度和容量分别为：`len = high-low`，`cap = cap(slice)-low`。[参考](https://go.dev/play/p/9aVps3541w9)
- 使用 `make` 初始化切片时，需要补充 `len` 参数（`cap` 参数可选），否则无法编译（[参考](https://go.dev/play/p/xrKuOwRCQiU)）。当然，如果在能够确认的情况下，最好可以预先分配容量。
- 字面量初始化切片时候，可以指定索引，没有指定索引的元素会在前一个索引基础之上加一。[参考](https://go.dev/play/p/_hUqp3DqucC)
- `append()` 的第二个参数不能直接使用 `slice`，需使用 `...` 操作符来将一个切片追加到另一个切片上，或者直接跟上具体的元素。[参考](https://go.dev/play/p/lz7VtTQxQrl) 另外，尽量不要在复制时使用 `append()`，如在合并多个 slice 的时候。
- 对于空切片的判断，应该写成 `if slice != nil && len(slice) == 0` 这种方式而不是 `if len(slice) == 0`。
- 对切片截取后赋值给新的切片，则对新切片的修改会影响原来的切片，因为这两个切片共享同一个底层数组。而追加会新建一个切片，不会影响原有切片。[参考](https://go.dev/play/p/A0T0PVUWoUz)

### 7. map

- `map` 是线程不安全的，在并发中需要加锁，可以通过 `sync.RWMutex` 加锁或者使用线程安全的 `sync.Map` 来解决。同理，在查找、赋值、遍历、删除 `map` 的过程中都会检测写标志，一旦发现写标志置位（等于1），则直接 `panic`。

- `map[key]struct` 中 `struct` 是不可寻址的，因此无法直接赋值。若想知道 `struct` 中的地址，可以考虑使用临时变量（[参考](https://go.dev/play/p/l5dIxdOHpn5)）或者修改数据结构（[参考](https://go.dev/play/p/0-0Xr_ytXUt)）。[参考](https://go.dev/play/p/tNum3CIxW7l)

- `map` 必须初始化才能使用，且无法对 `map` 的 `key` 和 `value` 进行取地址，否则将无法编译。[参考](https://go.dev/play/p/1glmROGjWK5)

- 使用 `make` 创建 `map` 变量时可以指定第二个参数，不过会被忽略。

- 两个 `map` 是无法直接进行比较的，只能通过遍历 `map` 的每个元素，比较元素是否都深度相同。[参考](https://play.golang.com/p/bVT7Iynj4mG)

- 检查 `map` 中的 `key` 是否存在，可以使用返回的第二个参数 `ok` 来判断，如 `v, ok := map["hi"]` 中的 `v` 会返回 "hi" 在 `map` 中对应的值，如果 "hi" 不存在，则返回对应类型的零值，`ok` 会返回 "hi" 是否存在于 `map` 中。

- 删除 `map` 中不存在的键值对时，不会报错，相当于没有任何作用；获取、打印 `map` 中不存在的键值对时，返回对应类型的零值。[参考](https://go.dev/play/p/bYzDU1bPN7v)

- 尽量不要在 `map` 的键和值中使用指针，这样可以减少 GC 的开销。另外，字符串也是指针，若想要在 `map` 中使用，尽量使用 `[]byte` 而不是 `string`。

### 8. channel

- 使用 `chan struct{}` 来传递信号，尽量避免使用 `chan bool`。

- 读、写一个 `nil channel` 会造成永久阻塞；向已经关闭的 `channel` 发送数据，会造成 `panic`；从一个已经关闭的 `channel` 接收数据，如果缓冲区为空，则返回一个对应类型的零值，否则读取出对应的值；关闭一个已经关闭的 `channel` 会 `panic`。[参考](https://go.dev/play/p/h7NnRmXbtEA)

- 及时用 `close()` 关闭通道，否则可能会导致死锁。如在 `for range` 循环中遍历 `channel` 中的值时。

- 不能在单向通道上做逆向操作，也不能用 `close` 函数关闭接收端（`<-chan`），但可以关闭发送端。[参考](https://go.dev/play/p/UYJtCrteG_V)

- `select` 语句会随机选择一个可用通道做收发操作，如要等全部通道消息处理结束(`closed`)，可将已完成通道设置为 `nil`，这样它就会被阻塞，不再被 `select` 选中。当所有通道都不可用时，`select` 会执行 `default` 语句，避开 `select` 阻塞，也可用 `default` 处理一些默认逻辑。

### 9. func

- Go 中的可变参数作为函数参数时，必须放在最后一位。

- 函数只能与 `nil` 比较。[参考 1](https://go.dev/play/p/_vtECkR00ZZ) [参考 2](https://go.dev/play/p/lVSTjlMlLMR)

- 可变长参数作为函数的参数，传递的是指针，因此在函数内部修改可变长参数会修改原数据。[参考 1](https://go.dev/play/p/apu9JTmorrp) [参考 2](https://go.dev/play/p/NnGkzIPWDeD)

- 如果想实现函数或者方法的链式调用，则返回该函数或者方法的指针值即可（不过这样会引起变量逃逸）。

- 在函数有多个返回值时，只要有一个返回值是命名的，其他的也必须命名。如果有多个返回值，则必须加上括号 `()`；如果只有一个返回值且命名也必须加上括号 `()`。

- 在函数中，不要传递引用类型（`map, slice, chan, interface`）的指针。

### 10. method

- 使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响原来的值。如果想要在调用函数中修改原来的值，可以使用指针接收者定义的方法。
- **实现了接收者是值类型的方法，相当于自动实现了接收者是指针类型的方法；而实现了接收者是指针类型的方法，不会自动生成对应接收者是值类型的方法**。[参考](https://go.dev/play/p/UYla-H0C-0U)
- 当使用 `type` 声明一个新类型时，它不会继承原有类型的方法集。
- 非命名类型（`unamed type`，如 `struct{}`、`[]string`、`interface{}`、`map[string]bool` 等）不能作为方法的接收者。[参考](https://go.dev/play/p/Xbdnni_JasU)
- 当目标方法的接收者是指针类型时，那么被复制的就是指针。[参考](https://go.dev/play/p/ttoONtuoHan)
- 在方法中，指针类型的接收者必须是合法指针（包括 `nil`），或者能够获取到实例地址的表达式。[参考](https://go.dev/play/p/kW1kZeZMWBF)
- 不可寻址的结构体不能调用带结构体指针接收者的方法，但可以调用值接收者的方法。[参考](https://go.dev/play/p/OibZeTvBQJF)
- 基于类型创建的方法必须定义在同一个包内，或者定义该类型的一个新类型。[参考 1](https://go.dev/play/p/u5fYzh-7b72) [参考 2](https://go.dev/play/p/gWdbC0S_Z-d)

### 11. interface

- 如何检测某个类型是否实现了接口？
  
  - `var _ interface_name = (*type_name)(nil)` 检测 `*type_name` 类型是否实现了 `interface_name` 接口
  
  - `var _ interface_name = type_name{}` 检测 `type_name` 类型是否实现了 `interface_name` 接口

- `iface` 和 `eface` 都是 Go 中描述接口的底层结构体，区别在于 `iface` 描述的接口包含方法，而 `eface` 描述的接口是不包含任何方法的空接口（`interface{}`）。

- 当且仅当动态值和动态类型都为 `nil` 时，接口类型值才为 `nil`。[参考](https://go.dev/play/p/K-iX86rToeG)

- `interface{}` 作为函数参数时，可以接收任何类型的参数，包括用户自定义的类型以及指针类型，但不要使用 `*interface{}`。[参考](https://go.dev/play/p/9Ekk_CUlYdI)

- 将小整数转换为接口值不再需要进行内存分配（小整数是指 0 到 255 之间的数）。因此，一般来说接口意味着必须在堆中动态分配。

- 仅有接口类型的变量才能使用类型断言。[参考](https://go.dev/play/p/dGLVU2XEz5R)

### 12. struct

- 可在结构体中添加 `_ struct{}` 字段以防止结构体字段使用纯值方式初始化。

- 结构体中的私有属性不建议增加 `JSON` 标签，因为无法解析。

- 无法为函数返回的结构体中的字段赋值。[参考](https://go.dev/play/p/qaLDSDS2Udn)

- 空结构体 `struct{}` 实例不占据任何的内存空间。

- 使用 `&T{}` 代替 `new(T)`。

- 两个结构体在内部字段类型和字段顺序以及类型都为可比较类型时，才可以用 `==` 和 `!=` 进行比较。

- 嵌入式类型应位于结构体内字段列表的顶部，且必须有一个空行将嵌入式字段与常规字段分隔开。

- Go 中结构体里的成员变量最好要全部大写。

- 尽量避免复制较大（超过 4 个字段）的 `struct`，我们可以通过内存对齐来减小 `struct` 的大小。

### 13. goroutine

- 在未进行并发控制的代码中，如果存在多处 goroutine，则他们的运行顺序是不确定的。[参考](https://go.dev/play/p/07mnc88nAzD)
- 如果你不知道该如何关闭 goroutine，则不要使用它们。

### 14. context

- 不要将 `context` 放到结构体中，而是直接将 `context` 类型作为函数的第一参数，并且命名为 `ctx`。

- 不要向函数传入一个 `nil` 的 `context`，如果你实在不知道传什么，则用 `context.TODO()`。

- 不要把本应该作为函数参数的类型塞到 `context` 中，`context` 存储的应该是一些共同的数据。例如登陆的 session、cookie 等。

- `context` 是天然并发安全的。因此，同一个 `context` 可以被传递到多个 goroutine 中。

### 15. lock

- 对变量加锁后再进行复制，会将锁的状态一同复制。

- **将 `Mutex` 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。也可以通过嵌入 `*Mutex` 来避免复制的问题，但需要初始化**。[参考](https://go.dev/play/p/iL0qUgiiggH)

- 读写锁的存在是为了解决读多写少时的性能问题，读场景较多时，读写锁可有效地减少锁阻塞的时间。

### 16. init

- `init()` 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误。

### 17. itoa

- 在一个常量声明代码块中，如果 `iota` 没有出现在第一行，则常量的初始值就是非零值（即对应的行数）。[参考 1](https://go.dev/play/p/eR5v-srspmq) [参考 2](https://studygolang.com/articles/2192)

### 18. assign

- 在同一个作用域中，多次声明同一个变量名，后声明的变量仅在当前作用域生效。

- 不同类型的值是不能相互赋值的，即使底层类型一样。对于底层类型相同的变量可以相互赋值的一个重要条件是，至少有一个变量不是有名类型（`named type`，如内置类型和用 `type` 声明的类型）。[参考](https://go.dev/play/p/Dc6b2Ee5cPs)

- 多重赋值分为两个步骤，先分别计算等号左边的表达式和等号右边的表达式，然后再将右边的值赋值给左边。[参考](https://go.dev/play/p/PFj1vbJ1awa)

- `:=` 操作符不能用于结构体字段赋值（[参考 1](https://go.dev/play/p/6v8mUJzsP0M) [参考 2](https://go.dev/play/p/Ui-VB7OVGOl)）。并且，其必须在函数内部使用。[参考](https://go.dev/play/p/sZputp9qoSV)

- 不能在单独的声明中重复声明一个变量，但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的。如果出现作用域之后，则会导致变量隐藏的问题。[参考](https://go.dev/play/p/BBRMntynbLB)

### 19. pointer

- 不能使用多级指针调用方法。[参考](https://go.dev/play/p/L6jl8KpI2D7)
- 只要有一个指针指向一个引用的变量，那么这个变量就不会被释放。因此，在 Go 语言中返回函数参数或临时变量是安全的。
- 对空指针解引用会造成 `panic`。[参考](https://go.dev/play/p/MHb9socpfdk)
- 指针不支持索引。[参考](https://go.dev/play/p/5tsektd4No-)
- 永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。
- `new` 一个对象，返回的是该对象的指针类型，不能对指针执行 `append()` 操作。[参考](https://go.dev/play/p/8g8BLBNHH4b)
- 当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象并与其绑定，以便在稍后执行时能隐式传入接收者参数。[参考](https://go.dev/play/p/-n9RfLyVWVX)

### 20. operator

- 递增运算符 `++` 和递减运算符 `--` 的优先级低于解引用运算符 `*` 和取址运算符  `&`，解引用运算符和取址运算符的优先级低于选择器 `.` 中的属性选择操作符。[参考](https://go.dev/play/p/-bGBSCUGbyG)

- Go 语言中不支持 `++i` 和 `--i` 操作。另外，`i++` 和 `i--` 在 Go 语言中是语句，不是表达式，因此不能赋值给其他变量。

- `&^` 为按位置零。表达式 `z = x &^ y` 表示为如果 `y` 中的 bit 位为 1，则 `z` 对应的 bit 位为 0，否则 `z` 对应 bit 位等于 `x` 中相应的 bit 位的值。

- `|` 为或操作符。表达式 `z = x | y` 表示为如果 `y` 中的 bit 位为 1，则 `z` 对应 bit 位为 1，否则 `z` 对应 bit 位等于 `x` 中相应的 bit 位的值。

### 21. constant

- 同变量不同，Go 语言中常量未使用是能编译通过的，但不可以对常量取地址。[参考](https://go.dev/play/p/kCKr-8S_IUP)

- 常量组中如不指定类型和初始化值，则与上一行非空常量右值相同。[参考](https://go.dev/play/p/mkBzT7ttBfv)

### 22. error

- `error` 作为函数的返回值，必须要对其进行处理，或者赋值给 `_`。

- `error` 作为函数的多个返回值之一，必须是最后一个参数。

- `error` 字符串不应该大写开头或者在末尾加上标点符号。

### 23. misc

- Go 中仅有值传递。
- `%+d` 表示输出数值的符号。[参考](https://go.dev/play/p/MJjQAexpHV7)
- 两个不同类型的数值不能相加，否则会编译报错。[参考](https://go.dev/play/p/DehdvmqOjtc)
- Go 中不同类型是不能比较的（比如长度不同的两个数组），切片也是不能进行比较的。[参考](https://go.dev/play/p/hgAuoeiYg57)
- 常见的 `bool`、数值型、字符、指针、数组等类型是可以比较的，而切片、`map`、函数等是不可比较的。
- 类型转换、类型断言本质上都是把一个类型转换成另外一个，但类型断言是对接口变量进行的操作。
- 注意区分类型别名（有 `=`）和新建类型（无 `=`）的区别.前者和原类型一致，拥有原类型实现的方法，而后者是新的类型。[参考](https://go.dev/play/p/BxQsLa-zhnh)
- Go 语言中不存在引用变量，每个变量都占用一个唯一的内存位置。
- Go 中的预定义标识符（如 `string`、`len` 等）是可以作为变量使用的，但关键字不行（如 `default`）。
- 用字面量初始化数组、`slice` 和 `map` 时，最好是在每个元素后面加上逗号。[参考](https://go.dev/play/p/D9v3aFCTRL0)
- `cap()` 函数适用于数组、数组指针、`slice` 和 `channel`，不适用于 `map`，可以使用 `len()` 返回 `map` 的元素个数。当使用 `make` 创建 `map` 变量时指定第二个参数会被忽略。
- 如果有未使用的变量，代码将编译失败，但可以有未使用的全局变量。另外，函数的参数未使用也是可以的。当然，如无必要，可以注释掉或者移除未使用的变量。[参考](https://go.dev/play/p/xYxO9jOJNYg)
- Go 语言中，大括号不能放在单独的一行，否则会编译错误。
- Go 中存在断行规则，请在 `;` 之后断行。[参考](https://gfw.go101.org/article/line-break-rules.html)
- 请注意代码中 `println()` 和 `fmt.Println()` 的区别，后者会使得变量逃逸。[参考](https://go.dev/play/p/PNLMlw2nHn4)
- Go 语言中大多数数据类型都可以转化为有效的 `JSON` 文本，但 `channel`、`complex`、`func` 不行。[参考](https://go.dev/play/p/EPi1Y0YTNIn)
- 一个类型确定数字型常量所表示的值是不能溢出它的类型的表示范围的。一个类型不确定数字型常量所表示的值是可以溢出它的默认类型的表示范围的。当一个类型不确定数字常量值溢出它的默认类型的表示范围时，此数值不会被截断。

### 24. Other

- 格式化
  
  - 使用 `gofmt` 对代码进行格式化
  
  - 使用 `goimports` 对 `import` 部分格式化，且运算符与操作数之间留有一个空格
  
  - 优先使用 `strconv` 而不是 `fmt`

- 代码长度
  
  - 函数参数不宜过长，应控制在 5 个以内
  
  - 可将主要代码中的一行长度控制在 80-120 字符之间，超出的部分可以换行
  
  - 单个文件中代码行数最好不要超过 800 行
  
  - 函数长度最好不要超过 80 行，否则考虑重构

- 命名规则
  
  - 包名全部小写，不允许有大写或者下划线
  
  - 项目名可以使用中划线连接多个单词
  
  - 函数名要采用驼峰式
  
  - 文件名要小写，并使用下划线分隔单词
  
  - 结构体的命名要用驼峰式，且使用名词而不是动词
  
  - 接口命名规则与结构体命名规则基本一致，单个函数的接口使用 "er" 结尾，两个函数以两个函数名命名，三个以上类似于结构体名

- 注释
  
  - 每个可导出的命名都要有注释
  
  - 禁止使用多行注释
  
  - 多段注释之间使用空格进行分隔
  
  - 注释掉的代码在提交前应该删除，否则说明不删的理由和后续处理建议

## Reference

- [https://studygolang.com](https://studygolang.com)

- [https://dave.cheney.net/practical-go/presentations/qcon-china.html](https://dave.cheney.net/practical-go/presentations/qcon-china.html)

- [https://golang.design/go-questions/](https://golang.design/go-questions/)

- [https://geektutu.com/post/high-performance-go.html](https://geektutu.com/post/high-performance-go.html)

- [https://gfw.go101.org/article/101.html](https://gfw.go101.org/article/101.html)

- [https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw](https://mp.weixin.qq.com/s/rEXhrAqEOg9Ja4wYomOsGw)

- [https://www.practical-go-lessons.com](https://www.practical-go-lessons.com)

- [https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/](https://tonybai.com/2015/09/17/7-things-you-may-not-pay-attation-to-in-go/)

- [https://mp.weixin.qq.com/s/QONfbKioFf6VqJE2OwP7Kw](https://mp.weixin.qq.com/s/QONfbKioFf6VqJE2OwP7Kw)

- [https://medium.com/scum-gazeta/golang-simple-optimization-notes-70bc64673980](https://medium.com/scum-gazeta/golang-simple-optimization-notes-70bc64673980)

## Credit

You All Guys!
