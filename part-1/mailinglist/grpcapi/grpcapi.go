package grpcapi

import (
	"context"
	"database/sql"
	"log"
	"mailinglist/mdb"
	"mailinglist/proto"
	"net"
	"time"

	"google.golang.org/grpc"
)

type MailServer struct {
	proto.UnimplementedMailingListServiceServer
	db *sql.DB
}

func fromProto(pbEntry *proto.Subscriber) mdb.Subscriber {
	return mdb.Subscriber{
		Id:          pbEntry.Id,
		Email:       pbEntry.Email,
		ConfirmedAt: time.Unix(pbEntry.ConfirmedAt, 0),
		OptOut:      pbEntry.OptOut,
	}
}

func returnProto(mdbEntry *mdb.Subscriber) proto.Subscriber {
	return proto.Subscriber{
		Id:          mdbEntry.Id,
		Email:       mdbEntry.Email,
		ConfirmedAt: mdbEntry.ConfirmedAt.Unix(),
		OptOut:      mdbEntry.OptOut,
	}
}

func subscriberResponce(db *sql.DB, email string) (*proto.SubscriberResponse, error) {
	entry, err := mdb.GetSubscriber(db, email)
	if err != nil {
		return &proto.SubscriberResponse{}, err
	}
	if entry == nil {
		return &proto.SubscriberResponse{}, nil
	}

	res := returnProto(entry)

	return &proto.SubscriberResponse{Subscriber: &res}, nil
}

func (s *MailServer) GetSubscriber(ctx context.Context, req *proto.GetSubscriberRequest) (*proto.SubscriberResponse, error) {
	log.Printf("gRPC GetSubscriber: %v\n", req)
	return subscriberResponce(s.db, req.EmailAddr)
}

func (s *MailServer) CreateSubscriber(ctx context.Context, req *proto.CreateSubscriberRequest) (*proto.SubscriberResponse, error) {
	log.Printf("gRPC CreateSubscriber: %v\n", req)

	if err := mdb.CreateSubscriber(s.db, req.EmailAddr); err != nil {
		return nil, err
	}
	return subscriberResponce(s.db, req.EmailAddr)
}

func (s *MailServer) UpdateSubscriber(ctx context.Context, req *proto.UpdateSubscriberRequest) (*proto.SubscriberResponse, error) {
	log.Printf("gRPC UpdateSubscriber: %v\n", req)
	entry := fromProto(req.Subscriber)
	if err := mdb.UpdateSubscriber(s.db, entry); err != nil {
		return &proto.SubscriberResponse{}, err
	}
	return subscriberResponce(s.db, entry.Email)
}

func (s *MailServer) CancelSubscription(ctx context.Context, req *proto.CancelSubscriptionRequest) (*proto.SubscriberResponse, error) {
	log.Printf("gRPC CancelSubscription: %v\n", req)
	if err := mdb.CancelSubscription(s.db, req.EmailAddr); err != nil {
		return &proto.SubscriberResponse{}, err
	}
	return subscriberResponce(s.db, req.EmailAddr)
}

func Serve(db *sql.DB, port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("gRPC server error: failure to bind %v\n", port)
	}

	grpcServer := grpc.NewServer()
	mailServer := MailServer{db: db}

	proto.RegisterMailingListServiceServer(grpcServer, &mailServer)

	log.Printf("gRPC API server listening on %v\n", port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC server error: %v\n", err)
	}
}
