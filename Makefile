CODE_COVERAGE_FILE = coverage.out

# code coverage
.PHONY: cc
cc:
	go test -coverprofile=$(CODE_COVERAGE_FILE)
	go tool cover -html=$(CODE_COVERAGE_FILE)
