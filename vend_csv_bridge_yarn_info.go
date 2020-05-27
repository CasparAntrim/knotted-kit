package main

import(
    "fmt"
    // "strings"
)

type VendCsvBridgeYarnInfo struct {
    Sku             string  `csv:"sku"`
    Country         string  `csv:"mfg_country"`
    Weight          string  `csv:"yarn_weight"`
    SkeinWeight     string  `csv:"skein_weight"`
    Length          string  `csv:"length"`
    KnitGauge       string  `csv:"knitting_gauge"`
    NeedleSize      string  `csv:"needle_size"`
    CrochetGauge    string  `csv:"crochet_gauge"`
    HookSize        string  `csv:"hook_size"`
    Fiber           string  `csv:"yarn_fiber"`
    WashingInfo     string  `csv:"washing_info"`
    Construction    string  `csv:"construction"`
}

func(sh VendCsvBridgeYarnInfo) Describe() string {
    // Return formatted string for "Description" in CSV
    // fieldMap := make(map[string]string)
    // fieldMap["Country"] = sh.Country
    // fieldMap["Weight"] = sh.Weight
    // fieldMap["Skein Weight"] = sh.SkeinWeight
    // fieldMap["Length"] = sh.Length
    // fieldMap["Knit Gauge"] = sh.KnitGauge
    // fieldMap["Needle Size"] = sh.NeedleSize
    // fieldMap["Crochet Gauge"] = sh.CrochetGauge
    // fieldMap["Hook Size"] = sh.HookSize
    // fieldMap["Fiber"] = sh.Fiber
    // fieldMap["Washing Info"] = sh.WashingInfo
    // fieldMap["Construction"] = sh.Construction

    // var stringSlice []string
    // for k, v := range fieldMap {
    //     s := sh.wrapField(k, v)
    //     stringSlice = append(stringSlice, s)
    // }
    //
    // final := strings.Join(stringSlice, "")

    final := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s%s",
        sh.wrapField("Country:", sh.Country),
        sh.wrapField("Weight:", sh.Weight),
        sh.wrapField("Length:", sh.Length),
        sh.wrapField("Fiber:", sh.Fiber),
        sh.wrapField("Knit Gauge", sh.KnitGauge),
        sh.wrapField("Needle Size:", sh.NeedleSize),
        sh.wrapField("Crochet Gauge:", sh.CrochetGauge),
        sh.wrapField("Hook Size:", sh.HookSize),
        sh.wrapField("Construction:", sh.Construction),
        sh.wrapField("Skein Weight:", sh.SkeinWeight),
        sh.wrapField("Washing Info:", sh.WashingInfo))

    return final
}

func(sh VendCsvBridgeYarnInfo) wrapField(k string, v string) string {

    s := fmt.Sprintf("<p><strong>%s</strong><em> %s</em></p>", k, v)

    return s
}
