server <- function(input, output, session) {
  output$card <- renderOutput({
    div(
      class = "b-w-~0.4 b-c-blue-200 t-s-4 shadow-lg hover:bk-red-500 hover:t-white m-3 b-r-2 bk-blue-100 t-white p-x-4 p-y-2",
      h1("Content", class = "t-s-6", span("hello", class = "badge bk-sky p-x-4"))
      h2("hello", class = "bk-white d-none"),
      h3("Hover test", class="md:hover:bk-red-400 text-white"),
      div(
        class = "md:d-f",
        div(
          class = "md:f-grow-1"
        ),
        div(
          class = "me:f-shrink-1"
        )
      ),
      div(
        class = "xl:d-f",
        div(
          class = "xl:f-grow-1"
        ),
        div(
          class = "xl:f-shrink-1"
        )
      )
    )
  })
}
