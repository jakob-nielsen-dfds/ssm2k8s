package main

import (
	"os"

	"github.com/thofisch/ssm2k8s/internal/config"
	"github.com/thofisch/ssm2k8s/internal/logging"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app           = kingpin.New(os.Args[0], "A command-line secret manager")
	globalDebug   = app.Flag("globalDebug", "Enable globalDebug mode.").Envar("DEBUG").Bool()
	globalRegion  = app.Flag("region", "AWS region").Envar("AWS_DEFAULT_REGION").String()
	putCmd        = app.Command("put", "Create/update a secret.")
	putOptions    = NewPutCommand(putCmd)
	listCmd       = app.Command("list", "List secrets")
	deleteCmd     = app.Command("delete", "Delete secrets")
	deleteOptions = NewDeleteCommand(deleteCmd)
)

func main() {
	app.Version(config.VersionString)

	command := kingpin.MustParse(app.Parse(os.Args[1:]))
	logger := logging.NewConsoleLogger(*globalDebug)

	switch command {
	case putCmd.FullCommand():
		ExecutePut(logger, putOptions)

	case listCmd.FullCommand():
		ExecuteList(logger)

	case deleteCmd.FullCommand():
		ExecuteDelete(logger, deleteOptions)

	default:
		kingpin.Usage()
		os.Exit(1)
	}

}
