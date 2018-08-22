package tasks

import (
	"fmt"

	printer "github.com/KablamoOSS/go-cli-printer"
	"github.com/KablamoOSS/kombustion/internal/cloudformation"
	"github.com/KablamoOSS/kombustion/internal/cloudformation/tasks"
	"github.com/KablamoOSS/kombustion/internal/manifest"
	"github.com/urfave/cli"
)

// PrintEvents outputs the events of a stack
func PrintEvents(c *cli.Context) {
	printer.Progress("Kombusting")

	manifestFile := manifest.FindAndLoadManifest()

	region := c.String("region")
	if region == "" {
		// If no region was provided by the cli flag, check for the default in the manifest
		if manifestFile.Region != "" {
			region = manifestFile.Region
		}
	}

	acctID, cfClient := tasks.GetCloudformationClient(
		c.GlobalString("profile"),
		region,
	)

	environment := c.String("environment")

	var stackName string

	fileName := c.Args().Get(0)
	if fileName == "" {
		printer.Fatal(
			fmt.Errorf("Can't find stack, no source template provided"),
			"Add the path to the source template file, or provide use --stack-name",
			"https://www.kombustion.io/api/cli/#stacks",
		)
	}

	if env, ok := manifestFile.Environments[environment]; ok {
		if !env.IsWhitelistedAccount(acctID) {
			printer.Fatal(
				fmt.Errorf("Account %s is not allowed for environment %s", acctID, environment),
				"Use whitelisted account, or add account to environment accounts in kombustion.yaml",
				"",
			)
		}
	}

	stackName = cloudformation.GetStackName(
		manifestFile,
		fileName,
		environment,
		c.String("stack-name"),
	)

	tasks.PrintStackEvents(cfClient, stackName)
}
