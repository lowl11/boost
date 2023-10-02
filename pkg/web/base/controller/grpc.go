package controller

import (
	"github.com/lowl11/boost/pkg/interfaces"
	"github.com/lowl11/lazylog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Grpc struct {
	//
}

func (controller Grpc) Error(err error) error {
	if err == nil {
		return nil
	}

	log.Error(err)
	boostError, ok := err.(interfaces.Error)
	if !ok {
		return status.Error(codes.Internal, err.Error())
	}

	return status.Error(boostError.GrpcCode(), boostError.Error())
}

func (controller Grpc) NotFound(err error) error {
	if err == nil {
		return nil
	}

	log.Error(err)
	boostError, ok := err.(interfaces.Error)
	if !ok {
		return status.Error(codes.Internal, err.Error())
	}

	return status.Error(codes.NotFound, boostError.Error())
}
