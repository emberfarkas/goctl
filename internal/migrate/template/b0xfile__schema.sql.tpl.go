// Code generaTed by fileb0x at "2021-11-22 18:10:37.789763796 +0800 CST m=+0.160127637" from config file "b0x.yaml" DO NOT EDIT.
// modified(2021-11-15 14:35:03.4687994 +0800 CST)
// original path: source/schema.sql.tpl

package template

import (
	"os"
)

// FileSchemaSQLTpl is "/schema.sql.tpl"
var FileSchemaSQLTpl = []byte("\x69\x6d\x70\x6f\x72\x74\x20\x72\x65\x71\x75\x65\x73\x74\x20\x66\x72\x6f\x6d\x20\x27\x40\x2f\x75\x74\x69\x6c\x73\x2f\x72\x65\x71\x75\x65\x73\x74\x27\x0d\x0a\x69\x6d\x70\x6f\x72\x74\x20\x71\x73\x20\x66\x72\x6f\x6d\x20\x27\x71\x73\x27\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe6\x9f\xa5\xe8\xaf\xa2\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\xe5\x88\x86\xe9\xa1\xb5\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x70\x61\x67\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x71\x75\x65\x72\x79\x29\x20\x7b\x0d\x0a\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x50\x61\x67\x65\x27\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x67\x65\x74\x27\x2c\x0d\x0a\x20\x20\x20\x20\x70\x61\x72\x61\x6d\x73\x3a\x20\x71\x75\x65\x72\x79\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe6\x9f\xa5\xe8\xaf\xa2\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\xe5\x88\x97\xe8\xa1\xa8\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x62\x61\x74\x63\x68\x67\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x71\x75\x65\x72\x79\x29\x20\x7b\x0d\x0a\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x4c\x69\x73\x74\x27\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x67\x65\x74\x27\x2c\x0d\x0a\x20\x20\x20\x20\x70\x61\x72\x61\x6d\x73\x3a\x20\x71\x75\x65\x72\x79\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe6\x9f\xa5\xe8\xaf\xa2\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\xe8\xaf\xa6\xe7\xbb\x86\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x67\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\x28\x7b\x7b\x2e\x50\x6b\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x29\x20\x7b\x0d\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x2f\x27\x20\x2b\x20\x7b\x7b\x2e\x50\x6b\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x67\x65\x74\x27\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe6\x96\xb0\xe5\xa2\x9e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x61\x64\x64\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x64\x61\x74\x61\x29\x20\x7b\x0d\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x27\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x70\x6f\x73\x74\x27\x2c\x0d\x0a\x20\x20\x20\x20\x64\x61\x74\x61\x3a\x20\x64\x61\x74\x61\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe4\xbf\xae\xe6\x94\xb9\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x75\x70\x64\x61\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x64\x61\x74\x61\x29\x20\x7b\x0d\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x27\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x70\x75\x74\x27\x2c\x0d\x0a\x20\x20\x20\x20\x64\x61\x74\x61\x3a\x20\x64\x61\x74\x61\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a\x2f\x2f\x20\xe5\x88\xa0\xe9\x99\xa4\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0d\x0a\x65\x78\x70\x6f\x72\x74\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x64\x65\x6c\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x7b\x7b\x2e\x50\x6b\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x29\x20\x7b\x0d\x0a\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x72\x65\x71\x75\x65\x73\x74\x28\x7b\x0d\x0a\x20\x20\x20\x20\x75\x72\x6c\x3a\x20\x27\x2f\x61\x64\x6d\x69\x6e\x2f\x76\x31\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x4e\x61\x6d\x65\x7d\x7d\x2f\x27\x20\x2b\x20\x7b\x7b\x2e\x50\x6b\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x2c\x0d\x0a\x20\x20\x20\x20\x6d\x65\x74\x68\x6f\x64\x3a\x20\x27\x64\x65\x6c\x65\x74\x65\x27\x0d\x0a\x20\x20\x7d\x29\x0d\x0a\x7d\x0d\x0a\x0d\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "/schema.sql.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileSchemaSQLTpl)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
