library(shiny)

ui <- fluidPage(
  includeCSS("style.min.css"),
  div(
    class = "display-flex padding-2 margin-bottom-2 width-40",
    div(
      class = "flex-grow-1",
      h1("Hello, world!", class = "color-red hover:color-green")
    ),
    div(
      class = "flex-shrink-1 md@display-none",
      h2("Hello, world!", class = "color-blue hover:color-green")
    )
  )
)

server <- \(input, output, session){}

shinyApp(ui, server)
