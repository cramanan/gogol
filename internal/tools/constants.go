package tools

const (
	GODEFAULT = `package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
`

	PYTHONDEFAULT = `if __name__ == "__main__":
	print("Hello World")
`

	HTMLDEFAULT = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1.0" />
        <title>Document</title>
        <link rel="stylesheet" href="style.css" />
        <script src="script.js" defer></script>
    </head>
    <body>
        <h1>Hello World</h1>
    </body>
</html>
`

	CSSDEFAULT = `*,
*::before,
*::after {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

`
)
