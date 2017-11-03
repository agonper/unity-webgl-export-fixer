# unity-webgl-export-fixer
A simple application that fixes *.gz files lookup for unity WebGL build exports

## Usage

Download latest version from releases

Just copy and execute `uwef.exe` (in Windows) or `uwef` (in macOS) to your build output folder and the tool will lookup for the build JSON config file and will fix its properties if needed.

The program behavior can be customized:
* `-folder`: You can pass by a different build output folder to fix it without the need of copying the program

## Usage Requisites

For Windows and macOS:
Nothing, since there's a compiled version of the application available in the repo

For Linux:
Golang 1.8+, since you'll need to compile the application (sorry, this will get fixed soon).

## Build

Run `build.bat` in Windows or `build.sh` in macOS / Linux

## Build requisites

Golang 1.8+