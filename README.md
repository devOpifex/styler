# Styler

Tool to generate CSS class but less strict.

```html
<!--src/file.html-->
<div class="m-2 b-r-2 b-c-red b-w-1 pos-rel">
    <h1 class="t-bold t-red-400">Hello</h1>
</div>
```

```go
./styler -dir=src -out=styles.css -warn=false
```
