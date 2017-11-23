package worker

import (
	"fmt"
	"strings"
)

var Datastore map[string][]*Statistic

type Statistic struct {
	ip        string
	request   string
	bytesSent string
}

func Parse(c chan string) {
	Datastore = make(map[string][]*Statistic, 2)
	for {
		select {
		case line := <-c:
			if len(line) < 1 {
				return
			}
			bySpace := strings.Split(line, " ")
			statistic := &Statistic{
				ip:        bySpace[0],
				request:   bySpace[6],
				bytesSent: bySpace[7],
			}
			stack := Datastore[bySpace[6]]
			stack = append(stack, statistic)
			Datastore[bySpace[6]] = stack
		}
	}
}

func (s *Statistic) toString() string {
	return fmt.Sprintf(`{"ip":"%s","request":"%s","bytesSent":"%s"}`, s.ip, s.request, s.bytesSent)
}
