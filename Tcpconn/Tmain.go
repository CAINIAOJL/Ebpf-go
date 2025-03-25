//package main

//import (
//    "bytes"
//    "encoding/binary"
//    "flag"
//   "fmt"
//    "log"
//    "os"
//    "os/signal"
//    "syscall"
//    "time"

//    "github.com/cilium/ebpf/link"
//    "github.com/cilium/ebpf/rlimit"
//	"github.com/cilium/ebpf/perf"
//    "golang.org/x/sys/unix"
////)
//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -tags linux bpf bpf.c -- -I/usr/src/linux-headers-6.11.0-17-generic/include
// 定义 eBPF 程序中的事件结构体
//type event struct {
//    skaddr   uint64
//    pid      uint32
//    task     [16]byte
//    saddr    [16]byte
//    sport    uint16
//    daddr    [16]byte
//    dport    uint16
//    oldstate uint8
//    newstate uint8
//    delta_us uint32
//    family   uint8
//}

// tcp_states 定义 TCP 连接状态
//var tcp_states = []string{
//    "",
//    "ESTABLISHED",
//    "SYN_SENT",
//    "SYN_RECV",
//    "FIN_WAIT1",
//    "FIN_WAIT2",
//    "TIME_WAIT",
//    "CLOSE",
//    "CLOSE_WAIT",
//    "LAST_ACK",
//    "LISTEN",
//    "CLOSING",
//    "NEW_SYN_RECV",
//    "UNKNOWN",
//}

//var (
//    verbose      = flag.Bool("verbose", false, "verbose debug output")
//    timestamp    = flag.Bool("timestamp", false, "include timestamp on output")
//    ipv4         = flag.Bool("4", false, "trace IPV4 family only")
//    ipv6         = flag.Bool("6", false, "trace IPV6 family only")
//    localports   = flag.String("L", "", "Comma-separated list of local ports to trace")
//    remoteports  = flag.String("D", "", "Comma-separated list of remote ports to trace")
//    wideOutput   = flag.Bool("W", false, "wide output")
//    exiting      = false
//    targetFamily = 0
//)

//func handleEvent(data []byte) {
//    var e event
//    err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &e)
//    if err != nil {
//        log.Printf("Error reading event: %v", err)
//        return
//    }

//    saddr := netIP(e.saddr[:], int(e.family))
//    daddr := netIP(e.daddr[:], int(e.family))

//    if *timestamp {
//        t := time.Now().Format("15:04:05")
//        fmt.Printf("%-8s ", t)
//    }

//    if *wideOutput {
//        family := 4
//        if e.family == unix.AF_INET6 {
//            family = 6
//        }
//        fmt.Printf("%-16x %-7d %-16s %-2d %-26s %-5d %-26s %-5d %-11s -> %-11s %.3f\n",
//            e.skaddr, e.pid, string(bytes.Trim(e.task[:], "\x00")), family, saddr, e.sport, daddr, e.dport,
//            tcp_states[e.oldstate], tcp_states[e.newstate], float64(e.delta_us)/1000)
//    } else {
//        fmt.Printf("%-16x %-7d %-10.10s %-15s %-5d %-15s %-5d %-11s -> %-11s %.3f\n",
//            e.skaddr, e.pid, string(bytes.Trim(e.task[:], "\x00")), saddr, e.sport, daddr, e.dport,
//            tcp_states[e.oldstate], tcp_states[e.newstate], float64(e.delta_us)/1000)
//    }
//}

//func netIP(b []byte, family int) string {
//    switch family {
//    case unix.AF_INET:
//        return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
//    case unix.AF_INET6:
//        return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x",
//            binary.BigEndian.Uint16(b[0:2]), binary.BigEndian.Uint16(b[2:4]),
//            binary.BigEndian.Uint16(b[4:6]), binary.BigEndian.Uint16(b[6:8]),
//            binary.BigEndian.Uint16(b[8:10]), binary.BigEndian.Uint16(b[10:12]),
//            binary.BigEndian.Uint16(b[12:14]), binary.BigEndian.Uint16(b[14:16]))
//    default:
//        return ""
//    }
//}

//func main() {
//    flag.Parse()

//    if *ipv4 {
//        targetFamily = unix.AF_INET
//    } else if *ipv6 {
//        targetFamily = unix.AF_INET6
//    }

    // 允许当前进程锁定内存以用于 eBPF 资源
//    if err := rlimit.RemoveMemlock(); err != nil {
//        log.Fatal(err)
//    }

    // 加载预编译的 eBPF 对象
//    objs := bpfObjects{}
//    if err := loadBpfObjects(&objs, nil); err != nil {
//        log.Fatalf("loading objects: %v", err)
//    }
//    defer objs.Close()

    // 设置 eBPF 程序的参数
