# Styler

Tool to generate CSS classes, a bit like tailwind but less strict.

```html
<!--src/file.html-->
<div class="m2 br-2 bc-red bw-1">
    <h1 class="text-bold">Hello</h1>
</div>
```

```go
./styler -dir=src -out=styles.css -warn=false
```
