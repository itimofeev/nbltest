package main

import (
	"bitbucket.org/Axxonsoft/bl/security"
	"context"
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
	rand.Seed(time.Now().Unix())
	conn, err := an.MakeConnection(ctx, nblAddr, "root", "root", time.Second*10)
	util.CheckErr(err)

	login, password := createAdminUser(conn)
	util.CheckErr(checkLogin(login, password))
}

//noinspection ALL
func createAdminUser(conn *grpc.ClientConn) (userLogin, password string) {
	client := security.NewSecurityServiceClient(conn)

	userLogin = util.GetRandomNameU()
	password = "password"

	_, err := client.ChangeConfig(ctx, &security.ChangeConfigRequest{AddedUsers: []*security.User{
		{
			Login:   userLogin,
			Name:    userLogin,
			Enabled: true,
			CloudId: util.RandInt64(),
		},
	}})
	util.CheckErr(err)

	configResponse, err := client.ListConfig(ctx, &security.ListConfigRequest{})
	util.CheckErr(err)

	user := findUserByLogin(configResponse.Users, userLogin)
	if user == nil {
		util.CheckOk(false, "user not found by login "+userLogin)
	}

	adminRole := findRoleByName(configResponse.Roles, "admin")
	if adminRole == nil {
		util.CheckOk(false, "admin role not found")
	}

	_, err = client.ChangeConfig(ctx, &security.ChangeConfigRequest{
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
	})
	util.CheckErr(err)

	return userLogin, password
}

func checkLogin(login, password string) error {
	_, err := an.MakeConnection(ctx, nblAddr, login, password, time.Second)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/video-origins", webserverAddr), nil)
	util.CheckErr(err)
	request.SetBasicAuth(login, password)

	resp, err := util.DefaultHTTPClient.Do(request)
	util.CheckErr(err)
	util.CheckOk(resp.StatusCode == http.StatusOK, fmt.Sprintf("got %d, expected 200", resp.StatusCode))

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
