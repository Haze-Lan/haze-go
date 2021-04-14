package event

import "sync"

type DataEvent struct {
	Data  interface{}
	Topic string
}

type Handler func(data interface{})

type DataChannel chan DataEvent

type DataChannelSlice []DataChannel

// EventBus 存储有关订阅者感兴趣的特定主题的信息
type EventBus struct {
	subscribers map[string]DataChannelSlice
	rm          sync.RWMutex
}

// EVENT_TOPIC_SERVER_QUIT 服务退出事件
const EVENT_TOPIC_SERVER_QUIT = "SERVER_QUIT"

// EVENT_TOPIC_SERVICE_CHANGE 服务实例变化事件
const EVENT_TOPIC_SERVICE_CHANGE = "SERVICE_CHANGE"
