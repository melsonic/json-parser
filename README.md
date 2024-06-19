### JSON PARSER (from the coding challenges)

- A JSON parser comprises mainly two parts:
    - **Lexical Analysis**: In this step, the input is converted into an array of tokens.
    - **Syntactic Analysis**: In this step, the tokens are parsed to check whether they align with the rules of the specified language.

### Application Details

- Currently, the software outputs whether the given JSON file is valid or not.
- To test the application yourself, you can put your JSON content inside the `demo.json` file and run `make run`. It will tell you whether the content is valid JSON or not.
- Additionally, there are plenty of test cases inside the `tests` directory. You can run them individually using the command `go run main.go relative-filepath`.
- I have also prepared a bash script for running all the test cases inside the `tests` directory. You can run it using either the `./tests.sh` command, but first make sure you run `make build`. You can also run `make test` which is just an addition of these two commands.

### Working of the Application

- In the first step, the input file content is broken into an array of tokens.
- In the next step, the tokens are cleaned up by trimming the whitespaces around them.
- The tokens are then passed into a parser function that parses the tokens to see if they align with the specified language rules.
