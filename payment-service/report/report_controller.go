package report

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"text/template"
	"time"

	"poc-golang/payment-service/webservices/middleware"
	"poc-golang/payment-service/webservices/product"
)

// SetupRoutes :
func SetupRoutes(contextBase string) {
	reportHandler := http.HandlerFunc(handleReport)
	http.Handle(fmt.Sprintf("%s/reports", contextBase), middleware.Middleware(reportHandler))
}

func handleReport(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var filter product.ProductReportFilter
		err := json.NewDecoder(r.Body).Decode(&filter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		products, err := product.SearchByFilter(filter)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.New("report.gotmpl").Funcs(template.FuncMap{"mod": func(i, x int) bool { return i%x == 0 }})
		t, err = t.ParseFiles(path.Join("report", "templates", "report.gotmpl"))
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Disposition", "Attachment")

		var tpl bytes.Buffer
		err = t.Execute(&tpl, products)
		rdr := bytes.NewReader(tpl.Bytes())
		http.ServeContent(w, r, "report.html", time.Now(), rdr)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
