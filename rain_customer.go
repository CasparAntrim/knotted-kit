package main

import(
)

type RainCustomer struct {
    RainId              string      `csv:"Customer ID"`
    LastName            string      `csv:"Last Name"`
    FirstName           string      `csv:"First Name"`
    Company             string      `csv:"Company"`
    Balance             string      `csv:"Balance"`
    RewardPoints        string      `csv:"Reward Pts"`
    Birthdate           string      `csv:"Birthday"`
    Address             string      `csv:"Address"`
    City                string      `csv:"City"`
    Postcode            string      `csv:"Zip"`
    State               string      `csv:"State"`
    Country             string      `csv:"Country"`
    Phone               string      `csv:"Phone"`
    Mobile              string      `csv:"Cell"`
    Email               string      `csv:"E-mail"`
    Notes               string      `csv:"Notes"`
}

func(rc RainCustomer) ToVend() VendCustomer {

    vc := VendCustomer{
        "", // Customer Code
        rc.FirstName,
        rc.LastName,
        rc.Company,
        rc.Balance,
        "", // Credit
        "", // Sex
        rc.Birthdate,
        "", // LoyaltyEnabled
        rc.RewardPoints, // Not sure if this is appropriate
        "", // Year-to-date
        rc.Address,
        rc.City,
        rc.Postcode,
        rc.State,
        rc.Country,
        rc.Phone,
        rc.Mobile,
        rc.Email,
        "Rain", // GroupName
        rc.Notes,
        rc.RainId,
    }

    return vc

}
