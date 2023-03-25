# mm-chatgpt-assignment
Miracle Mill ChatGPT Assignment

To trigger a code review by ChatGPT simply create a new pull request. The Github Action will install https://github.com/anc95/ChatGPT-CodeReview anbd execute the code review.

## Service Weaver `hello` example
> https://github.com/ServiceWeaver/weaver/tree/main/examples/hello

```
$ weaver generate .
$ go run .
``

## Github Actions definition
> `.github/workflows/main.yml`

```
name: ChatGPT Code Review

permissions:
  contents: read
  pull-requests: write

on:
  pull_request:
    types: [opened, reopened, synchronize]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: anc95/ChatGPT-CodeReview@main
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          OPENAI_API_KEY: ${{ secrets.CHATGPT_API_KEY }}
          # Optional
          LANGUAGE: English
          MODEL: gpt-3.5-turbo
          top_p: 1
          temperature: 1
```
