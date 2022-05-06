module github.com/encima/dgen/redis

go 1.16

require (
	github.com/encima/dgen/lib v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.5
	github.com/joho/godotenv v1.4.0
	github.com/myesui/uuid v1.0.0 // indirect
	github.com/prometheus/common v0.34.0 // indirect
	github.com/twinj/uuid v1.0.0
	gopkg.in/stretchr/testify.v1 v1.2.2 // indirect
)

replace github.com/encima/dgen/lib => ../lib
