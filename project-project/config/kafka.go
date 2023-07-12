package config

import (
	"context"
	"go.uber.org/zap"
	"pomfret.cn/project-project/internal/repo"
	"pomfret.cn/project_common/kk"
	"time"
)

var kw *kk.KafkaWriter

func InitKafkaWriter() func() {
	kw := kk.GetWriter("107.151.248.234:9092")
	return kw.Close
}
func SendLog(data []byte) {
	kw.Send(kk.LogData{
		Topic: "msproject_log",
		Data:  data,
	})
}

type KafkaCache struct {
	R     *kk.KafkaReader
	cache repo.Cache
}

func (c *KafkaCache) DeleteCache() {
	for {
		message, err := c.R.R.ReadMessage(context.Background())
		if err != nil {
			zap.L().Error("DeleteCache ReadMessage err", zap.Error(err))
			continue
		}
		if "task" == string(message.Value) {
			fields, err := c.cache.HKeys(context.Background(), "task")
			if err != nil {
				zap.L().Error("DeleteCache HKeys err", zap.Error(err))
				continue
			}
			time.Sleep(1 * time.Second)
			c.cache.Delete(context.Background(), fields)
		}
	}
}

func NewCacheReader() *KafkaCache {
	reader := kk.GetReader([]string{"localhost:9092"}, "cache_group", "msproject_log")
	return &KafkaCache{
		R: reader,
	}
}
