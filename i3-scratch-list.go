package main

import (
	"flag"
	"fmt"
	"go.i3wm.org/i3/v4"
)

func main() {
	noid := flag.Bool("noid", false, "Don't give me the window id")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: i3-scratch-list [-noid]\n\n")
		flag.PrintDefaults()

	}
	flag.Parse()

	ws, _ := i3.GetTree()
	sp := ws.Root.FindChild(
		func(n *i3.Node) bool {
			return n.Name == "__i3_scratch"
		},
	)

	for _, n := range sp.FloatingNodes {
		for _, wn := range n.Nodes {
			if !*noid {
				fmt.Println(wn.Window, " - ", wn.Name)
			} else {
				fmt.Println(wn.Name)
			}
		}
	}
}
