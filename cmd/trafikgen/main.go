// This is a helper CLI to download all the schemas from Trafikverket,
// uncompress them and then generate the code. Only people developing on
// trafikinfo itself should ever have a need for it, as the resulting code
// is included in the package.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	generatorCmd = newCmd("generate", "This command generates the Go code for each schema definition file.")
	downloadCmd  = newCmd("download", "This command downloads all known schemas from Trafikverket and unpacks them in the output directory. This directory can be passed as the '-schema-dir' to the generate command.")

	subnames = []string{
		generatorCmd.Name(),
		downloadCmd.Name(),
	}

	schemaDir   = generatorCmd.String("schema-dir", "", "path to directory with all the schemas")
	outputDir   = generatorCmd.String("output-dir", "", "path to the package root to output the Go code in")
	downloadDir = downloadCmd.String("download-dir", "", "path to directory to donwload the schemas to")
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	args := flag.Args()
	switch args[0] {
	case "generate":
		generatorCmd.Parse(args[1:])
		if *schemaDir == "" {
			log.Fatal("shcema-dir must be set")
		}
		if *outputDir == "" {
			log.Fatal("output-dir must be set")
		}
		generate(ctx, *schemaDir, *outputDir)
	case "download":
		downloadCmd.Parse(args[1:])
		if *downloadDir == "" {
			log.Fatalf("download-dir must be set")
		}
		tmpdir, err := os.MkdirTemp("", "trafikgen-*")
		if err != nil {
			log.Fatal(err)
		}
		log.Println("schema download starting...")
		if err := download(ctx, tmpdir); err != nil {
			log.Fatal(err)
		}
		log.Println("schema download completed")
		log.Println("schema extraction starting...")
		if err := extract(ctx, tmpdir, *downloadDir); err != nil {
			log.Fatal(err)
		}
		log.Println("schema extraction completed")
		if err := os.RemoveAll(tmpdir); err != nil {
			log.Printf("failed to cleanup temporary directory: %s\n", tmpdir)
		}
	default:
		log.Fatalf("[ERROR] unknown subcommand '%s', available subcommands are: %s", args[0], strings.Join(subnames, ", "))
	}
}

func newCmd(name, descr string) *flag.FlagSet {
	f := flag.NewFlagSet(name, flag.ExitOnError)
	f.Usage = func() {
		fmt.Fprintf(f.Output(), "Usage of: %s\n\n", fmt.Sprintf("%s %s", os.Args[0], name))
		fmt.Fprintf(f.Output(), descr+"\n")
		f.PrintDefaults()
	}

	return f
}
