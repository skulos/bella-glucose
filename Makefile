all:
	go run main.go
	wkhtmltoimage chart.html chart.png
	wkhtmltopdf chart.html chart.pdf