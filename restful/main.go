package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

func Say(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.URL.Host, "-", req.URL.Path, "-", req.Form)
	fmt.Fprintf(resp, "hello world")
}

type User struct {
	Name string
	Age  string
	ID   []int
}

type UserResource struct {
	// normally one would use DAO (data access object)
	users map[string]User
}

// WebService creates a new service that can handle REST requests for User resources.
func (u UserResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON) // you can specify this per route as well

	ws.Route(ws.GET("/").To(u.findAllUsers).
		// docs
		Doc("get all users").
		Writes([]User{}).
		Returns(200, "OK", []User{}))

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		// docs
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("integer").DefaultValue("1")).
		Writes(User{}). // on the response
		Returns(200, "OK", User{}).
		Returns(404, "Not Found", nil))

	return ws
}

// GET http://localhost:8080/users
//
func (u UserResource) findAllUsers(request *restful.Request, response *restful.Response) {
	list := []User{}
	for _, each := range u.users {
		list = append(list, each)
	}
	response.WriteEntity(list)
}

func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("user-id")
	usr := u.users[id]
	if len(usr.ID) == 0 {
		response.WriteErrorString(http.StatusNotFound, "User could not be found.")
	} else {
		response.WriteEntity(usr)
	}
}

func main() {
	//http.HandleFunc("/user/hello", Say)
	//http.ListenAndServe(":8080", nil)

	type APIServer struct {
		Container *restful.Container
	}
	u := UserResource{map[string]User{}}
	u.users["xiewei"] = User{
		Name: "xiewei",
		Age:  "20",
		ID:   []int{1, 2, 3, 4},
	}
	apiServer := &APIServer{
		Container: restful.DefaultContainer.Add(u.WebService()),
	}

	log.Printf("start listening on localhost:9990")
	log.Fatal(http.ListenAndServe(":9990", apiServer.Container))
}
