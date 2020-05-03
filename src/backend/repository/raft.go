package repository

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func Put(key string,value string,ctx context.Context) error{
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	_, puterr := cli.Put(ctx, key, value)
	if puterr != nil {
		log.Fatal(err)
		return puterr
	}else {
		return nil
	}

}

func Get(key string,ctx context.Context) ([]byte,error){
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	response, geterr := cli.Get(ctx, key)
	if geterr != nil {
		log.Fatal(err)
		return []byte(""),geterr
	}else {
		if len(response.Kvs) == 0{
			return []byte(""),nil
		}else{
			return response.Kvs[0].Value,nil
		}

	}

}

func GetAll(ctx context.Context) (string,error){
	cli, err := clientv3.New(clientv3.Config{
		DialTimeout: 5 * time.Second,
		Endpoints: []string{"localhost:2379", "localhost:22379", "localhost:32379"},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()
	opts := []clientv3.OpOption{
		clientv3.WithPrefix(),
	}
	response, geterr := cli.Get(ctx,"",opts...)
	if geterr != nil {
		log.Fatal(err)
		return "",geterr
	}else {
		ans := ""
		for _, item := range response.Kvs {
			ans += string(item.Key)[1:]+","
		}
		return ans,nil
	}

}