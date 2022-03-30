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
7. Update `KITT_REPO/build/common/docker/deps/Dockerfile`:
   ```
   ENV PROTOS_COMMON_VERSION v0.X.X
   ```
8. Rebuild the docker image:
   ```
   docker build --build-arg GITHUB_TOKEN -t gcr.io/kitt-220208/deps .
   ```
9. Push that docker image (optional):
   ```
   docker push gcr.io/kitt-220208/deps
   ```
10. Import the common package in a proto:
   ```
   import "github.com/kitt-technology/protos-common/common.proto";
   ```
11. Use the common object:
   ```protobuf
   common.Money price = 1;
   ```


## Useful Resources
Google's proto library:
https://github.com/googleapis/googleapis/blob/master/google/type/datetime.proto
