# api-parking-system

Parking System API developed using Golang

## Running Locally

### Setup Environment Variables

There are multiple environment variables that need to be setup in .env files

1. Mongo DB Credentials, including:
    - MONGO_URI : Connection String of form mongodb+srv://{username}:{password}@{cluster}
    - MONGO_DBNAME : Database which will be used
2. Google Cloud Storage
    - GCS_PROJECT_ID : The Project ID of GCP being used
    - GCS_BUCKET_NAME : The name of the bucket
    - GCS_UPLOAD_PATH : Path where images will be stored
3. Environment:
    - ENV: Set DEV for development environment
    - GOOGLE_APPLICATION_CREDENTIALS_DEV : Used for authorization to GCS with service account key, fill in service account .json file location
4. Utility:
    - JWT_SECRET: Jwt secret for signing jwt

### Connecting to GCP Environment

To connect to GCP, service account authentication is needed.

1. Download serviceAccount.json from GCP
2. Set GOOGLE_APPLICATION_CREDENTIALS_DEV as service account path locally, preferably in secrets/ folder

## Deployment To GCP

### Setup GCP

To upload to GCP we need to authenticate

1. Installing SDK https://cloud.google.com/sdk?hl=en
2. Authenticating
    ```
    gcloud auth login
    ```
3. Set Project
    ```
    gcloud config set project {PROJECT_ID}
    ```

### Building Docker Image

1. Make sure modules have been installed

    ```
    go mod tidy
    ```

2. Build the image

    ```
    docker build -t go-parking-system:$TAG .
    ```

    - **TAG** should be of form **[MAJOR]:[MINOR]**
    - **.** describes the location of the Dockerfile

### Dev: Running Dockerfile locally

1. Setup Env file and serviceAccount.json as in [Running Locally]
2. Run dockerfile:

    ```
    docker run --env-file .env -p 8080:8080 -v ${pwd}/secrets:/secrets go-parking-system:$TAG
    ```

    - --env-file to load environment variables from file (when running container, can ignore Error loading .env file error)
    - -p forwards port 8080 from host to 8080 in container
    - -v attaches volume from host (i.e. secrets folder) to secrets in container

### Pushing to Artifact Registry

0. Authenticate artifact registry

    ```
    gcloud auth configure-docker asia-southeast2-docker.pkg.dev

    gcloud artifacts repositories describe {REGISTRY_NAME} --project={PROJECT_ID} --location=asia-southeast2
    ```

1. Tag the image

    ```
    docker tag go-parking-system:$TAG asia-southeast2-docker.pkg.dev/parking-system-417615/docker/go-parking-system:$TAG
    ```

2. Push to Artifact Registry

    ```
    docker push asia-southeast2-docker.pkg.dev/parking-system-417615/docker/go-parking-system:$TAG
    ```
