package main

import()

type VendCustomer struct {
    Code                string      `csv:"customer_code"`
    FirstName           string      `csv:"first_name"`
    LastName            string      `csv:"last_name"`
    Company             string      `csv:"company_name"`
    Balance             string      `csv:"balance"`
    Credit              string      `csv:"store_credit_balance"`
    Sex                 string      `csv:"sex"`
    Birthdate           string      `csv:"date_of_birth"`
    LoyaltyEnabled      string      `csv:"enable_loyalty"`
    LoyaltyBalance      string      `csv:"loyalty_balance"`
    YearToDate          string      `csv:"year_to_date"`
    PhysicalAddress     string      `csv:"physical_address1"`
    PhysicalCity        string      `csv:"physical_city"`
    PhysicalPostcode    string      `csv:"physical_postcode"`
    PhysicalState       string      `csv:"physical_state"`
    PhysicalCountryId   string      `csv:"physical_country_id"`
    Phone               string      `csv:"phone"`
    Mobile              string      `csv:"mobile"`
    Email               string      `csv:"email"`
    GroupName           string      `csv:"customer_group_name"`
    Note                string      `csv:"note"`
    Custom1             string      `csv:"custom_field_1"`
}
