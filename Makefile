
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

SRC=cvis.go

all: cvis

cvis: 
#	$(GC) httppic.go
	$(GC) cvis.go
	$(GC) main.go
	$(LD) -o clusterVis  main.$O
	rm -f *.$O

test: $(TARG)
	./test.sh


