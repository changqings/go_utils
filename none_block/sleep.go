package none_block

import "time"

// Sleep n seconds
func Sleep(n int) {
	<-time.After(time.Duration(n) * time.Second)
}
