package main

const httpStatusOK = "200 OK"
const httpStatusNotFound = "404.1"


var requestTypes = []string{"/", "/2.jpg", "/1.png", "/index-html.html", "/123.txt", "/1.pdf", "/html5.png"}
var contentNames = []string{"index.html", "2.jpg", "1.png", "index-html.html", "123.txt", "1.pdf", "html5.png"}
var contentTypes = []string{"text/html", "image/jpg", "image/png", "text/html", "text/html", "application/pdf", "image/png"}