package io

const MQTT = "mqtt"

type global struct {
	Mode string
}

type mqtt struct {
	Host       string
	User       string
	Password   string
	ShareTopic string
}

type http struct {
	Host string
}

type config struct {
	Global global
	Mqtt   mqtt
	Http   http
}
