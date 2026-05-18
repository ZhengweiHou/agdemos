package benchtest

import (
	kdemob "agdemoA/internal/adpgen/kitex/demob"
	kdemobapi "agdemoB/api/demob"
	"context"
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/transport"
)

func TestKitexPoc0(t *testing.T) {
	c, err := kdemob.NewClient(
		"agdemob_kitex",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithTransportProtocol(transport.GRPC),
		// client.WithShortConnection(),
		// client.WithMuxConnection(1), // 连接多路复用，仅支持TTHead
		// client.WithLongConnection(connpool.IdleConfig{
		// 	MinIdlePerAddress: 1,
		// 	MaxIdlePerAddress: 1,
		// 	MaxIdleGlobal:     1,
		// 	MaxIdleTimeout:    time.Minute,
		// }),
		// client.WithConnReporterEnabled(),
	)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	req := &kdemobapi.BRequest{
		Name: "kitex-poc",
	}

	resp, err := c.Calldemob(context.Background(), req)
	if err != nil {
		t.Fatalf("Calldemob failed: %v", err)
	}
	fmt.Printf("resp: %v\n", resp)
	// time.Sleep(time.Second)

}

func TestKitexPoc1(t *testing.T) {
	c, err := kdemob.NewClient(
		"agdemob_kitex",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithTransportProtocol(transport.GRPC),
		// client.WithShortConnection(),
		// client.WithMuxConnection(1), // 连接多路复用，仅支持TTHead
		// client.WithLongConnection(connpool.IdleConfig{
		// 	MinIdlePerAddress: 1,
		// 	MaxIdlePerAddress: 1,
		// 	MaxIdleGlobal:     1,
		// 	MaxIdleTimeout:    time.Minute,
		// }),
		// client.WithConnReporterEnabled(),
	)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}

	req := &kdemobapi.BRequest{
		Name: "kitex-poc",
	}

	// start cpu profile
	cpuProfile, _ := os.Create("benchmark-cpu.pprof")
	defer cpuProfile.Close()
	_ = pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	// heap profile after finish
	heapProfile, _ := os.Create("benchmark-mem.pprof")
	defer func() {
		_ = pprof.WriteHeapProfile(heapProfile)
		heapProfile.Close()
	}()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10000; j++ {
				resp, err := c.Calldemob(context.Background(), req)
				if err != nil {
					t.Fatalf("Calldemob failed: %v", err)
				}
				fmt.Printf("resp: %v\n", resp)
				// time.Sleep(time.Second)
			}
		}()
	}

	wg.Wait()
}

// go test -bench=BenchmarkKitexPoc -run=^$ -v
func BenchmarkKitexPoc(b *testing.B) {

	c, err := kdemob.NewClient(
		"agdemob_kitex",
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithTransportProtocol(transport.GRPC),
		// client.WithShortConnection(),
	)
	if err != nil {
		b.Fatalf("NewClient failed: %v", err)
	}

	req := &kdemobapi.BRequest{
		Name: "kitex-poc",
	}

	// start cpu profile
	cpuProfile, _ := os.Create("benchmark-cpu.pprof")
	defer cpuProfile.Close()
	_ = pprof.StartCPUProfile(cpuProfile)
	defer pprof.StopCPUProfile()

	// heap profile after finish
	heapProfile, _ := os.Create("benchmark-mem.pprof")
	defer func() {
		_ = pprof.WriteHeapProfile(heapProfile)
		heapProfile.Close()
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		_, err := c.Calldemob(context.Background(), req)
		if err != nil {
			b.Fatalf("Calldemob failed: %v", err)
		}
	}
}
