# source this file to update settings

# WARNING
# treat this file as a template
# do not add it to version control if it contains any 
# private information such as passwords, addresses etc

# set the required environment variables

# base path of RETROSHEET project code
export RETROSHEET="${HOME}/projects/baseball/retrosheet"

# url to mongodb server
# local
export RETROSHEET_MONGO="mongodb://localhost:27017"
# example for mongodb atlas
# export RETROSHEET_MONGO="mongodb+srv://<username>:<password>@cluster0.<cluster id>.mongodb.net/<database name>?retryWrites=true&w=majority"

# based path to retrosheet data files
export RETROSHEET_DATA="${HOME}/projects/baseball/data"
