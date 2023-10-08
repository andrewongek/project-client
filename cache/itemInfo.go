package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"project-client/structs"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	itemCacheTTL = 1 * time.Hour
	keyTemplate  = "itemInfo_%d"
)

type ItemInfo struct {
	redis *redis.Client
}

func NewOrderInfo(address string) *ItemInfo {
	newRedisCli := redis.NewClient(&redis.Options{
		Addr: address,
	})

	return &ItemInfo{
		redis: newRedisCli,
	}
}

func (i *ItemInfo) GetItemInfo(itemid int64) (structs.ItemData, error) {
	res := structs.ItemData{}
	ctx := context.Background()
	val, err := i.redis.Get(ctx, fmt.Sprintf(keyTemplate, itemid)).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(val), res)
	if err != nil {
		return res, fmt.Errorf("unmarshal item info err=%s", err.Error())
	}

	return res, nil
}
