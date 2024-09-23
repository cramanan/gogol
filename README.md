# Gogol

<sub>Forked from [aquemaati/gogol](https://github.com/aquemaati/gogol)</sub>

Create projects faster than ever !

## Overview

-   [Description](#description)

-   [Requirements](#requirements)

-   [Installation](#installation)

-   [Usage](#usage)

    1. [Languages](#languages)

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

-   HTML/CSS/JS project with a README.md and a LICENSE.md:

    ```sh
    gogol html -rl
    ```

-   Golang project with a makefile, dockerfile and a tests folder:

    ```sh
    gogol go -tmd
    ```
