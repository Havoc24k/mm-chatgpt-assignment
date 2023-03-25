# mm-chatgpt-assignment
Miracle Mill ChatGPT Assignment

## Service Weaver `hello` example
> https://github.com/ServiceWeaver/weaver/tree/main/examples/hello

```
$ weaver generate .
$ go run .
```

## ChatGPT Code Review via Github Actions

Setup your ChatGPT API key in the project's secrets page and access it within the script like this: `CHATGPT_API_KEY: ${{ secrets.CHATGPT_API_KEY }}`

```
name: Code Review

on:
  pull_request:
    types: [opened, edited, reopened, synchronize]

jobs:
  code_review:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Install Python
      uses: actions/setup-python@v2
      with:
        python-version: '3.9'

    - name: Install dependencies
      run: pip install requests

    - name: Run code review
      env:
        PR_NUMBER: ${{ github.event.pull_request.number }}
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        CHATGPT_API_KEY: ${{ secrets.CHATGPT_API_KEY }}
      run: |
        set -e

        # Get the list of files changed in the pull request
        files=$(git diff --name-only HEAD~1)

        # Loop through each file and get its contents
        for file in $files; do
          contents=$(git show HEAD~1:$file)

          # Send the contents to ChatGPT for review
          response=$(curl -s -X POST -H "Authorization: Bearer $CHATGPT_API_KEY" -H "Content-Type: application/json" -d "{\"text\": \"$contents\"}" "https://api.openai.com/v1/engines/davinci-codex/completions?prompt=Please review the following code:&max_tokens=100&n=1")

          # Parse the response to get the review
          review=$(echo $response | jq -r '.choices[].text')

          # Post the review as a comment on the pull request
          curl -s -X POST -H "Authorization: Bearer $GITHUB_TOKEN" -H "Content-Type: application/json" -d "{\"body\": \"$review\"}" "https://api.github.com/repos/${{ github.repository }}/issues/${PR_NUMBER}/comments"
        done
```
