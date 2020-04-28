# point-to-point network check

> **WIP**

```console
$ p2pnc check tcp-endpoint  localhost:9090 localhost:9091 google.com:80 does-not-exists.com:80 google.com:443 127.0.0.1:9092

I0428 16:14:07.369824 1672695 cmd.go:66] Success | DNSDone     |     0ms | 🔍✔ Resolved host name localhost successfully
I0428 16:14:07.369882 1672695 cmd.go:73] Success | ConnectDone |     0ms | 🔌✔ TCP connection to localhost:9090 succeeded
I0428 16:14:07.370005 1672695 cmd.go:66] Success | DNSDone     |     0ms | 🔍✔ Resolved host name localhost successfully
I0428 16:14:07.370056 1672695 cmd.go:70] Failure | ConnectDone |     0ms | 🔌❌ Failed to establish a TCP connection to localhost:9091: dial tcp [::1]:9091: connect: connection refused
I0428 16:14:07.375280 1672695 cmd.go:62] Failure | DNSDone     |     6ms | 🔍❌ Failure looking up host does-not-exists.com: dial tcp: lookup does-not-exists.com on 127.0.0.1:53: no such host
I0428 16:14:07.403873 1672695 cmd.go:66] Success | DNSDone     |     0ms | 🔍✔ Resolved host name google.com successfully
I0428 16:14:07.403905 1672695 cmd.go:73] Success | ConnectDone |    34ms | 🔌✔ TCP connection to google.com:443 succeeded
I0428 16:14:07.403996 1672695 cmd.go:66] Success | DNSDone     |     1ms | 🔍✔ Resolved host name google.com successfully
I0428 16:14:07.404008 1672695 cmd.go:73] Success | ConnectDone |    34ms | 🔌✔ TCP connection to google.com:80 succeeded

. . .

```