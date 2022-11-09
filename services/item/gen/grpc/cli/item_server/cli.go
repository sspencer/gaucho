// Code generated by goa v3.10.2, DO NOT EDIT.
//
// ItemServer gRPC client CLI support package
//
// Command:
// $ goa gen gaucho/services/item/design -o services/item

package cli

import (
	"flag"
	"fmt"
	itemc "gaucho/services/item/gen/grpc/item/client"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `item (get|get-all|add|update|remove)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` item get --message '{
      "field": 893706942
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		itemFlags = flag.NewFlagSet("item", flag.ContinueOnError)

		itemGetFlags       = flag.NewFlagSet("get", flag.ExitOnError)
		itemGetMessageFlag = itemGetFlags.String("message", "", "")

		itemGetAllFlags = flag.NewFlagSet("get-all", flag.ExitOnError)

		itemAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		itemAddMessageFlag = itemAddFlags.String("message", "", "")

		itemUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		itemUpdateMessageFlag = itemUpdateFlags.String("message", "", "")

		itemRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		itemRemoveMessageFlag = itemRemoveFlags.String("message", "", "")
	)
	itemFlags.Usage = itemUsage
	itemGetFlags.Usage = itemGetUsage
	itemGetAllFlags.Usage = itemGetAllUsage
	itemAddFlags.Usage = itemAddUsage
	itemUpdateFlags.Usage = itemUpdateUsage
	itemRemoveFlags.Usage = itemRemoveUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "item":
			svcf = itemFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "item":
			switch epn {
			case "get":
				epf = itemGetFlags

			case "get-all":
				epf = itemGetAllFlags

			case "add":
				epf = itemAddFlags

			case "update":
				epf = itemUpdateFlags

			case "remove":
				epf = itemRemoveFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "item":
			c := itemc.NewClient(cc, opts...)
			switch epn {
			case "get":
				endpoint = c.Get()
				data, err = itemc.BuildGetPayload(*itemGetMessageFlag)
			case "get-all":
				endpoint = c.GetAll()
				data = nil
			case "add":
				endpoint = c.Add()
				data, err = itemc.BuildAddPayload(*itemAddMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = itemc.BuildUpdatePayload(*itemUpdateMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = itemc.BuildRemovePayload(*itemRemoveMessageFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// itemUsage displays the usage of the item command and its subcommands.
func itemUsage() {
	fmt.Fprintf(os.Stderr, `Service that provides item API
Usage:
    %[1]s [globalflags] item COMMAND [flags]

COMMAND:
    get: Retrieve item by the given id
    get-all: Retrieve all items
    add: Create an item
    update: Update an item
    remove: Remove item

Additional help:
    %[1]s item COMMAND --help
`, os.Args[0])
}
func itemGetUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item get -message JSON

Retrieve item by the given id
    -message JSON: 

Example:
    %[1]s item get --message '{
      "field": 893706942
   }'
`, os.Args[0])
}

func itemGetAllUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item get-all

Retrieve all items

Example:
    %[1]s item get-all
`, os.Args[0])
}

func itemAddUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item add -message JSON

Create an item
    -message JSON: 

Example:
    %[1]s item add --message '{
      "damage": 11,
      "description": "Enim eum est et commodi.",
      "healing": 20,
      "name": "pgd",
      "protection": 19
   }'
`, os.Args[0])
}

func itemUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item update -message JSON

Update an item
    -message JSON: 

Example:
    %[1]s item update --message '{
      "id": 1795403724,
      "item": {
         "damage": 10,
         "description": "Rem hic doloribus tempora placeat.",
         "healing": 27,
         "name": "phb",
         "protection": 4
      }
   }'
`, os.Args[0])
}

func itemRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item remove -message JSON

Remove item
    -message JSON: 

Example:
    %[1]s item remove --message '{
      "id": 3738918782
   }'
`, os.Args[0])
}