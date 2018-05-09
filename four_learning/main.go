package main

import "fmt"

// declare struct

type Information struct {
	Name    string
	Age     int
	School  string
	Married bool
	Habits  []string
	Infor   map[string]string
}

func exampleInformationDeclare() {
	// one
	var (
		exampleInformationOne Information
		Infor                 map[string]string
	)
	Infor = make(map[string]string)
	Infor["key"] = "hello-world"
	exampleInformationOne = Information{
		"xiewei", 18, "shanghaiUniversity", true,
		[]string{"1", "2", "3"}, Infor,
	}
	fmt.Println(exampleInformationOne, exampleInformationOne.Name)

	// two
	var exampleInformationTwo Information
	exampleInformationTwo.Name = "golang"
	exampleInformationTwo.School = "Google"
	exampleInformationTwo.Age = 10
	exampleInformationTwo.Infor = Infor

	fmt.Println(exampleInformationTwo, exampleInformationTwo.Married)

	// three
	var exampleInformationThree Information
	exampleInformationThree = Information{
		Name:    "Python",
		School:  "Gudio school",
		Age:     18,
		Habits:  []string{"1", "2"},
		Infor:   Infor,
		Married: true,
	}

	fmt.Println(exampleInformationThree, exampleInformationThree.School)

	// four

	var exampleInformationFour *Information
	exampleInformationFour = new(Information)
	exampleInformationFour.Name = "Java"

	fmt.Println(exampleInformationFour, exampleInformationFour.School)
}

// declare with label

type InformationWithLabel struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func exampleInformationWithLabelFiled() {
	var example InformationWithLabel
	example = InformationWithLabel{
		Name: "xiewei",
		Age:  18,
	}
	fmt.Println(example.Name, example.Age)
}

type Name struct {
}

func (n Name) ResponseNumber() int {
	return 1
}

type Response struct {
	Code    int
	Result  []byte
	Headers map[string]string
}

func (r Response) GetAttrCode() int {
	return r.Code
}
func (r Response) GetAttrResult() []byte {
	return r.Result
}

func (r Response) GetAttrHeader() map[string]string {
	return r.Headers
}

func (r *Response) SetCode(code int) {
	r.Code = code
}

func (r *Response) SetHeaders(key, value string) {
	r.Headers[key] = value
}

func exampleResponse() {
	var (
		response Response
		headers  map[string]string
	)
	headers = make(map[string]string)
	headers["Server"] = "GitHub.com"
	headers["Status"] = "Ok"
	response.Headers = headers
	response.Code = 200
	response.Result = []byte("hello world")

	fmt.Println(response.GetAttrCode())
	fmt.Println(response.GetAttrHeader())
	fmt.Println(response.GetAttrResult())

	response.SetCode(404)
	fmt.Println(response)

	response.SetHeaders("Status", "failed")
	fmt.Println(response)
}

func NormalFunc(arg int) int {
	return arg + 1
}

type RenameInt int

func (i RenameInt) Get() {
	fmt.Println(i)
}

type Requests struct {
	Url    string
	Params string
}

type CollectionRequests struct {
	CollectionNumber int
	Requests
	Response
}

func exampleCollectionRequests() {

	var collectionRequests CollectionRequests
	collectionRequests.CollectionNumber = 10
	collectionRequests.Url = "https://www.example.com"
	collectionRequests.Params = "name"
	collectionRequests.Code = 201
	collectionRequests.Result = []byte("hello Golang")

	var headers map[string]string
	headers = make(map[string]string)
	headers["status"] = "Good"
	collectionRequests.Headers = headers
	fmt.Println(collectionRequests)

	fmt.Println(collectionRequests.GetAttrCode())
}

type OtherRequests struct {
	Request Requests
	Resp    Response
	Code    int
}

func (o OtherRequests) String() string {
	return fmt.Sprintf("Request = %v , Response = %v , Code = %d", o.Request, o.Resp, o.Code)
}

func (o OtherRequests) GetAttrCode() {
	fmt.Println(fmt.Sprintf("Outer Code = %d", o.Code))
	fmt.Println(fmt.Sprintf("inner Code = %d", o.Resp.Code))
}

func exampleOtherRequests() {
	var other OtherRequests
	other.Code = 201
	other.Resp.Code = 202
	fmt.Println(other)
	other.GetAttrCode()
	fmt.Println(other.Resp.GetAttrCode())
}

func main() {
	exampleInformationDeclare()
	exampleInformationWithLabelFiled()
	var name Name
	name.ResponseNumber()

	exampleResponse()
	NormalFunc(12)

	exampleCollectionRequests()

	exampleOtherRequests()
}
