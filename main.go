package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Post("/create-blog-post", creatBlogPost)

	app.Get("/retrive-blog-post", retriveBlogPost)

	app.Get("/retrive-all-blog-post", retriveAllBlogs)

	app.Delete("/remove-blog-post", deleteBlog)

	app.Put("/update-blog-post", updateBlog)

	log.Fatal(app.Listen(":3000"))
}
