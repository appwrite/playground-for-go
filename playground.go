package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/appwrite/sdk-for-go/account"
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/client"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/functions"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/permission"
	"github.com/appwrite/sdk-for-go/storage"
	"github.com/appwrite/sdk-for-go/users"
)

var appwriteProject = "<PROJECT_ID>"
var appwriteApiKey = "<API_KEY>"

var (
	appwriteClient   client.Client
	appwriteDatabase *databases.Databases
	appwriteFunction *functions.Functions
	appwriteStorage  *storage.Storage
	appwriteAccount  *account.Account
	appwriteUsers    *users.Users
	databaseId       string
	collectionId     string
	documentId       string
	userId           string
	bucketId         string
	fileId           string
	functionId       string
)

func main() {
	appwriteClient = appwrite.NewClient(
		appwrite.WithProject(appwriteProject),
		appwrite.WithKey(appwriteApiKey),
		// appwrite.WithJWT("JWT") // Use this to authenticate with JWT instead of API_KEY
	)

	appwriteDatabase = appwrite.NewDatabases(appwriteClient)
	appwriteFunction = appwrite.NewFunctions(appwriteClient)
	appwriteUsers = appwrite.NewUsers(appwriteClient)
	appwriteStorage = appwrite.NewStorage(appwriteClient)
	appwriteAccount = appwrite.NewAccount(appwriteClient)

	// GetAccount() // Use this only with JWT
	CreateUser()
	ListUsers()
	DeleteUser()

	CreateDatabase()

	CreateCollection()
	ListCollection()

	CreateDocument()
	ListDocuments()

	DeleteDocument()
	DeleteCollection()
	DeleteDatabase()
	CreateBucket()
	ListBuckets()

	// UploadFile() TODO: Fix how we send content range etc in SDK Go
	ListFiles()
	DeleteFile()

	DeleteBucket()

	CreateFunction()
	ListFunctions()
	DeleteFunction()

	fmt.Println("Successfully ran playground!")
}

func print(data interface{}) {
	b, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(b))
}

func CreateUser() {
	fmt.Println("Running Create user API")

	var name = strconv.Itoa(int(time.Now().Unix()))

	user, _ := appwriteUsers.Create(
		id.Unique(),
		appwriteUsers.WithCreateEmail(name+"@example.com"),
		appwriteUsers.WithCreatePassword(name+"1234"),
		appwriteUsers.WithCreateName(name),
	)

	userId = user.Id

	print(user)
}

func ListUsers() {
	fmt.Println("Running List User API")

	users, _ := appwriteUsers.List()

	print(users)
}

func GetAccount() {
	fmt.Println("Running List Users API")

	response, _ := appwriteAccount.Get()

	print(response)
}

func DeleteUser() {
	fmt.Println("Running Delete User API")

	response, _ := appwriteUsers.Delete(userId)

	print(response)
}

func CreateDatabase() {
	fmt.Println("Running Create Database API")

	database, _ := appwriteDatabase.Create(id.Unique(), "Movies")

	databaseId = database.Id

	print(database)
}

func DeleteDatabase() {
	fmt.Println("Running Delete Database API")

	response, _ := appwriteDatabase.Delete(databaseId)

	print(response)
}

