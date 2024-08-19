package main

import (
	"fmt"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/file"
	"github.com/appwrite/sdk-for-go/functions"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/role"
	"github.com/appwrite/sdk-for-go/storage"
	"github.com/appwrite/sdk-for-go/users"
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

func main() {
	client := appwrite.NewClient(
		appwrite.WithEndpoint("https://cloud.appwrite.io/v1"), // Replace with your endpoint
		appwrite.WithProject("YOUR_PROJECT_ID"),               // Replace with your project ID
		appwrite.WithKey("YOUR_API_KEY"),                      // Replace with your API Key
	)

	databasesSdk := appwrite.NewDatabases(client)
	storageSdk := appwrite.NewStorage(client)
	usersSdk := appwrite.NewUsers(client)
	functionsSdk := appwrite.NewFunctions(client)

	dbId := createDatabase(databasesSdk)
	listDatabases(databasesSdk)
	getDatabase(databasesSdk, dbId)
	updateDatabase(databasesSdk, dbId)

	collectionId := createCollection(databasesSdk, dbId)
	listCollections(databasesSdk, dbId)
	getCollection(databasesSdk, dbId, collectionId)
	updateCollection(databasesSdk, dbId, collectionId)
	listAttributes(databasesSdk, dbId, collectionId)

	documentId := createDocument(databasesSdk, dbId, collectionId)
	getDocument(databasesSdk, dbId, collectionId, documentId)
	listDocuments(databasesSdk, dbId, collectionId)
	updateDocument(databasesSdk, dbId, collectionId, documentId)

	deleteDocument(databasesSdk, dbId, collectionId, documentId)
	deleteCollection(databasesSdk, dbId, collectionId)
	deleteDatabase(databasesSdk, dbId)

	bucketId := createBucket(storageSdk)
	listBuckets(storageSdk)
	getBucket(storageSdk, bucketId)
	updateBucket(storageSdk, bucketId)

	fileId := uploadFile(storageSdk, bucketId)
	listFiles(storageSdk, bucketId)
	getFile(storageSdk, bucketId, fileId)
	updateFile(storageSdk, bucketId, fileId)
	deleteFile(storageSdk, bucketId, fileId)
	deleteBucket(storageSdk, bucketId)

	userId := createUser(usersSdk)
	listUsers(usersSdk)
	getUser(usersSdk, userId)
	updateUserName(usersSdk, userId)
	deleteUser(usersSdk, userId)

	functionId := createFunction(functionsSdk)
	listFunctions(functionsSdk)
	getFunction(functionsSdk, functionId)
	uploadDeployment(functionsSdk, functionId)
	executeSync(functionsSdk, functionId)
	executeAsync(functionsSdk, functionId)
	deleteFunction(functionsSdk, functionId)
}

func createFunction(sdk *functions.Functions) string {
	Info("\nRunning Create Function API")

	response, err := sdk.Create(
		id.Unique(),
		"Go Hello World",
		"go-1.22",
		sdk.WithCreateExecute([]string{role.Any()}),
	)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("Function (%v)\"Go Hello World\" created", response.Id))

	return response.Id
}

func getFunction(sdk *functions.Functions, functionId string) {
	Info("\nRunning Get Function API")

	response, err := sdk.Get(functionId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Runtime))
}

func listFunctions(sdk *functions.Functions) {
	Info("\nRunning List Functions API")

	response, err := sdk.List()

	if err != nil {
		panic(err)
	}

	Success("Functions list:")
	for _, function := range response.Functions {
		Success(fmt.Sprintf("(%v) %v - %v", function.Id, function.Name, function.Runtime))
	}
}

