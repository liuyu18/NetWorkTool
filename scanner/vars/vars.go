package vars

import "sync"

var (
	ThreadNum = 5
	Result    *sync.Map
)

func init() {
	Result = &sync.Map{}
}
