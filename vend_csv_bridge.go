package main

import(
)

type VendCsvBridge struct {
    Id                      string  `csv:"id"`
    Handle                  string  `csv:"handle"`
    Sku                     string  `csv:"sku"`
    Name                    string  `csv:"name"`
    Description             string  `csv:"description"`
    Type                    string  `csv:"type"`
    VarOptOneName           string  `csv:"variant_option_one_name"`
    VarOptOneVal            string  `csv:"variant_option_one_value"`
    Brand                   string  `csv:"brand_name"`
    Vendor                  string  `csv:"supplier_name"`
    Active                  string  `csv:"active"`
    Inventory               string  `csv:"inventory_Main_Outlet"`
}
