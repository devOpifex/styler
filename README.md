# Styler

Tool to generate CSS classes, a bit like tailwind but less strict.

## Install

```bash
go install github.com/devOpifex/styler@latest
```

## Setup

Create a `.styler` file in your project root.

```bash
styler -create
```

## Usage

Add a `class` attribute to your HTML elements.

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

- Media queries are suffixed with `@` and prefixed with `md@`, `lg@` etc.
- States are suffixed with `:` and prefixed with `hover:`, `active:` etc.
- Numeric values are set as `unit` specified in the config. defaults to `rem`
and are divided by the `divider` specified in the config (`4` by default), 
e.g.: `padding-top-2` will result in `padding-top: 0.5rem`
