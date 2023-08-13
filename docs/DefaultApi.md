# \DefaultApi

All URIs are relative to *https://api.app.shortcut.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](DefaultApi.md#CreateCategory) | **Post** /api/v3/categories | Create Category
[**CreateEntityTemplate**](DefaultApi.md#CreateEntityTemplate) | **Post** /api/v3/entity-templates | Create Entity Template
[**CreateEpic**](DefaultApi.md#CreateEpic) | **Post** /api/v3/epics | Create Epic
[**CreateEpicComment**](DefaultApi.md#CreateEpicComment) | **Post** /api/v3/epics/{epic-public-id}/comments | Create Epic Comment
[**CreateEpicCommentComment**](DefaultApi.md#CreateEpicCommentComment) | **Post** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Create Epic Comment Comment
[**CreateGroup**](DefaultApi.md#CreateGroup) | **Post** /api/v3/groups | Create Group
[**CreateIteration**](DefaultApi.md#CreateIteration) | **Post** /api/v3/iterations | Create Iteration
[**CreateLabel**](DefaultApi.md#CreateLabel) | **Post** /api/v3/labels | Create Label
[**CreateLinkedFile**](DefaultApi.md#CreateLinkedFile) | **Post** /api/v3/linked-files | Create Linked File
[**CreateMilestone**](DefaultApi.md#CreateMilestone) | **Post** /api/v3/milestones | Create Milestone
[**CreateMultipleStories**](DefaultApi.md#CreateMultipleStories) | **Post** /api/v3/stories/bulk | Create Multiple Stories
[**CreateProject**](DefaultApi.md#CreateProject) | **Post** /api/v3/projects | Create Project
[**CreateStory**](DefaultApi.md#CreateStory) | **Post** /api/v3/stories | Create Story
[**CreateStoryComment**](DefaultApi.md#CreateStoryComment) | **Post** /api/v3/stories/{story-public-id}/comments | Create Story Comment
[**CreateStoryLink**](DefaultApi.md#CreateStoryLink) | **Post** /api/v3/story-links | Create Story Link
[**CreateStoryReaction**](DefaultApi.md#CreateStoryReaction) | **Post** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions | Create Story Reaction
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /api/v3/stories/{story-public-id}/tasks | Create Task
[**DeleteCategory**](DefaultApi.md#DeleteCategory) | **Delete** /api/v3/categories/{category-public-id} | Delete Category
[**DeleteCustomField**](DefaultApi.md#DeleteCustomField) | **Delete** /api/v3/custom-fields/{custom-field-public-id} | Delete Custom Field
[**DeleteEntityTemplate**](DefaultApi.md#DeleteEntityTemplate) | **Delete** /api/v3/entity-templates/{entity-template-public-id} | Delete Entity Template
[**DeleteEpic**](DefaultApi.md#DeleteEpic) | **Delete** /api/v3/epics/{epic-public-id} | Delete Epic
[**DeleteEpicComment**](DefaultApi.md#DeleteEpicComment) | **Delete** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Delete Epic Comment
[**DeleteFile**](DefaultApi.md#DeleteFile) | **Delete** /api/v3/files/{file-public-id} | Delete File
[**DeleteIteration**](DefaultApi.md#DeleteIteration) | **Delete** /api/v3/iterations/{iteration-public-id} | Delete Iteration
[**DeleteLabel**](DefaultApi.md#DeleteLabel) | **Delete** /api/v3/labels/{label-public-id} | Delete Label
[**DeleteLinkedFile**](DefaultApi.md#DeleteLinkedFile) | **Delete** /api/v3/linked-files/{linked-file-public-id} | Delete Linked File
[**DeleteMilestone**](DefaultApi.md#DeleteMilestone) | **Delete** /api/v3/milestones/{milestone-public-id} | Delete Milestone
[**DeleteMultipleStories**](DefaultApi.md#DeleteMultipleStories) | **Delete** /api/v3/stories/bulk | Delete Multiple Stories
[**DeleteProject**](DefaultApi.md#DeleteProject) | **Delete** /api/v3/projects/{project-public-id} | Delete Project
[**DeleteStory**](DefaultApi.md#DeleteStory) | **Delete** /api/v3/stories/{story-public-id} | Delete Story
[**DeleteStoryComment**](DefaultApi.md#DeleteStoryComment) | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Delete Story Comment
[**DeleteStoryLink**](DefaultApi.md#DeleteStoryLink) | **Delete** /api/v3/story-links/{story-link-public-id} | Delete Story Link
[**DeleteStoryReaction**](DefaultApi.md#DeleteStoryReaction) | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions | Delete Story Reaction
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Delete Task
[**DisableGroups**](DefaultApi.md#DisableGroups) | **Put** /api/v3/groups/disable | Disable Groups
[**DisableIterations**](DefaultApi.md#DisableIterations) | **Put** /api/v3/iterations/disable | Disable Iterations
[**DisableStoryTemplates**](DefaultApi.md#DisableStoryTemplates) | **Put** /api/v3/entity-templates/disable | Disable Story Templates
[**EnableGroups**](DefaultApi.md#EnableGroups) | **Put** /api/v3/groups/enable | Enable Groups
[**EnableIterations**](DefaultApi.md#EnableIterations) | **Put** /api/v3/iterations/enable | Enable Iterations
[**EnableStoryTemplates**](DefaultApi.md#EnableStoryTemplates) | **Put** /api/v3/entity-templates/enable | Enable Story Templates
[**GetCategory**](DefaultApi.md#GetCategory) | **Get** /api/v3/categories/{category-public-id} | Get Category
[**GetCurrentMemberInfo**](DefaultApi.md#GetCurrentMemberInfo) | **Get** /api/v3/member | Get Current Member Info
[**GetCustomField**](DefaultApi.md#GetCustomField) | **Get** /api/v3/custom-fields/{custom-field-public-id} | Get Custom Field
[**GetEntityTemplate**](DefaultApi.md#GetEntityTemplate) | **Get** /api/v3/entity-templates/{entity-template-public-id} | Get Entity Template
[**GetEpic**](DefaultApi.md#GetEpic) | **Get** /api/v3/epics/{epic-public-id} | Get Epic
[**GetEpicComment**](DefaultApi.md#GetEpicComment) | **Get** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Get Epic Comment
[**GetEpicWorkflow**](DefaultApi.md#GetEpicWorkflow) | **Get** /api/v3/epic-workflow | Get Epic Workflow
[**GetExternalLinkStories**](DefaultApi.md#GetExternalLinkStories) | **Get** /api/v3/external-link/stories | Get External Link Stories
[**GetFile**](DefaultApi.md#GetFile) | **Get** /api/v3/files/{file-public-id} | Get File
[**GetGroup**](DefaultApi.md#GetGroup) | **Get** /api/v3/groups/{group-public-id} | Get Group
[**GetIteration**](DefaultApi.md#GetIteration) | **Get** /api/v3/iterations/{iteration-public-id} | Get Iteration
[**GetLabel**](DefaultApi.md#GetLabel) | **Get** /api/v3/labels/{label-public-id} | Get Label
[**GetLinkedFile**](DefaultApi.md#GetLinkedFile) | **Get** /api/v3/linked-files/{linked-file-public-id} | Get Linked File
[**GetMember**](DefaultApi.md#GetMember) | **Get** /api/v3/members/{member-public-id} | Get Member
[**GetMilestone**](DefaultApi.md#GetMilestone) | **Get** /api/v3/milestones/{milestone-public-id} | Get Milestone
[**GetProject**](DefaultApi.md#GetProject) | **Get** /api/v3/projects/{project-public-id} | Get Project
[**GetRepository**](DefaultApi.md#GetRepository) | **Get** /api/v3/repositories/{repo-public-id} | Get Repository
[**GetStory**](DefaultApi.md#GetStory) | **Get** /api/v3/stories/{story-public-id} | Get Story
[**GetStoryComment**](DefaultApi.md#GetStoryComment) | **Get** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Get Story Comment
[**GetStoryLink**](DefaultApi.md#GetStoryLink) | **Get** /api/v3/story-links/{story-link-public-id} | Get Story Link
[**GetTask**](DefaultApi.md#GetTask) | **Get** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Get Task
[**GetWorkflow**](DefaultApi.md#GetWorkflow) | **Get** /api/v3/workflows/{workflow-public-id} | Get Workflow
[**ListCategories**](DefaultApi.md#ListCategories) | **Get** /api/v3/categories | List Categories
[**ListCategoryMilestones**](DefaultApi.md#ListCategoryMilestones) | **Get** /api/v3/categories/{category-public-id}/milestones | List Category Milestones
[**ListCustomFields**](DefaultApi.md#ListCustomFields) | **Get** /api/v3/custom-fields | List Custom Fields
[**ListEntityTemplates**](DefaultApi.md#ListEntityTemplates) | **Get** /api/v3/entity-templates | List Entity Templates
[**ListEpicComments**](DefaultApi.md#ListEpicComments) | **Get** /api/v3/epics/{epic-public-id}/comments | List Epic Comments
[**ListEpicStories**](DefaultApi.md#ListEpicStories) | **Get** /api/v3/epics/{epic-public-id}/stories | List Epic Stories
[**ListEpics**](DefaultApi.md#ListEpics) | **Get** /api/v3/epics | List Epics
[**ListFiles**](DefaultApi.md#ListFiles) | **Get** /api/v3/files | List Files
[**ListGroupStories**](DefaultApi.md#ListGroupStories) | **Get** /api/v3/groups/{group-public-id}/stories | List Group Stories
[**ListGroups**](DefaultApi.md#ListGroups) | **Get** /api/v3/groups | List Groups
[**ListIterationStories**](DefaultApi.md#ListIterationStories) | **Get** /api/v3/iterations/{iteration-public-id}/stories | List Iteration Stories
[**ListIterations**](DefaultApi.md#ListIterations) | **Get** /api/v3/iterations | List Iterations
[**ListLabelEpics**](DefaultApi.md#ListLabelEpics) | **Get** /api/v3/labels/{label-public-id}/epics | List Label Epics
[**ListLabelStories**](DefaultApi.md#ListLabelStories) | **Get** /api/v3/labels/{label-public-id}/stories | List Label Stories
[**ListLabels**](DefaultApi.md#ListLabels) | **Get** /api/v3/labels | List Labels
[**ListLinkedFiles**](DefaultApi.md#ListLinkedFiles) | **Get** /api/v3/linked-files | List Linked Files
[**ListMembers**](DefaultApi.md#ListMembers) | **Get** /api/v3/members | List Members
[**ListMilestoneEpics**](DefaultApi.md#ListMilestoneEpics) | **Get** /api/v3/milestones/{milestone-public-id}/epics | List Milestone Epics
[**ListMilestones**](DefaultApi.md#ListMilestones) | **Get** /api/v3/milestones | List Milestones
[**ListProjects**](DefaultApi.md#ListProjects) | **Get** /api/v3/projects | List Projects
[**ListRepositories**](DefaultApi.md#ListRepositories) | **Get** /api/v3/repositories | List Repositories
[**ListStories**](DefaultApi.md#ListStories) | **Get** /api/v3/projects/{project-public-id}/stories | List Stories
[**ListStoryComment**](DefaultApi.md#ListStoryComment) | **Get** /api/v3/stories/{story-public-id}/comments | List Story Comment
[**ListWorkflows**](DefaultApi.md#ListWorkflows) | **Get** /api/v3/workflows | List Workflows
[**Search**](DefaultApi.md#Search) | **Get** /api/v3/search | Search
[**SearchEpics**](DefaultApi.md#SearchEpics) | **Get** /api/v3/search/epics | Search Epics
[**SearchIterations**](DefaultApi.md#SearchIterations) | **Get** /api/v3/search/iterations | Search Iterations
[**SearchMilestones**](DefaultApi.md#SearchMilestones) | **Get** /api/v3/search/milestones | Search Milestones
[**SearchStories**](DefaultApi.md#SearchStories) | **Get** /api/v3/search/stories | Search Stories
[**SearchStoriesOld**](DefaultApi.md#SearchStoriesOld) | **Post** /api/v3/stories/search | Search Stories (Old)
[**StoryHistory**](DefaultApi.md#StoryHistory) | **Get** /api/v3/stories/{story-public-id}/history | Story History
[**UnlinkProductboardFromEpic**](DefaultApi.md#UnlinkProductboardFromEpic) | **Post** /api/v3/epics/{epic-public-id}/unlink-productboard | Unlink Productboard from Epic
[**UpdateCategory**](DefaultApi.md#UpdateCategory) | **Put** /api/v3/categories/{category-public-id} | Update Category
[**UpdateCustomField**](DefaultApi.md#UpdateCustomField) | **Put** /api/v3/custom-fields/{custom-field-public-id} | Update Custom Field
[**UpdateEntityTemplate**](DefaultApi.md#UpdateEntityTemplate) | **Put** /api/v3/entity-templates/{entity-template-public-id} | Update Entity Template
[**UpdateEpic**](DefaultApi.md#UpdateEpic) | **Put** /api/v3/epics/{epic-public-id} | Update Epic
[**UpdateEpicComment**](DefaultApi.md#UpdateEpicComment) | **Put** /api/v3/epics/{epic-public-id}/comments/{comment-public-id} | Update Epic Comment
[**UpdateFile**](DefaultApi.md#UpdateFile) | **Put** /api/v3/files/{file-public-id} | Update File
[**UpdateGroup**](DefaultApi.md#UpdateGroup) | **Put** /api/v3/groups/{group-public-id} | Update Group
[**UpdateIteration**](DefaultApi.md#UpdateIteration) | **Put** /api/v3/iterations/{iteration-public-id} | Update Iteration
[**UpdateLabel**](DefaultApi.md#UpdateLabel) | **Put** /api/v3/labels/{label-public-id} | Update Label
[**UpdateLinkedFile**](DefaultApi.md#UpdateLinkedFile) | **Put** /api/v3/linked-files/{linked-file-public-id} | Update Linked File
[**UpdateMilestone**](DefaultApi.md#UpdateMilestone) | **Put** /api/v3/milestones/{milestone-public-id} | Update Milestone
[**UpdateMultipleStories**](DefaultApi.md#UpdateMultipleStories) | **Put** /api/v3/stories/bulk | Update Multiple Stories
[**UpdateProject**](DefaultApi.md#UpdateProject) | **Put** /api/v3/projects/{project-public-id} | Update Project
[**UpdateStory**](DefaultApi.md#UpdateStory) | **Put** /api/v3/stories/{story-public-id} | Update Story
[**UpdateStoryComment**](DefaultApi.md#UpdateStoryComment) | **Put** /api/v3/stories/{story-public-id}/comments/{comment-public-id} | Update Story Comment
[**UpdateStoryLink**](DefaultApi.md#UpdateStoryLink) | **Put** /api/v3/story-links/{story-link-public-id} | Update Story Link
[**UpdateTask**](DefaultApi.md#UpdateTask) | **Put** /api/v3/stories/{story-public-id}/tasks/{task-public-id} | Update Task
[**UploadFiles**](DefaultApi.md#UploadFiles) | **Post** /api/v3/files | Upload Files



## CreateCategory

> Category CreateCategory(ctx).CreateCategory(createCategory).Execute()

Create Category



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createCategory := *openapiclient.NewCreateCategory("Name_example", "Type_example") // CreateCategory | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCategory(context.Background()).CreateCategory(createCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCategory** | [**CreateCategory**](CreateCategory.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEntityTemplate

> EntityTemplate CreateEntityTemplate(ctx).CreateEntityTemplate(createEntityTemplate).Execute()

Create Entity Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createEntityTemplate := *openapiclient.NewCreateEntityTemplate("Name_example", *openapiclient.NewCreateStoryContents()) // CreateEntityTemplate | Request paramaters for creating an entirely new entity template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEntityTemplate(context.Background()).CreateEntityTemplate(createEntityTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEntityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEntityTemplate`: EntityTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEntityTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEntityTemplate** | [**CreateEntityTemplate**](CreateEntityTemplate.md) | Request paramaters for creating an entirely new entity template. | 

### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEpic

> Epic CreateEpic(ctx).CreateEpic(createEpic).Execute()

Create Epic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createEpic := *openapiclient.NewCreateEpic("Name_example") // CreateEpic | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEpic(context.Background()).CreateEpic(createEpic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEpic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEpic`: Epic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEpic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEpic** | [**CreateEpic**](CreateEpic.md) |  | 

### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEpicComment

> ThreadedComment CreateEpicComment(ctx, epicPublicId).CreateEpicComment(createEpicComment).Execute()

Create Epic Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The ID of the associated Epic.
    createEpicComment := *openapiclient.NewCreateEpicComment("Text_example") // CreateEpicComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEpicComment(context.Background(), epicPublicId).CreateEpicComment(createEpicComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEpicComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEpicComment`: ThreadedComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEpicComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The ID of the associated Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEpicCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEpicComment** | [**CreateEpicComment**](CreateEpicComment.md) |  | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEpicCommentComment

> ThreadedComment CreateEpicCommentComment(ctx, epicPublicId, commentPublicId).CreateCommentComment(createCommentComment).Execute()

Create Epic Comment Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The ID of the associated Epic.
    commentPublicId := int64(789) // int64 | The ID of the parent Epic Comment.
    createCommentComment := *openapiclient.NewCreateCommentComment("Text_example") // CreateCommentComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEpicCommentComment(context.Background(), epicPublicId, commentPublicId).CreateCommentComment(createCommentComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEpicCommentComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEpicCommentComment`: ThreadedComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEpicCommentComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The ID of the associated Epic. | 
**commentPublicId** | **int64** | The ID of the parent Epic Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEpicCommentCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createCommentComment** | [**CreateCommentComment**](CreateCommentComment.md) |  | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).CreateGroup(createGroup).Execute()

Create Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createGroup := *openapiclient.NewCreateGroup("Name_example", "MentionName_example") // CreateGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateGroup(context.Background()).CreateGroup(createGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroup** | [**CreateGroup**](CreateGroup.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIteration

> Iteration CreateIteration(ctx).CreateIteration(createIteration).Execute()

Create Iteration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createIteration := *openapiclient.NewCreateIteration("Name_example", "StartDate_example", "EndDate_example") // CreateIteration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateIteration(context.Background()).CreateIteration(createIteration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIteration`: Iteration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateIteration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIteration** | [**CreateIteration**](CreateIteration.md) |  | 

### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> Label CreateLabel(ctx).CreateLabelParams(createLabelParams).Execute()

Create Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createLabelParams := *openapiclient.NewCreateLabelParams("Name_example") // CreateLabelParams | Request parameters for creating a Label on a Shortcut Story.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateLabel(context.Background()).CreateLabelParams(createLabelParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLabelParams** | [**CreateLabelParams**](CreateLabelParams.md) | Request parameters for creating a Label on a Shortcut Story. | 

### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLinkedFile

> LinkedFile CreateLinkedFile(ctx).CreateLinkedFile(createLinkedFile).Execute()

Create Linked File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createLinkedFile := *openapiclient.NewCreateLinkedFile("Name_example", "Type_example", "Url_example") // CreateLinkedFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateLinkedFile(context.Background()).CreateLinkedFile(createLinkedFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateLinkedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLinkedFile`: LinkedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateLinkedFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLinkedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLinkedFile** | [**CreateLinkedFile**](CreateLinkedFile.md) |  | 

### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMilestone

> Milestone CreateMilestone(ctx).CreateMilestone(createMilestone).Execute()

Create Milestone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createMilestone := *openapiclient.NewCreateMilestone("Name_example") // CreateMilestone | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMilestone(context.Background()).CreateMilestone(createMilestone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMilestone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMilestone** | [**CreateMilestone**](CreateMilestone.md) |  | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultipleStories

> []StorySlim CreateMultipleStories(ctx).CreateStories(createStories).Execute()

Create Multiple Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createStories := *openapiclient.NewCreateStories([]openapiclient.CreateStoryParams{*openapiclient.NewCreateStoryParams("Name_example")}) // CreateStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateMultipleStories(context.Background()).CreateStories(createStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateMultipleStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultipleStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateMultipleStories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipleStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStories** | [**CreateStories**](CreateStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateProject

> Project CreateProject(ctx).CreateProject(createProject).Execute()

Create Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createProject := *openapiclient.NewCreateProject("Name_example", int64(123)) // CreateProject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateProject(context.Background()).CreateProject(createProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createProject** | [**CreateProject**](CreateProject.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStory

> Story CreateStory(ctx).CreateStoryParams(createStoryParams).Execute()

Create Story



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createStoryParams := *openapiclient.NewCreateStoryParams("Name_example") // CreateStoryParams | Request parameters for creating a story.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStory(context.Background()).CreateStoryParams(createStoryParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStory`: Story
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStoryParams** | [**CreateStoryParams**](CreateStoryParams.md) | Request parameters for creating a story. | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStoryComment

> StoryComment CreateStoryComment(ctx, storyPublicId).CreateStoryComment(createStoryComment).Execute()

Create Story Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    createStoryComment := *openapiclient.NewCreateStoryComment("Text_example") // CreateStoryComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStoryComment(context.Background(), storyPublicId).CreateStoryComment(createStoryComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStoryComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStoryComment`: StoryComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStoryComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoryCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createStoryComment** | [**CreateStoryComment**](CreateStoryComment.md) |  | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStoryLink

> StoryLink CreateStoryLink(ctx).CreateStoryLink(createStoryLink).Execute()

Create Story Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    createStoryLink := *openapiclient.NewCreateStoryLink("Verb_example", int64(123), int64(123)) // CreateStoryLink | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStoryLink(context.Background()).CreateStoryLink(createStoryLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStoryLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStoryLink`: StoryLink
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStoryLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoryLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStoryLink** | [**CreateStoryLink**](CreateStoryLink.md) |  | 

### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStoryReaction

> []StoryReaction CreateStoryReaction(ctx, storyPublicId, commentPublicId).CreateOrDeleteStoryReaction(createOrDeleteStoryReaction).Execute()

Create Story Reaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    commentPublicId := int64(789) // int64 | The ID of the Comment.
    createOrDeleteStoryReaction := *openapiclient.NewCreateOrDeleteStoryReaction("Emoji_example") // CreateOrDeleteStoryReaction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateStoryReaction(context.Background(), storyPublicId, commentPublicId).CreateOrDeleteStoryReaction(createOrDeleteStoryReaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateStoryReaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStoryReaction`: []StoryReaction
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateStoryReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoryReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrDeleteStoryReaction** | [**CreateOrDeleteStoryReaction**](CreateOrDeleteStoryReaction.md) |  | 

### Return type

[**[]StoryReaction**](StoryReaction.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> Task CreateTask(ctx, storyPublicId).CreateTask(createTask).Execute()

Create Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Task will be in.
    createTask := *openapiclient.NewCreateTask("Description_example") // CreateTask | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateTask(context.Background(), storyPublicId).CreateTask(createTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Task will be in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTask** | [**CreateTask**](CreateTask.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCategory

> DeleteCategory(ctx, categoryPublicId).Execute()

Delete Category



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    categoryPublicId := int64(789) // int64 | The unique ID of the Category.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteCategory(context.Background(), categoryPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryPublicId** | **int64** | The unique ID of the Category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteCustomField(ctx, customFieldPublicId).Execute()

Delete Custom Field

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    customFieldPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the CustomField.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteCustomField(context.Background(), customFieldPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldPublicId** | **string** | The unique ID of the CustomField. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntityTemplate

> DeleteEntityTemplate(ctx, entityTemplatePublicId).Execute()

Delete Entity Template

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityTemplatePublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the entity template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteEntityTemplate(context.Background(), entityTemplatePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEntityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityTemplatePublicId** | **string** | The unique ID of the entity template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEpic

> DeleteEpic(ctx, epicPublicId).Execute()

Delete Epic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteEpic(context.Background(), epicPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEpic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEpicComment

> DeleteEpicComment(ctx, epicPublicId, commentPublicId).Execute()

Delete Epic Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The ID of the associated Epic.
    commentPublicId := int64(789) // int64 | The ID of the Comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteEpicComment(context.Background(), epicPublicId, commentPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteEpicComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The ID of the associated Epic. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEpicCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFile

> DeleteFile(ctx, filePublicId).Execute()

Delete File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    filePublicId := int64(789) // int64 | The File’s unique ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteFile(context.Background(), filePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filePublicId** | **int64** | The File’s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIteration

> DeleteIteration(ctx, iterationPublicId).Execute()

Delete Iteration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    iterationPublicId := int64(789) // int64 | The unique ID of the Iteration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteIteration(context.Background(), iterationPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iterationPublicId** | **int64** | The unique ID of the Iteration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, labelPublicId).Execute()

Delete Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    labelPublicId := int64(789) // int64 | The unique ID of the Label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteLabel(context.Background(), labelPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelPublicId** | **int64** | The unique ID of the Label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLinkedFile

> DeleteLinkedFile(ctx, linkedFilePublicId).Execute()

Delete Linked File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    linkedFilePublicId := int64(789) // int64 | The unique identifier of the linked file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteLinkedFile(context.Background(), linkedFilePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteLinkedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedFilePublicId** | **int64** | The unique identifier of the linked file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMilestone

> DeleteMilestone(ctx, milestonePublicId).Execute()

Delete Milestone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    milestonePublicId := int64(789) // int64 | The ID of the Milestone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteMilestone(context.Background(), milestonePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestonePublicId** | **int64** | The ID of the Milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMultipleStories

> DeleteMultipleStories(ctx).DeleteStories(deleteStories).Execute()

Delete Multiple Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    deleteStories := *openapiclient.NewDeleteStories([]int64{int64(123)}) // DeleteStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteMultipleStories(context.Background()).DeleteStories(deleteStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteMultipleStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMultipleStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteStories** | [**DeleteStories**](DeleteStories.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProject

> DeleteProject(ctx, projectPublicId).Execute()

Delete Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPublicId := int64(789) // int64 | The unique ID of the Project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteProject(context.Background(), projectPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPublicId** | **int64** | The unique ID of the Project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStory

> DeleteStory(ctx, storyPublicId).Execute()

Delete Story



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStory(context.Background(), storyPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoryComment

> DeleteStoryComment(ctx, storyPublicId, commentPublicId).Execute()

Delete Story Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    commentPublicId := int64(789) // int64 | The ID of the Comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStoryComment(context.Background(), storyPublicId, commentPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStoryComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoryCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoryLink

> DeleteStoryLink(ctx, storyLinkPublicId).Execute()

Delete Story Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyLinkPublicId := int64(789) // int64 | The unique ID of the Story Link.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStoryLink(context.Background(), storyLinkPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStoryLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyLinkPublicId** | **int64** | The unique ID of the Story Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoryLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStoryReaction

> DeleteStoryReaction(ctx, storyPublicId, commentPublicId).CreateOrDeleteStoryReaction(createOrDeleteStoryReaction).Execute()

Delete Story Reaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    commentPublicId := int64(789) // int64 | The ID of the Comment.
    createOrDeleteStoryReaction := *openapiclient.NewCreateOrDeleteStoryReaction("Emoji_example") // CreateOrDeleteStoryReaction | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteStoryReaction(context.Background(), storyPublicId, commentPublicId).CreateOrDeleteStoryReaction(createOrDeleteStoryReaction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteStoryReaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoryReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createOrDeleteStoryReaction** | [**CreateOrDeleteStoryReaction**](CreateOrDeleteStoryReaction.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> DeleteTask(ctx, storyPublicId, taskPublicId).Execute()

Delete Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The unique ID of the Story this Task is associated with.
    taskPublicId := int64(789) // int64 | The unique ID of the Task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DeleteTask(context.Background(), storyPublicId, taskPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The unique ID of the Story this Task is associated with. | 
**taskPublicId** | **int64** | The unique ID of the Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableGroups

> DisableGroups(ctx).Execute()

Disable Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DisableGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DisableGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableGroupsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableIterations

> DisableIterations(ctx).Execute()

Disable Iterations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DisableIterations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DisableIterations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableIterationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableStoryTemplates

> DisableStoryTemplates(ctx).Execute()

Disable Story Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.DisableStoryTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DisableStoryTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableStoryTemplatesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableGroups

> EnableGroups(ctx).Execute()

Enable Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.EnableGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EnableGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableGroupsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableIterations

> EnableIterations(ctx).Execute()

Enable Iterations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.EnableIterations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EnableIterations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableIterationsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableStoryTemplates

> EnableStoryTemplates(ctx).Execute()

Enable Story Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.EnableStoryTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EnableStoryTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableStoryTemplatesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategory

> Category GetCategory(ctx, categoryPublicId).Execute()

Get Category



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    categoryPublicId := int64(789) // int64 | The unique ID of the Category.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCategory(context.Background(), categoryPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryPublicId** | **int64** | The unique ID of the Category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentMemberInfo

> MemberInfo GetCurrentMemberInfo(ctx).Execute()

Get Current Member Info



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCurrentMemberInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCurrentMemberInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentMemberInfo`: MemberInfo
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCurrentMemberInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentMemberInfoRequest struct via the builder pattern


### Return type

[**MemberInfo**](MemberInfo.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomField

> CustomField GetCustomField(ctx, customFieldPublicId).Execute()

Get Custom Field

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    customFieldPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the CustomField.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCustomField(context.Background(), customFieldPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldPublicId** | **string** | The unique ID of the CustomField. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntityTemplate

> EntityTemplate GetEntityTemplate(ctx, entityTemplatePublicId).Execute()

Get Entity Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityTemplatePublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the entity template.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEntityTemplate(context.Background(), entityTemplatePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEntityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEntityTemplate`: EntityTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEntityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityTemplatePublicId** | **string** | The unique ID of the entity template. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpic

> Epic GetEpic(ctx, epicPublicId).Execute()

Get Epic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEpic(context.Background(), epicPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEpic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEpic`: Epic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEpic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpicComment

> ThreadedComment GetEpicComment(ctx, epicPublicId, commentPublicId).Execute()

Get Epic Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The ID of the associated Epic.
    commentPublicId := int64(789) // int64 | The ID of the Comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEpicComment(context.Background(), epicPublicId, commentPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEpicComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEpicComment`: ThreadedComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEpicComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The ID of the associated Epic. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpicWorkflow

> EpicWorkflow GetEpicWorkflow(ctx).Execute()

Get Epic Workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetEpicWorkflow(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetEpicWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEpicWorkflow`: EpicWorkflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetEpicWorkflow`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicWorkflowRequest struct via the builder pattern


### Return type

[**EpicWorkflow**](EpicWorkflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExternalLinkStories

> []StorySlim GetExternalLinkStories(ctx).GetExternalLinkStoriesParams(getExternalLinkStoriesParams).Execute()

Get External Link Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    getExternalLinkStoriesParams := *openapiclient.NewGetExternalLinkStoriesParams("ExternalLink_example") // GetExternalLinkStoriesParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetExternalLinkStories(context.Background()).GetExternalLinkStoriesParams(getExternalLinkStoriesParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetExternalLinkStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalLinkStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetExternalLinkStories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalLinkStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getExternalLinkStoriesParams** | [**GetExternalLinkStoriesParams**](GetExternalLinkStoriesParams.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> UploadedFile GetFile(ctx, filePublicId).Execute()

Get File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    filePublicId := int64(789) // int64 | The File’s unique ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetFile(context.Background(), filePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFile`: UploadedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filePublicId** | **int64** | The File’s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx, groupPublicId).Execute()

Get Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    groupPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the Group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetGroup(context.Background(), groupPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPublicId** | **string** | The unique ID of the Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIteration

> Iteration GetIteration(ctx, iterationPublicId).Execute()

Get Iteration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    iterationPublicId := int64(789) // int64 | The unique ID of the Iteration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetIteration(context.Background(), iterationPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIteration`: Iteration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetIteration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iterationPublicId** | **int64** | The unique ID of the Iteration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabel

> Label GetLabel(ctx, labelPublicId).Execute()

Get Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    labelPublicId := int64(789) // int64 | The unique ID of the Label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetLabel(context.Background(), labelPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelPublicId** | **int64** | The unique ID of the Label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedFile

> LinkedFile GetLinkedFile(ctx, linkedFilePublicId).Execute()

Get Linked File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    linkedFilePublicId := int64(789) // int64 | The unique identifier of the linked file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetLinkedFile(context.Background(), linkedFilePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetLinkedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLinkedFile`: LinkedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetLinkedFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedFilePublicId** | **int64** | The unique identifier of the linked file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMember

> Member GetMember(ctx, memberPublicId).GetMember(getMember).Execute()

Get Member



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    memberPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Member's unique ID.
    getMember := *openapiclient.NewGetMember() // GetMember | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetMember(context.Background(), memberPublicId).GetMember(getMember).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMember`: Member
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberPublicId** | **string** | The Member&#39;s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getMember** | [**GetMember**](GetMember.md) |  | 

### Return type

[**Member**](Member.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMilestone

> Milestone GetMilestone(ctx, milestonePublicId).Execute()

Get Milestone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    milestonePublicId := int64(789) // int64 | The ID of the Milestone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetMilestone(context.Background(), milestonePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestonePublicId** | **int64** | The ID of the Milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, projectPublicId).Execute()

Get Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPublicId := int64(789) // int64 | The unique ID of the Project.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetProject(context.Background(), projectPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPublicId** | **int64** | The unique ID of the Project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> Repository GetRepository(ctx, repoPublicId).Execute()

Get Repository



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    repoPublicId := int64(789) // int64 | The unique ID of the Repository.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRepository(context.Background(), repoPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRepository`: Repository
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repoPublicId** | **int64** | The unique ID of the Repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Repository**](Repository.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStory

> Story GetStory(ctx, storyPublicId).Execute()

Get Story



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStory(context.Background(), storyPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStory`: Story
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoryComment

> StoryComment GetStoryComment(ctx, storyPublicId, commentPublicId).Execute()

Get Story Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    commentPublicId := int64(789) // int64 | The ID of the Comment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoryComment(context.Background(), storyPublicId, commentPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoryComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoryComment`: StoryComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoryComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoryCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStoryLink

> StoryLink GetStoryLink(ctx, storyLinkPublicId).Execute()

Get Story Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyLinkPublicId := int64(789) // int64 | The unique ID of the Story Link.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStoryLink(context.Background(), storyLinkPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStoryLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStoryLink`: StoryLink
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStoryLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyLinkPublicId** | **int64** | The unique ID of the Story Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoryLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx, storyPublicId, taskPublicId).Execute()

Get Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The unique ID of the Story this Task is associated with.
    taskPublicId := int64(789) // int64 | The unique ID of the Task.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTask(context.Background(), storyPublicId, taskPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The unique ID of the Story this Task is associated with. | 
**taskPublicId** | **int64** | The unique ID of the Task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> Workflow GetWorkflow(ctx, workflowPublicId).Execute()

Get Workflow



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    workflowPublicId := int64(789) // int64 | The ID of the Workflow.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetWorkflow(context.Background(), workflowPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflow`: Workflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowPublicId** | **int64** | The ID of the Workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Workflow**](Workflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> []Category ListCategories(ctx).Execute()

List Categories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: []Category
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


### Return type

[**[]Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategoryMilestones

> []Milestone ListCategoryMilestones(ctx, categoryPublicId).Execute()

List Category Milestones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    categoryPublicId := int64(789) // int64 | The unique ID of the Category.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCategoryMilestones(context.Background(), categoryPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCategoryMilestones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategoryMilestones`: []Milestone
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCategoryMilestones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryPublicId** | **int64** | The unique ID of the Category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoryMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFields

> []CustomField ListCustomFields(ctx).Execute()

List Custom Fields

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCustomFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCustomFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomFields`: []CustomField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCustomFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldsRequest struct via the builder pattern


### Return type

[**[]CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntityTemplates

> []EntityTemplate ListEntityTemplates(ctx).Execute()

List Entity Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEntityTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEntityTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEntityTemplates`: []EntityTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEntityTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEntityTemplatesRequest struct via the builder pattern


### Return type

[**[]EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpicComments

> []ThreadedComment ListEpicComments(ctx, epicPublicId).Execute()

List Epic Comments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEpicComments(context.Background(), epicPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEpicComments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEpicComments`: []ThreadedComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEpicComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEpicCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpicStories

> []StorySlim ListEpicStories(ctx, epicPublicId).GetEpicStories(getEpicStories).Execute()

List Epic Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.
    getEpicStories := *openapiclient.NewGetEpicStories() // GetEpicStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEpicStories(context.Background(), epicPublicId).GetEpicStories(getEpicStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEpicStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEpicStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEpicStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEpicStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getEpicStories** | [**GetEpicStories**](GetEpicStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEpics

> []EpicSlim ListEpics(ctx).ListEpics(listEpics).Execute()

List Epics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    listEpics := *openapiclient.NewListEpics() // ListEpics | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListEpics(context.Background()).ListEpics(listEpics).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListEpics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEpics`: []EpicSlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListEpics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listEpics** | [**ListEpics**](ListEpics.md) |  | 

### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiles

> []UploadedFile ListFiles(ctx).Execute()

List Files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListFiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiles`: []UploadedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


### Return type

[**[]UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupStories

> []StorySlim ListGroupStories(ctx, groupPublicId).ListGroupStories(listGroupStories).Execute()

List Group Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    groupPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the Group.
    listGroupStories := *openapiclient.NewListGroupStories() // ListGroupStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListGroupStories(context.Background(), groupPublicId).ListGroupStories(listGroupStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGroupStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGroupStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPublicId** | **string** | The unique ID of the Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listGroupStories** | [**ListGroupStories**](ListGroupStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> []Group ListGroups(ctx).Execute()

List Groups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: []Group
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


### Return type

[**[]Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIterationStories

> []StorySlim ListIterationStories(ctx, iterationPublicId).GetIterationStories(getIterationStories).Execute()

List Iteration Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    iterationPublicId := int64(789) // int64 | The unique ID of the Iteration.
    getIterationStories := *openapiclient.NewGetIterationStories() // GetIterationStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIterationStories(context.Background(), iterationPublicId).GetIterationStories(getIterationStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIterationStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIterationStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIterationStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iterationPublicId** | **int64** | The unique ID of the Iteration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListIterationStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getIterationStories** | [**GetIterationStories**](GetIterationStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIterations

> []IterationSlim ListIterations(ctx).Execute()

List Iterations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListIterations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListIterations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIterations`: []IterationSlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListIterations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIterationsRequest struct via the builder pattern


### Return type

[**[]IterationSlim**](IterationSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabelEpics

> []EpicSlim ListLabelEpics(ctx, labelPublicId).Execute()

List Label Epics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    labelPublicId := int64(789) // int64 | The unique ID of the Label.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLabelEpics(context.Background(), labelPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLabelEpics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabelEpics`: []EpicSlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLabelEpics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelPublicId** | **int64** | The unique ID of the Label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabelStories

> []StorySlim ListLabelStories(ctx, labelPublicId).GetLabelStories(getLabelStories).Execute()

List Label Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    labelPublicId := int64(789) // int64 | The unique ID of the Label.
    getLabelStories := *openapiclient.NewGetLabelStories() // GetLabelStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLabelStories(context.Background(), labelPublicId).GetLabelStories(getLabelStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLabelStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabelStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLabelStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelPublicId** | **int64** | The unique ID of the Label. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLabelStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getLabelStories** | [**GetLabelStories**](GetLabelStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

> []Label ListLabels(ctx).ListLabels(listLabels).Execute()

List Labels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    listLabels := *openapiclient.NewListLabels() // ListLabels | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLabels(context.Background()).ListLabels(listLabels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLabels`: []Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listLabels** | [**ListLabels**](ListLabels.md) |  | 

### Return type

[**[]Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLinkedFiles

> []LinkedFile ListLinkedFiles(ctx).Execute()

List Linked Files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListLinkedFiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListLinkedFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLinkedFiles`: []LinkedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListLinkedFiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLinkedFilesRequest struct via the builder pattern


### Return type

[**[]LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembers

> []Member ListMembers(ctx).ListMembers(listMembers).Execute()

List Members



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    listMembers := *openapiclient.NewListMembers() // ListMembers | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMembers(context.Background()).ListMembers(listMembers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: []Member
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listMembers** | [**ListMembers**](ListMembers.md) |  | 

### Return type

[**[]Member**](Member.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMilestoneEpics

> []EpicSlim ListMilestoneEpics(ctx, milestonePublicId).Execute()

List Milestone Epics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    milestonePublicId := int64(789) // int64 | The ID of the Milestone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMilestoneEpics(context.Background(), milestonePublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMilestoneEpics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMilestoneEpics`: []EpicSlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMilestoneEpics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestonePublicId** | **int64** | The ID of the Milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMilestoneEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EpicSlim**](EpicSlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMilestones

> []Milestone ListMilestones(ctx).Execute()

List Milestones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListMilestones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListMilestones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMilestones`: []Milestone
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListMilestones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMilestonesRequest struct via the builder pattern


### Return type

[**[]Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProjects

> []Project ListProjects(ctx).Execute()

List Projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListProjects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProjects`: []Project
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListProjects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProjectsRequest struct via the builder pattern


### Return type

[**[]Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepositories

> []Repository ListRepositories(ctx).Execute()

List Repositories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRepositories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRepositories`: []Repository
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoriesRequest struct via the builder pattern


### Return type

[**[]Repository**](Repository.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStories

> []StorySlim ListStories(ctx, projectPublicId).GetProjectStories(getProjectStories).Execute()

List Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPublicId := int64(789) // int64 | The unique ID of the Project.
    getProjectStories := *openapiclient.NewGetProjectStories() // GetProjectStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListStories(context.Background(), projectPublicId).GetProjectStories(getProjectStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPublicId** | **int64** | The unique ID of the Project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getProjectStories** | [**GetProjectStories**](GetProjectStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStoryComment

> []StoryComment ListStoryComment(ctx, storyPublicId).Execute()

List Story Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListStoryComment(context.Background(), storyPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListStoryComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStoryComment`: []StoryComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListStoryComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStoryCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflows

> []Workflow ListWorkflows(ctx).Execute()

List Workflows



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListWorkflows(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflows`: []Workflow
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


### Return type

[**[]Workflow**](Workflow.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> SearchResults Search(ctx).Search(search).Execute()

Search



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := *openapiclient.NewSearch("Query_example") // Search | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Search(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Search`: SearchResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchEpics

> EpicSearchResults SearchEpics(ctx).Search(search).Execute()

Search Epics



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := *openapiclient.NewSearch("Query_example") // Search | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchEpics(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchEpics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchEpics`: EpicSearchResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchEpics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**EpicSearchResults**](EpicSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchIterations

> IterationSearchResults SearchIterations(ctx).Search(search).Execute()

Search Iterations



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := *openapiclient.NewSearch("Query_example") // Search | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchIterations(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchIterations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchIterations`: IterationSearchResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchIterations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIterationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**IterationSearchResults**](IterationSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMilestones

> MilestoneSearchResults SearchMilestones(ctx).Search(search).Execute()

Search Milestones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := *openapiclient.NewSearch("Query_example") // Search | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchMilestones(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchMilestones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchMilestones`: MilestoneSearchResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchMilestones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMilestonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**MilestoneSearchResults**](MilestoneSearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStories

> StorySearchResults SearchStories(ctx).Search(search).Execute()

Search Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    search := *openapiclient.NewSearch("Query_example") // Search | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchStories(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchStories`: StorySearchResults
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchStories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | [**Search**](Search.md) |  | 

### Return type

[**StorySearchResults**](StorySearchResults.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchStoriesOld

> []StorySlim SearchStoriesOld(ctx).SearchStories(searchStories).Execute()

Search Stories (Old)



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    searchStories := *openapiclient.NewSearchStories() // SearchStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SearchStoriesOld(context.Background()).SearchStories(searchStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SearchStoriesOld``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchStoriesOld`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SearchStoriesOld`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchStoriesOldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchStories** | [**SearchStories**](SearchStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoryHistory

> []History StoryHistory(ctx, storyPublicId).Execute()

Story History

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.StoryHistory(context.Background(), storyPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.StoryHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoryHistory`: []History
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.StoryHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoryHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]History**](History.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnlinkProductboardFromEpic

> UnlinkProductboardFromEpic(ctx, epicPublicId).Execute()

Unlink Productboard from Epic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultApi.UnlinkProductboardFromEpic(context.Background(), epicPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UnlinkProductboardFromEpic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnlinkProductboardFromEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCategory

> Category UpdateCategory(ctx, categoryPublicId).UpdateCategory(updateCategory).Execute()

Update Category



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    categoryPublicId := int64(789) // int64 | The unique ID of the Category you wish to update.
    updateCategory := *openapiclient.NewUpdateCategory() // UpdateCategory | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCategory(context.Background(), categoryPublicId).UpdateCategory(updateCategory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**categoryPublicId** | **int64** | The unique ID of the Category you wish to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCategory** | [**UpdateCategory**](UpdateCategory.md) |  | 

### Return type

[**Category**](Category.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> CustomField UpdateCustomField(ctx, customFieldPublicId).UpdateCustomField(updateCustomField).Execute()

Update Custom Field



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    customFieldPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the CustomField.
    updateCustomField := *openapiclient.NewUpdateCustomField() // UpdateCustomField | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCustomField(context.Background(), customFieldPublicId).UpdateCustomField(updateCustomField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customFieldPublicId** | **string** | The unique ID of the CustomField. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomField** | [**UpdateCustomField**](UpdateCustomField.md) |  | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntityTemplate

> EntityTemplate UpdateEntityTemplate(ctx, entityTemplatePublicId).UpdateEntityTemplate(updateEntityTemplate).Execute()

Update Entity Template



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    entityTemplatePublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the template to be updated.
    updateEntityTemplate := *openapiclient.NewUpdateEntityTemplate() // UpdateEntityTemplate | Request parameters for changing either a template's name or any of   the attributes it is designed to pre-populate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEntityTemplate(context.Background(), entityTemplatePublicId).UpdateEntityTemplate(updateEntityTemplate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEntityTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEntityTemplate`: EntityTemplate
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEntityTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityTemplatePublicId** | **string** | The unique ID of the template to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntityTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEntityTemplate** | [**UpdateEntityTemplate**](UpdateEntityTemplate.md) | Request parameters for changing either a template&#39;s name or any of   the attributes it is designed to pre-populate. | 

### Return type

[**EntityTemplate**](EntityTemplate.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEpic

> Epic UpdateEpic(ctx, epicPublicId).UpdateEpic(updateEpic).Execute()

Update Epic



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The unique ID of the Epic.
    updateEpic := *openapiclient.NewUpdateEpic() // UpdateEpic | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEpic(context.Background(), epicPublicId).UpdateEpic(updateEpic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEpic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEpic`: Epic
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEpic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The unique ID of the Epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateEpic** | [**UpdateEpic**](UpdateEpic.md) |  | 

### Return type

[**Epic**](Epic.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEpicComment

> ThreadedComment UpdateEpicComment(ctx, epicPublicId, commentPublicId).UpdateComment(updateComment).Execute()

Update Epic Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    epicPublicId := int64(789) // int64 | The ID of the associated Epic.
    commentPublicId := int64(789) // int64 | The ID of the Comment.
    updateComment := *openapiclient.NewUpdateComment("Text_example") // UpdateComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateEpicComment(context.Background(), epicPublicId, commentPublicId).UpdateComment(updateComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateEpicComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEpicComment`: ThreadedComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateEpicComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicPublicId** | **int64** | The ID of the associated Epic. | 
**commentPublicId** | **int64** | The ID of the Comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEpicCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateComment** | [**UpdateComment**](UpdateComment.md) |  | 

### Return type

[**ThreadedComment**](ThreadedComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFile

> UploadedFile UpdateFile(ctx, filePublicId).UpdateFile(updateFile).Execute()

Update File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    filePublicId := int64(789) // int64 | The unique ID assigned to the file in Shortcut.
    updateFile := *openapiclient.NewUpdateFile() // UpdateFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateFile(context.Background(), filePublicId).UpdateFile(updateFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFile`: UploadedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filePublicId** | **int64** | The unique ID assigned to the file in Shortcut. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFile** | [**UpdateFile**](UpdateFile.md) |  | 

### Return type

[**UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> Group UpdateGroup(ctx, groupPublicId).UpdateGroup(updateGroup).Execute()

Update Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    groupPublicId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the Group.
    updateGroup := *openapiclient.NewUpdateGroup() // UpdateGroup | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateGroup(context.Background(), groupPublicId).UpdateGroup(updateGroup).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupPublicId** | **string** | The unique ID of the Group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGroup** | [**UpdateGroup**](UpdateGroup.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIteration

> Iteration UpdateIteration(ctx, iterationPublicId).UpdateIteration(updateIteration).Execute()

Update Iteration

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    iterationPublicId := int64(789) // int64 | The unique ID of the Iteration.
    updateIteration := *openapiclient.NewUpdateIteration() // UpdateIteration | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateIteration(context.Background(), iterationPublicId).UpdateIteration(updateIteration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateIteration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIteration`: Iteration
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateIteration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iterationPublicId** | **int64** | The unique ID of the Iteration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIterationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIteration** | [**UpdateIteration**](UpdateIteration.md) |  | 

### Return type

[**Iteration**](Iteration.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> Label UpdateLabel(ctx, labelPublicId).UpdateLabel(updateLabel).Execute()

Update Label



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    labelPublicId := int64(789) // int64 | The unique ID of the Label you wish to update.
    updateLabel := *openapiclient.NewUpdateLabel() // UpdateLabel | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateLabel(context.Background(), labelPublicId).UpdateLabel(updateLabel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLabel`: Label
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelPublicId** | **int64** | The unique ID of the Label you wish to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLabel** | [**UpdateLabel**](UpdateLabel.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLinkedFile

> LinkedFile UpdateLinkedFile(ctx, linkedFilePublicId).UpdateLinkedFile(updateLinkedFile).Execute()

Update Linked File



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    linkedFilePublicId := int64(789) // int64 | The unique identifier of the linked file.
    updateLinkedFile := *openapiclient.NewUpdateLinkedFile() // UpdateLinkedFile | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateLinkedFile(context.Background(), linkedFilePublicId).UpdateLinkedFile(updateLinkedFile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateLinkedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLinkedFile`: LinkedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateLinkedFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkedFilePublicId** | **int64** | The unique identifier of the linked file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLinkedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLinkedFile** | [**UpdateLinkedFile**](UpdateLinkedFile.md) |  | 

### Return type

[**LinkedFile**](LinkedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMilestone

> Milestone UpdateMilestone(ctx, milestonePublicId).UpdateMilestone(updateMilestone).Execute()

Update Milestone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    milestonePublicId := int64(789) // int64 | The ID of the Milestone.
    updateMilestone := *openapiclient.NewUpdateMilestone() // UpdateMilestone | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMilestone(context.Background(), milestonePublicId).UpdateMilestone(updateMilestone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMilestone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMilestone`: Milestone
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMilestone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**milestonePublicId** | **int64** | The ID of the Milestone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMilestoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMilestone** | [**UpdateMilestone**](UpdateMilestone.md) |  | 

### Return type

[**Milestone**](Milestone.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMultipleStories

> []StorySlim UpdateMultipleStories(ctx).UpdateStories(updateStories).Execute()

Update Multiple Stories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    updateStories := *openapiclient.NewUpdateStories([]int64{int64(123)}) // UpdateStories | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateMultipleStories(context.Background()).UpdateStories(updateStories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateMultipleStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMultipleStories`: []StorySlim
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateMultipleStories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMultipleStoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStories** | [**UpdateStories**](UpdateStories.md) |  | 

### Return type

[**[]StorySlim**](StorySlim.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProject

> Project UpdateProject(ctx, projectPublicId).UpdateProject(updateProject).Execute()

Update Project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    projectPublicId := int64(789) // int64 | The unique ID of the Project.
    updateProject := *openapiclient.NewUpdateProject() // UpdateProject | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateProject(context.Background(), projectPublicId).UpdateProject(updateProject).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectPublicId** | **int64** | The unique ID of the Project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateProject** | [**UpdateProject**](UpdateProject.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStory

> Story UpdateStory(ctx, storyPublicId).UpdateStory(updateStory).Execute()

Update Story



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The unique identifier of this story.
    updateStory := *openapiclient.NewUpdateStory() // UpdateStory | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateStory(context.Background(), storyPublicId).UpdateStory(updateStory).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStory`: Story
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The unique identifier of this story. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStory** | [**UpdateStory**](UpdateStory.md) |  | 

### Return type

[**Story**](Story.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStoryComment

> StoryComment UpdateStoryComment(ctx, storyPublicId, commentPublicId).UpdateStoryComment(updateStoryComment).Execute()

Update Story Comment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The ID of the Story that the Comment is in.
    commentPublicId := int64(789) // int64 | The ID of the Comment
    updateStoryComment := *openapiclient.NewUpdateStoryComment("Text_example") // UpdateStoryComment | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateStoryComment(context.Background(), storyPublicId, commentPublicId).UpdateStoryComment(updateStoryComment).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStoryComment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStoryComment`: StoryComment
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateStoryComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The ID of the Story that the Comment is in. | 
**commentPublicId** | **int64** | The ID of the Comment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoryCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateStoryComment** | [**UpdateStoryComment**](UpdateStoryComment.md) |  | 

### Return type

[**StoryComment**](StoryComment.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStoryLink

> StoryLink UpdateStoryLink(ctx, storyLinkPublicId).UpdateStoryLink(updateStoryLink).Execute()

Update Story Link



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyLinkPublicId := int64(789) // int64 | The unique ID of the Story Link.
    updateStoryLink := *openapiclient.NewUpdateStoryLink() // UpdateStoryLink | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateStoryLink(context.Background(), storyLinkPublicId).UpdateStoryLink(updateStoryLink).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateStoryLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStoryLink`: StoryLink
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateStoryLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyLinkPublicId** | **int64** | The unique ID of the Story Link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStoryLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateStoryLink** | [**UpdateStoryLink**](UpdateStoryLink.md) |  | 

### Return type

[**StoryLink**](StoryLink.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTask

> Task UpdateTask(ctx, storyPublicId, taskPublicId).UpdateTask(updateTask).Execute()

Update Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    storyPublicId := int64(789) // int64 | The unique identifier of the parent Story.
    taskPublicId := int64(789) // int64 | The unique identifier of the Task you wish to update.
    updateTask := *openapiclient.NewUpdateTask() // UpdateTask | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateTask(context.Background(), storyPublicId, taskPublicId).UpdateTask(updateTask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**storyPublicId** | **int64** | The unique identifier of the parent Story. | 
**taskPublicId** | **int64** | The unique identifier of the Task you wish to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTask** | [**UpdateTask**](UpdateTask.md) |  | 

### Return type

[**Task**](Task.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFiles

> []UploadedFile UploadFiles(ctx).File0(file0).StoryId(storyId).File1(file1).File2(file2).File3(file3).Execute()

Upload Files



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file0 := os.NewFile(1234, "some_file") // *os.File | A file upload. At least one is required.
    storyId := int64(789) // int64 | The story ID that these files will be associated with. (optional)
    file1 := os.NewFile(1234, "some_file") // *os.File | Optional additional files. (optional)
    file2 := os.NewFile(1234, "some_file") // *os.File | Optional additional files. (optional)
    file3 := os.NewFile(1234, "some_file") // *os.File | Optional additional files. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UploadFiles(context.Background()).File0(file0).StoryId(storyId).File1(file1).File2(file2).File3(file3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UploadFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFiles`: []UploadedFile
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UploadFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file0** | ***os.File** | A file upload. At least one is required. | 
 **storyId** | **int64** | The story ID that these files will be associated with. | 
 **file1** | ***os.File** | Optional additional files. | 
 **file2** | ***os.File** | Optional additional files. | 
 **file3** | ***os.File** | Optional additional files. | 

### Return type

[**[]UploadedFile**](UploadedFile.md)

### Authorization

[api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

