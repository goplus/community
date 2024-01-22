# run commnuity
cd cmd/gopcomm && cp /config/.community-env .env && go mod tidy && gop run . &

# run account
cd ../account/cmd && cp /config/.account-env .env && go mod tidy && gop run .