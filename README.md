

<p align="center">

  <h2 align="center">ron</h2>
  <p align="center">
</p>

![GitHub](https://img.shields.io/badge/golang%20->=1.15.x-blue.svg)
![test cases](https://github.com/9bany/ron/actions/workflows/ci.yml/badge.svg)

Used to build and restart a multiple programing when it crashes or some watched file changes.
<b>Aimed to be used in development only.</b>

## Summary

- [Install](#install)
- [Ron File](#ron-file)
- [Example](#example)
- [Implemented Languages](#implemented-languages)

## Install
- Linux
```
$ curl -k -L -s https://github.com/9bany/ron/releases/download/0.0.1/ron > ron;chmod +x ron;sudo mv ron /usr/local/bin
```

## Ron file
At the root of the project you want to observe, create new file with name `ron.yaml`.


**Properties**

|key |type| description|
|-|-|-|
|`root_path`| string | Path to the directory to be observed
|`exec_path`| string | Path to application entry point
|`language`| [string] | Execution language. Check [here](#implemented-languages-and-commands) if your favorite language has already been implemented
|`watch_extension`| Watch | List of extensions to be observed.
|`ignore_path`| [string] | List of paths not to be ignore.

## Example

Below is a simple example of the `ron.yaml` file

```yaml
root_path: "./"
exec_path: "index.js"
language: "node"
watch_extension:
  - js
  - ts
ignore_path:
  - "./tests/"
```

Once you have the Ron cli set up, just run the command in the folder containing the `ron.yaml` file.

```sh
ron
```
## Implemented languages 

| Languages | Name | Ex |
| - | - | - |
| GO | go | [example](https://github.com/9bany/ron/tree/master/examples/go)
| NodeJs | node |[example](https://github.com/9bany/ron/tree/master/examples/nodejs)
| Typescript | ts-node |[example](https://github.com/9bany/ron/tree/master/examples/ts-node)

