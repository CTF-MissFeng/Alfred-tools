package main

import (
	aw "github.com/deanishe/awgo"
	"tools/util/decode"
	"tools/util/ip"

	"strconv"
	"time"

	"tools/util/consts"
	"tools/util/datetimes"
	"tools/util/encode"
	"tools/util/hash"
)

func run() {
	var err error
	args := consts.Workflow.Args()
	if len(args) == 0 {
		return
	}

	defer func() {
		if err == nil {
			consts.Workflow.SendFeedback()
			return
		}
	}()



	if args[0] == "time"{
		if len(args) < 2{
			return
		}
		if args[1] == "now"{
			datetimes.ProcessNow()
			return
		}
		if datetimes.RegexpTimestamp.MatchString(args[1]){
			v, e := strconv.ParseInt(args[1], 10, 32)
			if e == nil {
				datetimes.ProcessTimestamp(time.Unix(v, 0))
				return
			}
			err = e
			return
		}
		// 处理时间字符串
		err = datetimes.ProcessTimeStr(args[1])
	}else if args[0] == "hash"{
		if len(args) < 2{
			return
		}
		hash.HashMain(args[1])
		return
	}else if args[0] == "encode"{
		if len(args) < 2{
			return
		}
		encode.Main(args[1])
		return
	}else if args[0] == "decode"{
		if len(args) < 2{
			return
		}
		decode.Main(args[1])
		return
	}else if args[0] == "ip"{
		if len(args) < 2{
			return
		}
		ip.Main(args[1])
		return
	}

}

func main() {
	consts.Workflow = aw.New()
	consts.Workflow.Run(run)
}

