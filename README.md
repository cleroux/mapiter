# mapiter

mapiter is a Go static analysis tool to find instances where `range` is used to iterate over a map.

## Installation

    go get -u github.com/cleroux/mapiter
    
## Usage

    make install
    go vet -vettool=$(which mapiter) ./...

### Flags

- **-tests** (default true) - Include test files in analysis
