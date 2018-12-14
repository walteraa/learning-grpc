package main

import (
	"context"
	"flag"
	"fmt"
	"net"
    "log"
	"google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "gopkg.in/mgo.v2/bson"
	timestamp "github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
    pb "user_service/proto"
    db "user_service/server/database"
    models "user_service/server/models"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = "50052"
    address =  "0.0.0.0:" + port
)

type UserServer struct{
    userDao *db.UserDAO
}

func (s *UserServer) GetUser(context context.Context, in *pb.UserRequest) ( *pb.UserResponse, error){
    // TODO: Figure out how to bind DAO response directly to UserResponse
    // TODO: Add error response in protobuf instead breaking the server
    
    user, err := s.userDao.FindById("5c130697e1054e8727cbb3a9")
    if err != nil{
        log.Fatal("User not found, id =",in.UserId, err )
    }
    return &pb.UserResponse{Id: user.Id.Hex(), FirstName: user.FirstName, LastName: user.LastName, BirthDate: &tspb.Timestamp{Seconds: user.BirthDate.Unix()}}, err
}

func (s *UserServer) CreateUser(context context.Context, in *pb.CreateUserRequest) (*pb.UserRequest, error){
    // TODO: Add error response in protobuf

    var user models.User
    birthDate,err := timestamp.Timestamp(in.BirthDate)
    user = models.User{
        Id: bson.NewObjectId(),
        FirstName: in.FirstName, 
        LastName: in.LastName,
        BirthDate: birthDate }
    err = s.userDao.Insert(user)
    return &pb.UserRequest{UserId: user.Id.Hex()}, err
}

func main(){
    listen, err := net.Listen("tcp", address)
    if err != nil {
        fmt.Printf("Error when listening: %v\n", err)
    }

    var userDao *db.UserDAO
    userDao,err = db.GetUserDAOInstance()
    if err != nil{
        log.Fatal("Couldnt't create the database access layer")
        return
    }
    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &UserServer{userDao})
    reflection.Register(grpcServer)
    fmt.Printf("Server listening on: %v\n", address)
    grpcServer.Serve(listen)
}
