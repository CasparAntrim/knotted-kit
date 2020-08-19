package main

import(
    "io/ioutil"
    "log"
    "os"
    "encoding/csv"
    "io"

    "github.com/jszwec/csvutil"
)

type VendCustomerCollection struct {
    Customers       []VendCustomer
}

func(vc VendCustomerCollection) Create(filepath string) []VendCustomer {

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

    var vendCusts []VendCustomer
    for {
        var vendCust VendCustomer
        if err := dec.Decode(&vendCust); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        vendCusts = append(vendCusts, vendCust)
    }

    return vendCusts
}

func(vc VendCustomerCollection) WriteCsv(destiny string) {

    b, err := csvutil.Marshal(vc.Customers)
    if err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile(destiny, b, 0666); err != nil {
        log.Fatal(err)
    }

}
