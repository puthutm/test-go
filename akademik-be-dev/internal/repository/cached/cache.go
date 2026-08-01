package cached

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	msg "unsia.ac.id/akademic_be/internal/config/message"
)

type CacheRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
	SetDefaultEx(key string, entry []byte) error
	Delete(key ...string) error
	SetEx(key string, entry []byte, expired time.Duration) error
	PublishMessage(channel, message string) error
	SubscribeToChannel(ctx context.Context, log *logrus.Logger, channel string)
	DeleteKeysByPattern(pattern string) error
}

type RedisCacheRepository struct {
	client *redis.Client
}

func NewRedisCacheRepository(
	client *redis.Client,
) CacheRepository {
	return &RedisCacheRepository{
		client: client,
	}
}

func (r *RedisCacheRepository) Get(key string) ([]byte, error) {
	val, err := r.client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, msg.ErrCacheMiss
	} else if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r *RedisCacheRepository) Set(key string, entry []byte) error {
	return r.client.Set(context.Background(), key, entry, 0).Err()
}

func (r *RedisCacheRepository) SetDefaultEx(key string, entry []byte) error {
	return r.client.Set(context.Background(), key, entry, 15*time.Minute).Err()
}

func (r *RedisCacheRepository) Delete(key ...string) error {
	err := r.client.Del(context.Background(), key...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCacheRepository) SetEx(key string, entry []byte, expired time.Duration) error {
	return r.client.Set(context.Background(), key, entry, expired).Err()
}

func (r *RedisCacheRepository) PublishMessage(channel, message string) error {
	err := r.client.Publish(context.Background(), channel, message).Err()
	if err != nil {
		log.Println("Error publishing message:", err)
	}
	return err
}

func (r *RedisCacheRepository) SubscribeToChannel(ctx context.Context, log *logrus.Logger, channel string) {
	pubsub := r.client.Subscribe(ctx, channel)

	run := true
	for run {
		select {
		case <-ctx.Done():
			log.Info("Shutting down subscription gracefully")
			run = false

		case msg, ok := <-pubsub.Channel():
			if !ok {
				log.Warn("Channel closed unexpectedly")
				run = false
			}
			fmt.Println("Received message:", msg.Payload)
		}
	}

	log.Infof("Closing consumer for channel : %s", channel)
	err := pubsub.Close()
	if err != nil {
		panic(err)
	}
}

func (r *RedisCacheRepository) DeleteKeysByPattern(pattern string) error {
	var cursor uint64
	keysToDelete := []string{}

	ctx := context.Background()
	for {
		keys, newCursor, err := r.client.Scan(ctx, cursor, pattern, 0).Result()
		if err != nil {
			return err
		}

		keysToDelete = append(keysToDelete, keys...)

		cursor = newCursor

		if cursor == 0 {
			break
		}
	}

	if len(keysToDelete) > 0 {
		_, err := r.client.Del(ctx, keysToDelete...).Result()
		if err != nil {
			return err
		}
		fmt.Println("Deleted keys:", keysToDelete)
	}

	return nil
}
