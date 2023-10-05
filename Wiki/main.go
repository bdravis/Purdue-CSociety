package main

import (
  "fmt" //for formatting messages to the console
  "net/http" //for web service
  //"log" //logging errors
  "errors" //creating new errors
  "os" //reading and writing files
  "html/template" //for generating the page HTML
  //"strconv" //converting status codes to string
)

//Structure to represent each wiki page
type WikiPage struct {
  DocName string //public
  path string
  statusCode int
  Content []byte //io libraries use []byte
  HTMLContent template.HTML //public
}

var FileNotFound *WikiPage // 404
var ErrorOccur *WikiPage // 500 status

//returns the WikiPage associated with the provided document name
//If the DocName does not coorspond to a valid file, return the 404 WikiPage
func getWikiPage( DocName string ) (*WikiPage, error) {
  return nil, nil
}

//returns the WikiPage for error files
//these exist in a different directory to not conflict with regular file names
func getErrorPage( statusCode int ) (*WikiPage, error) {
  return nil, nil
}

//returns a WikiPage* associated with the provided file path
// if the provided path does not exist, FileNotFound is returned instead
func getFile( filePath string ) (*WikiPage, error) {
  body, err := os.ReadFile(filePath)
  if ( err != nil ) {
    return FileNotFound, err
  }
  return &WikiPage{ DocName: "temp", path: filePath, statusCode: 200, Content: body}, err
}

//writes the Content of the page to the file structure
// errors.New("Page is empty")
func writeWikiPage( page *WikiPage ) error {
  if ( page == nil ) {
    return errors.New("Page is nil")
  }

  return os.WriteFile(page.path, page.Content, 0666)
}

//you can leave this alone
//simply grabs the file specified by the request
//needed for just general files needed by the web browser
func getFileContents(writer http.ResponseWriter, request *http.Request) {
  var path string = request.URL.Path[1:]
  http.ServeFile( writer, request, path)
}

//viewing a wiki page
func viewHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
}


//editing a wiki page
func editHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
}

//creating a new wiki page
func newHandler(writer http.ResponseWriter, request *http.Request) {
  getFileContents( writer, request )
}

//saving a wiki page
func saveHandler( writer http.ResponseWriter, request *http.Request ) {
  //redirect if trying to use anything other than a POST
  return
}


//home screen
func homeHandler( writer http.ResponseWriter, request *http.Request ) {
  return
}

func main() {
  //Initialize error files
  FileNotFound, _ = getErrorPage( 404 )
  ErrorOccur, _ = getErrorPage( 500 )


  fmt.Printf("Running...\n")
  //TODO:
  // Create handlers
  // Create listen and serve to start server
}
