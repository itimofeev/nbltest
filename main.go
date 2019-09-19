package main

import (
	"bitbucket.org/Axxonsoft/bl/security"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	"nbltest/an"
	"nbltest/util"
	"net/http"
	"time"
)

var ctx = context.Background()
var nblAddr = "192.168.56.101:50051"
var webserverAddr = "http://192.168.56.101:80"

func main() {
	initConfig()

	rand.Seed(time.Now().Unix())
	conn, err := an.MakeConnection(ctx, nblAddr, "root", "root", time.Second*10)
	util.CheckErr(err)

	login, password := createAdminUser(conn)
	util.CheckErr(checkLogin(login, password))
}

func initConfig() {
	nblAddrPtr := flag.String("nblAddr", "127.0.0.1:50051", "specify address to nbl with port")
	webserverAddrPtr := flag.String("webAddr", "http://127.0.0.1:80", "specify address to webserver with port and protocol")
	flag.Parse()
	nblAddr = *nblAddrPtr
	webserverAddr = *webserverAddrPtr
}

//noinspection ALL
func createAdminUser(conn *grpc.ClientConn) (userLogin, password string) {
	client := security.NewSecurityServiceClient(conn)

	userLogin = util.GetRandomNameU()
	password = "password"

	request := &security.ChangeConfigRequest{AddedUsers: []*security.User{
		{
			Login:   userLogin,
			Name:    userLogin,
			Enabled: true,
			CloudId: util.RandInt64(),
		},
	}}
	util.Log.WithField("req", request).Debug("sending request")
	_, err := client.ChangeConfig(ctx, request)
	util.CheckErr(err)

	configResponse, err := client.ListConfig(ctx, &security.ListConfigRequest{})
	util.CheckErr(err)

	util.Log.WithField("config", configResponse).Debug("loaded config")
	user := findUserByLogin(configResponse.Users, userLogin)
	if user == nil {
		util.CheckOk(false, "user not found by login "+userLogin)
	}

	adminRole := findRoleByName(configResponse.Roles, "admin")
	if adminRole == nil {
		util.CheckOk(false, "admin role not found")
	}

	request = &security.ChangeConfigRequest{
		ModifiedUserPasswords: []*security.UserPasswordAssignment{
			{
				UserIndex: user.Index,
				Password:  password,
			},
		},
		AddedUsersAssignments: []*security.UserAssignment{
			{
				UserId: user.Index,
				RoleId: adminRole.Index,
			},
		},
	}
	util.Log.WithField("req", request).Debug("sending request")
	_, err = client.ChangeConfig(ctx, request)
	util.CheckErr(err)

	return userLogin, password
}

func checkLogin(login, password string) error {
	log := util.Log.WithField("login", login).WithField("password", password)
	log.Debug("check login started")

	_, err := an.MakeConnection(ctx, nblAddr, login, password, time.Second)
	if err != nil {
		log.WithError(err).Error("nbl login failed")
		return err
	}
	log.Error("nbl login success")

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/video-origins", webserverAddr), nil)
	util.CheckErr(err)
	request.SetBasicAuth(login, password)

	resp, err := util.DefaultHTTPClient.Do(request)
	util.CheckErr(err)
	if resp.StatusCode != http.StatusOK {
		log.WithError(err).Errorf("web login failed, get not OK status but %d", resp.StatusCode)
		util.CheckOk(resp.StatusCode == http.StatusOK, fmt.Sprintf("got %d, expected 200", resp.StatusCode))
	}
	log.Error("web login success")

	return nil
}

func findUserByLogin(users []*security.User, login string) *security.User {
	for _, user := range users {
		if user.Login == login {
			return user
		}
	}
	return nil
}

func findRoleByName(roles []*security.Role, name string) *security.Role {
	for _, role := range roles {
		if role.Name == name {
			return role
		}
	}
	return nil
}
