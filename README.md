# Protos-common

The Kitt common proto library

## Adding a new common message
1. Add the message to `./common.proto`
2. If you want some custom logic:
   1. Add the message option:
    ```protobuf
    option (graphql.skip_message) = true;
    ```
   2. Create a file for your custom logic: `./common/<message_name>.graphql.go`
   3. To work with the generation library it needs to implement 3 things:
    ```go
    func <Message_name>FromArgs(field map[string]interface{}) *<Message_name> {}
    var <Message_name>GraphqlInputType gql.InputObject{}
    var <Message_name>GraphqlType gql.Object{}
    ```
   4. See `./common/money.graphql.go` for an example
3. Run `make generate`
4. Commit, merge, and pull
5. Tag the new version: `git tag v0.X.X`
6. Push the tag: `git push --tags`
7. Update the version in the deps docker image (`KITT_REPO/build/common/docker/deps/Dockerfile`):
   ```
   RUN git clone --depth 1 --branch v0.X.X https://github.com/kitt-technology/protos-common.git /protos/github.com/kitt-technology/protos-common
   ```
8. Rebuild that docker image:
   ```
   docker build --build-arg GITHUB_TOKEN -t gcr.io/kitt-220208/deps .
   ```
9. Import the common package in a proto:
   ```
   import "github.com/kitt-technology/protos-common/common.proto";
   ```
10. Use the common object:
   ```protobuf
   common.Money price = 1;
   ```
