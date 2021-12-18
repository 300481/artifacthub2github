# artifacthub2github

Google Cloud Function to receive a webhook from Artifact Hub and forward it as a GitHub repository dispatch event 

## needed environment variables

|variable|description          |
|--------|---------------------|
|TOKEN   |The GitHub auth token|
|OWNER   |The GitHub owner     |
|REPO    |The GitHub repository|