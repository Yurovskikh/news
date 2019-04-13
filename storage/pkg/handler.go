package pkg

import (
	"context"
	"fmt"
	"github.com/Yurovskikh/news/pb"
	"github.com/Yurovskikh/news/storage/config"

	"github.com/Yurovskikh/news/storage/pkg/model"
	"github.com/Yurovskikh/news/storage/pkg/service"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
	stdlog "log"
	"time"
)

type MsgHandler interface {
	Start()
}
type msgHandler struct {
	nc          *nats.Conn
	newsService service.NewsService
	logger      log.Logger
}

func NewMsgHandler(cfg *config.Config, newsService service.NewsService, logger log.Logger) MsgHandler {
	nc, err := nats.Connect(fmt.Sprintf("nats://%s:%d", cfg.NatsHost, cfg.NatsPort))
	if err != nil {
		stdlog.Fatal(err)
	}
	return &msgHandler{
		nc:          nc,
		newsService: newsService,
		logger:      log.With(logger, "service", "MsgHandler"),
	}
}

func (m *msgHandler) Start() {
	m.startSubCreateNews()
	m.startSubGetNewsById()
}

func (m *msgHandler) startSubCreateNews() {
	m.nc.Subscribe("News.Create", func(msg *nats.Msg) {
		logger := log.With(m.logger, "method", "SubCreateNews")
		level.Info(logger).Log("message processing", msg)
		var req pb.CreateNewsRequest
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		create, err := m.newsService.Create(context.Background(), &model.News{
			Header: req.News.Header,
			Date:   time.Unix(req.News.Date, 0),
		})
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		data, err := proto.Marshal(&pb.CreateNewsResponse{
			News: &pb.News{
				Id:     uint64(create.ID),
				Header: create.Header,
				Date:   create.Date.Unix(),
			},
			Sucess: true,
		})
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		err = m.nc.Publish(msg.Reply, data)
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		level.Info(logger).Log("message processed", msg)
	})
}

func (m *msgHandler) startSubGetNewsById() {
	m.nc.Subscribe("News.GetById", func(msg *nats.Msg) {
		logger := log.With(m.logger, "method", "SubGetNewsById")
		level.Info(logger).Log("message processing", msg)
		var req pb.GetNewsByIdRequest
		err := proto.Unmarshal(msg.Data, &req)
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		news, err := m.newsService.GetById(context.Background(), uint(req.Id))
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		data, err := proto.Marshal(&pb.GetNewsByIdResponse{
			News: &pb.News{
				Id:     uint64(news.ID),
				Header: news.Header,
				Date:   news.Date.Unix(),
			},
			Sucess: true,
		})
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		err = m.nc.Publish(msg.Reply, data)
		if err != nil {
			level.Error(logger).Log("err", err)
			return
		}
		level.Info(logger).Log("message processed", msg)
	})
}
