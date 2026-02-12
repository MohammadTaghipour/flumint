> ⚠️ **This project is currently under active development and may contain bugs or incomplete features.**


# 🦋 Flumint

**Flumint** is a CLI tool for building multi-client Flutter projects. It allows you to build your Flutter app for multiple clients and platforms easily.


## ✨ Features

- Build Flutter projects for different clients or brandings
- Support for `Android` and `Web` platforms
- Manage multiple environments: `dev`, `staging`, `prod`
- `doctor` command to check system and dependencies health


## 🛠️ Installation

1. Download the executable from the [Releases](https://github.com/MohammadTaghipour/flumint/releases) page.
2. Add the file to your **System Environment Path** so it can be run from anywhere.
3. Requirements:
   - Flutter SDK installed
   - For Android builds: Java and Android SDK installed


## ⚡ Usage

### Build Project

```bash
flumint build --client <CLIENT_NAME> [--platform android|web] [--env dev|staging|prod]
```

Flags:

* `--client` (required): Name of the client
* `--platform` (optional): Target platform, default is `android`
* `--env` (optional): Environment, default is `dev`

Example:

```bash
flumint build --client client_a --platform web --env prod
```

### Doctor

Check the system and dependencies:

```bash
flumint doctor
```



## 🤝 Contributing

Contributions are welcome! If you'd like to contribute, feel free to open a pull request or submit
an issue.


## 🛡️ License

This project is licensed under the [MIT License](https://mit-license.org/).

## 📧 Contact

For questions, feedback, or support, please reach out:

- **Developer**: Mohammad Taghipour
- **Email**: taghipourmohammad7@gmail.com
