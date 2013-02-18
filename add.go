package sokeng

import (
    "database/sql"
    "html/template"
    "strings"
)

type add struct{
        content string
        whokeng string
        email string
        images []imgurl

}

func (this *add) Add(w http.ResponseWriter, r *http.Request){

}


