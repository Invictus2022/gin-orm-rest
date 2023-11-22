package controller

import (
	"fmt"
	"net/http"
	dbclient "prac/db_client"
	prac "prac/source"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create a Post.
func CreatePost(c *gin.Context) {
	var reqBody prac.Post
	if err := c.ShouldBindJSON(&reqBody); err != nil { // Call BindJSON to bind the received JSON to reqBody.
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}
	_, err := dbclient.DBClient.Exec("INSERT INTO  students(id,title,content) VALUES($1,$2,$3)",
		reqBody.ID, reqBody.Title, reqBody.Content)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	} else {
		fmt.Println("value inserted")
	}
	c.Status(http.StatusOK)
}

// Get all the posts
func GetPosts(c *gin.Context) {
	var posts []prac.Post // Variable to hold our post that we fetch from our database.

	rows, err := dbclient.DBClient.Query("SELECT id,title,content,created_at FROM students;")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying the database"})
		return
	}
	defer rows.Close()

	//Iterate through the rows returned from the database. rows.next return booleans and loops until false.
	for rows.Next() {
		// Populate it with one of the row.
		var singlepost prac.Post

		//rows.Scan scan the row into the the above struct.
		err := rows.Scan(&singlepost.ID, &singlepost.Title, &singlepost.Content, &singlepost.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
			return
		}
		//  Now after scan is done append the singlepost struct to the Post slice.
		posts = append(posts, singlepost)
	}

	c.JSON(http.StatusOK, posts)
}

// Get Individual Posts
func GetPost(c *gin.Context) {

	// Get the ID parameter from the request URL.
	idStr := c.Param("id")

	// Convert the ID string to an integer.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If there is an error converting the ID, return a bad request response.
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	// Query the database to get the post with the specified ID.
	row := dbclient.DBClient.QueryRow("SELECT id,title,content,created_at FROM students  WHERE id=$1;", id)

	var post prac.Post

	// Populate post with data from the retrieved row.
	if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   "Error scanning row",
			"message": err.Error(),
		})
		return
	}
	// Return the post as a JSON response.
	c.JSON(http.StatusOK, post)
}

//Delete Post

func DelPost(c *gin.Context) {

	// Get the ID parameter from the request URL.
	idstr := c.Param("id")

	// Convert the ID string to an integer.
	id, err := strconv.Atoi(idstr)

	// If there is an error converting the ID, return a bad request response.
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	// Execute the DELETE query to delete the post with the specified ID.
	_, err = dbclient.DBClient.Exec("DELETE FROM students WHERE id=$1;", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error deleting post",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

// Update Post

func UpdatePost(c *gin.Context) {

	// Get the ID parameter from the request URL.
	idstr := c.Param("id")

	// Convert the ID string to an integer.
	id, err := strconv.Atoi(idstr)

	if err != nil {
		// If there is an error converting the ID, return a bad request response.
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid ID",
			"message": err.Error(),
		})
		return
	}

	// Check if the post with the specified ID exists.
	row := dbclient.DBClient.QueryRow("SELECT id, title, content, created_at FROM students WHERE id=$1;", id)

	// Make a var existingPost with type prac.Post struct. It will temporarly store the retrieved row from db.
	var existingPost prac.Post

	// Attempt to populate the existing post with data from the retrieved row.
	if err := row.Scan(&existingPost.ID, &existingPost.Title, &existingPost.Content, &existingPost.CreatedAt); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Post not found",
			"message": err.Error(),
		})
		return
	}
	// Bind the request JSON to a new post object.
	var updatedPost prac.Post
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	// Execute the UPDATE query to update the post with the specified ID.
	_, err = dbclient.DBClient.Exec("UPDATE students SET title=$1, content=$2 WHERE id=$3;",
		updatedPost.Title, updatedPost.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error updating post",
			"message": err.Error(),
		})
		return
	}

	// Return the updated post as a JSON response.
	c.JSON(http.StatusOK, updatedPost)
}
