package io

const MQTT = "mqtt"
const HTTP = "http"
const REQUESTPATH = "/api/v1/occupied"

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
	Host                string
	PollIntervalSeconds int
}

type config struct {
	Global global
	Mqtt   mqtt
	Http   http
}
