go get github.com/spf13/cobra/cobra
go get -u github.com/spf13/cobra

golang.org/x/sys/unix
golang.org/x/text/transform
github.com/hashicorp/hcl/hcl/printer

cd ~/go/src
mkdir -p golang.org/x
git clone https://github.com/golang/sys.git
git clone https://gitee.com/mirrors/hcl.git





go mod init clid
go env -w GO111MODULE=on

cobra init --pkg-name clid
cobra init --pkg-name clid --author "TianChang" -l MIT
cobra init --pkg-name clid -a "TianChang" -l MIT


go build
   go get github.com/mitchellh/go-homedir
   go get github.com/spf13/cobra
   go get github.com/spf13/viper


# 添加子命令
    cobra add version

#   单词
package pkg
author
license
