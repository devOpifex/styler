# Styler

Tool to generate CSS classes, a bit like tailwind but less strict.

For cases where we cannot use tailwind because of whatever
constraints, e.g.: using Shiny forcing Bootstrap.

Styler will scan for `class=""`.

```html
<!--src/file.html-->
<div class="margin-2 border-radius-2 border-red-100 border-width-1 position-relative">
    <h1 class="font-bold text-red-400">Hello</h1>
</div>
```

```bash
styler -dir=src -out=styles.css -warn=false
```

## Install

```bash
go install github.com/devOpifex/styler
```

## Shorthands

You don't have to use the shorthands, 
e.g.: `d-f` and `display-flex` are both correct.
They simply translate to something else in the generated
CSS.

- `b`: `bottom`
- `t`: `top`
- `c`: `color`
- `s`: `size`
- `r`: `radius`
- `m`: `margin`
- `p`: `padding`
- `w`: `width`
- `f`: `flex`
- `a`: `align`
- `j`: `justify`
- `i`: `items`
- `bk`: `background` (bg taken in many frameworks)
- `d`: `display`
- `pos`: `position`
- `rel`: `relative`
- `abs`: `absolute`
- `full`: `100%`
- `ov`: `overflow`
- `sh`: `shadow`
- `text`: `font`

## Details

- Special case for shadow takes `sm`, `md`, or `lg` suffixes.
- Support for `x`, `y`, `t`, and `b`, e.g.: `p-x-2`
- Tailwind's `hover:`, and `focus:` pseudo classes
- All tailwind's colors included, e.g.: `t-red-100`
- The last item in the string of attributes separated by `-` is the value, 
e.g.: `background-color-red` translates to `background-color:red;`
- Numeric values are divided by 4 and treated as `rem`, e.g.: `border-r-2` translates
to `border-radius:.25rem`, except for values `50` and `100` which are
treated as percentages.
You can prevent that by prefixing a number with `~` in which case it is treated as
"strict," e.g.: `border-w-~0.1` leads to `border-width:0.1rem`;
- Support for responsive `sm:`, `md:`, `lg:`, `xl:`, and `2xl:` prefixes.