//	objs.bpfVariables.FilterByDport.Set(*remoteports)
//	objs.bpfVariables.FilterBySport.Set(*localports)
//	objs.bpfVariables.TargetFamily.Set(*&targetFamily)
	
    //objs.bpf.rodata.filter_by_sport = *localports != ""
    //objs.bpf.rodata.filter_by_dport = *remoteports != ""
    //objs.bpf.rodata.target_family = uint8(targetFamily)

    // 附加 eBPF 程序
    //if err := objs.bpf.attach(); err != nil {
        //log.Fatalf("failed to attach BPF programs: %v", err)
    //}
//	rlink, err := link.AttachTracing(link.TracingOptions{
//		Program: objs.bpfPrograms.HandleSetState,
//	})
//	if err != nil {
//		log.Fatalf("failed to attach bpf program!")
//	}
//	defer rlink.Close()
    // 创建性能缓冲区
    //pb, err := link.PerfBuf(objs.maps.events, link.PerfBufOptions{
        //BPFMapOpts: ebpf.MapOptions{
            //MaxEntries: 1024,
        //},
        //HandleFunc: handleEvent,
    //})
	//reader, err := perf.NewReader(objs.bpfMaps.Events,2)
//    reader, err := perf.NewReader(objs.bpfMaps.Events, os.Getpagesize()*2)
//	if err != nil {
//        log.Fatalf("failed to open perf buffer: %v", err)
//    }
//    defer reader.Close()

    // 设置信号处理
//    sigCh := make(chan os.Signal, 1)
//    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
//    go func() {
//        <-sigCh
//        exiting = true
//    }()
//	rec := perf.Record
    // 打印表头
//    if *timestamp {
//        fmt.Printf("%-8s ", "TIME(s)")
//   }
//    if *wideOutput {
//        fmt.Printf("%-16s %-7s %-16s %-2s %-26s %-5s %-26s %-5s %-11s -> %-11s %s\n",
//            "SKADDR", "PID", "COMM", "IP", "LADDR", "LPORT", "RADDR", "RPORT",
//            "OLDSTATE", "NEWSTATE", "MS")
//    } else {
//        fmt.Printf("%-16s %-7s %-10s %-15s %-5s %-15s %-5s %-11s -> %-11s %s\n",
//            "SKADDR", "PID", "COMM", "LADDR", "LPORT", "RADDR", "RPORT",
//            "OLDSTATE", "NEWSTATE", "MS")
//    }

    // 主循环
//    for!exiting {
//        err := reader.ReadInto(rec)
//        if err != nil && err != os.ErrDeadlineExceeded {
//            log.Printf("error polling perf buffer: %v", err)
//        }
//    }
//}

////////////////////////////////////////////////////////////////////

//go:build linux
package main

import (
    "bytes"
    "encoding/binary"
    "flag"
    "fmt"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/cilium/ebpf/link"
    "github.com/cilium/ebpf/perf"
    "github.com/cilium/ebpf/rlimit"
    "golang.org/x/sys/unix"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -tags linux bpf bpf.c -- -I/usr/src/linux-headers-6.11.0-17-generic/include

// 定义 eBPF 程序中的事件结构体
type event struct {
    skaddr   uint64
    pid      uint32
    task     [16]byte
    saddr    [16]byte
    sport    uint16
    daddr    [16]byte
    dport    uint16
    oldstate uint8
    newstate uint8
    delta_us uint32
    family   uint8
}

// tcp_states 定义 TCP 连接状态
var tcp_states = []string{
    "",
    "ESTABLISHED",
    "SYN_SENT",
    "SYN_RECV",
    "FIN_WAIT1",
    "FIN_WAIT2",
    "TIME_WAIT",
    "CLOSE",
    "CLOSE_WAIT",
    "LAST_ACK",
    "LISTEN",
    "CLOSING",
    "NEW_SYN_RECV",
    "UNKNOWN",
}

var (
    verbose      = flag.Bool("verbose", false, "verbose debug output")
    timestamp    = flag.Bool("timestamp", false, "include timestamp on output")
    ipv4         = flag.Bool("4", false, "trace IPV4 family only")
    ipv6         = flag.Bool("6", false, "trace IPV6 family only")
    localports   = flag.String("L", "", "Comma-separated list of local ports to trace")
    remoteports  = flag.String("D", "", "Comma-separated list of remote ports to trace")
    wideOutput   = flag.Bool("W", false, "wide output")
    exiting      = false
    targetFamily = 0
)

func handleEvent(data []byte) {
    var e event
    err := binary.Read(bytes.NewBuffer(data), binary.LittleEndian, &e)
    if err != nil {
        log.Printf("Error reading event: %v", err)
        return
    }

    saddr := netIP(e.saddr[:], int(e.family))
    daddr := netIP(e.daddr[:], int(e.family))

    if *timestamp {
        t := time.Now().Format("15:04:05")
        fmt.Printf("%-8s ", t)
    }

    if *wideOutput {
        family := 4
        if e.family == unix.AF_INET6 {
            family = 6
        }
        fmt.Printf("%-16x %-7d %-16s %-2d %-26s %-5d %-26s %-5d %-11s -> %-11s %.3f\n",
            e.skaddr, e.pid, string(bytes.Trim(e.task[:], "\x00")), family, saddr, e.sport, daddr, e.dport,
            tcp_states[e.oldstate], tcp_states[e.newstate], float64(e.delta_us)/1000)
    } else {
        fmt.Printf("%-16x %-7d %-10.10s %-15s %-5d %-15s %-5d %-11s -> %-11s %.3f\n",
            e.skaddr, e.pid, string(bytes.Trim(e.task[:], "\x00")), saddr, e.sport, daddr, e.dport,
            tcp_states[e.oldstate], tcp_states[e.newstate], float64(e.delta_us)/1000)
    }
}

func netIP(b []byte, family int) string {
    switch family {
    case unix.AF_INET:
        return fmt.Sprintf("%d.%d.%d.%d", b[0], b[1], b[2], b[3])
    case unix.AF_INET6:
        return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x",
            binary.BigEndian.Uint16(b[0:2]), binary.BigEndian.Uint16(b[2:4]),
            binary.BigEndian.Uint16(b[4:6]), binary.BigEndian.Uint16(b[6:8]),
            binary.BigEndian.Uint16(b[8:10]), binary.BigEndian.Uint16(b[10:12]),
            binary.BigEndian.Uint16(b[12:14]), binary.BigEndian.Uint16(b[14:16]))
    default:
        return ""
    }
}

