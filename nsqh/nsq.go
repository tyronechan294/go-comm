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
