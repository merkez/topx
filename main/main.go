package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mrturkmencom/topx"
)

func main() {
	top := flag.Int("top", 10, "usage -top=10")
	flag.Parse()
	prList := usage.GetAll()
	if len(prList) < *top {
		fmt.Printf("Error: There is no that amount of processes \nTotal num of processes is: %d\n", len(prList))
		return
	}
	prList = prList[:*top]
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 2, '\t', 0)
	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "User", "PID", "Name", "VMemory", "MemPerentage", "NumThreads", "Swap", "CreateTime")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t", "-----", "-----", "-----", "-----", "------", "-----", "-------------", "----------")
	for _, pr := range prList {
		fmt.Fprintf(w, "\n %s\t%d\t%s\t%s\t%f\t%d\t%d\t%s\t", pr.User, pr.Id, pr.Name, usage.ByteCountSI(pr.VMS), pr.MemPercentage, pr.NumberOfThreads, pr.Swap, pr.CreateTime)
	}
	fmt.Fprintf(w, "\n")
}
