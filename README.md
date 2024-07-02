# Gogol

<sub>Forked from [aquemaati/gogol](https://github.com/aquemaati/gogol)</sub>

Create projects faster than ever !

## Overview

- [Description](#description)

- [Requirements](#requirements)

- [Installation](#installation)

- [Usage](#usage)

  1. [Languages](#languages)
  2. [Miscellaneous](#miscellaneous)

### Description

This project shares the same name and idea as [aquemaati/gogol](https://github.com/aquemaati/gogol). It is a fork with different features and syntax.

### Requirements

[Golang](https://go.dev/) 1.22 or later.

### Installation

After installing Golang, you will need to install Gogol using:

```sh
go install github.com/cramanan/gogol@latest
```

### Usage

#### Languages

To create a new project, use:

```sh
gogol [language] [flags]
```

Example:

- HTML/CSS/JS project with a README.md and a LICENSE.md:

  ```sh
  gogol html -rl
  ```

- Golang project with a makefile, dockerfile and a tests folder:

  ```sh
  gogol go -tmd
  ```

Some languages have subcommands for specific types of projects:

- Golang web project with an HTTP server:

  ```sh
    gogol go web
  ```

#### Miscellaneous

export / import : The export function exports the inputed directory and saves it into a .json file (e.g.)

```sh
gogol export example
```

creates the file:

```json
{
  "name": "untitled",
  "directories": {},
  "files": {
    "go.mod": {
      "name": "go.mod",
      "content": "bW9kdWxlIHVudGl0bGVkCgpnbyAxLjE5Cg=="
    },
    "main.go": {
      "name": "main.go",
      "content": "cGFja2FnZSBtYWluCgppbXBvcnQgImZtdCIKCQkJCmZ1bmMgbWFpbigpewoJZm10LlByaW50bG4oIkhlbGxvIFdvcmxkIikKfQ=="
    }
  }
}
```
