// goWEB project main.go
package port_scan

import (
	"github.com/alecthomas/log4go"
	"math/rand"
	"net"
	"strconv"
	"time"
	"strings"
	"flag"
	"os"
)

const (
	startPort = 10000
	endPort   = 15000
)

// 生成一个范围内未被占用的端口
func RandomPort() (okPort int) {
	// listener, err := net.Listen("tcp", "0.0.0.0:0")
	// 轮询方式测试端口tcp是否连接正常
	/*for i := startPort; i <= endPort; i++ {
		port, err := testListen(i)
		if err != nil {
			continue
		}
		okPort = port
		break
	}*/
	for {
		rPort := randInt(startPort, endPort)
		port, err := testListen(rPort)
		if err != nil {
			continue
		}
		okPort = port
		break
	}
	return
}

func testListen(port int) (okPort int, err error) {
	//ip, _ := net.ResolveIPAddr("ip", "127.0.0.1")
	//ip := net.ParseIP("127.0.0.1")
	//_, err = net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log4go.Error("端口%d 已被占用", port)
		return 0, err
	}
	listener.Close()
	okPort = port
	return
}

// 生成指定范围内的随机数
func randInt(min, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return r.Intn(max-min) + min
}

func FlagRunmode() (str string) {
	mode := flag.String("runmode", "dev", "runmode exist 'dev,test,prod'")
	flag.Parse()
	arr := []string{"dev", "prod", "test"}
	bTag := false
	for _, v := range arr {
		if strings.Contains(v, *mode) {
			str = v
			bTag = true
			break
		}
	}
	if bTag {
		log4go.Stdout("runmode =", *mode)

	} else {
		log4go.Error("runmode not exist 'dev,test,prod'")
		os.Exit(1)
	}
	return
}
