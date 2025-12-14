/// AUTO-GENERATED FILE
/// DO NOT MODIFY

class AppConfig {
  static const config = {
    'app': {
      'name': 'Application Name',
      'version': '1.0.0+1',
      'description': 'A short description about app',
    },
    'platforms': {
      'android': {'packageName': 'com.example.customerapp'},
      'ios': {'bundleIdentifier': 'com.example.customerapp'},
      'web': {
        'name': 'App name for web',
        'shortName': 'short name',
        'description': 'App description for web',
      },
    },
    'assets': {
      'iconPath': 'path/to/customer/icon.png',
      'sourceDir': 'path/to/customer/assets/',
    },
    'build': {
      'outputs': [
        {
          'platform': 'android',
          'type': 'apk',
          'mode': 'release',
          'flutterArgs': ['--split-per-abi'],
        },
        {'platform': 'android', 'type': 'appbundle', 'mode': 'release'},
      ],
      'outputDir': 'build/outputs/',
    },
  };
}
