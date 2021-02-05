package main

import (
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs" // 根据容器配额设置 maxprocs
	_ "net/http/pprof"           // 注册 pprof 接口

	"moocss.com/gaea/cmd/job"
	"moocss.com/gaea/cmd/server"
)

func main() {
	root := cobra.Command{Use: "gaea"}

	root.AddCommand(
		server.Cmd,
		job.Cmd,
	)

	root.Execute()
}
