module example.com/main

go 1.15

require (
	example.com/data v0.0.0-00010101000000-000000000000 // indirect
	example.com/handlers v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
)

replace example.com/handlers => ../handlers

replace example.com/data => ../data
