> ⚠️ **This project is currently under active development and may contain bugs or incomplete features.**


# 🦋 Flumint

**Flumint** - a CLI tool for building multi-client Flutter projects. It allows you to build your Flutter app for multiple clients and platforms easily.

<div align="start">
  <img src="https://github.com/MohammadTaghipour/flumint/blob/master/flumint-logo.png" height="250" alt="Flumint Logo"/>
</div>

## ✨ Features

- Build and deploy Flutter projects for multiple clients with customized branding.
- Support for multiple platforms: `Android` (APK & AppBundle) and `Web`.
- Manage multiple environments seamlessly: `dev`, `staging`, and `prod`.
- `doctor` command to verify `Flutter`, `Dart`, `DevTools`, and system dependencies.
- `network` command to check connectivity and latency to all required repositories (`pub.dev`, `Flutter Storage`, `Maven`, `CocoaPods`).
- checkout command to quickly switch between client configurations, including app name, package name, and assets

## ⚙ Platform Support

| 🍏 MacOS | 🐧 Linux | 💻 Windows |
|:--------:|:--------:|:----------:|
|    ✅    |    ✅     |     ✅      |

## 🛠️ Installation

1. Download the executable from the [Releases](https://github.com/MohammadTaghipour/flumint/releases) page.
2. Rename the downloaded file to `flumint` and locate where you want.
3. Add the file to your **System Environment Path** so it can be run from anywhere.
4. Requirements:
   - Flutter SDK installed
   - For Android builds: Java and Android SDK installed


## ✈️ Getting Started with Flumint

To make a Flutter project compatible with Flumint:

1. Create a `clients` folder in the root of your project.
2. Inside `clients`, create a folder for each client (e.g., `client_a`, `client_b`).
3. Each client folder must have a `config.json` file with the following fields:

```json
{
  "app_name": "CLIENT_A_APP",
  "app_description": "Example app description",
  "package_name": "com.example.client_a"
}
```

4. Keep client-specific files inside each client folder (icons, assets, `google-services.json`, JKS, etc.) while preserving the folder structure.
5. When running `flumint build` or `flumint checkout`, the selected client’s folder contents are replaced to the project.
6. After building your Flutter project, you can get the current client using:

```dart
// This allows your app to detect which client it was built for.
const String client = String.fromEnvironment('client', defaultValue: 'client_a');

void main() {
  print('Current client: $client');
}
```


**Note:** The `clients` folder can be empty, but each client at least must have its `config.json` file.


**Example folder structure:**

```
my_flutter_project/
├─ lib/
├─ pubspec.yaml
├─ clients/
│  ├─ client_a/
│  │  ├─ config.json
│  │  ├─ assets/
│  │  │  ├─ icons/
│  │  │  ├─ ...
│  │  │  └─ images/
│  │  ├─ android/
│  │  │  ├─ app/
│  │  │  │  ├─ src/
│  │  │  │  │  ├─ ...
│  │  │  │  │  ├─ main/
│  │  │  │  │  │  ├─ ...
│  │  │  │  │  │  └─ res/
│  │  │  │  │  │      ├─ drawable/
│  │  │  │  │  │      ├─ ...
│  │  │  │  │  │      └─ values-night/
│  │  │  │  └─ google-services.json
│  │  └─ web/
│  │     └─ favicon.png
│  ├─ client_b/
│  │  ├─ config.json
│  │  ├─ assets/
│  │  └─ android/
│  └─ ...
```

## ⚡ Usage

### Build Project

Build a Flutter project for a specific client and platform.

```bash
flumint build --client <CLIENT_NAME> --target <apk|appbundle|web> [--path <PROJECT_PATH>] [--env dev|staging|prod]
```

**Flags:**

* `--client` **(required)**: Name of the client.
* `--target` **(required)**: Target platform: `apk`, `appbundle`, or `web`.
* `--path` (optional): Path to the Flutter project, default is current directory.
* `--env` (optional): Environment, default is `prod`.

**Example:**

```bash
flumint build --client client_a --target web --env prod
```

### Doctor

Check Flumint health, Flutter/Dart versions, and system dependencies:

```bash
flumint doctor
```

### Network

Check network connectivity to all required repositories:

```bash
flumint network
```

### Checkout

Switch the project configuration to another client:

```bash
flumint checkout --client <CLIENT_NAME> [--path <PROJECT_PATH>]
```

**Example:**

```bash
flumint checkout --client client_b
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
