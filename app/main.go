package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type Report struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedAt string `json:"createdAt"`
}

var reports []Report





func main() {

	gofakeit.Seed(0)

	r := mux.NewRouter()

	r.HandleFunc("/api/reports", getReports).Methods("GET")
	r.HandleFunc("/api/reports", createReport).Methods("POST")
	r.HandleFunc("/api/reports/html", getHTMLReports).Methods("GET")
	r.HandleFunc("/api/reports/{id}", getReport).Methods("GET")  // New route for fetching a specific report

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})

	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	http.ListenAndServe(":8080", nil)
}






func getReports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}





func createReport(w http.ResponseWriter, r *http.Request) {
	var report Report
	report.ID = uuid.NewV4().String()
	report.Title = gofakeit.Word()
	report.Content = gofakeit.Sentence(5)
	report.Status = gofakeit.Word()
	report.CreatedAt = gofakeit.Date().Format("2006-01-02 15:04:05")

	reports = append(reports, report)

	log.Println("Response:", report)

	jsonResponse, err := json.Marshal(report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getHTMLReports(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := `
<!DOCTYPE html>
<html>
<head>
	<style>
		table {
			font-family: Arial, sans-serif;
			border-collapse: collapse;
			width: 100%;
		}

		th, td {
			border: 1px solid #dddddd;
			text-align: left;
			padding: 8px;
		}

		th {
			background-color: #f2f2f2;
		}
	</style>
</head>
<body>
	<h2>Report Table</h2>
	<table>
		<tr>
			<th>ID</th>
			<th>Title</th>
			<th>Content</th>
			<th>Status</th>
			<th>Created At</th>
		</tr>
		{{ range . }}
			<tr>
				<td>{{ .ID }}</td>
				<td>{{ .Title }}</td>
				<td>{{ .Content }}</td>
				<td>{{ .Status }}</td>
				<td>{{ .CreatedAt }}</td>
			</tr>
		{{ end }}
	</table>
</body>
</html>
`

	tmpl, err := template.New("report").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log reports for debugging
	log.Println("Reports:", reports)

	// Execute the template without jsonResponse
	err = tmpl.Execute(w, reports)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getReport(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reportID := params["id"]
	log.Println("Requested Report ID:", reportID)

	// Find the report with the specified ID
	for _, report := range reports {
			// Print each report ID in the list
			log.Println("Report ID in List:", report.ID)

			if report.ID == reportID {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(report)
					return
			}
	}

	// Report not found
	http.NotFound(w, r)
}

