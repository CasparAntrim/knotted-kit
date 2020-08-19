package main

import(
    "fmt"
    "os"
    "io"
    "log"
    "strings"
    //"path"
    "encoding/csv"

    "github.com/integrii/flaggy"
    "github.com/jszwec/csvutil"
)

// Subcommands and flags set for Flaggy
var subcommandRename                *flaggy.Subcommand
var subcommandLink                  *flaggy.Subcommand
var subcommandDescribe              *flaggy.Subcommand
var subcommandRelate                *flaggy.Subcommand
var subcommandVerify                *flaggy.Subcommand
var subcommandCombine               *flaggy.Subcommand
var subcommandUpdateCusts           *flaggy.Subcommand
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
var relaterFilepathFlag             string
var relaterDestinyFlag              string
var relaterModelFlag                string
var relaterNewModelFlag             string
var verifierFilepathFlag            string
var verifierDestinyFlag             string
var verifierCustModelFlag           string
var verifierVModelFlag              string
var verifierDatapathFlag            string
var verifierMatchFlag               string
var combinerSourceFilepathFlag      string
var combinerSourceModelFlag         string
var combinerSource2FilepathFlag     string
var combinerSource2ModelFlag        string
var combinerDestinyModelFlag        string
var combinerResultFilepathFlag      string
var updateCustSourceFilepathFlag    string
var updateCustTargetFilepathFlag    string
var updateCustDestinyFilepathFlag   string
var updateCustSourceModelFlag       string
var updateCustTargetModelFlag       string
var updateCustAddTarget1Flag        string
var updateCustAddTarget2Flag        string
var updateCustAddTarget3Flag        string
var updateCustMatchFlag             string

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

    subcommandRelate = flaggy.NewSubcommand("relate")
    subcommandRelate.String(&relaterFilepathFlag, "f", "filepath", "Filepath to csv feed")
    subcommandRelate.String(&relaterDestinyFlag, "d", "destiny", "Filepath and filename for new CSV")
    subcommandRelate.String(&relaterModelFlag, "m", "model", "Current model/format of customer data")
    subcommandRelate.String(&relaterNewModelFlag, "n", "new", "New model/format to which to convert customer data")
    flaggy.AttachSubcommand(subcommandRelate, 1)

    subcommandVerify = flaggy.NewSubcommand("verify")
    subcommandVerify.String(&verifierFilepathFlag, "f", "filepath", "Filepath to csv customer feed to verify")
    subcommandVerify.String(&verifierDestinyFlag, "d", "destiny", "Filepath and name for final product of verification")
    subcommandVerify.String(&verifierCustModelFlag, "m", "model", "Model format of data to be verified")
    subcommandVerify.String(&verifierVModelFlag, "v", "verifier", "Model format of data to be verifier against (ie. ConstantContact)")
    subcommandVerify.String(&verifierDatapathFlag, "p", "datapath", "Filepath to data to verify against")
    subcommandVerify.String(&verifierMatchFlag, "r", "match", "Fieldname to match customer data against")
    flaggy.AttachSubcommand(subcommandVerify, 1)

    subcommandCombine = flaggy.NewSubcommand("combine")
    subcommandCombine.String(&combinerSourceFilepathFlag, "s1", "source1", "Filepath to first source csv")
    subcommandCombine.String(&combinerSourceModelFlag, "s1m", "source1-model", "Model of first source data")
    subcommandCombine.String(&combinerSource2FilepathFlag, "s2", "source2", "Filepath to second source csv")
    subcommandCombine.String(&combinerSource2ModelFlag, "s2m", "source2-model", "Model of second source data")
    subcommandCombine.String(&combinerDestinyModelFlag, "m", "model", "Destiny model for new data")
    subcommandCombine.String(&combinerResultFilepathFlag, "d", "destiny", "Filepath to resulting csv")
    flaggy.AttachSubcommand(subcommandCombine, 1)

    subcommandUpdateCusts = flaggy.NewSubcommand("update-customers")
    subcommandUpdateCusts.String(&updateCustSourceFilepathFlag, "s", "source", "Filepath to source of new data")
    subcommandUpdateCusts.String(&updateCustTargetFilepathFlag, "t", "target", "Filepath to target data to be updated")
    subcommandUpdateCusts.String(&updateCustDestinyFilepathFlag, "d", "destiny", "Filepath for new CSV")
    subcommandUpdateCusts.String(&updateCustSourceModelFlag, "sm", "source-model", "Vendor model for source data")
    subcommandUpdateCusts.String(&updateCustTargetModelFlag, "tm", "target-model", "Vendor model for target data")
    subcommandUpdateCusts.String(&updateCustAddTarget1Flag, "a1", "add1", "Additional target filepath 1")
    subcommandUpdateCusts.String(&updateCustAddTarget2Flag, "a2", "add2", "Additional target filepath 2")
    subcommandUpdateCusts.String(&updateCustAddTarget3Flag, "a3", "add3", "Additional target filepath 3")
    subcommandUpdateCusts.String(&updateCustMatchFlag, "x", "match", "Field to use to match between vendor models")
    flaggy.AttachSubcommand(subcommandUpdateCusts, 1)

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

    if subcommandRelate.Used {
        relateFromCsv(relaterFilepathFlag, relaterDestinyFlag, relaterModelFlag, relaterNewModelFlag)
    }

    if subcommandVerify.Used {
        verifyFromCsv(verifierFilepathFlag, verifierDestinyFlag, verifierCustModelFlag, verifierVModelFlag, verifierDatapathFlag, verifierMatchFlag)
    }

    if subcommandCombine.Used {
        combineCustomerData(combinerSourceFilepathFlag, combinerSourceModelFlag, combinerSource2FilepathFlag, combinerSource2ModelFlag, combinerDestinyModelFlag, combinerResultFilepathFlag)
    }

    if subcommandUpdateCusts.Used {
        updateCustomers(updateCustSourceFilepathFlag, updateCustTargetFilepathFlag, updateCustDestinyFilepathFlag, updateCustSourceModelFlag, updateCustTargetModelFlag, updateCustAddTarget1Flag, updateCustAddTarget2Flag, updateCustAddTarget3Flag, updateCustMatchFlag)
    }

    fmt.Println("Finished all jobs, exiting...")

}

