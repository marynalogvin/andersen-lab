package main

import (
	"context"
	"log"
	"mailinglist/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func logResponce(res *proto.SubscriberResponse, err error) {
	if err != nil {
		log.Fatalf("error:%v", err)
	}

	if res.Subscriber == nil {
		log.Printf("email not found")

	}
	log.Printf("responce:%v", res.Subscriber)
}

func createSubscriber(client proto.MailingListServiceClient, email string) *proto.Subscriber {
	log.Println("create subscriber")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CreateSubscriber(ctx, &proto.CreateSubscriberRequest{EmailAddr: email})
	logResponce(res, err)
	return res.Subscriber
}

func getSubscriber(client proto.MailingListServiceClient, addr string) *proto.Subscriber {
	log.Println("get subscriber")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetSubscriber(ctx, &proto.GetSubscriberRequest{EmailAddr: addr})
	logResponce(res, err)
	return res.Subscriber
}

func updateSubscriber(client proto.MailingListServiceClient, sub *proto.Subscriber) *proto.Subscriber {
	log.Println("update subscriber")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.UpdateSubscriber(ctx, &proto.UpdateSubscriberRequest{Subscriber: sub})
	logResponce(res, err)
	return res.Subscriber
}

func cancelSubscription(client proto.MailingListServiceClient, addr string) *proto.Subscriber {
	log.Println("cancel subscription")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.CancelSubscription(ctx, &proto.CancelSubscriptionRequest{EmailAddr: addr})
	logResponce(res, err)
	return res.Subscriber
}

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewMailingListServiceClient(conn)

	newSubscriber := createSubscriber(client, "first@com")
	newSubscriber.ConfirmedAt = 10000
	updateSubscriber(client, newSubscriber)
	cancelSubscription(client, newSubscriber.Email)
}
