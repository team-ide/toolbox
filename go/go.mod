module toolbox

go 1.15

replace (
	base => ./base
	config => ./config
	database_ => ./database
	elasticsearch => ./elasticsearch
	kafka => ./kafka
	mongo => ./mongo
	redis => ./redis
	web => ./web
	worker => ./worker
	zookeeper => ./zookeeper
)

require (
	base v0.0.0-00010101000000-000000000000
	config v0.0.0-00010101000000-000000000000
	database_ v0.0.0-00010101000000-000000000000
	elasticsearch v0.0.0-00010101000000-000000000000
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	kafka v0.0.0-00010101000000-000000000000
	mongo v0.0.0-00010101000000-000000000000
	redis v0.0.0-00010101000000-000000000000
	web v0.0.0-00010101000000-000000000000
	worker v0.0.0-00010101000000-000000000000
	zookeeper v0.0.0-00010101000000-000000000000
)
