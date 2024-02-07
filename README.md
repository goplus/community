Go+ Community written in Go+
=====

[![Build Status](https://github.com/goplus/community/actions/workflows/go.yml/badge.svg)](https://github.com/goplus/community/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/goplus/community)](https://goreportcard.com/report/github.com/goplus/community)
[![GitHub release](https://img.shields.io/github/v/tag/goplus/community.svg?label=release)](https://github.com/goplus/community/releases)
[![Coverage Status](https://codecov.io/gh/goplus/community/branch/main/graph/badge.svg)](https://codecov.io/gh/goplus/community)
[![GoDoc](https://pkg.go.dev/badge/github.com/goplus/community.svg)](https://pkg.go.dev/github.com/goplus/community)

### 🧳 Features

Go+ Community written in Go+ and Yap, support Go+ Community development.


### 🚀 How to run

1. **Prepare golang and goplus environment.**

- golang version >= 1.19
  - https://go.dev/dl/ (golang download, Find the golang 1.19 version)
- goplus version == v1.2.0-pre.1
  - Ref: https://github.com/goplus/gop/releases
  - How to install goplus: https://github.com/goplus/gop#how-to-install

2. **Prepare for your own config**

- Clone a file named `.env` from `.env_temp` in the `cmd/gopcomm` directory of the project.
- Modify the `.env` file to your own configuration.

3. **Run the project**

- Run the project with the following command:

```shell
gop run .
```

- Open the browser and enter the address: `http://localhost:8080` or `GOP_COMMUNITY_ENDPOINT` in your `.env` file.


### 📦 Contribute

1. **Fork the repository to your local repo**

2. **Config code linter**
  
    ```bash
    // install pre-commit
    pip install pre-commit

    // config code linter in pre-commit hook
    pre-commit install
    ```

3. **Modify your own code**

4. **Commit your code**

5. **Create a pull request**

> **Note:** Please check the PR brach is `mvp-20240119` or not.