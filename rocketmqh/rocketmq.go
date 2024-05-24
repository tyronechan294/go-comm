package rocketmqh

import (
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tyronechan294/go-comm/valueh"
)

type DelayTimeLevel int

const (
	DELAY_ONE_SECOND DelayTimeLevel = iota + 1
	DELAY_FIVE_SECONDS
	DELAY_TEN_SECONDS
	DELAY_THIRTY_SECONDS
	DELAY_ONE_MINUTE
	DELAY_TWO_MINUTES
	DELAY_THREE_MINUTES
	DELAY_FOUR_MINUTES
	DELAY_FIVE_MINUTES
	DELAY_SIX_MINUTES
	DELAY_SEVEN_MINUTES
	DELAY_EIGHT_MINUTES
	DELAY_NINE_MINUTES
	DELAY_TEN_MINUTES
	DELAY_TWENTY_MINUTES
	DELAY_THIRTY_MINUTES
	DELAY_ONE_HOUR
	DELAY_TWO_HOURS
)

func (x DelayTimeLevel) Number() int {
	return int(x)
}

func MessageWithTag(topic string, tag string, body any) (*primitive.Message, error) {
	b, err := json.Marshal(body)
	if valueh.IsNotNil(err) {
		return nil, err
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  b,
	}
	msg.WithTag(tag)

	return msg, nil
}

// 发送
func SendSyncMessageWithTag(rocketProducer rocketmq.Producer, topic string, tag string, body any) (*primitive.SendResult, error) {
	msg, err := MessageWithTag(topic, tag, body)
	if valueh.IsNotNil(err) {
		return nil, err
	}

	return rocketProducer.SendSync(gctx.New(), msg)
}

////////////////

// 不发送 延时
func MessageWithTagAndDelayTimeLevel(topic string, tag string, level DelayTimeLevel, body any) (*primitive.Message, error) {
	b, err := json.Marshal(body)
	if valueh.IsNotNil(err) {
		return nil, err
	}

	msg := &primitive.Message{
		Topic: topic,
		Body:  b,
	}
	msg.WithTag(tag)
	msg.WithDelayTimeLevel(level.Number())

	return msg, nil
}

// 发送延时
func SendSyncMessageWithTagAndDelayTimeLevel(rocketProducer rocketmq.Producer, topic string, tag string, level DelayTimeLevel, body any) (*primitive.SendResult, error) {
	msg, err := MessageWithTagAndDelayTimeLevel(topic, tag, level, body)
	if valueh.IsNotNil(err) {
		return nil, err
	}

	return rocketProducer.SendSync(gctx.New(), msg)
}
