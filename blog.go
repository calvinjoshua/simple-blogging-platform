package main

import (
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

//for simplicity i have managed the Id in memory, for fault tolerance Id will stored in persistent state
var Id int

type response struct {
	Status  int
	Message string
	Data    interface{}
	Error   string
}

// type CreateBlogRequestPayload struct {

// }

type ManageBlogRequestPayload struct {
	BlogId int    `json:"blogId"`
	Blog   string `json:"blog"`
	Author string `json:"author"`
}

func creatBlogPost(c *fiber.Ctx) error {
	var mu sync.Mutex // Create a mutex

	var err error

	payloadHolder := ManageBlogRequestPayload{}

	err = c.BodyParser(&payloadHolder)
	if err != nil {
		return c.JSON(response{
			Status:  400,
			Message: "Invalid request payload",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	if payloadHolder.Blog == "" || payloadHolder.Author == "" {
		return c.JSON(response{
			Status:  400,
			Message: "Invalid request payload",
			Data:    nil,
			Error:   "One of required request payload is empty string",
		})
	}

	mu.Lock()
	defer mu.Unlock()
	Id++

	shortid := Id

	if err != nil {
		return c.JSON(response{
			Status:  500,
			Message: "Something went wrong, please try after sometime!",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	// _, err = insertBlogData(shortid, payloadHolder.Blog, payloadHolder.Author)

	// if err != nil {
	// 	return c.JSON(response{
	// 		Status:  500,
	// 		Message: "Something went wrong, please try after sometime!",
	// 		Data:    nil,
	// 		Error:   err.Error(),
	// 	})
	// }

	return c.JSON(response{
		Status:  200,
		Message: "Blog created successfully with ID in Data field",
		Data:    shortid,
		Error:   "nil",
	})
}

func retriveBlogPost(c *fiber.Ctx) error {

	var payload ManageBlogRequestPayload

	var err error

	var blog interface{}

	err = c.BodyParser(&payload)
	if err != nil {
		return c.JSON(response{
			Status:  400,
			Message: "Invalid request payload",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	// blog, err = retriveBlog(payload.BlogId)

	// if err != nil {
	// 	return c.JSON(response{
	// 		Status:  500,
	// 		Message: "Something went wrong, please try after sometime!",
	// 		Data:    nil,
	// 		Error:   err.Error(),
	// 	})
	// }

	return c.JSON(response{
		Status:  200,
		Message: "Blog retrived",
		Data:    blog,
		Error:   "",
	})

}

func retriveAllBlogs(c *fiber.Ctx) error {

	var blogs interface{}

	var err error

	blogs, err = getAllblogs()

	if err != nil {
		return c.JSON(response{
			Status:  500,
			Message: "Something went wrong, please try after sometime!",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(response{
		Status:  200,
		Message: "All blogs retrived",
		Data:    blogs,
		Error:   "",
	})

}

func deleteBlog(c *fiber.Ctx) error {
	var payload ManageBlogRequestPayload
	var err error

	err = c.BodyParser(&payload)
	if err != nil {
		return c.JSON(response{
			Status:  400,
			Message: "Invalid request payload",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	err = _deleteBlog(payload.BlogId)

	if err != nil {
		return c.JSON(response{
			Status:  500,
			Message: "Something went wrong, please try after sometime!",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(response{
		Status:  200,
		Message: "Blog deleted with Id provided in Data field",
		Data:    payload.BlogId,
		Error:   "",
	})

}

func updateBlog(c *fiber.Ctx) error {

	var payload ManageBlogRequestPayload

	var err error

	err = c.BodyParser(&payload)
	if err != nil {
		return c.JSON(response{
			Status:  400,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	_, err = _updateBlog(payload.BlogId, payload.Blog)

	if err != nil {
		return c.JSON(response{
			Status:  500,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.JSON(response{
		Status:  200,
		Message: "Blog updated with Id " + strconv.Itoa(payload.BlogId) + "provided",
		Data:    payload.Blog,
		Error:   "",
	})

}
