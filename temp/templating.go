package temp

import "fmt"

func CreateRudimentaryTemplate(name string) string {
	tmpl := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Document</title>
		</head>
		<body>
		<h1> %s </h1>
		</body>
		</html>`,
		name)
	return tmpl
}
