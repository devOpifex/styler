# Styler

Tool to generate CSS classes, a bit like tailwind but less strict.

It's for the cases where we cannot use tailwind because of whatever
constraints we have to use another framework, e.g.: Bootstrap.

```html
<!--src/file.html-->
<div class="m-2 b-r-2 b-c-red b-w-1 pos-rel">
    <h1 class="t-bold t-red-400">Hello</h1>
</div>
```

```go
./styler -dir=src -out=styles.css -warn=false
```

## Shorthands

You don't have to use the shorthands, 
e.g.: `d-f` and `display-flex` are both correct.
They simply translate to something else in the generated
CSS.

- `b`: `border`
- `c`: `color`
- `s`: `size`
- `r`: `radius`
- `m`: `margin`
- `p`: `padding`
- `w`: `width`
- `f`: `flex`
- `t`: `font`
- `bk`: `background`
- `d`: `display`
- `pos`: `position`
- `rel`: `relative`
- `abs`: `absolute`
- `full`: `100%`
- `ov`: `overflow`
- `sh`: `shadow`

Special case for shadow takes `sm`, `md`, or `lg` suffixes.

Support for `x`, `y`, `t`, and `b`, e.g.: `p-x-2`

Tailwind's `hover:` prefix supported

All tailwind's colors included, e.g.: `t-red-100`
