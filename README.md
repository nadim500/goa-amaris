![Goa logo](https://goa.design/img/goa-banner.png "Goa")
## Overview

This project is a technical test for the company Amaris Consulting that was developed with [goa](https://goa.design)

## Run

To compile and run the project:

```bash
cd cmd/amaris
go build
./amaris
[amarisapi] 14:31:28 HTTP "openapi3.json" mounted on GET /openapi.json
[amarisapi] 14:31:28 HTTP "SortNames" mounted on POST /sortnames
[amarisapi] 14:31:28 HTTP "PokemonName" mounted on GET /pokemonname/{id}
[amarisapi] 14:31:28 HTTP "FriendlyChains" mounted on POST /friendlychains
[amarisapi] 14:31:28 HTTP server listening on "localhost:8080"
```

Open a new console and compile the generated CLI tool:

```bash
cd cmd/amaris-cli
go build
```

and run it:

```bash
./amaris-cli amaris sort-names --body '{"text": "Luis,Camilo,Andres,Laura"}'
{
    "Quantity": 4,
    "Data": [
        "Andres",
        "Camilo",
        "Laura",
        "Luis"
    ]
}
```

```bash
./amaris-cli amaris pokemon-name --id 1
{
    "Name": "bulbasaur"
}
```

```bash
./amaris-cli amaris friendly-chains --body '{"a": "tokio", "b": "kyoto"}'
{
    "Result": "Las cadenas de texto tokio y kyoto no son amigas"
}
```

The tool includes contextual help:

``` bash
./amaris-cli --help
```

Help is also available on each command:

``` bash
./amaris-cli amaris sort-names --help
```

## Document

We have a self-documenting service.

Run it:

``` bash
curl localhost:8080/openapi.json
{"openapi":"3.0.3","info":{"title":"Amaris Service","description":...
```

## Resources

### Docs

The [goa.design](https://goa.design) website provides a high level overview of
Goa and the DSL.
