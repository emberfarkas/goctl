// Code generaTed by fileb0x at "2023-08-17 15:11:52.079616 +0800 CST m=+0.002121751" from config file "b0x.yaml" DO NOT EDIT.
// modified(2022-12-28 21:00:30.637204591 +0800 CST)
// original path: source/dao.go.tpl

package template

import (
	"os"
)

// FileDaoGoTpl is "/dao.go.tpl"
var FileDaoGoTpl = []byte("\x70\x61\x63\x6b\x61\x67\x65\x20\x64\x61\x6f\x0a\x0a\x69\x6d\x70\x6f\x72\x74\x20\x28\x0a\x20\x20\x20\x20\x22\x63\x6f\x6e\x74\x65\x78\x74\x22\x0a\x20\x20\x20\x20\x7b\x7b\x24\x69\x69\x20\x3a\x3d\x20\x31\x7d\x7d\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x7d\x7d\x0a\x20\x20\x20\x20\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x74\x69\x6d\x65\x2e\x54\x69\x6d\x65\x22\x20\x2d\x7d\x7d\x7b\x7b\x2d\x20\x69\x66\x20\x65\x71\x20\x24\x69\x69\x20\x31\x20\x2d\x7d\x7d\x20\x22\x74\x69\x6d\x65\x22\x20\x7b\x7b\x24\x69\x69\x20\x3d\x20\x32\x7d\x7d\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x0a\x20\x20\x20\x20\x22\x62\x6c\x73\x2f\x70\x6b\x67\x2f\x74\x6f\x6f\x6c\x73\x22\x0a\x20\x20\x20\x20\x22\x62\x6c\x73\x2f\x70\x6b\x67\x2f\x65\x63\x6f\x64\x65\x22\x0a\x20\x20\x20\x20\x22\x62\x6c\x73\x2f\x73\x65\x72\x76\x69\x63\x65\x2f\x7b\x7b\x2e\x4d\x6f\x64\x75\x6c\x65\x7d\x7d\x2f\x69\x6e\x74\x65\x72\x6e\x61\x6c\x2f\x6d\x6f\x64\x65\x6c\x22\x0a\x29\x0a\x0a\x2f\x2f\x50\x61\x67\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe8\x8e\xb7\xe5\x8f\x96\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\xe5\xb8\xa6\xe5\x88\x86\xe9\xa1\xb5\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x50\x61\x67\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x66\x69\x6c\x74\x65\x72\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x70\x61\x67\x65\x53\x69\x7a\x65\x20\x69\x6e\x74\x2c\x20\x70\x61\x67\x65\x49\x6e\x64\x65\x78\x20\x69\x6e\x74\x29\x20\x28\x64\x6f\x63\x73\x20\x5b\x5d\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x74\x6f\x74\x61\x6c\x20\x69\x6e\x74\x36\x34\x2c\x20\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x66\x69\x6c\x74\x65\x72\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x2e\x49\x73\x51\x75\x65\x72\x79\x20\x7d\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x73\x74\x72\x69\x6e\x67\x22\x20\x2d\x7d\x7d\x20\x22\x22\x20\x7b\x7b\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x69\x6e\x74\x22\x20\x2d\x7d\x7d\x20\x30\x20\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x4e\x61\x6d\x65\x7d\x7d\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x45\x51\x22\x7d\x7d\x20\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4e\x45\x22\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x22\x7d\x7d\x20\x3e\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x45\x22\x7d\x7d\x20\x3e\x3d\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x22\x7d\x7d\x20\x3c\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x45\x22\x7d\x7d\x20\x3c\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x20\x6c\x69\x6b\x65\x20\x7b\x7b\x65\x6e\x64\x7d\x7d\x3f\x22\x2c\x20\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x22\x25\x22\x2b\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x2b\x22\x25\x22\x7b\x7b\x65\x6c\x73\x65\x7d\x7d\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x7b\x7b\x65\x6e\x64\x7d\x7d\x29\x0a\x20\x20\x20\x20\x7d\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x60\x64\x65\x6c\x65\x74\x65\x64\x5f\x61\x74\x60\x20\x49\x53\x20\x4e\x55\x4c\x4c\x22\x29\x2e\x43\x6f\x75\x6e\x74\x28\x26\x74\x6f\x74\x61\x6c\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x4f\x66\x66\x73\x65\x74\x28\x28\x70\x61\x67\x65\x49\x6e\x64\x65\x78\x20\x2d\x20\x31\x29\x20\x2a\x20\x70\x61\x67\x65\x53\x69\x7a\x65\x29\x2e\x4c\x69\x6d\x69\x74\x28\x70\x61\x67\x65\x53\x69\x7a\x65\x29\x2e\x46\x69\x6e\x64\x28\x26\x64\x6f\x63\x73\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x42\x61\x74\x63\x68\x47\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe8\x8e\xb7\xe5\x8f\x96\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\xe5\x88\x97\xe8\xa1\xa8\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x42\x61\x74\x63\x68\x47\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x29\x20\x28\x64\x6f\x63\x73\x20\x5b\x5d\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x2e\x49\x73\x51\x75\x65\x72\x79\x20\x7d\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x73\x74\x72\x69\x6e\x67\x22\x20\x2d\x7d\x7d\x20\x22\x22\x20\x7b\x7b\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x69\x6e\x74\x22\x20\x2d\x7d\x7d\x20\x30\x20\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x4e\x61\x6d\x65\x7d\x7d\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x45\x51\x22\x7d\x7d\x20\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4e\x45\x22\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x22\x7d\x7d\x20\x3e\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x45\x22\x7d\x7d\x20\x3e\x3d\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x22\x7d\x7d\x20\x3c\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x45\x22\x7d\x7d\x20\x3c\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x20\x6c\x69\x6b\x65\x20\x7b\x7b\x65\x6e\x64\x7d\x7d\x3f\x22\x2c\x20\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x22\x25\x22\x2b\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x2b\x22\x25\x22\x7b\x7b\x65\x6c\x73\x65\x7d\x7d\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x7b\x7b\x65\x6e\x64\x7d\x7d\x29\x0a\x20\x20\x20\x20\x7d\x0a\x7b\x7b\x20\x65\x6e\x64\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x46\x69\x6e\x64\x28\x26\x64\x6f\x63\x73\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x47\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe8\x8e\xb7\xe5\x8f\x96\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x47\x65\x74\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x29\x20\x28\x72\x65\x74\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x7b\x7b\x20\x72\x61\x6e\x67\x65\x20\x2e\x43\x6f\x6c\x75\x6d\x6e\x73\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x24\x78\x20\x3a\x3d\x20\x2e\x50\x6b\x7d\x7d\x0a\x7b\x7b\x2d\x20\x69\x66\x20\x28\x24\x78\x29\x20\x7d\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x73\x74\x72\x69\x6e\x67\x22\x20\x2d\x7d\x7d\x20\x22\x22\x20\x7b\x7b\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x75\x69\x6e\x74\x36\x34\x22\x20\x2d\x7d\x7d\x20\x30\x20\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x4e\x61\x6d\x65\x7d\x7d\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x45\x51\x22\x7d\x7d\x20\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4e\x45\x22\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x22\x7d\x7d\x20\x3e\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x45\x22\x7d\x7d\x20\x3e\x3d\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x22\x7d\x7d\x20\x3c\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x45\x22\x7d\x7d\x20\x3c\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x20\x6c\x69\x6b\x65\x20\x7b\x7b\x65\x6e\x64\x7d\x7d\x3f\x22\x2c\x20\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x22\x25\x22\x2b\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x2b\x22\x25\x22\x7b\x7b\x65\x6c\x73\x65\x7d\x7d\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x7b\x7b\x65\x6e\x64\x7d\x7d\x29\x0a\x20\x20\x20\x20\x7d\x0a\x7b\x7b\x2d\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x2e\x49\x73\x51\x75\x65\x72\x79\x20\x7d\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x73\x74\x72\x69\x6e\x67\x22\x20\x2d\x7d\x7d\x20\x22\x22\x20\x7b\x7b\x20\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x47\x6f\x54\x79\x70\x65\x20\x22\x69\x6e\x74\x22\x20\x2d\x7d\x7d\x20\x30\x20\x7b\x7b\x2d\x20\x65\x6e\x64\x7d\x7d\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x43\x6f\x6c\x75\x6d\x6e\x4e\x61\x6d\x65\x7d\x7d\x7b\x7b\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x45\x51\x22\x7d\x7d\x20\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4e\x45\x22\x7d\x7d\x20\x21\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x22\x7d\x7d\x20\x3e\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x47\x54\x45\x22\x7d\x7d\x20\x3e\x3d\x20\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x22\x7d\x7d\x20\x3c\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x54\x45\x22\x7d\x7d\x20\x3c\x3d\x20\x7b\x7b\x65\x6c\x73\x65\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x20\x6c\x69\x6b\x65\x20\x7b\x7b\x65\x6e\x64\x7d\x7d\x3f\x22\x2c\x20\x7b\x7b\x20\x69\x66\x20\x65\x71\x20\x2e\x51\x75\x65\x72\x79\x54\x79\x70\x65\x20\x22\x4c\x49\x4b\x45\x22\x7d\x7d\x22\x25\x22\x2b\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x2b\x22\x25\x22\x7b\x7b\x65\x6c\x73\x65\x7d\x7d\x65\x2e\x7b\x7b\x2e\x47\x6f\x46\x69\x65\x6c\x64\x7d\x7d\x7b\x7b\x65\x6e\x64\x7d\x7d\x29\x0a\x20\x20\x20\x20\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x2d\x7d\x7d\x0a\x7b\x7b\x2d\x20\x65\x6e\x64\x20\x7d\x7d\x0a\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x64\x6f\x63\x20\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x46\x69\x72\x73\x74\x28\x26\x64\x6f\x63\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x20\x3d\x20\x26\x64\x6f\x63\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x43\x72\x65\x61\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe5\x88\x9b\xe5\xbb\xba\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x43\x72\x65\x61\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x29\x20\x28\x64\x6f\x63\x20\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x43\x72\x65\x61\x74\x65\x28\x26\x65\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x64\x6f\x63\x20\x3d\x20\x2a\x65\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x55\x70\x64\x61\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe6\x9b\xb4\xe6\x96\xb0\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x55\x70\x64\x61\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x69\x64\x20\x75\x69\x6e\x74\x36\x34\x29\x20\x28\x75\x70\x64\x61\x74\x65\x20\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x50\x6b\x43\x6f\x6c\x75\x6d\x6e\x7d\x7d\x20\x3d\x20\x3f\x22\x2c\x20\x69\x64\x29\x2e\x46\x69\x72\x73\x74\x28\x26\x75\x70\x64\x61\x74\x65\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x2f\x2f\xe5\x8f\x82\xe6\x95\xb0\x31\x3a\xe6\x98\xaf\xe8\xa6\x81\xe4\xbf\xae\xe6\x94\xb9\xe7\x9a\x84\xe6\x95\xb0\xe6\x8d\xae\x0a\x20\x20\x20\x20\x2f\x2f\xe5\x8f\x82\xe6\x95\xb0\x32\x3a\xe6\x98\xaf\xe4\xbf\xae\xe6\x94\xb9\xe7\x9a\x84\xe6\x95\xb0\xe6\x8d\xae\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x4d\x6f\x64\x65\x6c\x28\x26\x75\x70\x64\x61\x74\x65\x29\x2e\x55\x70\x64\x61\x74\x65\x73\x28\x26\x65\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x44\x65\x6c\x65\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe5\x88\xa0\xe9\x99\xa4\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x44\x65\x6c\x65\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x69\x64\x20\x75\x69\x6e\x74\x36\x34\x29\x20\x28\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x50\x6b\x43\x6f\x6c\x75\x6d\x6e\x7d\x7d\x20\x3d\x20\x3f\x22\x2c\x20\x69\x64\x29\x2e\x44\x65\x6c\x65\x74\x65\x28\x26\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x7b\x7d\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a\x0a\x2f\x2f\x42\x61\x74\x63\x68\x44\x65\x6c\x65\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x20\xe6\x89\xb9\xe9\x87\x8f\xe5\x88\xa0\xe9\x99\xa4\x0a\x66\x75\x6e\x63\x20\x28\x64\x20\x2a\x64\x61\x6f\x29\x20\x42\x61\x74\x63\x68\x44\x65\x6c\x65\x74\x65\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x28\x63\x74\x78\x20\x63\x6f\x6e\x74\x65\x78\x74\x2e\x43\x6f\x6e\x74\x65\x78\x74\x2c\x20\x65\x20\x2a\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x2c\x20\x69\x64\x20\x5b\x5d\x75\x69\x6e\x74\x36\x34\x29\x20\x28\x65\x72\x72\x20\x65\x72\x72\x6f\x72\x29\x20\x7b\x0a\x20\x20\x20\x20\x74\x61\x62\x6c\x65\x20\x3a\x3d\x20\x64\x2e\x6f\x72\x6d\x2e\x57\x69\x74\x68\x43\x6f\x6e\x74\x65\x78\x74\x28\x63\x74\x78\x29\x2e\x54\x61\x62\x6c\x65\x28\x65\x2e\x54\x61\x62\x6c\x65\x4e\x61\x6d\x65\x28\x29\x29\x0a\x20\x20\x20\x20\x69\x66\x20\x65\x72\x72\x20\x3d\x20\x74\x61\x62\x6c\x65\x2e\x57\x68\x65\x72\x65\x28\x22\x7b\x7b\x2e\x50\x6b\x43\x6f\x6c\x75\x6d\x6e\x7d\x7d\x20\x69\x6e\x20\x28\x3f\x29\x22\x2c\x20\x69\x64\x29\x2e\x44\x65\x6c\x65\x74\x65\x28\x26\x6d\x6f\x64\x65\x6c\x2e\x7b\x7b\x2e\x43\x6c\x61\x73\x73\x4e\x61\x6d\x65\x7d\x7d\x7b\x7d\x29\x2e\x45\x72\x72\x6f\x72\x3b\x20\x65\x72\x72\x20\x21\x3d\x20\x6e\x69\x6c\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x0a\x7d\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "/dao.go.tpl", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileDaoGoTpl)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
