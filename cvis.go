
package cvis

import (
	"os";
	"bytes";
	"io";
	"fmt";
	/*"./pics"*/
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
    for i := 0; i < len(v); i++ {
        if i%7 == 0 {
			if i == 0 {
                fmt.Fprintf(w, "<tr><td width=\"5%%\">0-0:</td>\n")
            } else {
				if (i/7)%7 ==0 {
					fmt.Fprintf(w, "</tr><tr></tr><tr></tr><tr></tr>");
				} else  {
					fmt.Fprintf(w, "</tr><tr></tr>");
				}
                fmt.Fprintf(w, "<tr><td  width=\"5%%\">%d-%d:</td>\n", (i/49), (i/7)%7)
            }
        }
		fmt.Fprintf(w, "<td border=\"1\">  </td>")
        if(v[i] == 0xFF) {
            fmt.Fprintf(w, "<td align=\"center\" bgcolor= #FF0000 border= \"1\"> %d </td>\n", i%7) /* red */
        } else if  v[i] != 0x00 {
            fmt.Fprintf(w, "<td align=\"center\" bgcolor= #32DF00 border= \"1\"> %d </td>\n", i%7) /* green */
        } else {
            fmt.Fprintf(w, "<td align=\"center\" bgcolor= #FFFF00 border= \"1\"> %d </td>\n", i%7) /* yellow */
        }
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
