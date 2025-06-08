package uuid

import (
	"fmt"
	"github.com/puresnr/go-cell/cast"
	"go.uber.org/atomic"
	"os"
	"strings"
	"time"
)

const (
	UuidExpire = 7 * 86400
)

var (
	hostName string
	pid      int
	incr     atomic.Int64
)

func init() {
	var err error
	hostName, err = os.Hostname()
	if err != nil {
		panic(err)
	}

	pid = os.Getpid()
}

func Uuid() string {
	return fmt.Sprintf("%d-%d-%d-%s", time.Now().Unix(), incr.Add(1), pid, hostName)
}

func IsUuidTimeout(uuid string) bool {
	idx := strings.Index(uuid, "-")
	if idx == -1 {
		return true
	}
	return time.Now().Unix() > cast.Stoi_64(uuid[:idx])+UuidExpire
}
