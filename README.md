[![Build Status](https://github.com/ru84/triceped/actions/workflows/build.yml/badge.svg)](https://github.com/ru84/triceped/actions/workflows/build.yml)
[![Open Issues](https://img.shields.io/github/issues/ru84/triceped?label=issues)](https://github.com/ru84/triceped/issues)

# Triceped

**triceped** is a command-line tool written in Go that converts Terraform configurations into Azure Bicep templates. Designed to simplify the migration process, triceped helps organizations transition from Terraform to Bicep for Azure infrastructure deployments with minimal effort.

By parsing Terraform's HCL code, triceped generates equivalent Bicep files, enabling you to leverage Azure's native Infrastructure as Code (IaC) language while preserving your existing infrastructure definitions.

## Goals

1. **Simplify Migration**: Provide an easy and automated path to convert existing Terraform configurations to Bicep templates.
2. **Maintain Accuracy**: Ensure the generated Bicep code accurately reflects the original Terraform infrastructure.
3. **User-Friendly**: Offer a straightforward CLI that integrates seamlessly into your workflow.
4. **Extensible**: Allow for customization to support various Terraform and Bicep features.

### Non-Goals

1. **Replace Terraform**: triceped is not intended to replace Terraform but to assist in migration efforts.
2. **Full HCL Parsing**: Complete support for all HCL features is not guaranteed, especially for non-Azure providers.
3. **Provider-Agnostic Conversion**: Focus is on Azure resources; other cloud providers are out of scope.

## Get Started with triceped

To start using triceped:

1. **Install triceped**: Download the latest release from the [releases page](https://github.com/ru84/triceped/releases).
2. **Convert Your Configuration**: Run `triceped convert main.tf` to generate the equivalent `main.bicep` file.
3. **Review and Deploy**: Examine the generated Bicep code and deploy using Azure CLI or PowerShell.

## How Does triceped Work?

triceped reads your Terraform `.tf` files and translates the resources, variables, and modules into Bicep syntax. It maps Terraform resource blocks to Bicep resource declarations, handling properties and dependencies accordingly.

**Example Conversion Command:**

```bash
triceped -i main.tf -o main.bicep
```

This command converts `main.tf` into `main.bicep`, ready for deployment.

## Known Limitations

- **Incomplete HCL Support**: Some advanced HCL features may not be supported.
- **Custom Modules**: Manual intervention might be required for complex or custom modules.
- **State Management**: triceped does not handle Terraform state files; Bicep manages state differently.

## FAQ

**Why should I use triceped?**

If you're migrating to Azure Bicep for its native support and features, triceped automates the conversion of your existing Terraform configurations, saving time and reducing errors.

**Is triceped production-ready?**

Triceped is currently in **beta**. While we strive for reliability, we recommend thorough testing before using it in production environments.

**Does triceped support all Terraform features?**

Triceped focuses on core features relevant to Azure resources. Some Terraform-specific functionalities may not have direct equivalents in Bicep.

**How can I contribute?**

We welcome contributions! Please see our [Contributing Guidelines](https://github.com/ru84/triceped/blob/main/CONTRIBUTING.md) for details.

## Get Help or Report an Issue

- **Issues**: Report bugs or request features on our [GitHub Issues](https://github.com/ru84/triceped/issues) page.
- **Discussions**: Join the conversation on [GitHub Discussions](https://github.com/ru84/triceped/discussions) for questions and community support.

## Reference

- [Terraform Documentation](https://www.terraform.io/docs)
- [Azure Bicep Documentation](https://docs.microsoft.com/azure/azure-resource-manager/bicep/)

## License

This project is licensed under the [MIT License](https://github.com/ru84/triceped/blob/main/LICENSE).

## Contributing

We appreciate your interest in contributing! Please read our [Contributing Guidelines](https://github.com/ru84/triceped/blob/main/CONTRIBUTING.md) to get started.
