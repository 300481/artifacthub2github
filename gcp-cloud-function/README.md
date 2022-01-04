# artifacthub2github

Google Cloud Function to receive a webhook from Artifact Hub and forward it as a GitHub repository dispatch event 

![](artifacthub2github-drawio.png)

## needed environment variables

|variable|description                   |
|--------|------------------------------|
|TOKEN   |The GitHub auth token         |
|OWNER   |The GitHub owner              |
|REPO    |The GitHub repository         |
|SECRET  |The ArtifactHub Webhook Secret|
