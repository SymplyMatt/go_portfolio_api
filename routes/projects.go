package routes

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/SymplyMatt/go_portfolio_api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()
var projectCollection *mongo.Collection = OpenCollection(Client, "projects")

func AddProject(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var project models.Project

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	validationErr := validate.Struct(project)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}
	project.ID = primitive.NewObjectID()
	result, insertErr := projectCollection.InsertOne(ctx, project)
	if insertErr != nil {
		msg := "order item was not created"
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		fmt.Println(insertErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func GetProjects(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var entries []bson.M
	cursor, err := projectCollection.Find(ctx, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if err = cursor.All(ctx, &entries); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println(entries)
	c.JSON(http.StatusOK, entries)

}

func UpdateProject(c *gin.Context) {
	projectID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(projectID)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var project models.Project

	if err := c.BindJSON(&project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	validationErr := validate.Struct(project)
	if validationErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": validationErr.Error()})
		fmt.Println(validationErr)
		return
	}

	result, err := projectCollection.ReplaceOne(
		ctx,
		bson.M{"_id": docID},
		bson.M{
			"name":        project.Name,
			"intro":       project.Intro,
			"image":       project.Image,
			"description": project.Description,
			"items":       project.Items,
			"images":      project.Images,
		},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, result.ModifiedCount)

}

func DeleteProject(c *gin.Context) {
	projectID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(projectID)

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	result, err := projectCollection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, result.DeletedCount)
}
