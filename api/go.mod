module github.com/Yurovskikh/news/api

require (
	github.com/Yurovskikh/news/pb v0.0.0
	github.com/eclipse/paho.mqtt.golang v1.1.1 // indirect
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/protobuf v1.2.0
	github.com/gorilla/mux v1.7.1
	github.com/nats-io/go-nats v1.7.2
	github.com/nats-io/go-nats-streaming v0.4.2 // indirect
	github.com/nats-io/nkeys v0.0.2 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/spf13/viper v1.3.2
)

replace github.com/Yurovskikh/news/pb => ../pb
