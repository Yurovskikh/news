package service

import (
	"context"
	"fmt"
	"github.com/Yurovskikh/news/api/config"
	"github.com/Yurovskikh/news/pb"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	"log"
	"time"
)

type NewsService interface {
	//
	Create(ctx context.Context, news *pb.News) (*pb.News, error)
	//
	GetById(ctx context.Context, id uint64) (*pb.News, error)
}

type newsService struct {
	conn *nats.Conn
}

func NewNewsService(cfg *config.Config) NewsService {
	conn, err := nats.Connect(fmt.Sprintf("nats://%s:%d", cfg.NatsHost, cfg.NatsPort))
	if err != nil {
		log.Fatal(err)
	}
	return &newsService{
		conn: conn,
	}
}

func (s *newsService) Create(ctx context.Context, news *pb.News) (*pb.News, error) {
	data, err := proto.Marshal(&pb.CreateNewsRequest{News: news})
	if err != nil {
		return nil, err
	}
	msg, err := s.conn.Request("News.Create", data, 500*time.Millisecond)
	if err != nil {
		return nil, err
	}
	var resp pb.CreateNewsResponse
	err = proto.Unmarshal(msg.Data, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Sucess {
		return nil, fmt.Errorf(resp.Err)
	}
	return resp.News, nil
}

func (s *newsService) GetById(ctx context.Context, id uint64) (*pb.News, error) {
	data, err := proto.Marshal(&pb.GetNewsByIdRequest{Id: id})
	if err != nil {
		return nil, err
	}
	msg, err := s.conn.Request("News.GetById", data, 500*time.Millisecond)
	if err != nil {
		return nil, err
	}
	var resp pb.GetNewsByIdResponse
	err = proto.Unmarshal(msg.Data, &resp)
	if err != nil {
		return nil, err
	}
	if !resp.Sucess {
		return nil, fmt.Errorf(resp.Err)
	}
	return resp.News, nil
}
