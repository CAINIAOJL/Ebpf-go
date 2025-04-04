//go:build ignore

#include "vmlinux.h"
#include "common.h"
#include <stdbool.h>
#define TASK_COMM_LEN 16
#define IPPROTO_TCP 0


struct event {
    unsigned __int128 saddr;    //源IP地址
    unsigned __int128 daddr;    //目的IP地址
    __u64 skaddr;       //socket地址
    __u64 ts_us;        //时间戳
    __u64 delta_us;     //时间差
    __u32 pid;          //进程ID
    int oldstate;        //旧状态   
    int newstate;        //新状态
    __u16 family;        //协议族
    __u16 sport;         //源端口
    __u16 dport;         //目的端口

    char task[TASK_COMM_LEN];    //进程名
};


#define MAX_ENTRIES 10240
#define AF_INET 2
#define AF_INET6 10

const volatile bool filter_by_sport = false;
const volatile bool filter_by_dport = false;
const volatile short target_family = 0;

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, MAX_ENTRIES);
    __type(key, __u16);
    __type(value, __u16);
}sports SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, MAX_ENTRIES);
    __type(key, __u16);
    __type(value, __u16);
} dports SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, MAX_ENTRIES);
    __type(key, struct sock*);
    __type(value, __u64);
} timestamps SEC(".maps");


struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    //__uint(key.size, sizeof(__u32));
    //__uint(value.size, sizeof(__u32));
    __uint(key_size, sizeof(__u32));
	__uint(value_size, sizeof(__u32));
} events SEC(".maps");

SEC("tracepoint/sock/inet_sock_set_state")
int handle_set_state(struct trace_event_raw_inet_sock_set_state *ctx) {
    struct sock *sk = (struct sock*) ctx->skaddr;
    __u16 family = ctx->family;
    __u16 sport = ctx->sport;
    __u16 dport = ctx->dport;
    __u64 *tsp, delta_us, ts;
    struct event event = {};

    if(ctx->protocol != IPPROTO_TCP) {
        return 0;
    }


    if(target_family && target_family != family) {
        return 0;
    }

    if(filter_by_sport && !bpf_map_lookup_elem(&sports, &sport)) {
        return 0;
    }

    if(filter_by_dport && !bpf_map_lookup_elem(&dports, &dport)) {
        return 0;
    }

    tsp = bpf_map_lookup_elem(&timestamps, &sk);
    ts = bpf_ktime_get_ns();
    if(!tsp) {
        delta_us = 0;
    } else {
        delta_us = (ts - *tsp) / 1000;
    }

    event.skaddr = (__u64)sk;
    event.ts_us = ts / 1000;
    event.delta_us = delta_us;
    event.pid = bpf_get_current_pid_tgid() >> 32;
    event.oldstate = ctx->oldstate;
    event.newstate = ctx->newstate;
    event.family = family;
    event.sport = sport;
    event.dport = dport;
    bpf_get_current_comm(&event.task, sizeof(event.task));

    if(family == AF_INET) {
        //ipv4
        bpf_probe_read_kernel(&event.saddr, sizeof(event.saddr), (struct sock*)&sk->__sk_common.skc_rcv_saddr);
        bpf_probe_read_kernel(&event.daddr, sizeof(event.daddr), &sk->__sk_common.skc_daddr);
    } else if(family == AF_INET6) {
        //ipv6
        bpf_probe_read_kernel(&event.saddr, sizeof(event.saddr), &sk->__sk_common.skc_v6_rcv_saddr.in6_u.u6_addr32);
        bpf_probe_read_kernel(&event.daddr, sizeof(event.daddr), &sk->__sk_common.skc_v6_daddr.in6_u.u6_addr32);
    }

    bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &event, sizeof(event));
    
    if(ctx->newstate == TCP_CLOSE) {
        bpf_map_delete_elem(&timestamps, &sk);
    } else {
        bpf_map_update_elem(&timestamps, &sk, &ts, BPF_ANY);
    }

    return 0;
}

char LICENSE[] SEC("license") = "Dual BSD/GPL";