ui <- function(req) {
  fluidPage(
    title = "Testing styler",
    div(
      class = "p-4 m-2 bg-danger",
      uiOutput("card")
    )
  )
}
