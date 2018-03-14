
package main

import (
	"html/template"
	//"fmt"
	"net/http"
	"strconv"
	//"strings"
)




func main() {
	var n1,n2,n3 int64
	var a,b string
	var i int
	tmpl := template.Must(template.ParseFiles("re.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		
		s := r.FormValue("ans")
		//fmt.Fprintf(w,"%s", tmpl)
		
		for i=0; i<len(s); i++ {
                    if s[i]=='+' || s[i]=='-' || s[i]=='*' || s[i]=='/' {
                       a = s[0:i]
                       b = s[i+1:len(s)]
		       n1, _ = strconv.ParseInt(a, 10, 64)
		       n2, _ = strconv.ParseInt(b, 10, 64)
		       	
			break
                     
			 
                       }
                 } 
		
	if i!=len(s)  {	
		
		if s[i]=='+' {
                     n3 = n1+n2
		     a = strconv.FormatInt(n3, 10)    
      	}	else if s[i]=='-'  {
                    n3 = n1-n2
		    a = strconv.FormatInt(n3, 10)             
	}	else if s[i]=='*'   {
                     n3 = n1*n2
			a = strconv.FormatInt(n3, 10)
	}	else if s[i]=='/' {
                      if n2!=0 {
			n3 = n1/n2
			a = strconv.FormatInt(n3, 10)
        }        else  {a = "division by zero not possible"}
        }   }    else {
                  
		  n3, _ = strconv.ParseInt(s, 10, 64)
		  a = strconv.FormatInt(n3, 10)	 
		}

               
		
	
          tmpl.Execute(w, a)

	

    
		
	})

	http.ListenAndServe(":8080", nil)
}
