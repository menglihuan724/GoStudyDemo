package main

import (
	"net/http"
	"log"
	"flag"
	"strings"
	"github.com/bogdanovich/dns_resolver"
	"net/http/httputil"
	"net/url"
	"net"
	"time"
	"context"
)

type Cmd struct {
	bind string
	remote string
	ip string
}

func parseCmd() Cmd {
	var cmd Cmd
	flag.StringVar(&cmd.bind, "l", "0.0.0.0:8888", "listen on ip:port")
	flag.StringVar(&cmd.remote, "r", "http://192.168.42.116:0", "reverse proxy addr")
	flag.StringVar(&cmd.ip, "ip", "", "reverse proxy addr server ip")
	flag.Parse()
	return cmd
}


type handle struct {
reverseProxy string
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		remote := strings.Split(addr, ":")
		if cmd.ip == "" {
			resolver := dns_resolver.New([]string{"114.114.114.114", "114.114.115.115", "119.29.29.29", "223.5.5.5", "8.8.8.8", "208.67.222.222", "208.67.220.220"})
			resolver.RetryTimes = 5
			ip, err := resolver.LookupHost(remote[0])
			if err != nil {
				log.Println(err)
			}
			cmd.ip = ip[0].String()
		}
		addr = cmd.ip + ":" + remote[1]
		return dialer.DialContext(ctx, network, addr)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host

	proxy.ServeHTTP(w, r)
}






var cmd Cmd
var srv http.Server

func StartServer(bind string, remote string)  {
	log.Printf("Listening on %s, forwarding to %s", bind, remote)
	h := &handle{reverseProxy: remote}
	srv.Addr = bind
	srv.Handler = h
	//go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln("ListenAndServe: ", err)
		}
	//}()
}

func StopServer()  {
	if err := srv.Shutdown(nil) ; err != nil {
		log.Println(err)
	}
}

//url 要请求的URL
// ipaddr 当前网卡绑定的IP(一般都是网外IP)
func HttpGetFromIP(url, ipaddr string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				//本地地址  ipaddr是本地外网IP
				lAddr, err := net.ResolveTCPAddr(netw, ipaddr+":0")
				if err != nil {
					return nil, err
				}
				//被请求的地址
				rAddr, err := net.ResolveTCPAddr(netw, addr)
				if err != nil {
					return nil, err
				}
				conn, err := net.DialTCP(netw, lAddr,rAddr)
				if err != nil {
					return nil, err
				}
				deadline := time.Now().Add(35 * time.Second)
				conn.SetDeadline(deadline)
				return conn, nil
			},
		},
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	return client.Do(req)
}


func main() {
	//cmd = parseCmd()
	//StartServer("0.0.0.0:8888", "http://192.168.42.116:0")

	res,err:=HttpGetFromIP("http://wwww.baidu.com:80","192.168.42.116")
	if err !=nil{
		log.Print(err)
	}else {
		log.Print(res.Body)
	}
}


