Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:

```
0
1
2
3
4
5
6
7
8
9
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        /home/user/Works/wildberries-l2/listing/listing04.go:12 +0xa8
exit status 2
```

```go

```
