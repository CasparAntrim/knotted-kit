package main

import(
    "fmt"
    "log"
    "path"
    "strings"
    "io/ioutil"
)

type Linker struct {
    model       string
    prefix      string
    ref         string
    image_map   map[string]string
}

func(l *Linker) SetImageList() {

    l.image_map = l.scanForImageMap()
    fmt.Printf("\n%d images found in (%s)", len(l.image_map), l.ref)

}

func(l *Linker) ChangeLinks(linkable Linkable, sku string) {

    if _, ok := l.image_map[strings.ToUpper(sku)]; ok {
        link := fmt.Sprintf("%s%s", l.prefix, l.image_map[strings.ToUpper(sku)])
        linkable.ChangeLink(link)
    }

}

func(l Linker) scanForImageMap() map[string]string {

    var image_map map[string]string
    if isCsv(l.ref) {
        // Handle scanning csv for names
        image_map = l.scanCsv(l.ref)
    } else if path.IsAbs(l.ref) {
        // Handle scanning directory for names
        image_map = l.scanLocale(l.ref)
    } else {
        log.Fatal("You fucked up there, Jim.")
    }

    return image_map

}

func(l Linker) scanLocale(dir string) map[string]string {

    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    var names map[string]string
    names = make(map[string]string)
    for _, file := range files {

        names[undress(file.Name())] = file.Name()

    }

    return names

}

func(l Linker) scanCsv(ref string) map[string]string {

    var names map[string]string
    names = make(map[string]string)

    return names
}

func isCsv(p string) bool {

    ext := fmt.Sprintf("%s", path.Ext(p))
    if ext != ".csv" {
        return false
    } else {
        return true
    }

}
