# playground-for-go
Simple examples that help you get started with Appwrite + Go (=❤️)

**Work in progress**

## Get Started

This playground doesn't include any Go best practices but rather it's intended to show some of the most simple examples and use cases of using the Appwrite API in your Go application.

## Requirements
* A system with Go or Docker installed.
* An Appwrite instance.
* An Appwrite project created in the console.
* An Appwrite API key created in the console.

## Installation

1. Clone this repository.
2. `cd` into the repository.
3. Open the `playground.go` file found in the root of the cloned repository.
4. Copy Project ID, endpoint and API key from Appwrite console into `playground.go`
5. Run the playground:
    - Go:
        - Install dependencies `go get`
        - Execute the command `go run .`
    - Docker:
        - Execute the command `docker compose up`
6. You will see the JSON response in the console.

## API's Covered

- Database
    * Create Collection
    * List Collections
    * Delete Collection
    * Create Document
    * List Documents
    * Delete Document

- Storage
    * Create Bucket
    * List Buckets
    * Delete Bucket
    * Upload File
    * List Files
    * Delete File

- Users
    * Create User
    * List Users
    * Delete User

- Functions
    * Create Function
    * List Functions
    * Delete Function
    * Upload Deployment
    * Execute function (sync)
    * Execute function (async)

## Contributing

All code contributions - including those of people having commit access - must go through a pull request and approved by a core developer before being merged. This is to ensure proper review of all the code.

We truly ❤️ pull requests! If you wish to help, you can learn more about how you can contribute to this project in the [contribution guide](https://github.com/appwrite/appwrite/blob/master/CONTRIBUTING.md).

## Security

For security issues, kindly email us [security@appwrite.io](mailto:security@appwrite.io) instead of posting a public issue in GitHub.

## Follow Us

Join our growing community around the world! Follow us on [X](https://x.com/appwrite), [LinkedIn](https://www.linkedin.com/company/appwrite/mycompany/), or join our [Discord Server](https://appwrite.io/discord) for more help, ideas and discussions.