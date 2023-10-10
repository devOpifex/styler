server <- function(input, output, session) {
  output$card <- renderOutput({
    div(
      class = "rounded shadow bg-info text-white px-4 py-2",
      h1("Content", class = "text-xl")
    )
  })
}
