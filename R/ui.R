ui <- function(req) {
  fluidPage(
    title = "Testing styler",
    class = "m-3",
    div(
      class = "p-4 m-2 bg-danger w-100 bw-2 bs-solid br-2",
      uiOutput("card")
    )
  )
}
