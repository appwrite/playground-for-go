package main

import (
	"fmt"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/functions"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/storage"
	"github.com/appwrite/sdk-for-go/users"
)

var appwrite_project = "<PROJECT_ID>"
var appwrite_api_key = "<API_KEY>"

var (
	appwriteClient   client.Client
	appwriteDatabase *databases.Databases
	appwriteFunction *functions.Functions
	appwriteUsers    *users.Users
	appwriteStorage  *storage.Storage
	database_id      string
	collection_id    string
	document_id      string
	user_id          string
	bucket_id        string
	file_id          string
	function_id      string
)

func main() {
	appwriteClient = appwrite.NewClient(
		appwrite.WithProject(appwrite_project),
		appwrite.WithKey(appwrite_api_key),
	)

	appwriteDatabase = appwrite.NewDatabases(appwriteClient)
	appwriteFunction = appwrite.NewFunctions(appwriteClient)
	appwriteUsers = appwrite.NewUsers(appwriteClient)
	appwriteStorage = appwrite.NewStorage(appwriteClient)

	CreateUser()
	ListUsers()
	DeleteUsers()

	CreateDatabase()
	DeleteDatabase()

	CreateCollection()
	ListCollection()
	DeleteCollection()

	CreateDocument()
	ListDocuments()
	DeleteDocument()

	CreateBucket()
	ListBuckets()
	DeleteBucket()

	// UploadFile() TODO: Fix how we send content range etc in SDK Go
	ListFiles()

	CreateFunction()
	ListFunctions()
	DeleteFunction()

	fmt.Println("Successfully ran playground!")
}

func CreateUser() {
	fmt.Println("Running Create user API")

	user, _ := appwriteUsers.Create(
		id.Unique(),
		appwriteUsers.WithCreateEmail("test@example.com"),
		appwriteUsers.WithCreatePhone("+12065550100"),
		appwriteUsers.WithCreatePassword("test1234"),
		appwriteUsers.WithCreateName("test"),
	)

	user_id = user.Id

	fmt.Println(user)
}

func ListUsers() {
	fmt.Println("Running List User API")

	users, _ := appwriteUsers.List()

	fmt.Println(users)
}

func DeleteUser() {
	fmt.Println("Running Delete User API")

	response, _ := appwriteUsers.Delete(user_id)

	fmt.Println(response)
}

func CreateDatabase() {
	fmt.Println("Running Create Database API")

	database, _ := appwriteDatabase.Create(id.Unique(), "Movies")

	database_id = database.Id

	fmt.Println(database)
}

func DeleteDatabase() {
	fmt.Println("Running Delete Database API")

	response, _ := appwriteDatabase.Delete(database_id)

	fmt.Println(response)
}

func CreateCollection() {
	fmt.Println("Running Create Collection API")

	appwriteCollection, _ := appwriteDatabase.CreateCollection(
		database_id,
		id.Unique(),
		"Movies",
		appwriteDatabase.WithCreateCollectionPermissions([]string{"any"}),
		appwriteDatabase.WithCreateCollectionDocumentSecurity(true),
	)

	collection_id = appwriteCollection.Id

	appwriteDatabase.CreateStringAttribute(
		database_id,
		collection_id,
		"name",
		255,
		true,
	)

	appwriteDatabase.CreateIntegerAttribute(
		database_id,
		collection_id,
		"release_year",
		true,
		appwriteDatabase.WithCreateIntegerAttributeMin(0),
		appwriteDatabase.WithCreateIntegerAttributeMax(9999),
	)

	appwriteDatabase.CreateFloatAttribute(
		database_id,
		collection_id,
		"rating",
		true,
		appwriteDatabase.WithCreateFloatAttributeMin(0),
		appwriteDatabase.WithCreateFloatAttributeMax(99.99),
	)

	appwriteDatabase.CreateBooleanAttribute(
		database_id,
		collection_id,
		"kids",
		true,
	)

	appwriteDatabase.CreateEmailAttribute(
		database_id,
		collection_id,
		"email",
		false,
		appwriteDatabase.WithCreateEmailAttributeDefault("example@email.com"),
	)

	appwriteDatabase.CreateIndex(
		database_id,
		collection_id,
		"name_email_idx",
		"fulltext",
		[]string{"email"},
	)
}

