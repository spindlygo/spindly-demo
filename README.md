# Spindly demo app

This is still a prototype

#### To run,
1. Install Go : https://go.dev/doc/install
2. Install Node : https://nodejs.org/en/download
3. Clone this repo : `git clone https://github.com/spindlygo/spindly-demo`
4. Open repo folder in Visual Studio Code and install `Svelte`, `Go` extentions.
5. Open vscode terminal and run : 
```
npm install
npm run dev
```
6. Open a browser for the displayed URL. (`http://localhost:42635`)
7. To open it as a normal app window, change property `devdriver` to `chromeapp` in file `SpindlyConfigs.json`

#### To build packages for Linux, Windows and Mac : 
```
npm run build
```

#### To use [webview](https://github.com/webview/webview) as a driver,
1. Install libwebkit2gtk : `sudo apt-get install -y libwebkit2gtk-4.0-dev`
2. Add `adaptive` into property `driver` in file `SpindlyConfigs.json` to use webview in production.
3. (Optional) Change property `devdriver` to `adaptive` in file `SpindlyConfigs.json` to use webview while in development.
4. Cross compiling isn't supported for webview. Use [this github actions workflow](https://github.com/spindlygo/spindly-demo/blob/main/.github/workflows/build-all.yml) to compile your desktop webview app for other operating systems.

Mobile app support with Flutter coming soon.