type Collection interface {
    WriteCsv(filepath string)
}

func updateCustomers(sourceFilepath string, targetFilepath string, destinyFilepath string, sourceModel string, targetModel string, addTarget1 string, addTarget2 string, addTarget3 string, match string) {

    targets := determineAdditionalTargets(addTarget1, addTarget2, addTarget3)
    sourceTargetCode := determineSourceTargetCode(sourceModel, targetModel)

    switch sourceTargetCode {
    case "rain:vend":
        rCustCollection := new(RainCustomerCollection)
        rCusts := rCustCollection.Create(sourceFilepath)
        rCustCollection.Customers = rCusts

        vCustCollection := new(VendCustomerCollection)
        vCusts := vCustCollection.Create(targetFilepath)

        for _, t := range targets {
            custs := vCustCollection.Create(t)
            for _, c := range custs {
                vCusts = append(vCusts, c)
            }
        }

        vCustCollection.Customers = vCusts

        compareMap := rCustCollection.Map("rainId")

        var updatedCusts []VendCustomer
        for _, vc := range vCustCollection.Customers {
            var newVc VendCustomer
            var matchString string
            switch match {
            case "rainId":
                matchString = vc.Custom1
            }

            if _, ok := compareMap[matchString]; ok {
                oldRc := compareMap[matchString]
                newVc = oldRc.ToVend()
                updatedCusts = append(updatedCusts, newVc)
            }
        }

        vCustCollection.Customers = updatedCusts
        vCustCollection.WriteCsv(destinyFilepath)

    case "vend:vend":
        //
    }
}

func determineSourceTargetCode(m1 string, m2 string) string {
    codeString := fmt.Sprintf("%s:%s", m1, m2)
    return codeString
}

func determineAdditionalTargets(t1 string, t2 string, t3 string) []string {

    targets := [3]string{t1, t2, t3}
    var paths []string
    // count := 0

    for _, t := range targets {
        if t != "" && t != " " {
            // count += 1
            paths = append(paths, t)
        }
    }

    return paths

}

