server <- function(input, output, session) {
  output$card <- renderOutput({
    div(
      class = "hover:bk-red-500 hover:t-white m-3 b-r-2 shadow bk-blue-100 t-white p-x-4 p-y-2",
      h1("Content", class = "t-s-6", span("hello", class = "badge bk-sky p-x-4"))
      h2("hello", class = "bk-white d-none"),
      div(
        class = "d-f",
        div(
          class = "f-grow-1"
        ),
        div(
          class = "f-shrink-1"
        )
      )
    )
  })
}
