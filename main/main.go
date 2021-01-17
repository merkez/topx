package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrturkmencom/topx"
	tw "github.com/olekukonko/tablewriter"
)

var (
	HEADER = []string{"User", "PID", "Name", "VMemory", "MemPercentage", "Swap", "CreateTime"}
)

func main() {
	top := flag.Int("top", 10, "usage -top=10")
	flag.Parse()
	var procData [][]string
	table := tw.NewWriter(os.Stdout)
	table.SetHeader(HEADER)
	prList := usage.GetAll()
	if len(prList) < *top {
		fmt.Printf("Error: There is no that amount of processes \nTotal num of processes is: %d\n", len(prList))
		return
	}
	prList = prList[:*top]
	for _, pr := range prList {
		procData = append(procData, []string{pr.User, fmt.Sprintf("%d", pr.Id), pr.Name, fmt.Sprintf("%s", usage.ByteCountSI(pr.VMS)), fmt.Sprintf("%v", pr.MemPercentage), fmt.Sprintf("%d", pr.Swap), pr.CreateTime.String()})
	}
	table.AppendBulk(procData)
	table.Render()
}
