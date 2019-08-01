// Command-line utility to work with '\*.properties' file
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	flags "github.com/jessevdk/go-flags"
	"github.com/magiconair/properties"
)

var options struct {
	Get struct {
		File string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		Key  string `short:"k" long:"key" description:"Key" required:"true"`
	} `command:"get" description:"Reads property value by passed 'key' from specified file and prints it. If property is not found then empty string is printed"`

	Has struct {
		File string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		Key  string `short:"k" long:"key" description:"Key" required:"true"`
	} `command:"has" description:"Check if 'key' exists in specified file. If key exists the 'true' is printed, otherise 'false' is printed"`

	Put struct {
		File            string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		Key             string `short:"k" long:"key" description:"Key" required:"true"`
		Value           string `short:"v" long:"value" description:"Value" required:"true"`
		OnlyIfKeyExists bool   `long:"only-if-key-exists" description:"Don't create new keys"`
	} `command:"put" description:"Sets property 'key' to equal 'value' in specified file"`

	Delete struct {
		File string `short:"f" long:"file" description:"Path to properties file" value-name:"FILE" required:"true"`
		Key  string `short:"k" long:"key" description:"Key" required:"true"`
	} `command:"del" description:"Delete 'key' in specified file"`
}

func main() {
	parser := flags.NewParser(&options, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		flagsErr, ok := err.(*flags.Error)
		if ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else if ok && flagsErr.Type == flags.ErrCommandRequired {
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		} else {
			os.Exit(1)
		}
	}

	commandName := parser.Command.Active.Name

	switch commandName {
	case "get":
		p := properties.MustLoadFile(options.Get.File, properties.UTF8)
		value := p.GetString(options.Get.Key, "")
		fmt.Println(value)
	case "has":
		p := properties.MustLoadFile(options.Has.File, properties.UTF8)
		_, ok := p.Get(options.Has.Key)
		fmt.Println(ok)
	case "put":
		p := properties.MustLoadFile(options.Put.File, properties.UTF8)
		if _, ok := p.Get(options.Put.Key); options.Put.OnlyIfKeyExists && !ok {
			return
		}
		if _, _, err := p.Set(options.Put.Key, options.Put.Value); err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(options.Put.File, []byte(p.String()), 644); err != nil {
			panic(err)
		}
	case "del":
		p := properties.MustLoadFile(options.Delete.File, properties.UTF8)
		p.Delete(options.Delete.Key)
		if err := ioutil.WriteFile(options.Delete.File, []byte(p.String()), 644); err != nil {
			panic(err)
		}
	default:
		fmt.Println(fmt.Errorf("command '%s' is not implemented yet", commandName))
		os.Exit(1)
	}
}
