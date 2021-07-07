package main

import (
	"log"

	"github.com/quadtrix/zanecluster/backend/cluster"
	"github.com/quadtrix/zanecluster/backend/types"
)

func main() {
	log.Println("ZaneCluster Cluster Manager starting up...")
	cluster, err := cluster.New(true, "", types.Network{})
}
