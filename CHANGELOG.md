# 2.0.0

- Added `--path` argument to specify the path to the Flutter project.
- Added `checkout` command to switch between clients.

  - Example: `flumint checkout --client <CLIENT_NAME> [--path <PROJECT_PATH>]`
- Added `network` command to check connectivity to required repositories.

  - Example: `flumint network`
- Added support for Android `AppBundle`.

  - Example: `flumint build --client <CLIENT_NAME> --target <apk|appbundle|web>`
- Support for client-specific configurations via `clients/<CLIENT_NAME>/config.json`.
- Assets, configs, and resources are now automatically updated during build/checkout.
- Improved Android configuration changes during build/checkout.
- Added feature to update Android and Web app names.
- Improved command-line interface.



# 1.0.0

Initial release of **Flumint** – a multi-client Flutter build tool.

- CLI tool to build Flutter projects for multiple clients and brandings.
- Supported platforms: `Android` and `Web`.
- Manage multiple environments: `dev`, `staging`, `prod`.
- Commands:
    - `flumint build --client <CLIENT_NAME> [--platform android|web] [--env dev|staging|prod]` → Build project for a specific client and environment.
    - `flumint doctor` → Check system health and dependencies.
- Requirements:
    - Flutter SDK installed.
    - For Android builds: Java and Android SDK installed.
- Designed to be used as a standalone executable added to system PATH.
- Open for contributions with MIT license.
