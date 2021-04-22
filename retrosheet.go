package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dmh2000/retrosheet/src/loader"
)

type options struct {
	populate bool
	query bool
	help bool
	mongodb string
	dirname string
}

func getOptions() options {
	var opt options

	// command line flags
	flag.BoolVar(&opt.help,"h", false, "print help")
	flag.BoolVar(&opt.help,"help", false, "print help")
	flag.BoolVar(&opt.populate,"p", false, "repopulate the mongodb database")
	flag.BoolVar(&opt.populate,"-populate", false, "repopulate the mongodb database")
	flag.BoolVar(&opt.query, "q", false, "execute the sample queries")
	flag.BoolVar(&opt.query, "-query", false, "execute the sample queries")
	flag.StringVar(&opt.mongodb, "m", "mongodb://localhost:27017", "-m [mongodb uri]")
	flag.StringVar(&opt.mongodb, "-mongo", "mongodb://localhost:27017", "--mongo [mongodb uri]")
	flag.StringVar(&opt.dirname, "d", "", "-d [path to retrosheet data directory] ")
	flag.StringVar(&opt.dirname, "-dir", "", "-dir [path to retrosheet data directory]")
	flag.Parse()

	return opt
}

func printHelp() {
	fmt.Println("Usage: retrosheet -[options]")
	flag.PrintDefaults()
	fmt.Println("path to retrosheet data directory is required for populate option")
	fmt.Println("path default can be set with environment variable RETROSHEET_DATA")
	fmt.Println("mongodb uri is required for populate and query options")
	fmt.Println("mongodb uri default can be set with environment variable RETROSHEET_MONGO")
}


func main() {
	var opt options
	var err error

	// get default path and uri from environment variables if set
	opt.dirname = os.Getenv("RETROSHEET_DATA")
	opt.mongodb = os.Getenv("RETROSHEET_MONGO")

	// set logging options
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	opt = getOptions()

	// if help requested, print and exit
	if opt.help {
		printHelp()
		return
	}

	// if no options, print help and exit
	if !opt.populate && !opt.query  {
		printHelp()
		return
	}

	// print environment variable data
	fmt.Printf("Mongodb URI : %v\n",opt.mongodb)

	if opt.populate {
		fmt.Printf("Data path   : %v\n",opt.dirname)
		if opt.dirname == "" {
			log.Println("missing retrosheet data directory path")
			printHelp()
			os.Exit(1)
			return
		}

		fmt.Println("Populating Database")
		err = loader.PopulateRetrosheet(opt.dirname,opt.mongodb)
		if err != nil {
			log.Fatal(err)
		}
	}	

	if opt.query {
		fmt.Println("Running Queries")
	}
}