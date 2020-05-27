package main

import(
    "fmt"
    "log"
    "os"
    "strings"
    "io/ioutil"
)

type Renamer struct {
    location_dir    string
    destination_dir string
    name_refs       []NameRef
}

func(r Renamer) ExecuteAndReport() int {

    var name_map map[string]string
    name_map = make(map[string]string)

    for _, nr := range r.name_refs {
        name_map[nr.Src] = nr.Dst
    }


    matchCount := 0
    for _, filename := range r.scanLocale() {
        if _, ok := name_map[undress(filename)]; ok {
            fmt.Printf("OK! Found a match for (%s)\n", filename)
            matchCount = matchCount + 1
            oldpath := fmt.Sprintf("%s/%s", r.location_dir, filename)
            newpath := fmt.Sprintf("%s/%s", r.destination_dir, dress(name_map[undress(filename)], "jpg"))

            err := os.Rename(oldpath, newpath)
            if err != nil {
                log.Fatal(err)
            }
        } else {
            continue
        }
    }

    return matchCount

}

func(r Renamer) scanLocale() []string {

    files, err := ioutil.ReadDir(r.location_dir)
    if err != nil {
        log.Fatal(err)
    }

    var names []string
    for _, file := range files {
        // fmt.Println(file.Name())
        names = append(names, file.Name())
    }

    return names

}

func undress(s string) string {

    ns := strings.Split(s, ".")
    final := strings.ToUpper(ns[0])
    // fmt.Println(final)

    return final
}

func dress(s string, ext string) string {

    ns := strings.ToLower(s)
    final := fmt.Sprintf("%s.%s", ns, ext)

    return final

}
