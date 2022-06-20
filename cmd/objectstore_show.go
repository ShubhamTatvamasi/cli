package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var objectStoreShowCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"get", "info"},
	Example: `civo objectstore show OBJECTSTORE_NAME`,
	Short:   "Prints information about an Object Store",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		objectStore, err := client.FindObjectStore(args[0])
		if err != nil {
			utility.Error("%s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriter()

		ow.StartLine()
		fmt.Println()
		ow.AppendDataWithLabel("id", objectStore.ID, "ID")
		ow.AppendDataWithLabel("name", objectStore.Name, "Name")
		ow.AppendDataWithLabel("generated_name", objectStore.GeneratedName, "Generated Name")
		ow.AppendDataWithLabel("size", objectStore.MaxSize, "Size")
		ow.AppendDataWithLabel("max_objects", strconv.Itoa(objectStore.MaxObjects), "Max Objects")
		ow.AppendDataWithLabel("objectstore_endpoint", objectStore.ObjectStoreEndpoint, "Object Store Endpoint")
		ow.AppendDataWithLabel("region", client.Region, "Region")
		ow.AppendDataWithLabel("accesskey", objectStore.AccessKeyID, "Access Key")
		ow.AppendDataWithLabel("status", objectStore.Status, "Status")

		switch outputFormat {
		case "json":
			ow.WriteMultipleObjectsJSON(prettySet)
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			ow.WriteKeyValues()
			fmt.Printf("To access the secret key run: civo objectstore credential secret --access-key=%s\n", utility.Green(objectStore.AccessKeyID))
		}
	},
}