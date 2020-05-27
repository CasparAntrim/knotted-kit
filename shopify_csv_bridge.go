package main

import(
)

type ShopifyCsvBridge struct {
    Handle                  string  `csv:"Handle"`
    Title                   string  `csv:"Title"`
    Description             string  `csv:"Description"`
    Vendor                  string  `csv:"Vendor"`
    Type                    string  `csv:"Type"`
    Published               string  `csv:"Published"`
    Sku                     string  `csv:"Variant SKU"`
    Price                   string  `csv:"Variant Price"`
    Image                   string  `csv:"Image Src"`
}

func(sh *ShopifyCsvBridge) ChangeLink(link string) {
    sh.Image = link
}
