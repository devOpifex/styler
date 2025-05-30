library(shiny)

ui <- fluidPage(
  includeCSS("style.min.css"),
  div(
    class = "display-flex padding-2 margin-bottom-2 width-40",
    div(
      class = "flex-grow~1 box-shadow-1-1-1~2-grey",
      h1("Hello, world!", class = "color-red-400 hover:color-cyan-500")
    ),
    div(
      class = "flex-shrink~1 md@display-none",
      h2(
        "This is hidden on medium screens and larger hello-world",
        class = "color-blue hover:color-green"
      )
    )
  )
)

server <- \(input, output, session) {
}

shinyApp(ui, server)
