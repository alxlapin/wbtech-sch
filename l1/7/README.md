### Предупреждение о состоянии гонки при выполнении команды `go run -race`

```
==================
WARNING: DATA RACE
Read at 0x00c0001220c0 by main goroutine:
runtime.mapassign_fast64ptr()
/Users/alxlapin/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.4.darwin-arm64/src/internal/runtime/maps/runtime_fast64_swiss.go:367 +0x39c
main.main()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:25 +0x124

Previous write at 0x00c0001220c0 by goroutine 696:
runtime.mapaccess2_faststr()
/Users/alxlapin/go/pkg/mod/golang.org/toolchain@v0.0.1-go1.24.4.darwin-arm64/src/internal/runtime/maps/runtime_faststr_swiss.go:162 +0x29c
main.main.func1()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:17 +0x80

Goroutine 696 (finished) created at:
main.main()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:14 +0x54
==================
==================
WARNING: DATA RACE
Read at 0x00c00008e018 by main goroutine:
main.main()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:25 +0x12c

Previous write at 0x00c00008e018 by goroutine 430:
main.main.func1()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:17 +0x8c

Goroutine 430 (finished) created at:
main.main()
/Users/alxlapin/Documents/go/wbtech/wbtech-sc/l1/7/main.go:14 +0x54
==================
```