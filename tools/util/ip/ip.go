package ip

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"tools/util/consts"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/yinheli/qqwry"
)


func QQwry(data string){
	ip := qqwry.NewQQwry("qqwry.dat")
	ip.Find(data)
	if len(ip.City) == 0{
		return
	}
	infoStr := fmt.Sprintf("%s %s", ip.Country, ip.City)
	consts.Workflow.NewItem(infoStr).
		Subtitle("纯真IP库").
		Icon(consts.IPIcon).
		Arg(infoStr).
		Valid(true)
}

func Ip2Region(data string){
	searcher, err := xdb.NewWithFileOnly("ip2region.xdb")
	if err != nil {
		return
	}
	defer func(){
		if searcher != nil{
			searcher.Close()
		}
	}()
	region, err := searcher.SearchByStr(data)
	if err != nil{
		return
	}
	consts.Workflow.NewItem(region).
		Subtitle("ip2region库").
		Icon(consts.IPIcon).
		Arg(region).
		Valid(true)
}

func Main(data string) {
	ip := strings.Trim(data," ")
	address := net.ParseIP(ip)
	if address == nil {
		if strings.Index(ip, "http") == -1{
			ip = "https://" + ip
		}
		host, err := url.ParseRequestURI(ip)
		if err != nil {
			return
		}
		ns, err := net.LookupHost(host.Hostname())
		if err != nil{
			return
		}
		if len(ns) == 0{
			return
		}
		ip = ns[0]
		QQwry(ip)
		Ip2Region(ip)
		return
	}
	QQwry(ip)
	Ip2Region(ip)
}