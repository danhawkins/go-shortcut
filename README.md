# Go API client for openapi

Shortcut API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.app.shortcut.com*

| Class        | Method                                                                          | HTTP request                                                                        | Description                   |
| ------------ | ------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------- |
| *DefaultApi* | [**CreateCategory**](docs/DefaultApi.md#createcategory)                         | **Post** /api/v3/categories                                                         | Create Category               |
| *DefaultApi* | [**CreateEntityTemplate**](docs/DefaultApi.md#createentitytemplate)             | **Post** /api/v3/entity-templates                                                   | Create Entity Template        |
| *DefaultApi* | [**CreateEpic**](docs/DefaultApi.md#createepic)                                 | **Post** /api/v3/epics                                                              | Create Epic                   |
| *DefaultApi* | [**CreateEpicComment**](docs/DefaultApi.md#createepiccomment)                   | **Post** /api/v3/epics/{epic-public-id}/comments                                    | Create Epic Comment           |
| *DefaultApi* | [**CreateEpicCommentComment**](docs/DefaultApi.md#createepiccommentcomment)     | **Post** /api/v3/epics/{epic-public-id}/comments/{comment-public-id}                | Create Epic Comment Comment   |
| *DefaultApi* | [**CreateGroup**](docs/DefaultApi.md#creategroup)                               | **Post** /api/v3/groups                                                             | Create Group                  |
| *DefaultApi* | [**CreateIteration**](docs/DefaultApi.md#createiteration)                       | **Post** /api/v3/iterations                                                         | Create Iteration              |
| *DefaultApi* | [**CreateLabel**](docs/DefaultApi.md#createlabel)                               | **Post** /api/v3/labels                                                             | Create Label                  |
| *DefaultApi* | [**CreateLinkedFile**](docs/DefaultApi.md#createlinkedfile)                     | **Post** /api/v3/linked-files                                                       | Create Linked File            |
| *DefaultApi* | [**CreateMilestone**](docs/DefaultApi.md#createmilestone)                       | **Post** /api/v3/milestones                                                         | Create Milestone              |
| *DefaultApi* | [**CreateMultipleStories**](docs/DefaultApi.md#createmultiplestories)           | **Post** /api/v3/stories/bulk                                                       | Create Multiple Stories       |
| *DefaultApi* | [**CreateProject**](docs/DefaultApi.md#createproject)                           | **Post** /api/v3/projects                                                           | Create Project                |
| *DefaultApi* | [**CreateStory**](docs/DefaultApi.md#createstory)                               | **Post** /api/v3/stories                                                            | Create Story                  |
| *DefaultApi* | [**CreateStoryComment**](docs/DefaultApi.md#createstorycomment)                 | **Post** /api/v3/stories/{story-public-id}/comments                                 | Create Story Comment          |
| *DefaultApi* | [**CreateStoryLink**](docs/DefaultApi.md#createstorylink)                       | **Post** /api/v3/story-links                                                        | Create Story Link             |
| *DefaultApi* | [**CreateStoryReaction**](docs/DefaultApi.md#createstoryreaction)               | **Post** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions   | Create Story Reaction         |
| *DefaultApi* | [**CreateTask**](docs/DefaultApi.md#createtask)                                 | **Post** /api/v3/stories/{story-public-id}/tasks                                    | Create Task                   |
| *DefaultApi* | [**DeleteCategory**](docs/DefaultApi.md#deletecategory)                         | **Delete** /api/v3/categories/{category-public-id}                                  | Delete Category               |
| *DefaultApi* | [**DeleteCustomField**](docs/DefaultApi.md#deletecustomfield)                   | **Delete** /api/v3/custom-fields/{custom-field-public-id}                           | Delete Custom Field           |
| *DefaultApi* | [**DeleteEntityTemplate**](docs/DefaultApi.md#deleteentitytemplate)             | **Delete** /api/v3/entity-templates/{entity-template-public-id}                     | Delete Entity Template        |
| *DefaultApi* | [**DeleteEpic**](docs/DefaultApi.md#deleteepic)                                 | **Delete** /api/v3/epics/{epic-public-id}                                           | Delete Epic                   |
| *DefaultApi* | [**DeleteEpicComment**](docs/DefaultApi.md#deleteepiccomment)                   | **Delete** /api/v3/epics/{epic-public-id}/comments/{comment-public-id}              | Delete Epic Comment           |
| *DefaultApi* | [**DeleteFile**](docs/DefaultApi.md#deletefile)                                 | **Delete** /api/v3/files/{file-public-id}                                           | Delete File                   |
| *DefaultApi* | [**DeleteIteration**](docs/DefaultApi.md#deleteiteration)                       | **Delete** /api/v3/iterations/{iteration-public-id}                                 | Delete Iteration              |
| *DefaultApi* | [**DeleteLabel**](docs/DefaultApi.md#deletelabel)                               | **Delete** /api/v3/labels/{label-public-id}                                         | Delete Label                  |
| *DefaultApi* | [**DeleteLinkedFile**](docs/DefaultApi.md#deletelinkedfile)                     | **Delete** /api/v3/linked-files/{linked-file-public-id}                             | Delete Linked File            |
| *DefaultApi* | [**DeleteMilestone**](docs/DefaultApi.md#deletemilestone)                       | **Delete** /api/v3/milestones/{milestone-public-id}                                 | Delete Milestone              |
| *DefaultApi* | [**DeleteMultipleStories**](docs/DefaultApi.md#deletemultiplestories)           | **Delete** /api/v3/stories/bulk                                                     | Delete Multiple Stories       |
| *DefaultApi* | [**DeleteProject**](docs/DefaultApi.md#deleteproject)                           | **Delete** /api/v3/projects/{project-public-id}                                     | Delete Project                |
| *DefaultApi* | [**DeleteStory**](docs/DefaultApi.md#deletestory)                               | **Delete** /api/v3/stories/{story-public-id}                                        | Delete Story                  |
| *DefaultApi* | [**DeleteStoryComment**](docs/DefaultApi.md#deletestorycomment)                 | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id}           | Delete Story Comment          |
| *DefaultApi* | [**DeleteStoryLink**](docs/DefaultApi.md#deletestorylink)                       | **Delete** /api/v3/story-links/{story-link-public-id}                               | Delete Story Link             |
| *DefaultApi* | [**DeleteStoryReaction**](docs/DefaultApi.md#deletestoryreaction)               | **Delete** /api/v3/stories/{story-public-id}/comments/{comment-public-id}/reactions | Delete Story Reaction         |
| *DefaultApi* | [**DeleteTask**](docs/DefaultApi.md#deletetask)                                 | **Delete** /api/v3/stories/{story-public-id}/tasks/{task-public-id}                 | Delete Task                   |
| *DefaultApi* | [**DisableGroups**](docs/DefaultApi.md#disablegroups)                           | **Put** /api/v3/groups/disable                                                      | Disable Groups                |
| *DefaultApi* | [**DisableIterations**](docs/DefaultApi.md#disableiterations)                   | **Put** /api/v3/iterations/disable                                                  | Disable Iterations            |
| *DefaultApi* | [**DisableStoryTemplates**](docs/DefaultApi.md#disablestorytemplates)           | **Put** /api/v3/entity-templates/disable                                            | Disable Story Templates       |
| *DefaultApi* | [**EnableGroups**](docs/DefaultApi.md#enablegroups)                             | **Put** /api/v3/groups/enable                                                       | Enable Groups                 |
| *DefaultApi* | [**EnableIterations**](docs/DefaultApi.md#enableiterations)                     | **Put** /api/v3/iterations/enable                                                   | Enable Iterations             |
| *DefaultApi* | [**EnableStoryTemplates**](docs/DefaultApi.md#enablestorytemplates)             | **Put** /api/v3/entity-templates/enable                                             | Enable Story Templates        |
| *DefaultApi* | [**GetCategory**](docs/DefaultApi.md#getcategory)                               | **Get** /api/v3/categories/{category-public-id}                                     | Get Category                  |
| *DefaultApi* | [**GetCurrentMemberInfo**](docs/DefaultApi.md#getcurrentmemberinfo)             | **Get** /api/v3/member                                                              | Get Current Member Info       |
| *DefaultApi* | [**GetCustomField**](docs/DefaultApi.md#getcustomfield)                         | **Get** /api/v3/custom-fields/{custom-field-public-id}                              | Get Custom Field              |
| *DefaultApi* | [**GetEntityTemplate**](docs/DefaultApi.md#getentitytemplate)                   | **Get** /api/v3/entity-templates/{entity-template-public-id}                        | Get Entity Template           |
| *DefaultApi* | [**GetEpic**](docs/DefaultApi.md#getepic)                                       | **Get** /api/v3/epics/{epic-public-id}                                              | Get Epic                      |
| *DefaultApi* | [**GetEpicComment**](docs/DefaultApi.md#getepiccomment)                         | **Get** /api/v3/epics/{epic-public-id}/comments/{comment-public-id}                 | Get Epic Comment              |
| *DefaultApi* | [**GetEpicWorkflow**](docs/DefaultApi.md#getepicworkflow)                       | **Get** /api/v3/epic-workflow                                                       | Get Epic Workflow             |
| *DefaultApi* | [**GetExternalLinkStories**](docs/DefaultApi.md#getexternallinkstories)         | **Get** /api/v3/external-link/stories                                               | Get External Link Stories     |
| *DefaultApi* | [**GetFile**](docs/DefaultApi.md#getfile)                                       | **Get** /api/v3/files/{file-public-id}                                              | Get File                      |
| *DefaultApi* | [**GetGroup**](docs/DefaultApi.md#getgroup)                                     | **Get** /api/v3/groups/{group-public-id}                                            | Get Group                     |
| *DefaultApi* | [**GetIteration**](docs/DefaultApi.md#getiteration)                             | **Get** /api/v3/iterations/{iteration-public-id}                                    | Get Iteration                 |
| *DefaultApi* | [**GetLabel**](docs/DefaultApi.md#getlabel)                                     | **Get** /api/v3/labels/{label-public-id}                                            | Get Label                     |
| *DefaultApi* | [**GetLinkedFile**](docs/DefaultApi.md#getlinkedfile)                           | **Get** /api/v3/linked-files/{linked-file-public-id}                                | Get Linked File               |
| *DefaultApi* | [**GetMember**](docs/DefaultApi.md#getmember)                                   | **Get** /api/v3/members/{member-public-id}                                          | Get Member                    |
| *DefaultApi* | [**GetMilestone**](docs/DefaultApi.md#getmilestone)                             | **Get** /api/v3/milestones/{milestone-public-id}                                    | Get Milestone                 |
| *DefaultApi* | [**GetProject**](docs/DefaultApi.md#getproject)                                 | **Get** /api/v3/projects/{project-public-id}                                        | Get Project                   |
| *DefaultApi* | [**GetRepository**](docs/DefaultApi.md#getrepository)                           | **Get** /api/v3/repositories/{repo-public-id}                                       | Get Repository                |
| *DefaultApi* | [**GetStory**](docs/DefaultApi.md#getstory)                                     | **Get** /api/v3/stories/{story-public-id}                                           | Get Story                     |
| *DefaultApi* | [**GetStoryComment**](docs/DefaultApi.md#getstorycomment)                       | **Get** /api/v3/stories/{story-public-id}/comments/{comment-public-id}              | Get Story Comment             |
| *DefaultApi* | [**GetStoryLink**](docs/DefaultApi.md#getstorylink)                             | **Get** /api/v3/story-links/{story-link-public-id}                                  | Get Story Link                |
| *DefaultApi* | [**GetTask**](docs/DefaultApi.md#gettask)                                       | **Get** /api/v3/stories/{story-public-id}/tasks/{task-public-id}                    | Get Task                      |
| *DefaultApi* | [**GetWorkflow**](docs/DefaultApi.md#getworkflow)                               | **Get** /api/v3/workflows/{workflow-public-id}                                      | Get Workflow                  |
| *DefaultApi* | [**ListCategories**](docs/DefaultApi.md#listcategories)                         | **Get** /api/v3/categories                                                          | List Categories               |
| *DefaultApi* | [**ListCategoryMilestones**](docs/DefaultApi.md#listcategorymilestones)         | **Get** /api/v3/categories/{category-public-id}/milestones                          | List Category Milestones      |
| *DefaultApi* | [**ListCustomFields**](docs/DefaultApi.md#listcustomfields)                     | **Get** /api/v3/custom-fields                                                       | List Custom Fields            |
| *DefaultApi* | [**ListEntityTemplates**](docs/DefaultApi.md#listentitytemplates)               | **Get** /api/v3/entity-templates                                                    | List Entity Templates         |
| *DefaultApi* | [**ListEpicComments**](docs/DefaultApi.md#listepiccomments)                     | **Get** /api/v3/epics/{epic-public-id}/comments                                     | List Epic Comments            |
| *DefaultApi* | [**ListEpicStories**](docs/DefaultApi.md#listepicstories)                       | **Get** /api/v3/epics/{epic-public-id}/stories                                      | List Epic Stories             |
| *DefaultApi* | [**ListEpics**](docs/DefaultApi.md#listepics)                                   | **Get** /api/v3/epics                                                               | List Epics                    |
| *DefaultApi* | [**ListFiles**](docs/DefaultApi.md#listfiles)                                   | **Get** /api/v3/files                                                               | List Files                    |
| *DefaultApi* | [**ListGroupStories**](docs/DefaultApi.md#listgroupstories)                     | **Get** /api/v3/groups/{group-public-id}/stories                                    | List Group Stories            |
| *DefaultApi* | [**ListGroups**](docs/DefaultApi.md#listgroups)                                 | **Get** /api/v3/groups                                                              | List Groups                   |
| *DefaultApi* | [**ListIterationStories**](docs/DefaultApi.md#listiterationstories)             | **Get** /api/v3/iterations/{iteration-public-id}/stories                            | List Iteration Stories        |
| *DefaultApi* | [**ListIterations**](docs/DefaultApi.md#listiterations)                         | **Get** /api/v3/iterations                                                          | List Iterations               |
| *DefaultApi* | [**ListLabelEpics**](docs/DefaultApi.md#listlabelepics)                         | **Get** /api/v3/labels/{label-public-id}/epics                                      | List Label Epics              |
| *DefaultApi* | [**ListLabelStories**](docs/DefaultApi.md#listlabelstories)                     | **Get** /api/v3/labels/{label-public-id}/stories                                    | List Label Stories            |
| *DefaultApi* | [**ListLabels**](docs/DefaultApi.md#listlabels)                                 | **Get** /api/v3/labels                                                              | List Labels                   |
| *DefaultApi* | [**ListLinkedFiles**](docs/DefaultApi.md#listlinkedfiles)                       | **Get** /api/v3/linked-files                                                        | List Linked Files             |
| *DefaultApi* | [**ListMembers**](docs/DefaultApi.md#listmembers)                               | **Get** /api/v3/members                                                             | List Members                  |
| *DefaultApi* | [**ListMilestoneEpics**](docs/DefaultApi.md#listmilestoneepics)                 | **Get** /api/v3/milestones/{milestone-public-id}/epics                              | List Milestone Epics          |
| *DefaultApi* | [**ListMilestones**](docs/DefaultApi.md#listmilestones)                         | **Get** /api/v3/milestones                                                          | List Milestones               |
| *DefaultApi* | [**ListProjects**](docs/DefaultApi.md#listprojects)                             | **Get** /api/v3/projects                                                            | List Projects                 |
| *DefaultApi* | [**ListRepositories**](docs/DefaultApi.md#listrepositories)                     | **Get** /api/v3/repositories                                                        | List Repositories             |
| *DefaultApi* | [**ListStories**](docs/DefaultApi.md#liststories)                               | **Get** /api/v3/projects/{project-public-id}/stories                                | List Stories                  |
| *DefaultApi* | [**ListStoryComment**](docs/DefaultApi.md#liststorycomment)                     | **Get** /api/v3/stories/{story-public-id}/comments                                  | List Story Comment            |
| *DefaultApi* | [**ListWorkflows**](docs/DefaultApi.md#listworkflows)                           | **Get** /api/v3/workflows                                                           | List Workflows                |
| *DefaultApi* | [**Search**](docs/DefaultApi.md#search)                                         | **Get** /api/v3/search                                                              | Search                        |
| *DefaultApi* | [**SearchEpics**](docs/DefaultApi.md#searchepics)                               | **Get** /api/v3/search/epics                                                        | Search Epics                  |
| *DefaultApi* | [**SearchIterations**](docs/DefaultApi.md#searchiterations)                     | **Get** /api/v3/search/iterations                                                   | Search Iterations             |
| *DefaultApi* | [**SearchMilestones**](docs/DefaultApi.md#searchmilestones)                     | **Get** /api/v3/search/milestones                                                   | Search Milestones             |
| *DefaultApi* | [**SearchStories**](docs/DefaultApi.md#searchstories)                           | **Get** /api/v3/search/stories                                                      | Search Stories                |
| *DefaultApi* | [**SearchStoriesOld**](docs/DefaultApi.md#searchstoriesold)                     | **Post** /api/v3/stories/search                                                     | Search Stories (Old)          |
| *DefaultApi* | [**StoryHistory**](docs/DefaultApi.md#storyhistory)                             | **Get** /api/v3/stories/{story-public-id}/history                                   | Story History                 |
| *DefaultApi* | [**UnlinkProductboardFromEpic**](docs/DefaultApi.md#unlinkproductboardfromepic) | **Post** /api/v3/epics/{epic-public-id}/unlink-productboard                         | Unlink Productboard from Epic |
| *DefaultApi* | [**UpdateCategory**](docs/DefaultApi.md#updatecategory)                         | **Put** /api/v3/categories/{category-public-id}                                     | Update Category               |
| *DefaultApi* | [**UpdateCustomField**](docs/DefaultApi.md#updatecustomfield)                   | **Put** /api/v3/custom-fields/{custom-field-public-id}                              | Update Custom Field           |
| *DefaultApi* | [**UpdateEntityTemplate**](docs/DefaultApi.md#updateentitytemplate)             | **Put** /api/v3/entity-templates/{entity-template-public-id}                        | Update Entity Template        |
| *DefaultApi* | [**UpdateEpic**](docs/DefaultApi.md#updateepic)                                 | **Put** /api/v3/epics/{epic-public-id}                                              | Update Epic                   |
| *DefaultApi* | [**UpdateEpicComment**](docs/DefaultApi.md#updateepiccomment)                   | **Put** /api/v3/epics/{epic-public-id}/comments/{comment-public-id}                 | Update Epic Comment           |
| *DefaultApi* | [**UpdateFile**](docs/DefaultApi.md#updatefile)                                 | **Put** /api/v3/files/{file-public-id}                                              | Update File                   |
| *DefaultApi* | [**UpdateGroup**](docs/DefaultApi.md#updategroup)                               | **Put** /api/v3/groups/{group-public-id}                                            | Update Group                  |
| *DefaultApi* | [**UpdateIteration**](docs/DefaultApi.md#updateiteration)                       | **Put** /api/v3/iterations/{iteration-public-id}                                    | Update Iteration              |
| *DefaultApi* | [**UpdateLabel**](docs/DefaultApi.md#updatelabel)                               | **Put** /api/v3/labels/{label-public-id}                                            | Update Label                  |
| *DefaultApi* | [**UpdateLinkedFile**](docs/DefaultApi.md#updatelinkedfile)                     | **Put** /api/v3/linked-files/{linked-file-public-id}                                | Update Linked File            |
| *DefaultApi* | [**UpdateMilestone**](docs/DefaultApi.md#updatemilestone)                       | **Put** /api/v3/milestones/{milestone-public-id}                                    | Update Milestone              |
| *DefaultApi* | [**UpdateMultipleStories**](docs/DefaultApi.md#updatemultiplestories)           | **Put** /api/v3/stories/bulk                                                        | Update Multiple Stories       |
| *DefaultApi* | [**UpdateProject**](docs/DefaultApi.md#updateproject)                           | **Put** /api/v3/projects/{project-public-id}                                        | Update Project                |
| *DefaultApi* | [**UpdateStory**](docs/DefaultApi.md#updatestory)                               | **Put** /api/v3/stories/{story-public-id}                                           | Update Story                  |
| *DefaultApi* | [**UpdateStoryComment**](docs/DefaultApi.md#updatestorycomment)                 | **Put** /api/v3/stories/{story-public-id}/comments/{comment-public-id}              | Update Story Comment          |
| *DefaultApi* | [**UpdateStoryLink**](docs/DefaultApi.md#updatestorylink)                       | **Put** /api/v3/story-links/{story-link-public-id}                                  | Update Story Link             |
| *DefaultApi* | [**UpdateTask**](docs/DefaultApi.md#updatetask)                                 | **Put** /api/v3/stories/{story-public-id}/tasks/{task-public-id}                    | Update Task                   |
| *DefaultApi* | [**UploadFiles**](docs/DefaultApi.md#uploadfiles)                               | **Post** /api/v3/files                                                              | Upload Files                  |


## Documentation For Models

 - [BasicWorkspaceInfo](docs/BasicWorkspaceInfo.md)
 - [Branch](docs/Branch.md)
 - [Category](docs/Category.md)
 - [Commit](docs/Commit.md)
 - [CreateCategory](docs/CreateCategory.md)
 - [CreateCategoryParams](docs/CreateCategoryParams.md)
 - [CreateCommentComment](docs/CreateCommentComment.md)
 - [CreateEntityTemplate](docs/CreateEntityTemplate.md)
 - [CreateEpic](docs/CreateEpic.md)
 - [CreateEpicComment](docs/CreateEpicComment.md)
 - [CreateGroup](docs/CreateGroup.md)
 - [CreateIteration](docs/CreateIteration.md)
 - [CreateLabelParams](docs/CreateLabelParams.md)
 - [CreateLinkedFile](docs/CreateLinkedFile.md)
 - [CreateMilestone](docs/CreateMilestone.md)
 - [CreateOrDeleteStoryReaction](docs/CreateOrDeleteStoryReaction.md)
 - [CreateProject](docs/CreateProject.md)
 - [CreateStories](docs/CreateStories.md)
 - [CreateStoryComment](docs/CreateStoryComment.md)
 - [CreateStoryCommentParams](docs/CreateStoryCommentParams.md)
 - [CreateStoryContents](docs/CreateStoryContents.md)
 - [CreateStoryLink](docs/CreateStoryLink.md)
 - [CreateStoryLinkParams](docs/CreateStoryLinkParams.md)
 - [CreateStoryParams](docs/CreateStoryParams.md)
 - [CreateTask](docs/CreateTask.md)
 - [CreateTaskParams](docs/CreateTaskParams.md)
 - [CustomField](docs/CustomField.md)
 - [CustomFieldEnumValue](docs/CustomFieldEnumValue.md)
 - [CustomFieldValueParams](docs/CustomFieldValueParams.md)
 - [DataConflictError](docs/DataConflictError.md)
 - [DeleteStories](docs/DeleteStories.md)
 - [EntityTemplate](docs/EntityTemplate.md)
 - [EntityTemplateTask](docs/EntityTemplateTask.md)
 - [Epic](docs/Epic.md)
 - [EpicAssociatedGroup](docs/EpicAssociatedGroup.md)
 - [EpicSearchResult](docs/EpicSearchResult.md)
 - [EpicSearchResults](docs/EpicSearchResults.md)
 - [EpicSlim](docs/EpicSlim.md)
 - [EpicState](docs/EpicState.md)
 - [EpicStats](docs/EpicStats.md)
 - [EpicWorkflow](docs/EpicWorkflow.md)
 - [GetEpicStories](docs/GetEpicStories.md)
 - [GetExternalLinkStoriesParams](docs/GetExternalLinkStoriesParams.md)
 - [GetIterationStories](docs/GetIterationStories.md)
 - [GetLabelStories](docs/GetLabelStories.md)
 - [GetMember](docs/GetMember.md)
 - [GetProjectStories](docs/GetProjectStories.md)
 - [Group](docs/Group.md)
 - [History](docs/History.md)
 - [HistoryActionBranchCreate](docs/HistoryActionBranchCreate.md)
 - [HistoryActionBranchMerge](docs/HistoryActionBranchMerge.md)
 - [HistoryActionBranchPush](docs/HistoryActionBranchPush.md)
 - [HistoryActionLabelCreate](docs/HistoryActionLabelCreate.md)
 - [HistoryActionLabelDelete](docs/HistoryActionLabelDelete.md)
 - [HistoryActionLabelUpdate](docs/HistoryActionLabelUpdate.md)
 - [HistoryActionProjectUpdate](docs/HistoryActionProjectUpdate.md)
 - [HistoryActionPullRequest](docs/HistoryActionPullRequest.md)
 - [HistoryActionStoryCommentCreate](docs/HistoryActionStoryCommentCreate.md)
 - [HistoryActionStoryCreate](docs/HistoryActionStoryCreate.md)
 - [HistoryActionStoryDelete](docs/HistoryActionStoryDelete.md)
 - [HistoryActionStoryLinkCreate](docs/HistoryActionStoryLinkCreate.md)
 - [HistoryActionStoryLinkDelete](docs/HistoryActionStoryLinkDelete.md)
 - [HistoryActionStoryLinkUpdate](docs/HistoryActionStoryLinkUpdate.md)
 - [HistoryActionStoryUpdate](docs/HistoryActionStoryUpdate.md)
 - [HistoryActionTaskCreate](docs/HistoryActionTaskCreate.md)
 - [HistoryActionTaskDelete](docs/HistoryActionTaskDelete.md)
 - [HistoryActionTaskUpdate](docs/HistoryActionTaskUpdate.md)
 - [HistoryActionWorkspace2BulkUpdate](docs/HistoryActionWorkspace2BulkUpdate.md)
 - [HistoryChangesStory](docs/HistoryChangesStory.md)
 - [HistoryChangesStoryLink](docs/HistoryChangesStoryLink.md)
 - [HistoryChangesTask](docs/HistoryChangesTask.md)
 - [HistoryReferenceBranch](docs/HistoryReferenceBranch.md)
 - [HistoryReferenceCommit](docs/HistoryReferenceCommit.md)
 - [HistoryReferenceCustomFieldEnumValue](docs/HistoryReferenceCustomFieldEnumValue.md)
 - [HistoryReferenceEpic](docs/HistoryReferenceEpic.md)
 - [HistoryReferenceGeneral](docs/HistoryReferenceGeneral.md)
 - [HistoryReferenceGroup](docs/HistoryReferenceGroup.md)
 - [HistoryReferenceIteration](docs/HistoryReferenceIteration.md)
 - [HistoryReferenceLabel](docs/HistoryReferenceLabel.md)
 - [HistoryReferenceProject](docs/HistoryReferenceProject.md)
 - [HistoryReferenceStory](docs/HistoryReferenceStory.md)
 - [HistoryReferenceStoryTask](docs/HistoryReferenceStoryTask.md)
 - [HistoryReferenceWorkflowState](docs/HistoryReferenceWorkflowState.md)
 - [Icon](docs/Icon.md)
 - [Identity](docs/Identity.md)
 - [Iteration](docs/Iteration.md)
 - [IterationAssociatedGroup](docs/IterationAssociatedGroup.md)
 - [IterationSearchResults](docs/IterationSearchResults.md)
 - [IterationSlim](docs/IterationSlim.md)
 - [IterationStats](docs/IterationStats.md)
 - [Label](docs/Label.md)
 - [LabelSlim](docs/LabelSlim.md)
 - [LabelStats](docs/LabelStats.md)
 - [LinkedFile](docs/LinkedFile.md)
 - [ListEpics](docs/ListEpics.md)
 - [ListGroupStories](docs/ListGroupStories.md)
 - [ListLabels](docs/ListLabels.md)
 - [ListMembers](docs/ListMembers.md)
 - [MaxSearchResultsExceededError](docs/MaxSearchResultsExceededError.md)
 - [Member](docs/Member.md)
 - [MemberInfo](docs/MemberInfo.md)
 - [Milestone](docs/Milestone.md)
 - [MilestoneSearchResult](docs/MilestoneSearchResult.md)
 - [MilestoneSearchResults](docs/MilestoneSearchResults.md)
 - [MilestoneStats](docs/MilestoneStats.md)
 - [Profile](docs/Profile.md)
 - [Project](docs/Project.md)
 - [ProjectStats](docs/ProjectStats.md)
 - [PullRequest](docs/PullRequest.md)
 - [PullRequestLabel](docs/PullRequestLabel.md)
 - [Repository](docs/Repository.md)
 - [Search](docs/Search.md)
 - [SearchResults](docs/SearchResults.md)
 - [SearchStories](docs/SearchStories.md)
 - [Story](docs/Story.md)
 - [StoryComment](docs/StoryComment.md)
 - [StoryContents](docs/StoryContents.md)
 - [StoryContentsTask](docs/StoryContentsTask.md)
 - [StoryCustomField](docs/StoryCustomField.md)
 - [StoryHistoryChangeAddsRemovesInt](docs/StoryHistoryChangeAddsRemovesInt.md)
 - [StoryHistoryChangeAddsRemovesUuid](docs/StoryHistoryChangeAddsRemovesUuid.md)
 - [StoryHistoryChangeOldNewBool](docs/StoryHistoryChangeOldNewBool.md)
 - [StoryHistoryChangeOldNewInt](docs/StoryHistoryChangeOldNewInt.md)
 - [StoryHistoryChangeOldNewStr](docs/StoryHistoryChangeOldNewStr.md)
 - [StoryHistoryChangeOldNewUuid](docs/StoryHistoryChangeOldNewUuid.md)
 - [StoryLink](docs/StoryLink.md)
 - [StoryReaction](docs/StoryReaction.md)
 - [StorySearchResult](docs/StorySearchResult.md)
 - [StorySearchResults](docs/StorySearchResults.md)
 - [StorySlim](docs/StorySlim.md)
 - [StoryStats](docs/StoryStats.md)
 - [SyncedItem](docs/SyncedItem.md)
 - [Task](docs/Task.md)
 - [ThreadedComment](docs/ThreadedComment.md)
 - [TypedStoryLink](docs/TypedStoryLink.md)
 - [UnusableEntitlementError](docs/UnusableEntitlementError.md)
 - [UpdateCategory](docs/UpdateCategory.md)
 - [UpdateComment](docs/UpdateComment.md)
 - [UpdateCustomField](docs/UpdateCustomField.md)
 - [UpdateCustomFieldEnumValue](docs/UpdateCustomFieldEnumValue.md)
 - [UpdateEntityTemplate](docs/UpdateEntityTemplate.md)
 - [UpdateEpic](docs/UpdateEpic.md)
 - [UpdateFile](docs/UpdateFile.md)
 - [UpdateGroup](docs/UpdateGroup.md)
 - [UpdateIteration](docs/UpdateIteration.md)
 - [UpdateLabel](docs/UpdateLabel.md)
 - [UpdateLinkedFile](docs/UpdateLinkedFile.md)
 - [UpdateMilestone](docs/UpdateMilestone.md)
 - [UpdateProject](docs/UpdateProject.md)
 - [UpdateStories](docs/UpdateStories.md)
 - [UpdateStory](docs/UpdateStory.md)
 - [UpdateStoryComment](docs/UpdateStoryComment.md)
 - [UpdateStoryContents](docs/UpdateStoryContents.md)
 - [UpdateStoryLink](docs/UpdateStoryLink.md)
 - [UpdateTask](docs/UpdateTask.md)
 - [UploadedFile](docs/UploadedFile.md)
 - [Workflow](docs/Workflow.md)
 - [WorkflowState](docs/WorkflowState.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### api_token

- **Type**: API key
- **API key parameter name**: Shortcut-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Shortcut-Token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author