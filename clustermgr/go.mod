module github.com/quadtrix/zanecluster/clustermgr

go 1.16

replace github.com/quadtrix/zanecluster/backend/cluster => ../backend/cluster

replace github.com/quadtrix/zanecluster/backend/types => ../backend/types

require (
	github.com/quadtrix/zanecluster/backend/cluster v0.0.0-00010101000000-000000000000
	github.com/quadtrix/zanecluster/backend/types v0.0.0-00010101000000-000000000000
)