func CreateCollection() {
	fmt.Println("Running Create Collection API")

	appwriteCollection, _ := appwriteDatabase.CreateCollection(
		databaseId,
		id.Unique(),
		"Movies",
		appwriteDatabase.WithCreateCollectionPermissions([]string{permission.Create("any")}),
		appwriteDatabase.WithCreateCollectionDocumentSecurity(true),
	)

	collectionId = appwriteCollection.Id
	print(appwriteCollection)

	nameAttributeResponse, _ := appwriteDatabase.CreateStringAttribute(
		databaseId,
		collectionId,
		"name",
		255,
		true,
	)

	print(nameAttributeResponse)

	yearAttributeResponse, _ := appwriteDatabase.CreateIntegerAttribute(
		databaseId,
		collectionId,
		"release_year",
		true,
		appwriteDatabase.WithCreateIntegerAttributeMin(0),
		appwriteDatabase.WithCreateIntegerAttributeMax(9999),
	)

	print(yearAttributeResponse)

	ratingAttributeResponse, _ := appwriteDatabase.CreateFloatAttribute(
		databaseId,
		collectionId,
		"rating",
		true,
		appwriteDatabase.WithCreateFloatAttributeMin(0),
		appwriteDatabase.WithCreateFloatAttributeMax(99.99),
	)

	print(ratingAttributeResponse)

	kidsAttributeResponse, _ := appwriteDatabase.CreateBooleanAttribute(
		databaseId,
		collectionId,
		"kids",
		true,
	)

	print(kidsAttributeResponse)

	emailAttributeResponse, _ := appwriteDatabase.CreateEmailAttribute(
		databaseId,
		collectionId,
		"email",
		false,
		appwriteDatabase.WithCreateEmailAttributeDefault("example@email.com"),
	)

	print(emailAttributeResponse)

	// Wait for attributes to be created
	time.Sleep(2 * time.Second)

	appwriteDatabase.CreateIndex(
		databaseId,
		collectionId,
		"name_email_idx",
		"fulltext",
		[]string{"email"},
	)
}

func ListCollection() {
	fmt.Println("Running List Collection API")

	var collections, _ = appwriteDatabase.ListCollections(
		databaseId,
	)

	print(collections)
}

func DeleteCollection() {
	fmt.Println("Running Delete Collection API")

	response, _ := appwriteDatabase.DeleteCollection(
		databaseId,
		collectionId,
	)

	print(response)
}

func CreateDocument() {
	fmt.Println("Running Create Document API")

	document, _ := appwriteDatabase.CreateDocument(
		databaseId,
		collectionId,
		id.Unique(),
		map[string]interface{}{
			"name":         "Spider Man",
			"release_year": 1920,
			"rating":       99,
			"kids":         false,
			"email":        "example@email.com",
		})

	documentId = document.Id

	print(document)
}

func ListDocuments() {
	fmt.Println("Running List Document API")

	documents, _ := appwriteDatabase.ListDocuments(databaseId, collectionId)

	print(documents)
}

func DeleteDocument() {
	fmt.Println("Running Delete Document API")

	response, _ := appwriteDatabase.DeleteDocument(
		databaseId,
		collectionId,
		documentId,
	)

	print(response)
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

	bucketId = bucket.Id

	print(bucket)
}

func ListBuckets() {
	fmt.Println("Running List Buckets API")

	buckets, _ := appwriteStorage.ListBuckets()

	print(buckets)
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

// 	fileId = file.Id

// 	print(file)
// }

func ListFiles() {
	fmt.Println("Running List Files API")

	var files, _ = appwriteStorage.ListFiles(bucketId)

	print(files)
}

func DeleteFile() {
	fmt.Println("Running Delete File API")

	var response, _ = appwriteStorage.DeleteFile(bucketId, fileId)

	print(response)
}

func DeleteBucket() {
	fmt.Println("Running Delete Bucket API")

	var response, _ = appwriteStorage.DeleteBucket(bucketId)

	print(response)
}

func CreateFunction() {
	fmt.Println("Running Create Function API")

	var function, _ = appwriteFunction.Create(
		id.Unique(),
		"Test Function",
		"python-3.9",
		appwriteFunction.WithCreateExecute(
			[]string{
				"any",
			}))

	functionId = function.Id

	print(function)
}

func ListFunctions() {
	fmt.Println("Runnnig List Functions API")

	var functions, _ = appwriteFunction.List()

	print(functions)
}

func DeleteFunction() {
	fmt.Println("Running Delete Function API")

	var response, _ = appwriteFunction.Delete(functionId)

	print(response)
}
