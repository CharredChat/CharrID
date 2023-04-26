package CharredProcess

import (
	"time"

	"github.com/CharredChat/charrid/CharrID"
)

type CharredProcess struct {
	ProcessID uint8
	LastTime  time.Time
	Incr      uint8
}

func New(id uint8) CharredProcess {
	return CharredProcess{id, time.Now(), 0}
}

func (pr *CharredProcess) Generate() (CharrID.CharrID, error) {
	var since = time.Since(pr.LastTime)
	if since.Milliseconds() == 0 {
		if pr.Incr == 255 {
			time.Sleep(time.Millisecond)
			pr.Incr = 0
		} else {
			pr.Incr++
		}
	} else {
		pr.Incr = 0
	}
	pr.LastTime = time.Now()
	return CharrID.New(pr.ProcessID, pr.Incr)
}
