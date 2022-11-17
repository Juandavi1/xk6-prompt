# xk6-prompt

![k6 version](https://img.shields.io/badge/K6-v0.41.0-7d64ff)
![xk6 version](https://img.shields.io/badge/Xk6-v0.8.1-7d64ff)
![xk6 version](https://img.shields.io/badge/Go-v1.19-79d4fd)

![Alt text](prompt.svg)

k6 extension that adds support for input arguments via UI.


#### Install 

1. Install [xk6](https://github.com/grafana/xk6)
```shell
go install go.k6.io/xk6/cmd/xk6@latest
```
2. Build the extension using:

```shell
xk6 build --with github.com/Juandavi1/xk6-prompt
```

#### Import
```js
import prompt from 'k6/x/prompt';
```

#### Input select
```js
const options = ["smoke", "load"]
const selected = prompt.select("kind of test", ...options)
```

#### Read string
```js
const inputString = prompt.readString("type a string")
console.log(typeof inputString)
```

#### Read int
```js
const inputNumber = prompt.readInt("Type a number")
console.log(typeof inputNumber)
```
