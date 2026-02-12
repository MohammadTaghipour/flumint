# 0.1.0

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
