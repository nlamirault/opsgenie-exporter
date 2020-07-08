# Opsgenie Prometheus Exporter

A Prometheus exporter for the Bbox Miami, a Set-Top-Box (TV box) provided by French Internet Service Provider Bouygues Telecom.

[![License Apache 2][badge-license]](LICENSE)

Metrics are :

| Name                                               | Exposed informations                                  | Labels               |
| -------------------------------------------------- | ------------------------------------------------------| ---------------------|
| `opsgenie_up`                                      | Was the last query of Opsgenie successful.            |


## Installation

You can download the binaries :

* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/opsgenie-exporter-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/opsgenie-exporter-0.1.0_darwin_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/opsgenie-exporter-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/opsgenie-exporter-0.1.0_linux_arm) ]
* Architecture arm64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/opsgenie-exporter-0.1.0_linux_arm) ]


## Usage

Launch the Prometheus exporter :

        $ opsgenie-exporter

## Development

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).

## License

See [LICENSE](LICENSE) for the complete license.

## Changelog

A [changelog](ChangeLog.md) is available

## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat