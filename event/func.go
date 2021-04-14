package event

import "google.golang.org/grpc/grpclog"

var log = grpclog.Component("event")

func (eb *EventBus) Subscribe(topic string, handler Handler) {
	eb.rm.Lock()
	ch := make(chan DataEvent)
	go func(dc DataChannel) {
		de := <-dc
		log.Infof("receive the event %s ,data: %s", topic, de)
		handler(de.Data)
	}(ch)
	if prev, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(prev, ch)
	} else {
		eb.subscribers[topic] = append([]DataChannel{}, ch)
	}
	eb.rm.Unlock()
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[topic]; found {
		channels := append(DataChannelSlice{}, chans...)
		go func(data DataEvent, dataChannelSlices DataChannelSlice) {
			for _, ch := range dataChannelSlices {
				ch <- data
			}
		}(DataEvent{Data: data, Topic: topic}, channels)
	}
	eb.rm.RUnlock()
}

var GlobalEventBus = &EventBus{subscribers: map[string]DataChannelSlice{}}
