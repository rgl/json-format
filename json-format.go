// indent JSON from STDIN into STDOUT.
//
// compile with:
//
//  go build -ldflags "-s" json-format.go
//
// try with:
//
//  echo '{"Hello":1,"World":2}' | ./json-format
//
// NB for a Python version see https://gist.github.com/rgl/7895875

package main

import "bytes"
import "encoding/json"
import "io/ioutil"
import "log"
import "os"

func main() {
    unformatted, err := ioutil.ReadAll(os.Stdin)

    if err != nil {
    	log.Fatalf("ERROR Failed to read from STDIN: %s\n", err)
    }

    var formatted bytes.Buffer

    err = json.Indent(&formatted, unformatted, "", "  ")

    if err != nil {
    	log.Fatalf("ERROR %s\n", err)
    }

    formatted.WriteTo(os.Stdout)
}