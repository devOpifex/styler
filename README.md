# Styler

Tool to generate CSS classes, a bit like tailwind but less strict.

## Install

If you have Go installed, you can install styler with:

```bash
go install github.com/devOpifex/styler@latest
```

Otherwise, you can download the binary from the [releases](https://github.com/devOpifex/styler/releases) page.

## Setup

Create a `.styler` file in your project root.

```bash
styler -create
```

## Usage

First, edit `.styler` config file, then add a `class` attribute to your HTML elements.

```r
ui <- fluidPage(
  div(
    class = "display-flex padding-2 margin-bottom-2 width-40",
    div(
      class = "flex-grow-1",
      h1("Hello, world!", class = "color-red-400 hover:color-cyan-500")
    ),
    div(
      class = "flex-shrink-1 md@display-none",
      h2("This is hidden on medium and larger screens", class = "color-blue hover:color-green")
    )
  )
)
```

Call `styler` to generate the CSS.

```bash
styler
```

- Media queries are suffixed with `@` e.g.: `md@`, `lg@`
- States are suffixed with `:`, e.g.: `hover:`, `active:`
- Numeric values are set as `unit` specified in the config (defaults to `rem`)
and are divided by the `divider` specified in the config (defaults to `4`), 
e.g.: `padding-top-2` will result in `padding-top: 0.5rem`
- Numerics precded by `~` are treated as strict and are not divided or suffixed with `unit`
e.g.: `flex~1` results in `flex: 1`
- Media queries can be edited in `.styler` config file
- Colors can be edited in `.styler` config file
- Extracted CSS properties are checked against the [W3C CSS Properties](https://www.w3.org/Style/CSS/all-properties.en.html) table

## Advanced Property Values

Styler now supports complex CSS property values with multiple components:

- **Multiple Value Parts**: Create properties with multiple values like `box-shadow-1-2-4-red`
  which becomes `box-shadow: 0.25rem 0.5rem 1rem red`

- **Selective Strict Values**: Use `~` to mark specific parts as "strict" (no division or unit):
  - `border-1-solid-red` → `border: 0.25rem solid red` (number is divided by 4, unit added)
  - `border-~1-solid-red` → `border: 1 solid red` (strict number, no division or unit)
  - `border-1~solid-red` → `border: 0.25rem solid red` (only affects the part with ~)
  - `box-shadow~1-2-2-grey` → `box-shadow: 1 0.5rem 0.5rem grey` (strict property name)

This approach gives you fine-grained control over how numeric values are processed.
