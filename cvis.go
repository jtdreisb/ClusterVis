package main

import (
	"os";
	"bytes";
	"io";
	"fmt";
	"template";
)

type mytable struct{
	v []byte
}
var fmap = template.FormatterMap{
    "build": buildTable,
}

func buildTable(w io.Writer, data interface{}, formatter string) {
    v,ok := data.([]byte)
	if !ok {
		fmt.Printf("unable to get a byte array!\n")
		os.Exit(1);
	}
    fmt.Fprintf(w, "<table height=\"100%%\" width=\"100%%\">\n")
	var iovero int
	for i := 0; i < len(v)/49; i++ {
		fmt.Fprintf(w, "<tr><td width=\"5%%\">%d:</td>\n",i)
		for j:= 0; j < 7; j++ { // 7 overos at position j 
			fmt.Fprintf(w, "</tr><tr><td width=\"5%%\">\n")
			for k:=0; k < 7; k++ { // index through each card on level 
				iovero = 49*i + 7*k + j

				if(v[iovero] == 0xFF) { // Oh shit this indexing is a hack
					fmt.Fprintf(w, "<td align=\"center\" bgcolor= #FF0000 border= \"1\"> %d </td>\n", iovero+1) /* red */
				} else if  v[iovero] != 0x00 {
					fmt.Fprintf(w, "<td align=\"center\" bgcolor= #32DF00 border= \"1\"> %d </td>\n", iovero+1) /* green */
				} else {
					fmt.Fprintf(w, "<td align=\"center\" bgcolor= #FFFF00 border= \"1\"> %d </td>\n", iovero+1) /* yellow */
				}

			}
		}
		fmt.Fprintf(w, "</tr><tr></tr>")
	}
		fmt.Fprintf(w, "</table\n")
}

func Update(b []byte) (f *bytes.Buffer) {
	templ,_ :=  template.Parse(templatestr, fmap)
		/*	
	templ.rdelim = []byte{ '}', '}'}
	templ.ldelim = []byte{ '{', '{'}
	*/
	var table mytable
	f = bytes.NewBuffer(nil)
	table.v = b
	templ.Execute(table,f)
	return
}


const templatestr = `
<html>
<head> <meta http-equiv="refresh" content="0.1"> </head>
{v|build}
</html>
`
