package main

import (
	"fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"time"
)

type ExampleDocument struct {
	models.Document
	Name        string `json:"name"`
	ReleaseYear int    `json:"release_year"`
}
type ExampleDocuments struct {
	models.DocumentList
	Documents []ExampleDocument `json:"documents"`
}

var client = appwrite.NewClient(
	appwrite.WithEndpoint("https://v16.appwrite.org/v1"), // Replace with your endpoint
	appwrite.WithProject("6697f76e00251b11c823"),         // Replace with your project ID
	appwrite.WithKey("standard_c321c05842acd04ec9b826eaa6fa5a58adb1c116ceff3915210a4d4dbfc16f0a346da8b5160a364a20f0770d725e70df45782693e01ff0fae2d7078405d0e4c755b4353175d515d6668bf1899df771a88ee80e43805e3a9916f6160534b9c8c21ed5cbe806ee947eb4ef2345bb700779f8b7daa440b68659b255e1e59c2bcaaf"), // Replace with your API Key
)

var databasesSdk = appwrite.NewDatabases(client)
var storageSdk = appwrite.NewStorage(client)
var usersSdk = appwrite.NewUsers(client)
var functionsSdk = appwrite.NewFunctions(client)

var dbId string
var collectionId string
var documentId string
var bucketId string
var fileId string
var userId string
var functionId string

func main() {

	createDatabase()
	listDatabases()
	getDatabase()
	updateDatabase()

	createCollection()
	listCollections()
	getCollection()
	updateCollection()
	listAttributes()

	createDocument()
	getDocument()
	listDocuments()
	updateDocument()

	deleteDocument()
	deleteCollection()
	deleteDatabase()

	createBucket()
	listBuckets()
	getBucket()
	updateBucket()

	uploadFile()
	listFiles()
	getFile()
	updateFile()
	deleteFile()
	deleteBucket()

	createUser()
	listUsers()
	getUser()
	updateUserName()
	deleteUser()

	createFunction()
	listFunctions()
	getFunction()
	uploadDeployment()
	executeSync()
	executeAsync()
	deleteFunction()
}

func createFunction() {
	Info("\nRunning Create Function API")

	response, err := functionsSdk.Create(
		id.Unique(),
		"Go Hello World",
		"go-1.22",
		functionsSdk.WithCreateExecute([]string{role.Any()}),
	)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("Function (%v)\"Go Hello World\" created", response.Id))

	functionId = response.Id
}

func getFunction() {
	Info("\nRunning Get Function API")

	response, err := functionsSdk.Get(functionId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Runtime))
}

func listFunctions() {
	Info("\nRunning List Functions API")

	response, err := functionsSdk.List()

	if err != nil {
		panic(err)
	}

	Success("Functions list:")
	for _, function := range response.Functions {
		Success(fmt.Sprintf("(%v) %v - %v", function.Id, function.Name, function.Runtime))
	}
}

func uploadDeployment() {
	Info("\nRunning Upload Deployment  API")
	deployment, err := functionsSdk.CreateDeployment(
		functionId,
		file.NewInputFile("./files/gocode.tar.gz", "code.tar.gz"),
		true,
		functionsSdk.WithCreateDeploymentEntrypoint("main.go"),
	)

	if err != nil {
		panic(err)
	}
	Info("Waiting for deployment to be ready...")
	for {
		response, err := functionsSdk.GetDeployment(functionId, deployment.Id)
		if err != nil {
			panic(err)
		}
		if response.Status == "ready" {
			break
		}

		if response.Status == "failed" {
			panic("Function build failed")
		}

		time.Sleep(5 * time.Second)
	}

	Success("Deployment uploaded")

}

func executeSync() {
	Info("\nRunning Execute Function API (sync)")

	execution, err := functionsSdk.CreateExecution(
		functionId,
	)
	if err != nil {
		panic(err)
	}

	Success(execution.ResponseBody)
}

func executeAsync() {
	Info("\nRunning Execute Function API (async)")

	execution, err := functionsSdk.CreateExecution(
		functionId,
		functionsSdk.WithCreateExecutionAsync(true),
	)
	if err != nil {
		panic(err)
	}
	Info("Waiting a little to ensure execution is finished ...")
	time.Sleep(2 * time.Second)

	executionResults, err := functionsSdk.GetExecution(functionId, execution.Id)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("%d", executionResults.ResponseStatusCode))
}

func deleteFunction() {
	Info("\nRunning Delete Function API")

	_, err := functionsSdk.Delete(functionId)
	if err != nil {
		panic(err)
	}

	Success("Function deleted")
}

func createUser() {
	Info("\nRunning Create User API")

	response, err := usersSdk.Create(
		id.Unique(),
		usersSdk.WithCreateEmail("test@example.com"),
		usersSdk.WithCreatePassword("password"),
		usersSdk.WithCreateName("Some User"),
	)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("User (%v)\"Some User\" created", response.Id))

	userId = response.Id
}

