package main

import(
    "os"
    "io"
    "io/ioutil"
    "fmt"
    "log"
    "strings"
    "encoding/csv"

    "github.com/jszwec/csvutil"
)

type RainCustomerCollection struct {
    Customers       []RainCustomer
}

func(rc RainCustomerCollection) Create(filepath string) []RainCustomer {

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

    var rainCusts []RainCustomer
    for {
        var rainCust RainCustomer
        if err := dec.Decode(&rainCust); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        rainCusts = append(rainCusts, rainCust)
    }

    return rainCusts

}

func(rc RainCustomerCollection) Map(fieldname string) map[string]RainCustomer {

    fieldMap := make(map[string]RainCustomer)

    switch fieldname {
    case "email":
        for _, c := range rc.Customers {
            fieldMap[strings.ToLower(c.Email)] = c
        }
    case "name":
        for _, c := range rc.Customers {
            firstName := cleanNonAlphanum(strings.ToLower(c.FirstName))
            lastName := cleanNonAlphanum(strings.ToLower(c.LastName))
            namestring := fmt.Sprintf("%s%s", firstName, lastName)
            fieldMap[namestring] = c
        }
    case "rainId":
        for _, c := range rc.Customers {
            rainId := c.RainId
            fieldMap[rainId] = c
        }
    }

    return fieldMap

}

func(rc RainCustomerCollection) Add(custs []RainCustomer) {

    for _, x := range custs {
        rc.Customers = append(rc.Customers, x)
    }

}

func(rc RainCustomerCollection) ToVend() []VendCustomer {

    var vendCusts []VendCustomer
    for _, rainCust := range rc.Customers {

        vendCust := rainCust.ToVend()

        vendCusts = append(vendCusts, vendCust)
    }

    return vendCusts

}

func(rc RainCustomerCollection) WriteCsv(destiny string) {

    b, err := csvutil.Marshal(rc.Customers)
    if err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile(destiny, b, 0666); err != nil {
        log.Fatal(err)
    }

}
