# File upload and download backend setup

Requirements: [link](https://github.com/vijaykrishnavanshi/assignment/blob/main/Backend-FullTime.md)


### Steps to run
- Initiate IPFS node in your machine
  - If IPFS cli is not installed in your machine, refer [here](https://github.com/ipfs/kubo/tree/master#install)
  - If you're using MACOS
    - ```$ brew install --formula ipfs```
    - ```$ ipfs init```
    - ```$ ipfs daemon```

- Once your IPFS mode is up and running, run the above program
  - ```go mod tidy```
  - ```go mod vendor```
  - ```go run main.go```

### Postman collection
Click [here](./postman_collection.json)

