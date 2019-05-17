//module github.com/phpli/golang
module phpli

go 1.12

replace golang.org/x/net v0.0.0-20181114220301-adae6a3d119a => github.com/golang/net v0.0.0-20181114220301-adae6a3d119a

replace google.golang.org/grpc v1.16.0 => github.com/grpc/grpc-go v1.16.0

replace golang.org/x/crypto v0.0.0-20181127143415-eb0de9b17e85 => github.com/golang/crypto v0.0.0-20181127143415-eb0de9b17e85

require github.com/astaxie/beego v1.11.1 // indirect
