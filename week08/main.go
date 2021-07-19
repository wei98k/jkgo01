package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jmcvetta/randutil"
	"golang.org/x/sync/errgroup"
	"strconv"
	"strings"
	"time"
)


func main() {

	ctx := context.Background()
	rdb := CreateClient()
	rdb.FlushAll(ctx)

	m1, err := GetMemory(ctx, rdb)
	if err != nil {
		panic(err)
	}
	option := &Option{
		Num:  10000, // 数量
		Size: 100, // 容量大小
	}
	err = CreateValue(ctx, rdb, option)
	if err != nil {
		panic(err)
	}
	m2, err := GetMemory(ctx, rdb)
	if err != nil {
		//t.Fatal(err)
	}
	fmt.Printf("key：%d ,key size：%d \n", option.Num, option.Size)
	res := Contrast(m1, m2, option)
	for _, v := range res {
		fmt.Printf("norm：%s,total： %d byte,avg：%d \n", v.Name, v.Total, v.Avg)
	}
    
}

type Option struct {
	Num  int
	Size int
}

// CreateValue 生成value值的大小
func CreateValue(ctx context.Context, rdb *redis.Client, option *Option) error {

	eg, ctx := errgroup.WithContext(ctx)
	keyPrefix, err := randutil.AlphaString(6)
	if err != nil {
		return err
	}
	for i := 0; i <= option.Num; i++ {
		key := fmt.Sprintf("%s-%d", keyPrefix, i)
		eg.Go(func() error {
			val, err := randutil.AlphaString(option.Size)
			if err != nil {
				return err
			}
			return rdb.Set(ctx, key, val, time.Minute).Err()
		})
	}
	return eg.Wait()
}

// GetMemory 获取内存信息
func GetMemory(ctx context.Context, rdb *redis.Client) (map[string]int, error) {
	memory := rdb.Info(ctx, "memory")
	ret := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(memory.String()))
	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), ":")
		if len(sp) != 2 {
			continue
		}
		switch sp[0] {
		case "used_memory":
			fallthrough
		case "used_memory_rss":
			//fallthrough
			//case "used_memory_peak":
			parseInt, err := strconv.ParseInt(sp[1], 10, 64)
			if err != nil {
				return nil, err
			}
			ret[sp[0]] = int(parseInt)
		}
	}
	if err := scanner.Err(); err != nil {
		return ret, err
	}
	return ret, nil
}

type Result struct {
	Name  string `json:"name"`
	Total int    `json:"total"`
	Avg   int    `json:"avg"`
}

func Contrast(before map[string]int, after map[string]int, option *Option) []Result {
	result := make([]Result, 0)
	for k, v := range before {
		total := after[k] - v
		result = append(result, Result{
			Name:  k,
			Total: total,
			Avg:   total / option.Num,
		})
	}
	return result
}

func CreateClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "192.168.2.27:6379",
		Password: "",
		DB:       0,
	})
}


