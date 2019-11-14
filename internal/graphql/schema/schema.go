// Package schema transforms *.graphql files into a single string
// Do not change go:generate... comment
package schema

import "bytes"

//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...

// GetRootSchema reads the .graphql schema files from the generated _bindata.go file, concatenating the
// files together into one string.
func GetRootSchema() string {
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		b := MustAsset(name)
		buf.Write(b)

		// Add a newline if the file does not end in a newline.
		if len(b) > 0 && b[len(b)-1] != '\n' {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