func uploadDeployment(sdk *functions.Functions, functionId string) {
	Info("\nRunning Upload Deployment  API")
	deployment, err := sdk.CreateDeployment(
		functionId,
		file.NewInputFile("./files/gocode.tar.gz", "code.tar.gz"),
		true,
		sdk.WithCreateDeploymentEntrypoint("main.go"),
	)

	if err != nil {
		panic(err)
	}
	Info("Waiting for deployment to be ready...")
	for {
		response, err := sdk.GetDeployment(functionId, deployment.Id)
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

func executeSync(sdk *functions.Functions, functionId string) {
	Info("\nRunning Execute Function API (sync)")

	execution, err := sdk.CreateExecution(
		functionId,
	)
	if err != nil {
		panic(err)
	}

	Success(execution.ResponseBody)
}

func executeAsync(sdk *functions.Functions, functionId string) {
	Info("\nRunning Execute Function API (async)")

	execution, err := sdk.CreateExecution(
		functionId,
		sdk.WithCreateExecutionAsync(true),
	)
	if err != nil {
		panic(err)
	}
	Info("Waiting a little to ensure execution is finished ...")
	time.Sleep(2 * time.Second)

	executionResults, err := sdk.GetExecution(functionId, execution.Id)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("%d", executionResults.ResponseStatusCode))
}

func deleteFunction(sdk *functions.Functions, functionId string) {
	Info("\nRunning Delete Function API")

	_, err := sdk.Delete(functionId)
	if err != nil {
		panic(err)
	}

	Success("Function deleted")
}

func createUser(users *users.Users) string {
	Info("\nRunning Create User API")

	response, err := users.Create(
		id.Unique(),
		users.WithCreateEmail("test@example.com"),
		users.WithCreatePassword("password"),
		users.WithCreateName("Some User"),
	)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("User (%v)\"Some User\" created", response.Id))

	return response.Id
}

func listUsers(users *users.Users) {
	Info("\nRunning List Users API")

	response, err := users.List()

	if err != nil {
		panic(err)
	}
	Success("Users list:")
	for _, user := range response.Users {
		Success(fmt.Sprintf("(%v) %v - %v", user.Id, user.Name, user.Email))
	}
}

func getUser(users *users.Users, userId string) {
	Info("\nRunning Get User API")

	response, err := users.Get(userId)

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Email))
}

func updateUserName(users *users.Users, userId string) {
	Info("\nRunning Update User Name API")

	response, err := users.UpdateName(userId, "Updated Name")

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("(%v) %v - %v", response.Id, response.Name, response.Email))
}

func deleteUser(users *users.Users, userId string) {
	Info("\nRunning Delete User API")

	_, err := users.Delete(userId)

	if err != nil {
		panic(err)
	}
	Success("User deleted")

}

func createBucket(storage *storage.Storage) string {
	Info("\nRunning Create Bucket API")

	response, err := storage.CreateBucket(
		id.Unique(),
		"All Files",
		storage.WithCreateBucketPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("")),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}))

	if err != nil {
		panic(err)
	}
	Success(fmt.Sprintf("Bucket (%v) \"All Files\" created", response.Id))

	return response.Id
}

func listBuckets(storage *storage.Storage) {
	Info("\nRunning List Bucket API")

	response, err := storage.ListBuckets()

	if err != nil {
		panic(err)
	}

	Success("Buckets list:")
	for _, bucket := range response.Buckets {
		Success(fmt.Sprintf("(%v) %v", bucket.Id, bucket.Name))
	}
}

