package main

import(
    "fmt"
    "os"
    "io"
    "log"
    //"path"
    "encoding/csv"

    "github.com/integrii/flaggy"
    "github.com/jszwec/csvutil"
)

// Subcommands and flags set for Flaggy
var subcommandRename                *flaggy.Subcommand
var subcommandLink                  *flaggy.Subcommand
var subcommandDescribe              *flaggy.Subcommand
var renamerFilepathFlag             string
var renamerDirectoryFlag            string
var renamerDestinationFlag          string
var linkerReferenceFlag             string
var linkerFilepathFlag              string
var linkerDestinationFlag           string
var linkerPrefixFlag                string
var linkerModelFlag                 string
var describerFilepathFlag           string
var describerDestinyFlag            string

type Linkable interface {
    ChangeLink(link string)
}

type Describable interface {
    Describe()
    Sku() string
}

func init() {

    subcommandRename = flaggy.NewSubcommand("rename")
    subcommandRename.String(&renamerFilepathFlag, "r", "reference", "Filepath to csv reference feed")
    subcommandRename.String(&renamerDirectoryFlag, "l", "location", "Path to directory where the files are located")
    subcommandRename.String(&renamerDestinationFlag, "d", "destination", "Destination directory for renamed files")
    flaggy.AttachSubcommand(subcommandRename, 1)

    subcommandLink = flaggy.NewSubcommand("link")
    subcommandLink.String(&linkerReferenceFlag, "r", "reference", "Filepath to either csv reference feed or directory containing images")
    subcommandLink.String(&linkerFilepathFlag, "f", "filepath", "Filepath of csv to be edited")
    subcommandLink.String(&linkerPrefixFlag, "p", "prefix", "Uri to be prefixed before image name in link")
    subcommandLink.String(&linkerDestinationFlag, "d", "destination", "Filepath to write new csv")
    subcommandLink.String(&linkerModelFlag, "m", "model", "Format of the csv product data (ie shopify)")
    flaggy.AttachSubcommand(subcommandLink, 1)

    subcommandDescribe = flaggy.NewSubcommand("describe")
    subcommandDescribe.String(&describerFilepathFlag, "f", "filepath", "Filepath to csv feed")
    subcommandDescribe.String(&describerDestinyFlag, "d", "destiny", "Filepath and filename for new CSV")
    flaggy.AttachSubcommand(subcommandDescribe, 1)

    flaggy.Parse()

}

func main() {

    fmt.Println("Welcome to the Knotted Multi-tool")

    if subcommandRename.Used {
        renameFromCsv(renamerFilepathFlag, renamerDirectoryFlag, renamerDestinationFlag)
    }

    if subcommandLink.Used {
        linkCsv(linkerModelFlag, linkerReferenceFlag, linkerFilepathFlag, linkerPrefixFlag, linkerDestinationFlag)
    }

    if subcommandDescribe.Used {
        describeFromCsv(describerFilepathFlag, describerDestinyFlag)
    }

    fmt.Println("Finished all jobs, exiting...")

}

func renameFromCsv(filepath string, location_dir string, destination_dir string) {

    f, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    r := csv.NewReader(f)

    dec, err := csvutil.NewDecoder(r)
    if err != nil {
        log.Fatal(err)
    }

    var name_refs []NameRef
    for {
        var nm NameRef
        if err := dec.Decode(&nm); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        name_refs = append(name_refs, nm)
    }

    renamer := Renamer{
        location_dir,
        destination_dir,
        name_refs,}

    report := renamer.ExecuteAndReport()

    fmt.Printf("Found and renamed %d images to (%s)", report, destination_dir)
}

func linkCsv(model string, ref string, filepath string, prefix string, des_dir string) {

    linker := new(Linker)
    linker.model = model
    linker.prefix = prefix
    linker.ref = ref
    linker.SetImageList()

    switch model {
    case "shopify":
        sc := new(ShopifyCollection)
        products := sc.Create(filepath)

        for _, product := range products {
            linker.ChangeLinks(product, product.Sku)
        }

        sc.WriteCsv(products, des_dir)

    }
}

func describeFromCsv(filepath string, destiny string) {

    vc := new(VendCollection)
    products := vc.Create(filepath)
    yarnInfo := vc.CreateYarnInfo(filepath)

    descMap := make(map[string]string)
    for _, p := range yarnInfo {
        descMap[p.Sku] = p.Describe()
    }

    for _, p := range products {
        p.Description = descMap[p.Sku]
    }

    vc.WriteCsv(products, destiny)
}
