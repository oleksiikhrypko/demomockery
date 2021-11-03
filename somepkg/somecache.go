package somepkg

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"demomockery/somepkg/models"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type Cache struct {
	cl *redis.Client
}

func NewCache(cl *redis.Client) *Cache {
	return &Cache{
		cl: cl,
	}
}

func (c *Cache) GetData(key int) ([]models.DataRec, error) {
	// get data
	rkey := fmt.Sprintf("pfx_%d", key)
	data, err := c.cl.HGetAll(context.Background(), rkey).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}
	// build result
	res := make([]models.DataRec, 0, len(data))
	for k, v := range data {
		id, idErr := strconv.Atoi(k)
		if idErr != nil {
			return nil, errors.Wrap(idErr, "invalid id")
		}
		res = append(res, models.DataRec{
			ID:  id,
			Val: v,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].ID < res[j].ID
	})
	return res, nil
}
