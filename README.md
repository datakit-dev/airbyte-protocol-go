# Airbyte Protocol Go Types

![GitHub tag](https://img.shields.io/github/tag/datakit-dev/airbyte-protocol-go.svg)
![License](https://img.shields.io/github/license/datakit-dev/airbyte-protocol-go.svg)

This module, `github.com/datakit-dev/airbyte-protocol-go`, provides comprehensive Go types for the entire Airbyte Protocol, auto-generated from the official JSON Schema. With this, Go developers can seamlessly interact with the Airbyte Protocol without having to worry about type mismatches or manual conversions.

## Specification Details

This module is based on the **v1.0.0** specification of the Airbyte Protocol, which was obtained from [here](https://raw.githubusercontent.com/airbytehq/airbyte/v0.40.23/airbyte-protocol/protocol-models/src/main/resources/airbyte_protocol/airbyte_protocol.yaml). For ease of reference, the specification is saved locally in `assets/airbyte_protocol_v1.0.0.yaml`.

The YAML specification was converted to JSON using the following command:
```
yq -o=json > assets/airbyte_protocol_v1.0.0.json
```

The Go code for this module was generated using [Quicktype](https://app.quicktype.io/).

## Features

- **Complete Coverage**: Covers the entire Airbyte Protocol specification.
- **Auto-generated**: Code is generated from the official JSON Schema, ensuring accuracy and up-to-date representation.
- **Strongly Typed**: Take advantage of Go's strong typing to avoid common errors and improve code clarity.

## Installation

To install the module, use the go get command:
```
go get github.com/datakit-dev/airbyte-protocol-go
```

## Usage

After installation, import the module to your Go code:
```go
import airbyte "github.com/datakit-dev/airbyte-protocol-go"
```

## Contribute

Contributions are always welcome! If you have suggestions, improvements, or bug fixes, please [create an issue](https://github.com/datakit-dev/airbyte-protocol-go/issues) or submit a pull request.

For major changes, please open an issue first to discuss what you'd like to change. Ensure that tests and linting pass before submitting your pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
