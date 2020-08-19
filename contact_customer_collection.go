package main

import(
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "encoding/csv"
    "io"
    "strings"
    "regexp"

    "github.com/jszwec/csvutil"
)

type ContactCustomerCollection struct {
    Customers       []ContactCustomer
}

func(cc ContactCustomerCollection) Create(filepath string) []ContactCustomer {

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

    var cCusts []ContactCustomer
    for {
        var cCust ContactCustomer
        if err := dec.Decode(&cCust); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }

        cCusts = append(cCusts, cCust)
    }

    return cCusts
}

func(cc ContactCustomerCollection) ToRain() []RainCustomer {

    var rainCusts []RainCustomer
    for _, cCust := range cc.Customers {

        rainCust := cCust.ToRain()

        rainCusts = append(rainCusts, rainCust)
    }

    return rainCusts

}

func(cc ContactCustomerCollection) WriteCsv(destiny string) {

    b, err := csvutil.Marshal(cc.Customers)
    if err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile(destiny, b, 0666); err != nil {
        log.Fatal(err)
    }

}

func(cc ContactCustomerCollection) Map(fieldname string) map[string]ContactCustomer {

    fieldMap := make(map[string]ContactCustomer)

    switch fieldname {
    case "email":
        for _, c := range cc.Customers {
            fieldMap[strings.ToLower(c.Email)] = c
        }
    case "name":
        for _, c := range cc.Customers {
            firstName := cleanNonAlphanum(strings.ToLower(c.FirstName))
            lastName := cleanNonAlphanum(strings.ToLower(c.LastName))
            namestring := fmt.Sprintf("%s%s", firstName, lastName)
            fieldMap[namestring] = c
        }
    }

    return fieldMap

}

func cleanNonAlphanum(str string) string {

    reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
    if err != nil {
        log.Fatal(err)
    }

    processedStr := reg.ReplaceAllString(str, "")

    return processedStr

}
