package cluster

import "project/internal/domain/house"

type Cluster struct {
	ID     string
	Name   string
	Houses []house.House
}
