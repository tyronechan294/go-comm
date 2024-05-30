package nsqh

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/nsqio/go-nsq"
	"github.com/tyronechan294/go-comm/valueh"
	"time"
)

func Publish(nsqProducer *nsq.Producer, topic string, body any) error {
	b, err := gjson.Marshal(body)
	if valueh.IsNotEmpty(err) {
		return err
	}

	return nsqProducer.Publish(
		topic,
		b,
	)
}

func PublishWithTag(nsqProducer *nsq.Producer, topic string, tag string, body any) error {
	return Publish(
		nsqProducer,
		topic+"_"+tag,
		body,
	)
}

func DeferredPublish(nsqProducer *nsq.Producer, topic string, delay time.Duration, body any) error {
	b, err := gjson.Marshal(body)
	if valueh.IsNotEmpty(err) {
		return err
	}

	return nsqProducer.DeferredPublish(
		topic,
		delay,
		b,
	)
}

func DeferredPublishWithTag(nsqProducer *nsq.Producer, topic string, tag string, delay time.Duration, body any) error {
	return DeferredPublish(
		nsqProducer,
		topic+"_"+tag,
		delay,
		body,
	)
}

//////////

func ConsumerConnectConcurrentHandler(
	topic string,
	channel string,
	config *nsq.Config,
	addr string,
	concurrency int,
	handler nsq.Handler,
) (consumer *nsq.Consumer, err error) {

	consumer, err = nsq.NewConsumer(topic, channel, config)
	if valueh.IsNotEmpty(err) {
		return
	}

	consumer.AddConcurrentHandlers(handler, concurrency)

	err = consumer.ConnectToNSQD(addr)

	return
}

func ConsumerConnectConcurrentHandlerWithTag(
	topic string,
	tag string,
	channel string,
	config *nsq.Config,
	addr string,
	concurrency int,
	handler nsq.Handler,
) (consumer *nsq.Consumer, err error) {
	return ConsumerConnectConcurrentHandler(
		topic+"_"+tag,
		channel,
		config,
		addr,
		concurrency,
		handler,
	)
}
