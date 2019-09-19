module nbltest

go 1.12

require (
	bitbucket.org/Axxonsoft/bl v0.0.0
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873 // indirect
	google.golang.org/grpc v1.23.1
	google/api v0.0.0 // indirect
)

replace bitbucket.org/Axxonsoft/bl => ../nbltest/thirdparty/bitbucket.org/Axxonsoft/bl

replace google/api => ../nbltest/thirdparty/google/api
