### go工具链翻译

YY中...

### 文件命名
类似于文档翻译，或者直接与文档翻译合并为一个文件。

### 如何工作
如下代码：
```go
func fn() error {
    return fmt.Errorf("error")
}
```
经过工具处理之后，变成：
```go
func fn() error {
    return fmt.Errorf(cb5e100e5a9a3e7f6d1fd97512215282)
}
```
并产生一个symbol_zh.go的文件：
```go
const(
    cb5e100e5a9a3e7f6d1fd97512215282="error"
)
```
这样给build提供一个开关，即可编译相应语种的代码。

