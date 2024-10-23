FileUploadBot
Prerequisites:

Slack Account & Channel:
        You must have a Slack account.
        Ensure you have a channel in Slack where you want to upload files.

 Create a Slack App:
        
        api.slack.com/apps
 Click "Create New App" to start building your app.

 OAuth & Permissions:
        Once your app is created, navigate to the OAuth & Permissions section.
        Choose the required OAuth Scopes your bot will need (such as files:write to upload files, channels:read to read channel info).
        After selecting the scopes, click "Save Changes."

  Generate OAuth Token:
        Scroll down to OAuth Tokens for Your Workspace.
        Click "Install App to Workspace" to generate the OAuth token.
        Copy the OAuth token provided and add it to your code.

  Add OAuth Token to Code:
        Insert the generated OAuth token into the environment variable SLACK_BOT_TOKEN in your code.
