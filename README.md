# Codefresh Jira Issue Comment

This step will create a comment with build information on your Jira instance with the specified issue key. The Jira issue key can be set directly or by using a regex pattern against a source field. The inserted Jira comment id

## Prerequisites

- Please create an [API token](https://confluence.atlassian.com/cloud/api-tokens-938839638.html) for Jira

## Important Notes
- Only populated data will be added to the comment

### Step arguments

Name|Required|Description
---|---|---
ANNOTATION_NAME | No | Can customize the name of the build number annotation

### Codefresh.yml

```yaml
version: '0.1.0'
steps:
  JiraIssueComment:
    title: Jira Issue Comment
    type: jira-issue-comment
    arguments:
      ANNOTATION_NAME: '${{CF_BRANCH}}'
```