Project Structure:
go-bookstore/
│── backend/                # Go backend
│   │── main.go
│   │── go.mod              # Module file here
│   │── go.sum
│   │── pkg/
│   └── static/
│── frontend/               # HTML, CSS, JS files (if separate)
│── README.md

Control Flow
	user => routes => controllers => models (operations on database)