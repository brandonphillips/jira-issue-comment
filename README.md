# Codefresh Jira Issue Comment

This step will create a comment with build information on your Jira instance with the specified issue key. The Jira issue key can be set directly or by using a regex pattern against a source field. The inserted Jira comment id

## Jira Issue Regex Sample
JIRA_ISSUE_SOURCE_FIELD="branch/feature-SA-19/testing"
JIRA_ISSUE_SOURCE_FIELD_REGEX="[a-zA-Z]{2}-\d+"

This would set the Jira Issue Key to "SA-19"

## Prerequisites

- Create an [API token](https://confluence.atlassian.com/cloud/api-tokens-938839638.html) for Jira

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