func ListCollection() {
	fmt.Println("Running List Collection API")

	var collections, _ = appwriteDatabase.List()

	fmt.Println(collections)
}

func DeleteCollection() {
	fmt.Println("Running Delete Collection API")

	response, _ := appwriteDatabase.DeleteCollection(
		database_id,
		collection_id,
	)

	fmt.Println(response)
}

func CreateDocument() {
	fmt.Println("Running Create Document API")

	document, _ := appwriteDatabase.CreateDocument(
		database_id,
		collection_id,
		id.Unique(),
		map[string]interface{}{
			"name":         "Spider Man",
			"release_year": 1920,
			"rating":       99,
			"kids":         false,
			"email":        "example@email.com",
		})

	document_id = document.Id

	fmt.Println(document)
}

func ListDocuments() {
	fmt.Println("Running List Document API")

	documents, _ := appwriteDatabase.ListDocuments(database_id, collection_id)

	fmt.Println(documents)
}

func DeleteDocument() {
	fmt.Println("Running Delete Document API")

	response, _ := appwriteDatabase.DeleteDocument(
		database_id,
		collection_id,
		document_id,
	)

	fmt.Println(response)
}

func CreateBucket() {
	fmt.Println("Running Create Bucket API")

	bucket, _ := appwriteStorage.CreateBucket(
		id.Unique(),
		"awesome-bucket",
		appwriteStorage.WithCreateBucketFileSecurity(false),
		appwriteStorage.WithCreateBucketPermissions(
			[]string{
				permission.Read("any"),
				permission.Create("any"),
			}))

	bucket_id = bucket.Id

	fmt.Println(bucket)

}

func ListBuckets() {
	fmt.Println("Running List Buckets API")

	buckets, _ := appwriteStorage.ListBuckets()

	fmt.Println(buckets)
}

// func UploadFile() { // TODO: Fix how we send content range etc in SDK Go
// 	fmt.Println("Running Upload File API")

// 	var file, _ = appwriteStorage.CreateFile(
// 		bucket_id,
// 		id.Unique(),
// 		file.NewInputFile("./resources/nature.jpg", "file.jpg"),
// 		appwriteStorage.WithCreateFilePermissions(
// 			[]string{
// 				permission.Read("any"),
// 			}))

// 	file_id = file.Id

// 	fmt.Println(file)
// }

func ListFiles() {
	fmt.Println("Running List Files API")

	var files, _ = appwriteStorage.ListFiles(bucket_id)

	fmt.Println(files)
}

func DeleteFile() {
	fmt.Println("Running Delete File API")

	var response, _ = appwriteStorage.DeleteFile(bucket_id, file_id)

	fmt.Println(response)
}

func DeleteBucket() {
	fmt.Println("Running Delete Bucket API")

	var response, _ = appwriteStorage.DeleteBucket(bucket_id)

	fmt.Println(response)
}

func CreateFunction() {
	fmt.Println("Running Create Function API")

	var function, _ = appwriteFunction.Create(
		id.Unique(),
		"Test Function",
		"python-3.9",
		appwriteFunction.WithCreateExecute(
			[]string{
				permission.Read("any"),
			}))

	function_id = function.Id

	fmt.Println(function)
}

func ListFunctions() {
	fmt.Println("Runnnig List Functions API")

	var functions, _ = appwriteFunction.List()

	fmt.Println(functions)
}

func DeleteFunction() {
	fmt.Println("Running Delete Function API")

	var response, _ = appwriteFunction.Delete(function_id)

	fmt.Println(response)
}