func listUsers() {
	Info("\nRunning List Users API")

	response, err := usersSdk.List()

	if err != nil {
		panic(err)
	}
	Success("Users list:")
	for _, user := range response.Users {
		Success(fmt.Sprintf("(%v) %v - %v", user.Id, user.Name, user.Email))
	}
}

func getUser() {
	Info("\nRunning Get User API")

	response, err := usersSdk.Get(userId)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Email))
}

func updateUserName() {
	Info("\nRunning Update User Name API")

	response, err := usersSdk.UpdateName(userId, "Updated Name")

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Email))
}

func deleteUser() {
	Info("\nRunning Delete User API")

	_, err := usersSdk.Delete(userId)

	if err != nil {
		panic(err)
	}
	Success("User deleted")

}

func createBucket() {
	Info("\nRunning Create Bucket API")

	response, err := storageSdk.CreateBucket(
		id.Unique(),
		"All Files",
		storageSdk.WithCreateBucketPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("")),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}))

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("Bucket (%v) \"All Files\" created", response.Id))

	bucketId = response.Id
}

func listBuckets() {
	Info("\nRunning List Bucket API")

	response, err := storageSdk.ListBuckets()

	if err != nil {
		panic(err)
	}

	Success("Buckets list:")
	for _, bucket := range response.Buckets {
		Success(fmt.Sprintf("(%v) %v", bucket.Id, bucket.Name))
	}
}

