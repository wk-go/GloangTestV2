# README
执行`go build -gcflags '-m -l' main.go`可查看逃逸状态。

使用反汇编命令`go tool compile -S main.go`可以查看变量是否发生逃逸，可以看到调用了runtime.newobject()函数。