package main

import(
    "os"
    "io"
    "io/ioutil"
    "log"
    "encoding/csv"

    "github.com/jszwec/csvutil"
)

type ShopifyCollection struct {
}

func(sc ShopifyCollection) Create(filepath string) []*ShopifyCsvBridge {

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

    var products []*ShopifyCsvBridge
    for {
        var p ShopifyCsvBridge
        if err := dec.Decode(&p); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        products = append(products, &p)
    }

    return products

}

func(sc ShopifyCollection) WriteCsv(products []*ShopifyCsvBridge, destination string) {

    b, err := csvutil.Marshal(products)
    if err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile(destination, b, 0666); err != nil {
        log.Fatal(err)
    }

}
