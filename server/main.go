package main

import (
	"context"
	"log"
	"net"

	pb "grpc_donar_patient/Protos"
	"strconv"
	"sync"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDonarPatientServiceServer
}

var Users map[string]pb.Record
var id = 1

var mu sync.Mutex

const secretKey = "QWertY@966#Ub"

func (s *server) CreateUser(ctx context.Context, in *pb.Record) (*pb.Record, error) {

	user := *in

	mu.Lock()
	user.Id = strconv.Itoa(id)
	user.SecretCode = secretKey + strconv.Itoa(id)
	user.RequestedUsers = make(map[string]int32)
	user.PendingRequests = make(map[string]int32)
	user.ConnectedUsers = make(map[string]*pb.ShowUser)
	Users[strconv.Itoa(id)] = user
	id += 1
	mu.Unlock()

	return &user, nil
}

func (s *server) LoginUser(ctx context.Context, in *pb.Record) (*pb.Record, error) {
	user := *in

	_, prs := Users[user.Id]
	if prs {
		if Users[user.Id].SecretCode == user.SecretCode {
			b, _ := Users[user.Id]
			return &b, nil
		}
	}
	return nil, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.Record) (*pb.Success, error) {
	user := *in

	_, prs := Users[user.Id]
	if prs {
		if Users[user.Id].SecretCode == user.SecretCode {
			delete(Users, user.Id)
			return &pb.Success{Name: "Successfully Deleted!"}, nil
		} else {
			return &pb.Success{Name: "Invalid Code!"}, nil
		}
	} else {
		return &pb.Success{Name: "User Not Found!"}, nil
	}
}

func (s *server) GetUser(ctx context.Context, in *pb.Request) (*pb.ShowUser, error) {
	user := *in

	_, prs := Users[user.YourId]
	if prs {
		_, prs1 := Users[user.UserId]
		if prs1 {
			if Users[user.YourId].SecretCode == user.SecretCode {
				if (Users[user.YourId].UserType == "Patient" && Users[user.UserId].UserType == "Donar") ||
					(Users[user.YourId].UserType == "Donar" && Users[user.UserId].UserType == "Patient") {
					_, prs2 := Users[user.YourId].ConnectedUsers[user.UserId]
					show := pb.ShowUser{}
					if prs2 {
						show = *Users[user.YourId].ConnectedUsers[user.UserId]
					} else {
						show = pb.ShowUser{Id: user.UserId, Name: Users[user.UserId].Name}
					}
					return &show, nil
				}
			}
		}
	}
	return nil, nil
}

func (s *server) GetDonars(ctx context.Context, in *pb.Record) (*pb.RepShow, error) {

	user := *in

	_, prs := Users[user.Id]
	if prs {
		if Users[user.Id].SecretCode == user.SecretCode && Users[user.Id].UserType == "Patient" {
			show := make([]*pb.ShowUser, 0)
			for k, v := range Users {
				if v.UserType == "Donar" {
					_, prs1 := Users[user.Id].ConnectedUsers[k]
					if prs1 {
						show = append(show, Users[user.Id].ConnectedUsers[k])
					} else {
						show = append(show, &pb.ShowUser{Id: k, Name: v.Name})
					}
				}
			}
			return &pb.RepShow{User: show}, nil
		}
	}
	return nil, nil
}

func (s *server) GetPatients(ctx context.Context, in *pb.Record) (*pb.RepShow, error) {

	user := *in

	_, prs := Users[user.Id]
	if prs {
		if Users[user.Id].SecretCode == user.SecretCode && Users[user.Id].UserType == "Donar" {
			show := make([]*pb.ShowUser, 0)
			for k, v := range Users {
				if v.UserType == "Patient" {
					_, prs1 := Users[user.Id].ConnectedUsers[k]
					if prs1 {
						show = append(show, Users[user.Id].ConnectedUsers[k])
					} else {
						show = append(show, &pb.ShowUser{Id: k, Name: v.Name})
					}
				}
			}
			return &pb.RepShow{User: show}, nil
		}
	}
	return nil, nil
}

