July 27, 2024
# Log Something

Use `console.log`:

```js
console.log("Hello world");
```

# Log Errors:

Use `console.error`:

```js
console.error("Holy shittake!");
```

# Log Groups


```javascript
console.log('Begin:')

console.group()
console.log('Line 2 (console.log)')
console.log('Line 3 (console.log)')
console.log('Line 4 (console.log)')
console.group()
console.log('Line 5 (console.log)')
console.log('Line 6 (console.log)')
console.groupEnd()
console.log('Line 7 (console.log)')
console.groupEnd()

console.log('Done!')
```
```

Output:

```text
Begin:
  Line 2 (console.log)
  Line 3 (console.log)
  Line 4 (console.log)
    Line 5 (console.log)
    Line 6 (console.log)
  Line 7 (console.log)
Done!
```

# More Control

Use `process.stdout`

```javascript
var totalTime = 5000;
var waitInterval = totalTime / 10;
var currentInterval = 0;

function showPercentage(percentage){
    process.stdout.clearLine()
    process.stdout.cursorTo(0)
    console.log(`Processing ${percentage}%...` ) // Replace this line with process.stdout.write(`Processing ${percentage}%...`)
}

var interval = setInterval(function(){
    currentInterval += waitInterval
    showPercentage((currentInterval / totalTime) * 100)
}, waitInterval)

setTimeout(function(){
    clearInterval(interval)
}, totalTime + 100)
```