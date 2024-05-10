module github.com/PhilGruber/dimmy

go 1.21

require github.com/eclipse/paho.mqtt.golang v1.4.3

require (
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
)

replace github.com/PhilGruber/dimmy/devices => ./devices
replace github.com/PhilGruber/dimmy/core => ./core
