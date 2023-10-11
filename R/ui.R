ui <- function(req) {
  fluidPage(
    title = "Testing styler",
    class = "m-3",
    div(
      class = "p-4 m-2 bk-red-100 w-100 b-w-2 b-s-solid b-r-2",
      uiOutput("card")
    )
  )
}