func getBucket() {
	Info("\nRunning Get Bucket API")

	response, err := storageSdk.GetBucket(bucketId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func updateBucket() {
	Info("\nRunning Update Bucket API")

	response, err := storageSdk.UpdateBucket(bucketId, "Updated Bucket")

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func uploadFile() {
	Info("\nRunning Upload File API")

	response, err := storageSdk.CreateFile(
		bucketId,
		id.Unique(),
		file.NewInputFile("./files/nature.jpg", "nature.jpg"),
		storageSdk.WithCreateFilePermissions([]string{
			permission.Read(role.Any()),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}),
	)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))

	fileId = response.Id
}

func listFiles() {
	Info("\nRunning List Files API")

	response, err := storageSdk.ListFiles(bucketId)

	if err != nil {
		panic(err)
	}

	Success("Files list:")
	for _, localFile := range response.Files {
		Success(fmt.Sprintf("(%v) %v", localFile.Id, localFile.Name))
	}
}

func getFile() {
	Info("\nRunning Get File API")

	response, err := storageSdk.GetFile(bucketId, fileId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func updateFile() {
	Info("\nRunning Update File API")

	response, err := storageSdk.UpdateFile(
		bucketId,
		fileId,
		storageSdk.WithUpdateFilePermissions([]string{
			permission.Read(role.Any()),
			permission.Update(role.Any()),
			permission.Delete(role.Any()),
		}),
	)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func deleteFile() {
	Info("\nRunning Delete File API")

	_, err := storageSdk.DeleteFile(bucketId, fileId)
	if err != nil {
		panic(err)
	}

	Success("File deleted")
}

func deleteBucket() {
	Info("\nRunning Delete Bucket API")

	_, err := storageSdk.DeleteBucket(bucketId)
	if err != nil {
		panic(err)
	}

	Success("Bucket deleted")
}

func createDatabase() {
	Info("Running Create Database API")
	response, err := databasesSdk.Create(id.Unique(), "Default")
	if err != nil {
		panic(err)
	}

	Success("Created database")
	dbId = response.Id
}

func listDatabases() {
	Info("\nRunning List Databases API")
	response, err := databasesSdk.List()
	if err != nil {
		panic(err)
	}

	Success("Databases list:")
	for _, database := range response.Databases {
		Success(fmt.Sprintf("(%v) %v", database.Id, database.Name))
	}
}

func getDatabase() {
	Info("\nRunning Get Database API")

	response, err := databasesSdk.Get(dbId)
	if err != nil {
		panic(err)
	}

	Success("Database details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))
}

func updateDatabase() {
	Info("\nRunning Update Database API")

	response, err := databasesSdk.Update(dbId, "Updated Database")
	if err != nil {
		panic(err)
	}

	Success("Database details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))

}

func deleteDatabase() {
	Info("\nRunning Delete Database API")

	_, err := databasesSdk.Delete(dbId)
	if err != nil {
		panic(err)
	}

	Success("Database deleted")
}

func createCollection() {
	Info("\nRunning Create Collection API")

	response, err := databasesSdk.CreateCollection(
		dbId, id.Unique(),
		"Collection",
		databasesSdk.WithCreateCollectionPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("")),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}))

	if err != nil {
		panic(err)
	}
	Success("Collection created")

	collectionId = response.Id

	_, err = databasesSdk.CreateStringAttribute(
		dbId,
		collectionId,
		"name",
		255,
		false,
		databasesSdk.WithCreateStringAttributeDefault("Empty Name"),
		databasesSdk.WithCreateStringAttributeArray(false),
	)

	if err != nil {
		panic(err)
	}
	Success("Attribute `name` created")

	_, err = databasesSdk.CreateIntegerAttribute(
		dbId,
		collectionId,
		"release_year",
		false,
		databasesSdk.WithCreateIntegerAttributeMin(0),
		databasesSdk.WithCreateIntegerAttributeMax(5000),
		databasesSdk.WithCreateIntegerAttributeDefault(1970),
		databasesSdk.WithCreateIntegerAttributeArray(false),
	)

	if err != nil {
		panic(err)
	}

	Success("Attribute `release_year` created")

	Info("Waiting a little to ensure attributes are created ...")
	time.Sleep(time.Second * 2)

	_, err = databasesSdk.CreateIndex(
		dbId,
		collectionId,
		"key_release_year_asc",
		"key",
		[]string{"release_year"},
		databasesSdk.WithCreateIndexOrders([]string{"ASC"}),
	)

	if err != nil {
		panic(err)
	}

	Info("Waiting a little to ensure index is created  ...")
	time.Sleep(time.Second * 2)
	Success("Index `key_release_year_asc` created")

}

func deleteCollection() {
	Info("\nRunning Delete Collection API")

	_, err := databasesSdk.DeleteCollection(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Collection deleted")
}

func listCollections() {
	Info("\nRunning List Collections API")
	response, err := databasesSdk.ListCollections(dbId)
	if err != nil {
		panic(err)
	}

	Success("Collections list:")
	for _, collection := range response.Collections {
		Success(fmt.Sprintf("(%v) %v", collection.Id, collection.Name))
	}
}
func getCollection() {
	Info("\nRunning Get Collection API")

	response, err := databasesSdk.GetCollection(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Collection details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))
}

func updateCollection() {
	Info("\nRunning Update Collection API")

	response, err := databasesSdk.UpdateCollection(dbId, collectionId, "Updated Collection")
	if err != nil {
		panic(err)
	}

	Success("Collection details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))

}
func listAttributes() {
	Info("\nRunning List Attributes API")
	response, err := databasesSdk.ListAttributes(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Attributes list:")
	for _, attribute := range response.Attributes {
		fmt.Printf("Key: %v, Type: %v ", green.Render(attribute["key"].(string)), green.Render(attribute["type"].(string)))
	}
}

func createDocument() {
	Info("\nRunning Add Document API")
	response, err := databasesSdk.CreateDocument(
		dbId,
		collectionId,
		id.Unique(),
		map[string]any{
			"name":         "Spider Man",
			"release_year": 1920,
		},
		databasesSdk.WithCreateDocumentPermissions([]string{
			permission.Read(role.Any()),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}),
	)
	if err != nil {
		panic(err)
	}

	Success("Document created")
	documentId = response.Id
}
func listDocuments() {
	Info("\nRunning List Documents API")
	response, err := databasesSdk.ListDocuments(dbId, collectionId)

	if err != nil {
		panic(err)
	}

	Success("Document list:")

	var documents ExampleDocuments
	err = response.Decode(&documents)

	if err != nil {
		panic(err)
	}

	for _, document := range documents.Documents {
		Success(fmt.Sprintf("(%v) %v - %v", document.Id, document.Name, document.ReleaseYear))
	}
}

func getDocument() {
	Info("\nRunning Get Document API")
	response, err := databasesSdk.GetDocument(dbId, collectionId, documentId)
	if err != nil {
		panic(err)
	}
	var document ExampleDocument
	err = response.Decode(&document)

	if err != nil {
		panic(err)
	}

	Success("Document Details:")
	Success(fmt.Sprintf("(%v) %v - %v", document.Id, document.Name, document.ReleaseYear))

}
func updateDocument() {
	Info("\nRunning Update Document API")

	response, err := databasesSdk.UpdateDocument(
		dbId,
		collectionId,
		documentId,
		databasesSdk.WithUpdateDocumentData(map[string]any{"release_year": 2005}),
	)
	if err != nil {
		panic(err)
	}
	var document ExampleDocument
	err = response.Decode(&document)

	Success("Document Details:")
	Success(fmt.Sprintf("(%v) release_year=%v", document.Id, document.ReleaseYear))
}

func deleteDocument() {
	Info("\nRunning Delete Document API")

	_, err := databasesSdk.DeleteDocument(dbId, collectionId, documentId)
	if err != nil {
		panic(err)
	}

	Success("Document deleted")
}
