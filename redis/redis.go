package redis

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type Cache struct {
	conn redis.Conn
}

func (this *Cache) Open(protocol string, host string, port int64) {
	var err error
	this.conn, err = redis.Dial(protocol, fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Println(err)
	}
}

func (this *Cache) Set(key string, data interface{}) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	value, err := this.encode(data)
	if err != nil {
		return err
	}
	_, err = this.conn.Do("SET", key, value)
	return err
}

func (this *Cache) Get(key string, to interface{}) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	data, err := this.conn.Do("GET", key)
	if err == nil && data != nil {
		this.decode(data.([]byte), to)
	}
	return err
}

func (this *Cache) Hset(key string, field string, data interface{}) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	value, err := this.encode(data)
	if err != nil {
		return err
	}
	_, err = this.conn.Do("HSET", key, field, value)
	return err
}

func (this *Cache) Hget(key string, field string, to interface{}) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	data, err := this.conn.Do("HGET", key, field)
	if err == nil && data != nil {
		this.decode(data.([]byte), to)
	}
	return err
}

func (this *Cache) Del(key string) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	_, err := this.conn.Do("DEL", key)
	return err
}

func (this *Cache) Hdel(key string, field string) error {
	if this.conn == nil {
		return errors.New("Connect Redis First!")
	}
	_, err := this.conn.Do("HDEL", key, field)
	return err
}

func (this *Cache) encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (this *Cache) decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
