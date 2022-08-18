package cmd

import (
	"os"

	"github.com/civo/cli/cmd/apikey"
	"github.com/civo/cli/cmd/diskimage"
	"github.com/civo/cli/cmd/domain"
	"github.com/civo/cli/cmd/firewall"
	"github.com/civo/cli/cmd/instance"
	"github.com/civo/cli/cmd/ip"
	"github.com/civo/cli/cmd/kubernetes"
	"github.com/civo/cli/cmd/loadbalancer"
	"github.com/civo/cli/cmd/network"
	"github.com/civo/cli/cmd/permission"
	"github.com/civo/cli/cmd/region"
	"github.com/civo/cli/cmd/size"
	"github.com/civo/cli/cmd/sshkey"
	"github.com/civo/cli/cmd/teams"
	"github.com/civo/cli/cmd/volume"
	"github.com/civo/cli/common"
	"github.com/civo/cli/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "civo",
	Short: "CLI to manage cloud resources at Civo.com",
	Long: `civo is a CLI library for managing cloud resources such
as instances and Kubernetes clusters at Civo.com.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.ReadConfig)

	rootCmd.PersistentFlags().StringVarP(&config.Filename, "config", "", "", "config file (default is $HOME/.civo.json)")
	rootCmd.PersistentFlags().StringVarP(&common.OutputFields, "fields", "f", "", "output fields for custom format output (use -h to determine fields)")
	rootCmd.PersistentFlags().StringVarP(&common.OutputFormat, "output", "o", "human", "output format (json/human/custom)")
	rootCmd.PersistentFlags().BoolVarP(&common.DefaultYes, "yes", "y", false, "Automatic yes to prompts; assume \"yes\" as answer to all prompts and run non-interactively")
	rootCmd.PersistentFlags().StringVarP(&common.RegionSet, "region", "", "", "Choose the region to connect to, if you use this option it will use it over the default region")
	rootCmd.PersistentFlags().BoolVarP(&common.PrettySet, "pretty", "", false, "Print pretty the json output")

	rootCmd.AddCommand(apikey.APIKeyCmd)
	rootCmd.AddCommand(diskimage.DiskImageCmd)
	rootCmd.AddCommand(domain.DomainCmd)
	rootCmd.AddCommand(firewall.FirewallCmd)
	rootCmd.AddCommand(instance.InstanceCmd)
	rootCmd.AddCommand(ip.IPCmd)
	rootCmd.AddCommand(kubernetes.KubernetesCmd)
	rootCmd.AddCommand(loadbalancer.LoadBalancerCmd)
	rootCmd.AddCommand(network.NetworkCmd)
	rootCmd.AddCommand(permission.PermissionsCmd)
	rootCmd.AddCommand(region.RegionCmd)
	rootCmd.AddCommand(size.SizeCmd)
	rootCmd.AddCommand(sshkey.SSHKeyCmd)
	rootCmd.AddCommand(teams.TeamsCmd)
	rootCmd.AddCommand(volume.VolumeCmd)

	// Add warning if the region is empty, for the user with the old config
	config.ReadConfig()
}
