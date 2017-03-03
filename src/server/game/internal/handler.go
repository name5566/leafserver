package internal

import (
	"github.com/name5566/leaf/log"
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
		n1 := int(args[0])
		n2 := int(args[1])
		return n1 + n2, nil
	})
}