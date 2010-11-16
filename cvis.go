
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
    fmt.Fprintf(w, "<table align=\"center\" width=\"100%%\">\n")
    for i := 0; i < len(v); i++ {
        if i%7 == 0 {
            if i == 0 {
                fmt.Fprintf(w, "<tr><td width=\"5%%\">%d:</td>\n",i)
            } else {
                fmt.Fprintf(w, "</tr><tr><td width=\"5%\">%d:</td>\n",i)
            }
        }
        if(v[i] == 0xFF) {
            fmt.Fprintf(w, "<td bgcolor= #FF0000 border= \"1\"> </td>\n") /* red */
        } else if  v[i] != 0x00 {
            fmt.Fprintf(w, "<td bgcolor= #32DF00 border= \"1\"> </td>\n") /* green */
        } else {
            fmt.Fprintf(w, "<td bgcolor= #FFFF00 border= \"1\"> </td>\n") /* yellow */
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
<head> <meta http-equiv="refresh" content="1"> </head>
{v|build}
</html>
`
