#!/bin/sh

(cd tools/ && go build data2golang.go)

(cd text/ && /bin/ls -1 *.txt | while read T; do
  O=${T%%.txt}
  ../tools/data2golang --pkgname gotexttestdata $T $O > ../${O}.go
done)
