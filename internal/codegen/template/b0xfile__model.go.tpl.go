// Code generaTed by fileb0x at "2023-08-17 15:11:52.08009 +0800 CST m=+0.002596460" from config file "b0x.yaml" DO NOT EDIT.
// modified(2022-12-28 21:00:30.637309259 +0800 CST)
// original path: source/model.go.tpl

package template

import (
	"os"
)

// FileModelGoTpl is "/model.go.tpl"
var FileModelGoTpl = []byte("\x70\x61\x63\x6b\x61\x67\x65\x20\x6d\x6f\x64\x65\x6c\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x28\x0a\x7b\x7b\x24\x69\x69\x20\x3a\x3d\x20\x31\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x74\x69\x6d\x65\x2e\x54\x69\x6d\x65\x22\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x24\x69\x69\x20\x31\x20\x2d\x7d\x7d\x20\x0a\x20\x20\x20\x20\x22\x74\x69\x6d\x65\x22\x20\x0a\x20\x20\x20\x20\x7b\x7b\x2d\x20\x24\x69\x69\x20\x3d\x20\x32\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x29\x0a\x0a\x74\x79\x70\x65\x20\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\x73\x74\x72\x75\x63\x74\x20\x7b\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x2d\x7d\x7d\x20\x0a\x7b\x7b\x2d\x20\x24\x78\x20\x3a\x3d\x20\x2e\x50\x6b\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x28\x24\x78\x29\x20\x7d\x7d\x0a\x20\x20\x20\x20\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x7b\x7b\x2e\x47\x6f\x54\x79\x70\x65\x7d\x7d\x20\x60\x6a\x73\x6f\x6e\x3a\x22\x7b\x7b\x2e\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x22\x20\x67\x6f\x72\x6d\x3a\x22\x74\x79\x70\x65\x3a\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x54\x79\x70\x65\x7d\x7d\x3b\x70\x72\x69\x6d\x61\x72\x79\x5f\x6b\x65\x79\x22\x60\x20\x2f\x2f\x20\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x43\x6f\x6d\x6d\x65\x6e\x74\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x46\x69\x65\x6c\x64\x20\x22\x43\x72\x65\x61\x74\x65\x64\x41\x74\x22\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x46\x69\x65\x6c\x64\x20\x22\x55\x70\x64\x61\x74\x65\x64\x41\x74\x22\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x46\x69\x65\x6c\x64\x20\x22\x44\x65\x6c\x65\x74\x65\x64\x41\x74\x22\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x7d\x7d\x20\x0a\x20\x20\x20\x20\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x7b\x7b\x2e\x47\x6f\x54\x79\x70\x65\x7d\x7d\x20\x60\x6a\x73\x6f\x6e\x3a\x22\x7b\x7b\x2e\x4a\x73\x6f\x6e\x46\x69\x65\x6c\x64\x7d\x7d\x22\x20\x67\x6f\x72\x6d\x3a\x22\x74\x79\x70\x65\x3a\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x54\x79\x70\x65\x7d\x7d\x3b\x22\x60\x20\x2f\x2f\x20\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x43\x6f\x6d\x6d\x65\x6e\x74\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x20\x20\x20\x20\x42\x61\x73\x65\x4d\x6f\x64\x65\x6c\x0a\x7d\x0a\x0a\x66\x75\x6e\x63\x20\x28\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x29\x20\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x20\x73\x74\x72\x69\x6e\x67\x20\x7b\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x22\x7b\x7b\x2e\x54\x42\x4e\x61\x6d\x65\x7d\x7d\x22\x0a\x7d\x0a\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "/model.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileModelGoTpl)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
