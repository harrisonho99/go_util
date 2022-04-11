package use_context

import (
	"fmt"
	"net"
	"net/http"

	//reimplement
	"github.com/hotsnow199/go_util/context"
)

func GetIPFromRequest(r *http.Request) (net.IP, error) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}

	userIP := net.ParseIP(ip)
	if userIP == nil {
		return nil, fmt.Errorf("userip: %q is not IP:port", r.RemoteAddr)
	}

	return userIP, nil
}

type key int

const userIPKey key = 0

func NewUserIPContext(ctx context.Context, userIP net.IP) context.Context {
	return context.WithValue(ctx, userIPKey, userIP)
}

func UserIPFromContext(ctx context.Context) (net.IP, bool) {
	userIp, ok := ctx.Value(userIPKey).(net.IP)
	return userIp, ok
}
