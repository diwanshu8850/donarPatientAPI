package main

import (
	"context"
	"log"
	"time"

	pb "grpc_donar_patient/Protos"

	"google.golang.org/grpc"
)

const (
	address = "localhost:3000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDonarPatientServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateUser(ctx, &pb.Record{
		Name:               "Diw",
		Address:            "SomeRandom",
		EmailId:            "some@gmail.com",
		UserType:           "Donar",
		PhoneNo:            "123456",
		DiseaseDescription: "some random disease"})
	if err != nil {
		log.Fatalf("could not creat user: %v", err)
	}
	log.Printf("User 1: %s\n", r)
	//
	//r1, err1 := c.CreateUser(ctx, &pb.Record{
	//	Name:     "Diwanshu",
	//	Address:  "Random",
	//	EmailId:  "random@gmail.com",
	//	UserType: "Patient",
	//	PhoneNo:  "555555555"})
	//if err1 != nil {
	//	log.Fatalf("could not creat user: %v", err1)
	//}
	//log.Printf("User 2: %s\n", r1)
	//
	//r1, err1 := c.GetUser(ctx, &pb.Request{
	//	YourId: "1",
	//	SecretCode: "QWertY@966#Ub1",
	//	UserId: "2"})
	//if err1 != nil {
	//	log.Fatalf("could not get user: %v", err1)
	//}
	//log.Printf("User: %s\n", r1)
	//r2, err2 := c.SendRequest(ctx, &pb.Request{
	//	YourId: "1",
	//SecretCode: "QWertY@966#Ub1",
	//UserId: "2"})
	//if err2 != nil {
	//	log.Fatalf("could not send request: %v", err2)
	//}
	//log.Printf("Request Status:  %s\n", r2)
	//
	//r3, err3 := c.LoginUser(ctx, &pb.Record{Id: "1", SecretCode: "QWertY@966#Ub1"})
	//if err3 != nil {
	//	log.Fatalf("could not login user: %v", err3)
	//}
	//log.Printf("User: %s", r3)
	//
	//r5, err5 := c.LoginUser(ctx, &pb.Record{Id: "2", SecretCode: "QWertY@966#Ub2"})
	//if err5 != nil {
	//	log.Fatalf("could not login user: %v", err5)
	//}
	//log.Printf("User: %s", r5)
	//
	//r4, err4 := c.CancelRequest(ctx, &pb.Request{
	//	YourId: "2",
	//	SecretCode: "QWertY@966#Ub2",
	//	UserId: "1"})
	//if err4 != nil {
	//	log.Fatalf("could not cancel request: %v", err4)
	//}
	//log.Printf("Cancel Status: %s", r4)
	//
	//r6, err6 := c.LoginUser(ctx, &pb.Record{Id: "1", SecretCode: "QWertY@966#Ub1"})
	//if err6 != nil {
	//	log.Fatalf("could not login user: %v", err6)
	//}
	//log.Printf("User: %s", r6)
	//
	//r7, err7 := c.SendRequest(ctx, &pb.Request{
	//	YourId: "1",
	//	SecretCode: "QWertY@966#Ub1",
	//	UserId: "2"})
	//if err7 != nil {
	//	log.Fatalf("could not send request: %v", err7)
	//}
	//log.Printf("Request Status:  %s\n", r7)
	//
	//r8, err8 := c.AcceptRequest(ctx, &pb.Request{
	//	YourId: "2",
	//	SecretCode: "QWertY@966#Ub2",
	//	UserId: "1"})
	//if err8 != nil {
	//	log.Fatalf("could not creat user: %v", err8)
	//}
	//log.Printf("Request Status:  %s\n", r8)
	//
	//r9, err9 := c.LoginUser(ctx, &pb.Record{Id: "1", SecretCode: "QWertY@966#Ub1"})
	//if err9 != nil {
	//	log.Fatalf("could not login user: %v", err9)
	//}
	//log.Printf("User: %s", r9)
	//
	//r9, err9 = c.LoginUser(ctx, &pb.Record{Id: "2", SecretCode: "QWertY@966#Ub2"})
	//if err9 != nil {
	//	log.Fatalf("could not login user: %v", err9)
	//}
	//log.Printf("User: %s", r9)
	//
	//r2, err2 := c.CancelConnection(ctx, &pb.Request{
	//	YourId: "1",
	//	SecretCode: "QWertY@966#Ub1",
	//	UserId: "2"})
	//if err2 != nil {
	//	log.Fatalf("could not cancel connection: %v", err2)
	//}
	//log.Printf("Connection Cancel Stutus: %s", r2)
	//
	//r9, err9 = c.LoginUser(ctx, &pb.Record{Id: "1", SecretCode: "QWertY@966#Ub1"})
	//if err9 != nil {
	//	log.Fatalf("could not login user: %v", err9)
	//}
	//log.Printf("User: %s", r9)
	//
	//r9, err9 = c.LoginUser(ctx, &pb.Record{Id: "2", SecretCode: "QWertY@966#Ub2"})
	//if err9 != nil {
	//	log.Fatalf("could not login user: %v", err9)
	//}
	//log.Printf("User: %s", r9)
	//
	//r1, err1 := c.GetPatients(ctx, &pb.Record{Id: "1", SecretCode: "QWertY@966#Ub1"})
	//if err1 != nil {
	//	log.Fatalf("could not get all patients: %v", err1)
	//}
	//log.Printf("All Patients: %s", r1)
	//
	//r1, err1 := c.GetDonars(ctx, &pb.Record{Id: "2", SecretCode: "QWertY@966#Ub2"})
	//if err1 != nil {
	//	log.Fatalf("could not get all donars: %v", err1)
	//}
	//log.Printf("All Donars: %s", r1)
}
