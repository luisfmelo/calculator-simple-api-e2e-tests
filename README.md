# Simple Calculator API (unit & e2e test coverage POC)

This is a project intended to make unit and e2e tests to an API, get the coverage and merge them.

### Requests to use

* Health `curl http://localhost:9999`
* Sum: `curl -X POST http://localhost:9999/calc -H "Content-Type: application/json" -d '{"operation": "sum", "operand1": 1, "operand2": 3}'`
* Subtract `curl -X POST http://localhost:9999/calc -H "Content-Type: application/json" -d '{"operation": "subtract", "operand1": 5, "operand2": 3}'`

### Run Unit tests

`gotestsum --junitfile reports/unit.xml -- -tags unit -covermode=count -coverprofile=reports/unit.out ./...`

### Run e2e tests

`go test -coverpkg="./..." -c -tags testrunmain -covermode=count -coverprofile=reports/e2e.out ./...`
`./calculator-api.test -test.run "^TestRunMain$" -test.coverprofile=reports/e2e.out`

On another terminal window run:
`sh ./e2e_tests.sh`

On the first window, terminate with: **CTRL + C**

### Merge Coverages
`gocovmerge reports/unit.out reports/e2e.out > reports/coverage.out`

### Display coverage

`go tool cover -func=reports/unit.out`

`go tool cover -func=reports/e2e.out`

`go tool cover -func=reports/coverage.out`
