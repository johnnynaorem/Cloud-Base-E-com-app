# Enable Secret Manager API if not enabled
gcloud services enable secretmanager.googleapis.com

# Upload app.env
gcloud secrets create app-env-config --replication-policy="automatic"
gcloud secrets versions add app-env-config --data-file=app.env

# Upload credentials.json
gcloud secrets create gcp-creds --replication-policy="automatic"
gcloud secrets versions add gcp-creds --data-file=johnny-projectt-6d79dbdc53bb.json