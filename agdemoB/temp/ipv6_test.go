package temp

import (
	"fmt"
	"net/http"
	"testing"
)

func TestIPv6(t *testing.T) {
	// 绑定所有 IPv6 地址的 8080 端口（:: 表示 IPv6 任意地址）
	addr := "[::]:8080"
	// 注意：Go 中 IPv6 地址格式为 [::]:8080，但直接写 :::8080 也兼容
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("IPv6 request received")
		fmt.Fprintf(w, "Hello IPv6! Client IP: %s", r.RemoteAddr)
	})

	fmt.Printf("Server listening on %s (IPv6/IPv4 dual stack)\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