func (s *server) SendRequest(ctx context.Context, in *pb.Request) (*pb.Success, error) {

	user := *in

	_, prs := Users[user.YourId]
	if prs {
		if Users[user.YourId].SecretCode == user.SecretCode {
			if Users[user.YourId].UserType == "Donar" {
				_, prs1 := Users[user.UserId]
				if prs1 {
					if Users[user.UserId].UserType == "Patient" {
						Users[user.YourId].RequestedUsers[user.UserId] = 1
						Users[user.UserId].PendingRequests[user.YourId] = 1
						return &pb.Success{Name: "Request Sent Successfully!"}, nil
					} else {
						return &pb.Success{Name: "Invalid Request!"}, nil
					}
				} else {
					return &pb.Success{Name: "Invalid Request!"}, nil
				}
			} else {
				_, prs1 := Users[user.UserId]
				if prs1 {
					if Users[user.UserId].UserType == "Donar" {
						Users[user.YourId].RequestedUsers[user.UserId] = 1
						Users[user.UserId].PendingRequests[user.YourId] = 1
						return &pb.Success{Name: "Request Sent Successfully!"}, nil
					} else {
						return &pb.Success{Name: "Invalid Request!"}, nil
					}
				} else {
					return &pb.Success{Name: "Invalid Request!"}, nil
				}
			}
		} else {
			return &pb.Success{Name: "Invalid Code!!"}, nil
		}
	} else {
		return &pb.Success{Name: "User Not Found!"}, nil
	}
}

func (s *server) CancelRequest(ctx context.Context, in *pb.Request) (*pb.Success, error) {

	user := *in

	_, prs := Users[user.YourId]
	if prs {
		if Users[user.YourId].SecretCode == user.SecretCode {
			_,prs1 := Users[user.YourId].RequestedUsers[user.UserId]
			if prs1 {
				delete(Users[user.YourId].RequestedUsers, user.UserId)
				delete(Users[user.UserId].PendingRequests, user.YourId)
				return &pb.Success{Name: "Request Cancelled Successfully!"}, nil
			} else {
				return &pb.Success{Name: "Request Not Found!"}, nil
			}
		} else {
			return &pb.Success{Name: "Invalid Code!!"}, nil
		}
	} else{
		return &pb.Success{Name: "User Not Found!"}, nil
	}
}

func (s *server) AcceptRequest(ctx context.Context, in *pb.Request) (*pb.Success, error) {

	user := *in

	_, prs := Users[user.YourId]
	if prs {
		if Users[user.YourId].SecretCode == user.SecretCode {
			_,prs1 := Users[user.YourId].PendingRequests[user.UserId]
			if prs1 {
				delete(Users[user.UserId].RequestedUsers, user.YourId)
				delete(Users[user.YourId].PendingRequests, user.UserId)

				a := pb.ShowUser{
					Id: user.YourId,
					Name: Users[user.YourId].Name,
					Address: Users[user.YourId].Address,
					EmailId: Users[user.YourId].EmailId,
					PhoneNo: Users[user.YourId].PhoneNo}
				Users[user.UserId].ConnectedUsers[user.YourId] = &a
				b := pb.ShowUser{
					Id: user.UserId,
					Name: Users[user.UserId].Name,
					Address: Users[user.UserId].Address,
					EmailId: Users[user.UserId].EmailId,
					PhoneNo: Users[user.UserId].PhoneNo}
				Users[user.YourId].ConnectedUsers[user.UserId] = &b
				return &pb.Success{Name: "Request Accepted Successfully!"}, nil
			} else {
				return &pb.Success{Name: "Request Not Found!"}, nil
			}
		} else {
			return &pb.Success{Name: "Invalid Code!!"}, nil
		}
	} else{
		return &pb.Success{Name: "User Not Found!"}, nil
	}
}

func (s *server) CancelConnection(ctx context.Context, in *pb.Request) (*pb.Success, error) {

	user := *in

	_, prs := Users[user.YourId]
	if prs {
		if Users[user.YourId].SecretCode == user.SecretCode {
			_,prs1 := Users[user.YourId].ConnectedUsers[user.UserId]
			if prs1 {
				delete(Users[user.YourId].ConnectedUsers, user.UserId)
				delete(Users[user.UserId].ConnectedUsers, user.YourId)
				return &pb.Success{Name: "Connection Cancelled Successfully!"}, nil
			} else {
				return &pb.Success{Name: "Connection Not Found!"}, nil
			}
		} else {
			return &pb.Success{Name: "Invalid Code!!"}, nil
		}
	} else{
		return &pb.Success{Name: "User Not Found!"}, nil
	}

}

func main() {

	Users = make(map[string]pb.Record)

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ser := grpc.NewServer()

	pb.RegisterDonarPatientServiceServer(ser, &server{})

	if err := ser.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
