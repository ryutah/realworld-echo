# realworld-echo

## Quickstart

```console
make init
go run .
```

## Libraries

| lib            | use                            |
| -------------- | ------------------------------ |
| zap            | logging                        |
| echo           | web framework                  |
| oapi-codegen   | schema base OAS code generator |
| eri            | error utility                  |
| go-cmp         | test compare                   |
| wire           | dependency injection           |
| opentelemetory | tracing                        |
| gomock         | mocking                        |

## About Component

| component      | usage                                 |
| -------------- | ------------------------------------- |
| api            | API handler                           |
| usecase        | Implements application business logic |
| domain         | Implements domain logic               |
| infrastructure | Implements technical logic            |

### Diagram

```mermaid
classDiagram
  api --> usecase : call
  usecase --> domain : call
  domain <|-- infrastructure : implements
```

## Testing

1. Should be cover 100% covarage
1. Should be separate test package as `xxx_test`
   - because it should be test how to use other package

## Tips

### Only first time

```console
go get -u github.com/golang/mock/mockgen/model github.com/google/wire/cmd/wire@v0.5.0
```
