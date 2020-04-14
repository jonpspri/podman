package containers

import (
	"fmt"

	"github.com/containers/libpod/cmd/podmanV2/parse"
	"github.com/containers/libpod/cmd/podmanV2/registry"
	"github.com/containers/libpod/cmd/podmanV2/utils"
	"github.com/containers/libpod/pkg/domain/entities"
	"github.com/spf13/cobra"
)

var (
	initDescription = `Initialize one or more containers, creating the OCI spec and mounts for inspection. Container names or IDs can be used.`

	initCommand = &cobra.Command{
		Use:     "init [flags] CONTAINER [CONTAINER...]",
		Short:   "Initialize one or more containers",
		Long:    initDescription,
		PreRunE: preRunE,
		RunE:    initContainer,
		Args: func(cmd *cobra.Command, args []string) error {
			return parse.CheckAllLatestAndCIDFile(cmd, args, false, false)
		},
		Example: `podman init --latest
  podman init 3c45ef19d893
  podman init test1`,
	}
)

var (
	initOptions entities.ContainerInitOptions
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Mode:    []entities.EngineMode{entities.ABIMode, entities.TunnelMode},
		Command: initCommand,
	})
	flags := initCommand.Flags()
	flags.BoolVarP(&initOptions.All, "all", "a", false, "Initialize all containers")
	flags.BoolVarP(&initOptions.Latest, "latest", "l", false, "Act on the latest container podman is aware of")
	_ = flags.MarkHidden("latest")
}

func initContainer(cmd *cobra.Command, args []string) error {
	var errs utils.OutputErrors
	report, err := registry.ContainerEngine().ContainerInit(registry.GetContext(), args, initOptions)
	if err != nil {
		return err
	}
	for _, r := range report {
		if r.Err == nil {
			fmt.Println(r.Id)
		} else {
			errs = append(errs, r.Err)
		}
	}
	return errs.PrintErrors()
}
