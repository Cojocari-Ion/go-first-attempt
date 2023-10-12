package routes

import {
    "github.com/gorilla/mux"
    "github.com/mrDublionka/go-first-attempt/pkg/controllers"
}

var RegisterBlogRoutes = func(router *mux.Router) {
    router.HandlerFunc("/post", controllers.CreatePost).Methods("POST")

}

