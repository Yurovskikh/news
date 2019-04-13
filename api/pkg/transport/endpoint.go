package transport

import (
	"context"
	"github.com/Yurovskikh/news/api/pkg/service"
	"github.com/Yurovskikh/news/pb"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"time"
)

type Endpoints struct {
	CreateNews  endpoint.Endpoint
	GetNewsById endpoint.Endpoint
}

func MakeEndpoints(newsService service.NewsService, logger log.Logger) Endpoints {
	var createNews endpoint.Endpoint
	{
		createNews = makeCreateNews(newsService, logger)
		// todo validation middleware
	}
	var getNewsById endpoint.Endpoint
	{
		getNewsById = makeGetNewsById(newsService, logger)
		// todo validation middleware
	}
	return Endpoints{
		CreateNews:  createNews,
		GetNewsById: getNewsById,
	}
}

type (
	CreateNewsRequest struct {
		Header string    `json:"header"`
		Date   time.Time `json:"date"`
	}
	CreateNewsResponse struct {
		Id     uint64    `json:"id"`
		Header string    `json:"header"`
		Date   time.Time `json:"date"`
	}
)

func makeCreateNews(newsService service.NewsService, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger := log.With(logger, "method", "CreateNews")
		req := request.(CreateNewsRequest)
		level.Info(logger).Log("req", req)
		news, err := newsService.Create(ctx, &pb.News{
			Header: req.Header,
			Date:   req.Date.Unix(),
		})
		if err != nil {
			level.Error(logger).Log("err", err)
			return nil, err
		}
		return CreateNewsResponse{
			Id:     news.Id,
			Header: news.Header,
			Date:   time.Unix(news.Date, 0),
		}, nil
	}
}

type (
	GetNewsByIdRequest struct {
		Id uint64
	}
	GetNewByIdResponse struct {
		Id     uint64    `json:"id"`
		Header string    `json:"header"`
		Date   time.Time `json:"date"`
	}
)

func makeGetNewsById(newsService service.NewsService, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logger := log.With(logger, "method", "GetNewsById")
		req := request.(GetNewsByIdRequest)
		level.Info(logger).Log("req", req)
		news, err := newsService.GetById(ctx, req.Id)
		if err != nil {
			level.Error(logger).Log("err", err)
			return nil, err
		}
		return GetNewByIdResponse{
			Id:     news.Id,
			Header: news.Header,
			Date:   time.Unix(news.Date, 0),
		}, nil
	}
}
