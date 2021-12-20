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
    func <Message_name>FromArgs(field map[string]interface{}) *<Message_name>
    var <Message_name>GraphqlInputType gql.InputObject{}
    var <Message_name>GraphqlType gql.Object{}
    ```
   4. See `./common/money.graphql.go` for an example
3. Run `make generate`
4. Commit and merge
