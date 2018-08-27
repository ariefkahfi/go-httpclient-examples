package main


import ( 
	"bytes"
	"io/ioutil"
	"fmt"
	"net/http"
	"strings"
	"net/url"
	"mime/multipart"
	"io-examples/utils"
	"encoding/json"
)


func httpMultipartFormData(){
	var bff = new(bytes.Buffer)
	mWriter := multipart.NewWriter(bff)

	err := mWriter.WriteField("name","Hello")
	myLogger.PrintError("err_writeField",err)

 	ioWriter , err := mWriter.CreateFormFile("file_upload","filetoupload")
	myLogger.PrintError("err_createFormFile",err)


	bts , err := ioutil.ReadFile("filetoupload")
	myLogger.PrintError("err_readFile",err)
	ioWriter.Write(bts)

	
	err = mWriter.Close()
	myLogger.PrintError("err_close",err)
	

	nReq , err := http.NewRequest("POST" , "http://localhost:9600/api/doUpload",bff)
	myLogger.PrintError("err_newRequest",err)
	nReq.Header.Set("content-type",mWriter.FormDataContentType())
	
	resp , err := http.DefaultClient.Do(nReq)
	myLogger.PrintError("err_response",err)
	defer resp.Body.Close()

	
	bts , err = ioutil.ReadAll(resp.Body)
	myLogger.PrintError("err_readAll",err)

	fmt.Println(string(bts))
}

func customHttpPostClient() {
	body := url.Values{
		"name":[]string{"Hello"},
		"msg":[]string{"Message from me !!!"},
	}
	formValues := strings.NewReader(body.Encode())
	nReq , err := http.NewRequest("POST","http://localhost:9600/api/doPost",formValues)
	myLogger.PrintError("err_newRequest",err)

	nReq.Header.Add("content-type","application/x-www-form-urlencoded")


	nClient  := http.DefaultClient
	resp , err := nClient.Do(nReq)
	myLogger.PrintError("err_doRequest",err)

	defer resp.Body.Close()


	bts , err := ioutil.ReadAll(resp.Body)
	myLogger.PrintError("err_readAll",err)

	fmt.Println(string(bts))
}

func apiJsonPlaceholderPost(){
	reqBody := map[string]interface{}{
		"title": "foo",
      	"body": "bar",
      	"userId": 1,
	}
	sReader := writeReqBody(reqBody)
	resp , err := http.Post("https://jsonplaceholder.typicode.com/posts","application/json",sReader)
	myLogger.PrintError("err_post",err)

	defer resp.Body.Close()


	bts , err := ioutil.ReadAll(resp.Body)
	myLogger.PrintError("err_readAll",err)

	fmt.Println("apiJsonPlaceholderPost",string(bts))
}

func apiJsonPlaceholderGet() {
	resp , err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	myLogger.PrintError("err_get",err)
	defer resp.Body.Close()

	bts , err := ioutil.ReadAll(resp.Body)
	myLogger.PrintError("readAll",err)

	fmt.Println("apiJsonPlaceholderGet",string(bts))
}


func writeReqBody(reqBody interface{}) *strings.Reader{
	bts , err := json.Marshal(reqBody)
	myLogger.PrintError("err_marshal",err)

	return strings.NewReader(string(bts))
}

func main() {
	// apiJsonPlaceholderGet()
	// apiJsonPlaceholderPost()
	// customHttpPostClient()
	httpMultipartFormData()
}