func main() {
    flag.Parse()

    if *ipv4 {
        targetFamily = unix.AF_INET
    } else if *ipv6 {
        targetFamily = unix.AF_INET6
    }

    // 允许当前进程锁定内存以用于 eBPF 资源
    if err := rlimit.RemoveMemlock(); err != nil {
        log.Fatal(err)
    }

    // 加载预编译的 eBPF 对象
    objs := bpfObjects{}
    if err := loadBpfObjects(&objs, nil); err != nil {
        log.Fatalf("loading objects: %v", err)
    }
    defer objs.Close()

    // 设置 eBPF 程序的参数
    //filterBySport := *localports != ""
    //filterByDport := *remoteports != ""
    //if err := objs.bpfMaps.Sports.Put(0, uint8(filterBySport)); err != nil {
    //    log.Fatalf("failed to set filter_by_sport: %v", err)
    //}
    //if err := objs.bpfMaps.Dports.Put(0, uint8(filterByDport)); err != nil {
    //    log.Fatalf("failed to set filter_by_dport: %v", err)
    //}
    //if err := objs.bpfMaps.Family.Put(0, uint8(targetFamily)); err != nil {
    //    log.Fatalf("failed to set target_family: %v", err)
    //}

	err := objs.bpfVariables.FilterByDport.Set(*remoteports)
	if err != nil {
		log.Fatalf("set variables 'remoteports' failed ")
	}
	err = objs.bpfVariables.FilterBySport.Set(*localports)
	if err != nil {
		log.Fatalf("set variables 'localports' failed ")
	}
    // 附加 eBPF 程序
    rlink, err := link.AttachTracing(link.TracingOptions{
        Program: objs.bpfPrograms.HandleSetState,
    })
    if err != nil {
        log.Fatalf("failed to attach bpf program: %v", err)
    }
    defer rlink.Close()

    // 创建性能缓冲区
    reader, err := perf.NewReader(objs.bpfMaps.Events, os.Getpagesize()*2)
    if err != nil {
        log.Fatalf("failed to open perf buffer: %v", err)
    }
    defer reader.Close()

    // 设置信号处理
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigCh
        exiting = true
    }()

    // 打印表头
    if *timestamp {
        fmt.Printf("%-8s ", "TIME(s)")
    }
    if *wideOutput {
        fmt.Printf("%-16s %-7s %-16s %-2s %-26s %-5s %-26s %-5s %-11s -> %-11s %s\n",
            "SKADDR", "PID", "COMM", "IP", "LADDR", "LPORT", "RADDR", "RPORT",
            "OLDSTATE", "NEWSTATE", "MS")
    } else {
        fmt.Printf("%-16s %-7s %-10s %-15s %-5s %-15s %-5s %-11s -> %-11s %s\n",
            "SKADDR", "PID", "COMM", "LADDR", "LPORT", "RADDR", "RPORT",
            "OLDSTATE", "NEWSTATE", "MS")
    }

    // 主循环
    for!exiting {
        record, err := reader.Read()
        if err != nil {
            if err == perf.ErrClosed {
                break
            }
            log.Printf("error reading perf buffer: %v", err)
            continue
        }

        if record.LostSamples > 0 {
            log.Printf("perf event ring buffer full, dropped %d samples", record.LostSamples)
            continue
        }

        handleEvent(record.RawSample)
    }
}