func getBucket(storage *storage.Storage, bucketId string) {
	Info("\nRunning Get Bucket API")

	response, err := storage.GetBucket(bucketId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func updateBucket(storage *storage.Storage, bucketId string) {
	Info("\nRunning Update Bucket API")

	response, err := storage.UpdateBucket(bucketId, "Updated Bucket")

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func uploadFile(storage *storage.Storage, bucketId string) string {
	Info("\nRunning Upload File API")

	response, err := storage.CreateFile(
		bucketId,
		id.Unique(),
		file.NewInputFile("./files/nature.jpg", "nature.jpg"),
		storage.WithCreateFilePermissions([]string{
			permission.Read(role.Any()),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}),
	)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))

	return response.Id
}

func listFiles(storage *storage.Storage, bucketId string) {
	Info("\nRunning List Files API")

	response, err := storage.ListFiles(bucketId)

	if err != nil {
		panic(err)
	}

	Success("Files list:")
	for _, file := range response.Files {
		Success(fmt.Sprintf("(%v) %v", file.Id, file.Name))
	}
}

func getFile(storage *storage.Storage, bucketId, fileId string) {
	Info("\nRunning Get File API")

	response, err := storage.GetFile(bucketId, fileId)

	if err != nil {
		panic(err)
	}

	Success(fmt.Sprintf("(%v) %v", response.Id, response.Name))
}

func updateFile(storage *storage.Storage, bucketId, fileId string) {
	Info("\nRunning Update File API")

	response, err := storage.UpdateFile(
		bucketId,
		fileId,
		storage.WithUpdateFilePermissions([]string{
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

func deleteFile(storage *storage.Storage, bucketId, fileId string) {
	Info("\nRunning Delete File API")

	_, err := storage.DeleteFile(bucketId, fileId)
	if err != nil {
		panic(err)
	}

	Success("File deleted")
}

func deleteBucket(storage *storage.Storage, bucketId string) {
	Info("\nRunning Delete Bucket API")

	_, err := storage.DeleteBucket(bucketId)
	if err != nil {
		panic(err)
	}

	Success("Bucket deleted")
}

func createDatabase(dbs *databases.Databases) string {
	Info("Running Create Database API")
	response, err := dbs.Create(id.Unique(), "Default")
	if err != nil {
		panic(err)
	}

	Success("Created database")
	return response.Id
}

func listDatabases(dbs *databases.Databases) {
	Info("\nRunning List Databases API")
	response, err := dbs.List()
	if err != nil {
		panic(err)
	}

	Success("Databases list:")
	for _, database := range response.Databases {
		Success(fmt.Sprintf("(%v) %v", database.Id, database.Name))
	}
}

func getDatabase(dbs *databases.Databases, id string) {
	Info("\nRunning Get Database API")

	response, err := dbs.Get(id)
	if err != nil {
		panic(err)
	}

	Success("Database details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))
}

func updateDatabase(dbs *databases.Databases, id string) {
	Info("\nRunning Update Database API")

	response, err := dbs.Update(id, "Updated Database")
	if err != nil {
		panic(err)
	}

	Success("Database details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))

}

func deleteDatabase(dbs *databases.Databases, id string) {
	Info("\nRunning Delete Database API")

	_, err := dbs.Delete(id)
	if err != nil {
		panic(err)
	}

	Success("Database deleted")
}

func createCollection(dbs *databases.Databases, dbId string) string {
	Info("\nRunning Create Collection API")

	response, err := dbs.CreateCollection(
		dbId, id.Unique(),
		"Collection",
		dbs.WithCreateCollectionPermissions([]string{
			permission.Read(role.Any()),
			permission.Create(role.Users("")),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}))

	if err != nil {
		panic(err)
	}
	Success("Collection created")

	collectionId := response.Id

	_, err = dbs.CreateStringAttribute(
		dbId,
		collectionId,
		"name",
		255,
		false,
		dbs.WithCreateStringAttributeDefault("Empty Name"),
		dbs.WithCreateStringAttributeArray(false),
	)

	if err != nil {
		panic(err)
	}
	Success("Attribute `name` created")

	_, err = dbs.CreateIntegerAttribute(
		dbId,
		collectionId,
		"release_year",
		false,
		dbs.WithCreateIntegerAttributeMin(0),
		dbs.WithCreateIntegerAttributeMax(5000),
		dbs.WithCreateIntegerAttributeDefault(1970),
		dbs.WithCreateIntegerAttributeArray(false),
	)

	if err != nil {
		panic(err)
	}

	Success("Attribute `release_year` created")

	Info("Waiting a little to ensure attributes are created ...")
	time.Sleep(time.Second * 2)

	_, err = dbs.CreateIndex(
		dbId,
		collectionId,
		"key_release_year_asc",
		"key",
		[]string{"release_year"},
		dbs.WithCreateIndexOrders([]string{"ASC"}),
	)

	if err != nil {
		panic(err)
	}

	Info("Waiting a little to ensure index is created  ...")
	time.Sleep(time.Second * 2)
	Success("Index `key_release_year_asc` created")

	return collectionId

}

func deleteCollection(dbs *databases.Databases, dbId string, collectionId string) {
	Info("\nRunning Delete Collection API")

	_, err := dbs.DeleteCollection(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Collection deleted")
}

func listCollections(dbs *databases.Databases, dbId string) {
	Info("\nRunning List Collections API")
	response, err := dbs.ListCollections(dbId)
	if err != nil {
		panic(err)
	}

	Success("Collections list:")
	for _, collection := range response.Collections {
		Success(fmt.Sprintf("(%v) %v", collection.Id, collection.Name))
	}
}
func getCollection(dbs *databases.Databases, dbId string, collectionId string) {
	Info("\nRunning Get Collection API")

	response, err := dbs.GetCollection(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Collection details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))
}

func updateCollection(dbs *databases.Databases, dbId string, collectionId string) {
	Info("\nRunning Update Collection API")

	response, err := dbs.UpdateCollection(dbId, collectionId, "Updated Collection")
	if err != nil {
		panic(err)
	}

	Success("Collection details")
	fmt.Printf("Id: %v\n", green.Render(response.Id))
	fmt.Printf("Name: %v\n", green.Render(response.Name))
	fmt.Printf("CreatedAt: %v\n", green.Render(response.CreatedAt))
	fmt.Printf("UpdatedAt: %v\n", green.Render(response.UpdatedAt))

}
func listAttributes(dbs *databases.Databases, dbId string, collectionId string) {
	Info("\nRunning List Attributes API")
	response, err := dbs.ListAttributes(dbId, collectionId)
	if err != nil {
		panic(err)
	}

	Success("Attributes list:")
	for _, attribute := range response.Attributes {
		fmt.Printf("Key: %v, Type: %v ", green.Render(attribute["key"].(string)), green.Render(attribute["type"].(string)))
	}
}

func createDocument(dbs *databases.Databases, dbId string, collectionId string) string {
	Info("\nRunning Add Document API")
	response, err := dbs.CreateDocument(
		dbId,
		collectionId,
		id.Unique(),
		map[string]any{
			"name":         "Spider Man",
			"release_year": 1920,
		},
		dbs.WithCreateDocumentPermissions([]string{
			permission.Read(role.Any()),
			permission.Update(role.Users("")),
			permission.Delete(role.Users("")),
		}),
	)
	if err != nil {
		panic(err)
	}

	Success("Document created")
	return response.Id
}
func listDocuments(dbs *databases.Databases, dbId string, collectionId string) {
	Info("\nRunning List Documents API")
	response, err := dbs.ListDocuments(dbId, collectionId)

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

func getDocument(dbs *databases.Databases, dbId string, collectionId string, documentId string) {
	Info("\nRunning Get Document API")
	response, err := dbs.GetDocument(dbId, collectionId, documentId)
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
func updateDocument(dbs *databases.Databases, dbId string, collectionId string, documentId string) {
	Info("\nRunning Update Document API")

	response, err := dbs.UpdateDocument(
		dbId,
		collectionId,
		documentId,
		dbs.WithUpdateDocumentData(map[string]any{"release_year": 2005}),
	)
	if err != nil {
		panic(err)
	}
	var document ExampleDocument
	err = response.Decode(&document)

	Success("Document Details:")
	Success(fmt.Sprintf("(%v) release_year=%v", document.Id, document.ReleaseYear))
}

func deleteDocument(dbs *databases.Databases, dbId string, collectionId string, documentId string) {
	Info("\nRunning Delete Document API")

	_, err := dbs.DeleteDocument(dbId, collectionId, documentId)
	if err != nil {
		panic(err)
	}

	Success("Document deleted")
}
