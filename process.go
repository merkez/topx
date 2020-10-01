package usage

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	ps "github.com/shirou/gopsutil/process"
)

type Process struct {
	Id              int32
	Name            string
	Swap            uint64
	VMS             int64
	NumberOfThreads int32
	MemPercentage   float32
	CreateTime      time.Time
	User            string
}

func GetAll() []*Process {
	var pr *Process
	var prList []*Process
	processors, err := ps.Processes()
	if err != nil {
		log.Printf("Failed to get processors %v", err)
	}
	for _, p := range processors {

		m, err := p.MemoryInfo()
		if err != nil {
			log.Printf("%v", err)
		}

		name, _ := p.Name()
		cT, _ := p.CreateTime()
		memPercentage, _ := p.MemoryPercentWithContext(context.Background())
		user, _ := p.Username()
		numThreads, _ := p.NumThreads()

		timestamp := time.Unix(0, cT*int64(time.Millisecond))
		pr = &Process{
			Id:              p.Pid,
			Name:            name,
			Swap:            m.Swap,
			NumberOfThreads: numThreads,
			VMS:             int64(m.VMS),
			MemPercentage:   memPercentage,
			CreateTime:      timestamp,
			User:            user,
		}

		prList = append(prList, pr)
	}
	sort.Slice(prList, func(i, j int) bool {
		return prList[i].VMS > prList[j].VMS
	})
	return prList
}

// ByteCountSI converts from kilobytes to Megabytes
func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}
