package internal

import (
	"github.com/name5566/leaf/log"
	"time"
)

var (
	lastTime = time.Now().Unix()
	qps      = 0
)

func init() {
	ChanRPC.Register("f0", func(args []interface{}) (err error) {
		log.Debug("f0 is call")
		return
	})

	ChanRPC.Register("f1", func(args []interface{}) (interface{}, error) {
		log.Debug("f1 is call")
		return 1, nil
	})

	ChanRPC.Register("fn", func(args []interface{}) ([]interface{}, error) {
		log.Debug("fn is call")
		return []interface{}{1, 2, 3}, nil
	})

	ChanRPC.Register("add", func(args []interface{}) (interface{}, error) {
		log.Debug("add is call")
		n1 := args[0].(int)
		n2 := args[1].(int)
		return n1 + n2, nil
	})

	ChanRPC.Register("qpsTest", func(args []interface{}) (interface{}, error) {
		qps += 1
		curTime := time.Now().Unix()
		if curTime - lastTime >= 1 {
			log.Debug("qps %v", qps)
			qps = 0
			lastTime = curTime
		}
		return nil, nil
	})
}