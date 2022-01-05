这个仓库用来演示 Yaegi 在处理 interface 时的一些问题。被我自己的 wiki 所引用。

问题是：

对于这样一个普通的 interface：

```go
type Cheater interface {
	Cheat()
}
```

定义一个普通的 struct `A` 实现该接口：

```go
type A struct{}

func (A) Cheat() {}

func main() {
	var i common.Cheater = (*A)(nil)

	fmt.Println("Using interface natively:")
	fmt.Printf("reflect.TypeOf(i): %s\n", reflect.TypeOf(i))
	fmt.Println("")
}
```

通过 `(*A)(nil)` 传给 `i` 时，`i` 这个 interface 值的底层 `(type, value)` 对应该是 `(*main.A, nil)`。使用 `reflect.TypeOf(i)` 拿到的应该是 `*main.A`。

但是在 yaegi 中运行时：

* 假如 `Cheater` 被定义为 [binary form](https://pkg.go.dev/github.com/traefik/yaegi@v0.11.2/interp#hdr-Importing_packages)，那在 source form 代码中调用 `reflect.TypeOf(i)` 时拿不到类型信息，只能看到 `<nil>`
* 假如 `Cheater` 也被定义在 source form 中，`reflect.TypeOf(i)` 可以拿到类型信息，但是丢失了 struct 名，只能展示 `*struct {}`

运行这些命令体现上述的问题：

```bash
$ go run cmd/normal/main.go 
Using interface natively:
reflect.TypeOf(i): *main.A

$ go run cmd/yaegirun/main/main.go
Using interface defined outside:
reflect.TypeOf(i): <nil>

Using interface defined inside:
reflect.TypeOf(i): *struct {}
Type assertion: <nil>, true
```

官方有个 [issue](https://github.com/traefik/yaegi/issues/947) 有类似的问题，但是其中 yaegi 的作者 [回复](https://github.com/traefik/yaegi/issues/947#issuecomment-737880201) 表示这种跨边界的 interface 使用是会出现这种问题，而且不好修复。