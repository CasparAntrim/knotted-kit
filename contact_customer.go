package main

import()

type ContactCustomer struct {
    Email               string  `csv:"Email address"`
    FirstName           string  `csv:"First name"`
    LastName            string  `csv:"Last name"`
    Lists               string  `csv:"Email Lists"`
}

func(cc ContactCustomer) ToRain() RainCustomer {

    rCust := RainCustomer{
        "",
        cc.LastName,
        cc.FirstName,
        "", // Company
        "", // Balance
        "", // Reward Points
        "", // Birthdate
        "", // Address
        "", // City
        "", // Postcode
        "", // State
        "", // Country
        "", // Phone
        "", // Mobile
        cc.Email,
        "", // Notes
    }

    return rCust

}
