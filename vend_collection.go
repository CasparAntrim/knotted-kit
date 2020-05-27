package main

import(
    "os"
    "io"
    "io/ioutil"
    "log"
    "encoding/csv"

    "github.com/jszwec/csvutil"
)

type VendCollection struct {
}

func(vc VendCollection) Create(filepath string) []*VendCsvBridge {

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

    var products []*VendCsvBridge
    for {
        var p VendCsvBridge
        if err := dec.Decode(&p); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        products = append(products, &p)
    }

    return products

}

func(vc VendCollection) CreateYarnInfo(filepath string) []VendCsvBridgeYarnInfo {

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

    var products []VendCsvBridgeYarnInfo
    for {
        var p VendCsvBridgeYarnInfo
        if err := dec.Decode(&p); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        products = append(products, p)
    }

    return products

}

func(vc VendCollection) WriteCsv(products []*VendCsvBridge, destination string) {

    b, err := csvutil.Marshal(products)
    if err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile(destination, b, 0666); err != nil {
        log.Fatal(err)
    }

}
