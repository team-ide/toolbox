module kafka

go 1.15

replace (
	base => ../base
	worker => ../worker
)

require (
	base v0.0.0-00010101000000-000000000000
	github.com/Shopify/sarama v1.29.1
	worker v0.0.0-00010101000000-000000000000
)