func combineCustomerData(s1Path string, s1Model string, s2Path string, s2Model string, resultModel string, resultPath string) {

    fmt.Println("Combining customer data")

    matchString := fmt.Sprintf("%s:%s:%s", s1Model, s2Model, resultModel)

    switch matchString {
    case "rain:contact:rain":
        rCustCollection := new(RainCustomerCollection)
        rCusts := rCustCollection.Create(s1Path)
        rCustCollection.Customers = rCusts

        cCustCollection := new(ContactCustomerCollection)
        cCusts := cCustCollection.Create(s2Path)
        cCustCollection.Customers = cCusts

        r2CustCollection := new(RainCustomerCollection)
        r2CustCollection.Customers = cCustCollection.ToRain()

        for _, cust := range r2CustCollection.Customers {
            rCustCollection.Customers = append(rCustCollection.Customers, cust)
        }

        rCustCollection.WriteCsv(resultPath)
    }

    // s1Data := resolveCsvData("customer", s1Model, s1Path)
    // s2Data := resolveCsvData("customer", s2Model, s2Path)
    //
    // combinedData := combineCollections(s1Data, s2Data)



}

// func resolveCsvData(type string, model string, filepath string) Collection {
//
//     switch type {
//     case "customer":
//
//         switch model {
//         case "rain":
//             rCustCollection := new(RainCustomerCollection)
//             rCusts := rCustCollection.Create(filepath)
//             rCustCollection.Customers = rCusts
//         case "vend":
//             vCustCollection := new(VendCustomerCollection)
//             vCusts := vCustCollection.Create(filepath)
//             vCustCollection.Customers = vCusts
//         case "shopify":
//             fmt.Println("We don't do Shopify here.")
//         case "contact":
//             cCustCollection := new(ContactCustomerCollection)
//             cCusts := cCustCollection.Create(datapath)
//             cCustCollection.Customers = cCusts
//         }
//
//     case "product":
//         fmt.Println("We don't do products, yet.")
//     }
//
// }

func verifyFromCsv(filepath string, destiny string, custModel string, vModel string, datapath string, match string) {

    switch custModel {
    case "rain":
        rCustCollection := new(RainCustomerCollection)
        rCusts := rCustCollection.Create(filepath)
        rCustCollection.Customers = rCusts

        cCustCollection := new(ContactCustomerCollection)
        cCusts := cCustCollection.Create(datapath)
        cCustCollection.Customers = cCusts

        compareMap := cCustCollection.Map(match)

        var verifiedCusts []RainCustomer
        for _, rCust := range rCustCollection.Customers {
            var matchString string
            switch match {
            case "email":
                matchString = fmt.Sprintf("%s", strings.ToLower(rCust.Email))
            case "name":
                firstName := cleanNonAlphanum(strings.ToLower(rCust.FirstName))
                lastName := cleanNonAlphanum(strings.ToLower(rCust.LastName))
                matchString = fmt.Sprintf("%s%s", firstName, lastName)
            }

            if _, ok := compareMap[matchString]; ok {
                verifiedCusts = append(verifiedCusts, rCust)
            }
        }

        verifiedCustCollection := new(RainCustomerCollection)
        verifiedCustCollection.Customers = verifiedCusts
        verifiedCustCollection.WriteCsv(destiny)

    case "vend":
        vCustCollection := new(VendCustomerCollection)
        vCusts := vCustCollection.Create(filepath)
        vCustCollection.Customers = vCusts
    }

}

func relateFromCsv(filepath string, destiny string, model string, newModel string) {

    switch model {
    case "rain":
        rCustCollection := new(RainCustomerCollection)
        rCusts := rCustCollection.Create(filepath)
        rCustCollection.Customers = rCusts

        switch newModel {
        case "vend":
            vCustCollection := new(VendCustomerCollection)
            vCusts := rCustCollection.ToVend()
            vCustCollection.Customers = vCusts

            vCustCollection.WriteCsv(destiny)

        }
    }

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
