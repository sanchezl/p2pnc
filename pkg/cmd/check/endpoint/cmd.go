package endpoint

import (
	"context"
	"net"
	"net/http/httptrace"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tcp-endpoint",
		Short: "Checks that a tcp connection can be opened to an endpoint.",
		Args:  validatePositionalArgs,
		Run:   run,
	}
	return cmd
}

func run(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	//TODO make this performant, cancelable, timeouts, etc.
	for _, address := range args {
		go checkAddress(ctx, address)
	}
	<-ctx.Done()
}

func checkAddress(ctx context.Context, address string) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		ctx, latencyInfo := WithLatencyInfoCapture(ctx)
		dialer := &net.Dialer{}
		conn, err := dialer.DialContext(ctx, "tcp", address)
		if err == nil {
			conn.Close()
		}
		logCheckActions(address, err, latencyInfo)
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func isDNSError(err error) bool {
	if opErr, ok := err.(*net.OpError); ok {
		if _, ok := opErr.Err.(*net.DNSError); ok {
			return true
		}
	}
	return false
}

func logCheckActions(address string, checkErr error, latency *LatencyInfo) {
	host, _, _ := net.SplitHostPort(address)
	if isDNSError(checkErr) {
		klog.Infof("%7s | %-11s | %6dms | 🔍❌ Failure looking up host %s: %v", "Failure", "DNSDone", latency.DNS.Milliseconds(), host, checkErr)
		return
	}
	if latency.DNS != 0 {
		klog.Infof("%7s | %-11s | %6dms | 🔍✔ Resolved host name %s successfully", "Success", "DNSDone", latency.DNS.Milliseconds(), host)
	}

	if checkErr != nil {
		klog.Infof("%7s | %-11s | %6dms | 🔌❌ Failed to establish a TCP connection to %s: %v", "Failure", "ConnectDone", latency.Connect.Milliseconds(), address, checkErr)
		return
	}
	klog.Infof("%7s | %-11s | %6dms | 🔌✔ TCP connection to %v succeeded", "Success", "ConnectDone", latency.Connect.Milliseconds(), address)
}

func validatePositionalArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}
	return nil
}

type LatencyInfo struct {
	DNS           time.Duration
	Connect       time.Duration
	_dnsStart     time.Time
	_connectStart time.Time
}

func (r *LatencyInfo) dnsStart() {
	r._dnsStart = time.Now()
}

func (r *LatencyInfo) connectStart(addr string) {
	if r._connectStart.IsZero() {
		r._connectStart = time.Now()
	}
}

func (r *LatencyInfo) dnsDone() {
	r.DNS = time.Now().Sub(r._dnsStart)
}

func (r *LatencyInfo) connectDone(addr string) {
	r.Connect = time.Now().Sub(r._connectStart)
}

func WithLatencyInfoCapture(ctx context.Context) (context.Context, *LatencyInfo) {
	trace := &LatencyInfo{}
	return httptrace.WithClientTrace(ctx, &httptrace.ClientTrace{
		DNSStart: func(info httptrace.DNSStartInfo) {
			trace.dnsStart()
			klog.V(5).Infof("DNSStart: %s\n", info.Host)
		},
		DNSDone: func(info httptrace.DNSDoneInfo) {
			trace.dnsDone()
			klog.V(5).Infof("DNSDone: %v\n", info)
		},
		ConnectStart: func(network, addr string) {
			trace.connectStart(addr)
			klog.V(5).Infof("ConnectStart: %s %s\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			trace.connectDone(addr)
			klog.V(5).Infof("ConnectDone: %s,%s,%v\n", network, addr, err)
		},
	}), trace
}
